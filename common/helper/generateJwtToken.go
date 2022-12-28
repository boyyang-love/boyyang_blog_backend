package helper

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type GenerateJwtStruct struct {
	Id       int
	Username string
	Password string
	jwt.RegisteredClaims
}

func GenerateJwtToken(g *GenerateJwtStruct, secretKey string, expire int64) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256,
		GenerateJwtStruct{
			Id:       g.Id,
			Username: g.Username,
			Password: g.Password,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expire * 1000 * 1000))),
			},
		},
	)

	token, err := claims.SignedString([]byte(secretKey))

	return token, err

	//claims := make(jwt.MapClaims)
	//claims["exp"] = iat + seconds
	//claims["iat"] = iat
	//claims["userId"] = g.Id
	//token := jwt.New(jwt.SigningMethodHS256)
	//token.Claims = claims
	//return token.SignedString([]byte(secretKey))
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
