package handler

import (
	"github.com/eugene-krivtsov/idler/internal/model/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) initUsersRoutes(router *gin.Engine) {
	auth := router.Group("/users")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
}

func (h *Handler) signUp(c *gin.Context) {
	var signUpDTO dto.SignUpDTO
	if err := c.BindJSON(&signUpDTO); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	var err = h.userService.RegisterUser(c.Request.Context(), signUpDTO)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.Status(http.StatusCreated)
}

func (h *Handler) signIn(c *gin.Context) {

}
