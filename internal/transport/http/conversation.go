package http

import (
	"fmt"
	"github.com/eugene-krivtsov/idler/internal/model/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
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

	if dialogs != nil {
		c.JSON(http.StatusOK, dialogs)
	}
}

func (h *Handler) getConversationById(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		errorMessage := fmt.Sprintf("Illegal format of ID [%s]", id)
		newErrorResponse(c, http.StatusBadRequest, errorMessage)
		return
	}

	dialog, err := h.dialogService.GetById(c, id)
	if err != nil && err.Error() == "sql: no rows in result set" {
		errorMessage := fmt.Sprintf("Conversation not found by ID [%s]", id)
		newErrorResponse(c, http.StatusNotFound, errorMessage)
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

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		errorMessage := fmt.Sprintf("Illegal format of ID [%s]", id)
		newErrorResponse(c, http.StatusBadRequest, errorMessage)
		return
	}

	if err := h.dialogService.RemoveById(c, id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}
