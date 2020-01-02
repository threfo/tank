package util

import (
	"net/http"
	"strings"
	jwt "github.com/dgrijalva/jwt-go"
)

//jwt token claims
type Claims struct {
	Id_      string `json:"user_id"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	UserType int    `json:"user_type"`
	Exp      int    `json:"exp"`
	jwt.StandardClaims
}

//get jwt token from request
func GetJwtTokenFromRequest(r *http.Request) string {
	var credential string
	credential = r.Header.Get("Authorization")
	if credential != "" {
		credential = strings.Trim(credential, " ")
		if strings.HasPrefix(credential, "token") || strings.HasPrefix(credential, "Bearer") {
			credential = strings.Split(credential, " ")[1]
		}
	}
	return credential
}

//parse jwt token
func ParseJwtToken(credential string) *jwt.Token {
	var token *jwt.Token
	jwtSecret := "bello_hunter"

	t, err := jwt.ParseWithClaims(credential, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err == nil && t != nil {
		if _, ok := t.Claims.(*Claims); ok && t.Valid {
			token = t
		}
	}
	return token
} 
