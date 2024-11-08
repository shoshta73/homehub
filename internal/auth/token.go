package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/shoshta73/homehub/internal/models/user"
)

type UserClaims struct {
	Username    string `json:"username"`
	Id          string `json:"id"`
	Permissions uint8  `json:"permissions"`
}

type userClaims struct {
	UserClaims
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

	if c.Permissions < 1 {
		return errors.New("invalid permissions")
	}

	if !e {
		return errors.New("id does not exist")
	}

	return nil
}

func generateToken(user *user.User) string {
	tn := time.Now()

	claims := userClaims{
		UserClaims: UserClaims{
			Username:    user.Username,
			Id:          user.ID,
			Permissions: user.Permissions,
		},
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
func extractClaims(tokenString string) (*userClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &userClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	c, ok := token.Claims.(*userClaims)
	if !token.Valid || !ok {
		return nil, errors.New("invalid token")
	}

	if err != nil {
		return nil, err
	}

	return c, nil
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

func GetClaims(tokenString string) (*UserClaims, error) {
	c, err := extractClaims(tokenString)

	if err != nil {
		return nil, err
	}

	return &c.UserClaims, nil
}
