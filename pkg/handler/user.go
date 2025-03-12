package handler

import (
	"net/http"

	"github.com/Nurt0re/chatik"
	"github.com/gin-gonic/gin"
)

func (h *Handler) updateUser(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input chatik.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Updater.UpdateUser(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{Status: "ok"})
}

func (h *Handler) deleteUser(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.services.Updater.DeleteUser(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{Status: "ok"})
}

func (h *Handler) getUser(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := h.services.Updater.GetUser(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":    user.ID,
		"username":  user.Username,
		"email": user.Email,
	})
}

type ForUser struct {
	ID    int
	Username  string
	Email string
}

func (h *Handler) getAllUsers(c *gin.Context) {
	users, err := h.services.Updater.GetAllUsers()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	forusers := make([]ForUser, len(users))
	for i, user := range users {
		forusers[i].ID = user.ID
		forusers[i].Username = user.Username
		forusers[i].Email = user.Email
	}
		
	
	c.JSON(http.StatusOK, forusers)
}
