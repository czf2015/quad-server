package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"goserver/libs/conf"
)

var jwtSecret = []byte(conf.GetSectionKey("app", "JWT_SECRET").String())

type Claims struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Persistence string `json:"persistence"`
	jwt.StandardClaims
}

func GenerateToken(id, name, persistence string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		id,
		name,
		persistence,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "lcdp",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func RefreshToken(claims Claims) string {
	refreshTime := time.Now().Add(15 * time.Minute).Unix()
	if claims.ExpiresAt < refreshTime {
		token, err := GenerateToken(claims.Id, claims.Name, claims.Persistence)
		if err == nil {
			return token
		}
	}
	return ""
}

func ExpireToken(claims Claims) string {
	expiresAt := time.Now().Add(-1 * time.Second).Unix()
	expiredClaims := Claims{
		claims.Id,
		claims.Name,
		claims.Persistence,
		jwt.StandardClaims{
			ExpiresAt: expiresAt,
			Issuer:    "lcdp",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, expiredClaims)
	if token, err := tokenClaims.SignedString(jwtSecret); err != nil {
		return ""
	} else {
		return token
	}
}

func ParseToken(token string) (*Claims, error) {
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
