package handler

import (
	_ "auth_service/api/docs"
	"auth_service/api/token"
	"auth_service/models"
	"database/sql"
	"net/http"

	pb "auth_service/genproto/auth"

	"github.com/gin-gonic/gin"
)

// ReserveDesk
// @Summary Register User
// @Description Registers user
// @Tags Auth
// @ID register
// @Accept json
// @Produce json
// @Param user body auth.User true "User information to create it"
// @Success 201
// @Failure 500 {object} models.Error "Something went wrong in server"
// @Router /register [post]
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
	c.JSON(http.StatusCreated, resp)
}

// ReserveDesk
// @Summary Login user
// @Description checks the user and returns tokens
// @Tags Auth
// @ID login
// @Accept json
// @Produce json
// @Param user body models.User true "User Information to log in"
// @Success 200 {object} auth.Tokens  "Returns access and refresh tokens"
// @Failure 401 {object} models.Error "if Access token fails it will returns this"
// @Failure 500 {object} models.Error "Something went wrong in server"
// @Router /login [post]
func (h *Hendler) Login(c *gin.Context) {
	req := &models.User{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.Auth.Login(req)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user does not exist"})
			h.Logger.Error(err.Error())
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	auth := token.GenerateJWT(resp)

	c.JSON(http.StatusOK, auth)
}

// ReserveDesk
// @Summary log outs user
// @Description removes refresh token gets token from header
// @Tags Auth
// @ID logout
// @Accept json
// @Produce json
// @Success 200 
// @Failure 500 {object} models.Error "Something went wrong in server"
// @Router /logout [post]
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
