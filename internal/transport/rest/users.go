package rest

import (
	"github.com/eugene-krivtsov/idler/internal/model/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) initUsersRoutes(router *gin.Engine) {
	users := router.Group("/users")
	{
		users.POST("/sign-up", h.signUp)
		users.POST("/sign-in", h.signIn)

		authenticated := users.Group("/", h.userIdentity)
		{
			authenticated.GET("/bb", h.getAll)
		}
	}
}

func (h *Handler) signUp(c *gin.Context) {
	var signUpDTO dto.SignUpDTO
	if err := c.BindJSON(&signUpDTO); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	var err = h.userService.SignUp(signUpDTO)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Handler) signIn(c *gin.Context) {
	var signInDTO dto.SignInDTO
	if err := c.BindJSON(&signInDTO); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := h.userService.SignIn(signInDTO)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"token": token,
	})
}

func (h *Handler) getAll(c *gin.Context) {
	//userId, err := getUserId(c)
	//if err != nil {
	//	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	//	return
	//}

	users, err := h.userService.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}
