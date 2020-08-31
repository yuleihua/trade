package biz

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func CheckJWTToken(c echo.Context) string {
	user := c.Get("_user").(*jwt.Token)
	uid := user.Claims.(jwt.MapClaims)["uid"].(string)

	return uid
}

func MakeJWTToken(uid, name string) (t string, e error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = uid
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 12).Unix()

	t, e = token.SignedString([]byte("hello-88773dy2"))

	return
}
