package token

import (
	"auth_service/config"
	pb "auth_service/genproto/auth"
	"auth_service/models"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func GenerateJWT(user *models.User) *pb.Tokens {
	accessToken := jwt.New(jwt.SigningMethodHS256)
	refreshToken := jwt.New(jwt.SigningMethodHS256)

	claims := accessToken.Claims.(jwt.MapClaims)
	claims["user_id"] = user.Id
	claims["full_name"] = user.FullName
	claims["is_admin"] = user.IsAdmin
	claims["email"] = user.Email
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	cfg := config.Load()
	access, err := accessToken.SignedString([]byte(cfg.SIGNING_KEY))
	if err != nil {
		log.Fatalf("Access token is not generated %v", err)
	}
	_, err = uuid.Parse("hg")

	rftClaims := refreshToken.Claims.(jwt.MapClaims)
	rftClaims["user_id"] = user.Id
	rftClaims["full_name"] = user.FullName
	rftClaims["is_admin"] = user.IsAdmin
	rftClaims["email"] = user.Email
	rftClaims["iat"] = time.Now().Unix()
	rftClaims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	refresh, err := accessToken.SignedString([]byte(cfg.SIGNING_KEY))
	if err != nil {
		log.Fatalf("Access token is not generated %v", err)
	}

	return &pb.Tokens{
		AccessToken:  access,
		RefreshToken: refresh,
	}
}

func ExtractClaims(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(config.Load().SIGNING_KEY), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
