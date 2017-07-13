package hjwt

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
	jwt "github.com/dgrijalva/jwt-go"
)

var (
	key []byte = []byte("apiServer@gmail.com")
)

type Claims struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

// 产生json web token
func CreateToken(id int, name string, isAdmin bool) string {
	claims := Claims{
		id,
		name,
		isAdmin,
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix()),
			ExpiresAt: int64(time.Now().Unix() + 1000),
			Issuer:    "apiServer",
		},
	}
	/*
		token := jwt.New(jwt.SigningMethodHS256)
		//Headers
		token.Header["alg"] = "HS256"
		token.Header["typ"] = "JWT"
		//Claims
		//token.Claims["name"]
	*/
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(key)
	if err != nil {
		logs.Error(err)
		return ""
	}
	return signedToken
}

// 校验token是否有效
func CheckToken(signedToken string) (valid, isAdmin bool) {
	/*
		_, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
			return key, nil
		})
		if err != nil {
			fmt.Println("parase with claims failed.", err)
			return false
		}
	*/
	token, err := jwt.ParseWithClaims(signedToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return false, false
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		fmt.Println(claims.Id, claims.Name, claims.Admin)
		fmt.Println(claims.Issuer)
		if claims.Admin {
			return true, true
		}
		return true, false
	}
	return false, false
}

func isTokenValid(signedToken string) bool {
	token, err := jwt.ParseWithClaims(signedToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return false
	}
	if _, ok := token.Claims.(*Claims); ok && token.Valid {
		return true
	}
	return false
}

func isAdmin(signedToken string) bool {
	token, err := jwt.ParseWithClaims(signedToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return false
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		if claims.Admin {
			return true
		}
	}
	return false
}
