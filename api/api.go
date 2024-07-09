package api

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	router := gin.Default()

	main := router.Group("/reservedesk.uz")

	users := main.Group("/users")

	users.POST("/register")
	users.POST("/login")
	users.POST("/logout")

	return router
}
