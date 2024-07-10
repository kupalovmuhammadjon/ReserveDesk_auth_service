package api

import (
	"auth_service/api/handler"
	"auth_service/api/middleware"
	"database/sql"

	"github.com/gin-gonic/gin"
)



func NewRouter(db *sql.DB) *gin.Engine{
	
	h := handler.NewHendler(db)

	router := gin.Default()

	router.Use(middleware.JWTMiddleware())

	users := router.Group("/auth")
	users.POST("/registor", h.Register)
	users.POST("/login", h.Login)

	return router
}