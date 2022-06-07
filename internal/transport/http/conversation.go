package http

import (
	"github.com/eugene-krivtsov/idler/internal/model/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) initConversationRoutes(api *gin.RouterGroup) {
	chat := api.Group("/conversation", h.userIdentity)
	{
		chat.POST("/", h.createConversation)
		chat.GET("/", h.getAllConversations)
		chat.GET("/:id", h.getConversationById)
		chat.DELETE("/:id", h.deleteConversationById)
	}
}

func (h *Handler) createConversation(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var conversationDTO dto.ConversationDTO
	if err := c.BindJSON(&conversationDTO); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err = h.dialogService.CreateConversation(c, id, conversationDTO)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Handler) getAllConversations(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	dialogs, err := h.dialogService.GetAll(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dialogs)
}

func (h *Handler) getConversationById(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	dialog, err := h.dialogService.GetById(c, id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, dialog)
}

func (h *Handler) deleteConversationById(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.dialogService.RemoveById(c, id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}
