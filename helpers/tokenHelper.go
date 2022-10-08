package helper

import (
	"fmt"
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type SignedDetails struct {
	Email string
	Name  string
	Phone string
	Id    string
	jwt.StandardClaims
}

var SECRET_KEY string = os.Getenv("JWT_SECRET_KEY")
var TOKEN_EXPIRES_AT = time.Now().Local().Add(time.Hour * time.Duration(24)).Unix()

func GenerateAllTokens(email string, name string, phone, id string) (signedToken string, err error) {
	claims := &SignedDetails{
		Email: email,
		Name:  name,
		Phone: phone,
		Id:    id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: TOKEN_EXPIRES_AT,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return
	}
	return token, err
}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)
	if err != nil {
		msg = err.Error()
		return
	}
	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = fmt.Sprintf("INVALID TOKEN")
		msg = err.Error()
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("TOKEN EXPIRED")
		msg = err.Error()
		return
	}
	return claims, msg
}
