package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/shoshta73/homehub/internal/models/user"
)

type userClaims struct {
	Username string `json:"username"`
	Id       string `json:"id"`
	jwt.RegisteredClaims
}

func generateToken(user *user.User) string {
	tn := time.Now()

	claims := userClaims{
		Username: user.Username,
		Id:       user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 3)),
			IssuedAt:  jwt.NewNumericDate(tn),
			NotBefore: jwt.NewNumericDate(tn),
			Issuer:    "homehub-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		logger.Error("Failed to generate token", err)
		return ""
	}

	return tokenString
}
