package rest

import (
	"github.com/eugene-krivtsov/idler/internal/service"
	"github.com/eugene-krivtsov/idler/pkg/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	userService  service.Users
	tokenManager auth.TokenManager
}

func NewHandler(userService service.Users, tokenManager auth.TokenManager) *Handler {
	return &Handler{
		userService:  userService,
		tokenManager: tokenManager,
	}
}

func (h *Handler) Init(host, port string) *gin.Engine {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	//docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", host, port) // TODO:
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // TODO:

	router.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	h.initUsersRoutes(router)

	return router
}
