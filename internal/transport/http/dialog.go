package http

import (
	"github.com/eugene-krivtsov/idler/internal/model/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) initDialogRoutes(router *gin.Engine) {
	chat := router.Group("/dialog", h.userIdentity)
	{
		chat.POST("/", h.createDialog)
		chat.GET("/", h.getAllDialogs)
		chat.GET("/:id", h.getDialogById)
		chat.DELETE("/:id", h.deleteDialogById)
	}
}

func (h *Handler) createDialog(c *gin.Context) {
	var dialogDTO dto.DialogDTO
	if err := c.BindJSON(&dialogDTO); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	var err = h.dialogService.CreateDialog(c, dialogDTO)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Handler) getAllDialogs(c *gin.Context) {
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

func (h *Handler) getDialogById(c *gin.Context) {
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

func (h *Handler) deleteDialogById(c *gin.Context) {
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
