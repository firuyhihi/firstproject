package utils

import "github.com/golang-jwt/jwt"

type MyClaims struct {
	jwt.StandardClaims
	UserId     string `json:"user_id"`
	Email      string `json:"Email"`
	AccessUUID string `json:"AccessUUID"`
}
