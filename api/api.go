package api

import (
    "auth_service/api/handler"
    "database/sql"
    "github.com/gin-gonic/gin"
    _ "auth_service/api/docs"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)

// @title ReserveDesk API
// @version 1.0
// @description ReserveDesk is program to book seats in restaurants order food before arrival.

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /auth

func NewRouter(db *sql.DB) *gin.Engine {
    h := handler.NewHendler(db)

    router := gin.Default()

    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    users := router.Group("/auth")
    users.POST("/register", h.Register)
    users.POST("/login", h.Login)
    users.POST("/logout", h.Logout)

    return router
}
