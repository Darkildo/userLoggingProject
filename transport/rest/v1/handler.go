package v1

import (
	"github.com/gin-gonic/gin"
	"userLoggingProject/internal/core/service_provider"
)

type Handler struct {
	services *service_provider.ServiceProvider
}

func NewHandler(services *service_provider.ServiceProvider) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initLogRoutes(v1)
	}

}
