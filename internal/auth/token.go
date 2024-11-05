package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/shoshta73/homehub/internal/models/user"
)

type userClaims struct {
	Username string `json:"username"`
	Id       string `json:"id"`
	jwt.RegisteredClaims
}

func (c *userClaims) Valid() error {
	e, err := user.UsernameExists(c.Username)
	if err != nil {
		return err
	}
	if !e {
		return errors.New("username does not exist")
	}

	e, err = user.IdExists(c.Id)
	if err != nil {
		return err
	}

	if !e {
		return errors.New("id does not exist")
	}

	return nil
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

func validateToken(tokenString string) bool {
	token, err := jwt.ParseWithClaims(tokenString, &userClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	_, ok := token.Claims.(*userClaims)
	if !token.Valid || !ok {
		return false
	}

	if err != nil {
		logger.Error("Failed to parse token", err)
		return false
	}

	return true
}
