package utils

import (
	"basics/internal/models"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(Cfg.JWTSecret)

func GenerateJWT(user *models.LoginResponse) (*models.LoginResponse, error) {
	claims := jwt.MapClaims{
		"email":      user.Email,
		"company_id": user.CompanyId,
		"plant_id":   user.PlantId,
		"machine_id": user.PlantId,
		"scope":      user.Scope,
		"user_id":    user.UserId,
		"exp":        time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtSecret)
	if err != nil {
		return nil, fmt.Errorf("token creation failed: %w", err)
	}

	user.Token = tokenStr
	return user, nil

}

func VerifyJwtToken(tokenStr string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		return claims, nil
	}

	return nil, errors.New("token invalid")

}
