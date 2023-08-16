package helper

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type GenerateJwtStruct struct {
	Uid      uint
	Username string
	jwt.RegisteredClaims
}

func GenerateJwtToken(g *GenerateJwtStruct, secretKey string, expire int64) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256,
		GenerateJwtStruct{
			Uid:      g.Uid,
			Username: g.Username,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expire * 1000 * 1000))),
			},
		},
	)

	token, err := claims.SignedString([]byte(secretKey))

	return token, err
}

func ParseJwtToken(tokenStr string, secretKey string) (*GenerateJwtStruct, error) {
	jwtStruct := GenerateJwtStruct{}
	_, err := jwt.ParseWithClaims(
		tokenStr,
		&jwtStruct,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		},
	)

	if err == nil {
		return &jwtStruct, nil
	} else {
		return nil, err
	}
}
