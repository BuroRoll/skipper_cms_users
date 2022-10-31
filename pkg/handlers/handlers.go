package handlers

import (
	"Skipper_cms_users/pkg/docs"
	service "Skipper_cms_users/pkg/servises"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() {
	router := gin.Default()
	router.Use(corsMiddleware())

	api_v1 := router.Group("/api/v1", h.userIdentity)
	{
		users := api_v1.Group("/users")
		{
			users.GET("/", h.getUsers)
			users.POST("/add-role", h.addRoleToUser)
			users.POST("/register", h.registerUser)
			users.DELETE("/delete-role", h.deleteUserRole)
		}
		roles := api_v1.Group("/roles")
		{
			roles.GET("/", h.getRoles)
		}
	}
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run()
}
