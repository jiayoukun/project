package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var JwtSecret = []byte("ABAB")

type Claims struct {
	Id       uint   `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	jwt.StandardClaims
}

//签发token
func GenerateToken(id uint, username, password string) (string, error) {
	notTime := time.Now()
	expireTime := notTime.Add(24 * time.Hour)
	claims := Claims{
		Id:       id,
		UserName: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "todo_list",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //加密token
	token, err := tokenClaims.SignedString(JwtSecret)
	return token, err
}

//验证token
func ParseToken(token string) (*Claims, error) {
	tokenClaim, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) { //解析token
		return JwtSecret, nil
	})
	if tokenClaim != nil {
		if claims, ok := tokenClaim.Claims.(*Claims); ok && tokenClaim.Valid {
			return claims, nil
		}
	}
	return nil, err
}
