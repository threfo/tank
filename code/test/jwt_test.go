package test

import (
	"fmt"
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Id_      string `json:"user_id"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	UserType int    `json:"user_type"`
	Exp      int    `json:"exp"`
	jwt.StandardClaims
}

func TestParseJwt(t *testing.T) {

	secret := "bello_hunter"

	ss := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoiNWRiODBhMjAxZjJjYmUwMDAxZTIwODg1IiwibW9iaWxlIjoiMTM2MTI5NjMxMDQiLCJlbWFpbCI6bnVsbCwicm9sZSI6bnVsbCwidXNlcl90eXBlIjoxLCJwZXJtaXNzaW9ucyI6W10sImV4cCI6MTU3NzQzMjg3NywiaXNfYXV0aF9jbG91ZF9yZXN1bWUiOmZhbHNlfQ.GkgcSdLzLGkjSEIq0ajJMQkH2gS4fB8_YWVfPT8cKm4"
	token, err := jwt.ParseWithClaims(ss, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err == nil && token != nil {
		if claim, ok := token.Claims.(*Claims); ok && token.Valid {
			fmt.Println(claim)
		} else {
			fmt.Println("0")
		}
	} else {
		fmt.Println("1")
	}
}
