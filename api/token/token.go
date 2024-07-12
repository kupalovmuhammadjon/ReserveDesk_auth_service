package token

import (
	"auth_service/config"
	pb "auth_service/genproto/auth"
	"auth_service/models"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(user *models.User) *pb.Tokens {

	accessToken := jwt.New(jwt.SigningMethodHS256)
	refreshToken := jwt.New(jwt.SigningMethodHS256)

	claims := accessToken.Claims.(jwt.MapClaims)
	claims["user_id"] = user.Id
	claims["fullname"] = user.FullName
	claims["is_admin"] = user.IsAdmin
	claims["email"] = user.Email
	claims["password"] = user.Password
	claims["iat"] = time.Now().Unix()
	claims["ext"] = time.Now().Add(time.Hour).Unix()

	cfg := config.Load()

	access, err := accessToken.SignedString([]byte(cfg.SIGNING_KEY))
	if err != nil {
		log.Fatalf("Access token is not generated %v", err)
	}

	rftClaims := refreshToken.Claims.(jwt.MapClaims)
	rftClaims["user_id"] = user.Id
	rftClaims["fullname"] = user.FullName
	rftClaims["is_admin"] = user.IsAdmin
	rftClaims["email"] = user.Email
	rftClaims["passowrd"] = user.Password
	rftClaims["iat"] = time.Now().Unix()
	rftClaims["ext"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	refresh, err := refreshToken.SignedString([]byte(cfg.REFRESH_SIGNING_KEY))
	if err != nil {
		log.Fatalf("Access token is not generated %v", err)
	}

	return &pb.Tokens{
		AccessToken:  access,
		RefreshToken: refresh,
	}
}

func GenerateAccessToken(user *jwt.MapClaims) *string {

	accessToken := jwt.New(jwt.SigningMethodHS256)

	claims := accessToken.Claims.(jwt.MapClaims)
	claims["user_id"] = (*user)["user_id"]
	claims["full_name"] = (*user)["full_name"]
	claims["is_admin"] = (*user)["is_admin"]
	claims["email"] = (*user)["email"]
	claims["password"] = (*user)["password"]
	claims["iat"] = time.Now().Unix()
	claims["ext"] = time.Now().Add(time.Hour).Unix()

	cfg := config.Load()

	access, err := accessToken.SignedString([]byte(cfg.SIGNING_KEY))
	if err != nil {
		log.Fatalf("Access token is not generated %v", err)
	}

	return &access
}

func ExtractClaims(tokenStr string, isRefresh bool) (jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		if isRefresh{
			return []byte(config.Load().REFRESH_SIGNING_KEY), nil
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
