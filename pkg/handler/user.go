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

	if id != input.Id {
		newErrorResponse(c, http.StatusForbidden, "forbidden")
		return
	}

	if err := h.services.Authorization.UpdateUser(id, input); err != nil {
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

	if err := h.services.Authorization.DeleteUser(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{Status: "ok"})
}
