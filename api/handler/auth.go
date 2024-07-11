package handler

import (
	"auth_service/api/token"
	"auth_service/models"
	"net/http"

	pb "auth_service/genproto/auth"

	"github.com/gin-gonic/gin"
)

func (h *Hendler) Register(c *gin.Context) {
	req := &pb.User{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.Auth.Register(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Hendler) Login(c *gin.Context) {
	req := &models.User{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.Auth.Login(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	auth := token.GenerateJWT(resp)

	c.JSON(http.StatusOK, auth)
}

func (h *Hendler) Logout(c *gin.Context) {
	token := c.GetHeader("Authorization")
	
	if len(token) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "token is empty",
		})
		return 
	}

	err := h.Auth.Logout(&pb.Token{Token: token})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}



