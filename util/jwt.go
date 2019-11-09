package util

import (
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pengchujin/subscribe_go/config"
)


type Claims struct {
	Email string `json:"email"`
	Username string `json:"username"`
	UUID string `json:"uuid"`
	jwt.StandardClaims
}

func GenerateToken(email, username string, uuid string) (string, error) {
	conf := config.Get()
  jwtSecret := []byte(conf.JwtSecret)
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		email,
		username,
		uuid,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "subscribe",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	conf := config.Get()
  jwtSecret := []byte(conf.JwtSecret)
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
