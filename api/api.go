package api

import (
	"auth_service/api/handler"
	"database/sql"

	_ "auth_service/api/docs" // Swagger documentation
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// @title ReserveDesk
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /auth

// NewRouter sets up the router
func NewRouter(db *sql.DB) *gin.Engine {
	h := handler.NewHendler(db)

	router := gin.Default()

	// Serve Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	users := router.Group("/auth")
	users.POST("/register", h.Register)
	users.POST("/login", h.Login)
	users.POST("/logout", h.Logout)

	return router
}
