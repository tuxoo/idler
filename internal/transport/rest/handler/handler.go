package handler

import (
	"github.com/eugene-krivtsov/idler/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	userService service.Users
}

func NewHandler(userService service.Users) *Handler {
	return &Handler{
		userService: userService,
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

	return router
}
