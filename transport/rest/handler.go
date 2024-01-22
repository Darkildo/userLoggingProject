package rest

import (
	//docs "delivery/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"strconv"
	"userLoggingProject/docs"
	"userLoggingProject/internal/core/config"
	"userLoggingProject/internal/core/service_provider"
	v1 "userLoggingProject/transport/rest/v1"
)

type Handler struct {
	services *service_provider.ServiceProvider
}

func NewHandler(service *service_provider.ServiceProvider) *Handler {
	return &Handler{
		services: service,
	}
}

func (h *Handler) Init(conf *config.LaunchConfig) *gin.Engine {
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
		corsMiddleware,
	)
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Host = "localhost:" + strconv.Itoa(conf.Port)
	// Init router
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.services)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
