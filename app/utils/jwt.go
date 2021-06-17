package utils

import (
	"fmt"
	"log"
	"os"
	"sas-backend/app/database/models"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// UserClaims struct
type UserClaims struct {
	models.UserModel
	jwt.StandardClaims
}

// GenerateToken func
func GenerateToken(modelUserType models.UserTypeModel, permission interface{}) string {
	jwtExpired := os.Getenv("JWT_EXPIRED")
	uom := jwtExpired[len(jwtExpired)-1:]
	expiredToken, _ := strconv.Atoi(strings.Replace(jwtExpired, uom, "", -1))

	timeHourMinutes := time.Hour
	if uom == "d" {
		timeHourMinutes = time.Hour * 24
	}

	expiredAt := time.Now().Add(timeHourMinutes * time.Duration(expiredToken)).Unix()

	claims := UserClaims{
		models.UserModel{
			Email:     modelUserType.Email,
			RoleID:    modelUserType.RoleID,
			RoleName:  modelUserType.RoleName,
			Username:  modelUserType.Username,
			Fullname:  modelUserType.Fullname,
			Photo:     modelUserType.Photo,
			Positions: modelUserType.Positions,
			Grade:     modelUserType.Grade,
			Permision: permission,
		},
		jwt.StandardClaims{
			ExpiresAt: expiredAt,
			Issuer:    modelUserType.Username,
		},
	}

	var signingKey = []byte(os.Getenv("JWT_KEY"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signingKey)

	if err != nil {
		log.Println("Generate Token Failed..")
		return ""
	}

	return tokenString
}

// JwtDecode func
func JwtDecode(token string) (*jwt.Token, error) {
	var signingKey = []byte(os.Getenv("JWT_KEY"))

	return jwt.ParseWithClaims(token, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
}

// IsAuthorized func
func IsAuthorized(tokenString string) bool {
	token, err := JwtDecode(tokenString)
	if err != nil {
		return false
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		if claims == nil {
			return false
		}
		return true
	}

	return false
}

// JwtClaim func
func JwtClaim(tokenString string) models.UserModel {
	var signingKey = []byte(os.Getenv("JWT_KEY"))

	users := models.UserModel{}
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})

	if err != nil {
		fmt.Printf("Error JWT Claim")
	}

	users.Username = claims["username"].(string)
	users.Email = claims["email"].(string)
	users.Positions = claims["positions"].(string)
	users.Grade = claims["grade"].(string)
	users.RoleID = claims["role_id"]
	users.RoleName = claims["role_name"].(string)

	return users
}
