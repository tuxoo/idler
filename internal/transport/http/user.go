package http

import (
	"github.com/eugene-krivtsov/idler/internal/model/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) initUserRoutes(router *gin.Engine) {
	user := router.Group("/user")
	{
		user.POST("/sign-up", h.signUp)
		user.POST("/sign-in", h.signIn)

		authenticated := user.Group("/", h.userIdentity)
		{
			authenticated.GET("/profile", h.getCurrentUser)
			authenticated.GET("/", h.getAllUsers)
			authenticated.GET("/:email", h.getUserByEmail)
		}
	}
}

// @Summary      User SignUp
// @Tags         user-auth
// @Description  registering new user
// @ID           userSignUp
// @Accept       json
// @Produce      json
// @Param        input    body      dto.SignUpDTO  true  "sign up info"
// @Success      201      {string}  string         "ok"
// @Failure      400,404  {object}  errorResponse
// @Failure      500      {object}  errorResponse
// @Failure      default  {object}  errorResponse
// @Router       /user/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var signUpDTO dto.SignUpDTO
	if err := c.BindJSON(&signUpDTO); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	var err = h.userService.SignUp(c.Request.Context(), signUpDTO)
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

	token, err := h.userService.SignIn(c.Request.Context(), signInDTO)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"token": token,
	})
}

func (h *Handler) getCurrentUser(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := h.userService.GetById(c, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) getAllUsers(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	users, err := h.userService.GetAll(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *Handler) getUserByEmail(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := h.userService.GetByEmail(c, c.Param("email"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}
