package auth

import (
	"time"

	"github.com/astaxie/beego/logs"
	jwt "github.com/dgrijalva/jwt-go"
)

var (
	key []byte = []byte("apiServer@gmail.com")
)

type Claims struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	IsAdmin bool   `json:"IsAdmin"`
	jwt.StandardClaims
}

// 产生json web token
func CreateToken(id int, name string, isAdmin bool) string {
	claims := Claims{id, name, isAdmin, jwt.StandardClaims{
		NotBefore: int64(time.Now().Unix()),
		ExpiresAt: int64(time.Now().Unix() + 10000),
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

//check if token is valid
func IsTokenValid(signedToken string) bool {
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

//check if user is admin
func IsAdmin(signedToken string) bool {
	token, err := jwt.ParseWithClaims(signedToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return false
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		if claims.IsAdmin {
			return true
		}
	}
	return false
}

//Get id and name from cookie
func GetIdAndName(signedToken string) (int, string) {
	token, err := jwt.ParseWithClaims(signedToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return -1, ""
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.Id, claims.Name
	}
	return -1, ""
}
