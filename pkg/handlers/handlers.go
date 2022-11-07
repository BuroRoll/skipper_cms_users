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
			users.GET("/", Authorize("/users", "read", h), h.getUsers)
			users.GET("/info", Authorize("/info", "read", h), h.getUserInfo)

			users.POST("/register", Authorize("/register", "write", h), h.registerUser)

			users.PUT("/add-role", Authorize("/add-role", "write", h), h.addRoleToUser)
			users.DELETE("/delete-role", Authorize("/delete-role", "delete", h), h.deleteUserRole)
		}
		roles := api_v1.Group("/roles")
		{
			roles.GET("/", Authorize("/roles", "read", h), h.getRoles)
		}
		password := api_v1.Group("/password")
		{
			password.PUT("/change", h.changePassword)
		}
	}
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run()
}
