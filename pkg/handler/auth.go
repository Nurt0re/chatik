package handler

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/Nurt0re/chatik"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	oauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURL:  "http://localhost:8080/auth/callback",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	oauthStateString = "random"
)

func (h *Handler) signUp(c *gin.Context) {
	var input chatik.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	// Показывает эти данные в алфавитном порядке, нужно пофиксить
	c.JSON(http.StatusOK, map[string]interface{}{
		"id":       id,
		"username": input.Username,
		"email":    input.Email,
	})

}

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {

	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Set the token in the Authorization header
	c.Header("Authorization", "Bearer "+token)

	c.SetCookie("jwt", token, 60*60*12, "/", "localhost", false, true)
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
		"email": input.Email,
	})
}
func (h *Handler) Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}

func (h *Handler) oAuth(c *gin.Context) {
	url := oauthConfig.AuthCodeURL(oauthStateString)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *Handler) callback(c *gin.Context) {
	state := c.Query("state")
	if state != oauthStateString {
		newErrorResponse(c, http.StatusUnauthorized, "invalid oauth state")
		return
	}

	code := c.Query("code")
	token, err := oauthConfig.Exchange(c, code)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "code exchange failed")
		return
	}

	client := oauthConfig.Client(c, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to get user info")
		return
	}
	defer resp.Body.Close()

	// Handle user info response
	userInfo := struct {
		ID    string `json:"id"`
		Email string `json:"email"`
	}{}

	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to decode user info")
		return
	}

	// Use the user info as needed, for example, create a user session
	c.JSON(http.StatusOK, userInfo)
}
