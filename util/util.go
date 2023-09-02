package util

import (
	"github.com/dgrijalva/jwt-go/v4"
	"fmt"
	"errors"
	"time"
  "crypto/md5"
)



type UserClaim struct {
	Id             uint
	Name           string
	jwt.StandardClaims
}

func GenerateToken(id uint, name string) (string, error) {
	uc := UserClaim{
		Id:   id,
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Second * time.Duration(1800))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)

	tokenString, err := token.SignedString([]byte("dousheng"))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return tokenString, nil
}

// 解析token
func AnalyzeToken(token string) (*UserClaim, error) {
	uc := new(UserClaim)

	claim, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte("dousheng"), nil
	})
	if err != nil {
		return nil, err
	}
	if !claim.Valid {
		return uc, errors.New("token is invalid")
	}
	return uc, nil
}

func MD5(str string)string{
  return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
