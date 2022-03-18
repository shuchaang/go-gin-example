package util

import (
	"github.com/dgrijalva/jwt-go"
	"go-gin-example/pkg/setting"
	"time"
)

var jwtSecret = []byte(setting.AppSetting.JwtSecret)

type Claims struct {
	UserName string `json:username`
	PassWord string `json:password`
	jwt.StandardClaims
}

func GenerateToken(username,password string)(string,error){
	expiredTime := time.Now().Add(24 * time.Hour)
	claims:=Claims{
		UserName: username,
		PassWord: password,
		StandardClaims: jwt.StandardClaims {
			Issuer: "gin-blog",
			ExpiresAt: expiredTime.Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string)(*Claims,error){
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		return claims, nil
	} else {
		return nil, err
	}

}