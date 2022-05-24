package common

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/stereon/aivin.com/model"
	"time"
)
var jwtKkey  = []byte("a_secret_crect")

type Cliams struct {
	UserId string `json:"user_id,omitempty"`
	jwt.StandardClaims    `json:"jwt,omitempty"`
}

func  ReleaseToken(user *model.User) (string, error) {
	expairationTime := time.Now().Add(7*24*time.Hour)
	cliams := &Cliams{
		UserId: user.Ftelphone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expairationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "aivinli.com",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,cliams)
	tokenString,err := token.SignedString(jwtKkey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenSting string) (*jwt.Token,*Cliams,error) {
	claims := &Cliams{}
	token ,err := jwt.ParseWithClaims(tokenSting, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKkey,nil
	})

	return token, claims,err
}