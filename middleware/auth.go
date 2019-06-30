package middleware

import (
	"encoding/hex"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type AuthContext struct {
	echo.Context
	UserId int64
}

func Auth(disAllowAnon bool) echo.MiddlewareFunc {
	if disAllowAnon {
		return verifyAuth
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := getToken(c)
			cc, _ := verifyToken(c, token)
			writeCookie(cc)
			return next(cc)
		}
	}
}

func writeCookie(c echo.Context) {
	key := "AID"
	cookie, err := c.Cookie(key)
	if err != nil {
		newCookie := new(http.Cookie)
		id, _ := uuid.NewRandom()
		newCookie.Name = key
		newCookie.Value = hex.EncodeToString(id[:])
		newCookie.Expires = time.Now().Add(24 * time.Hour)
		c.SetCookie(newCookie)
	} else {
		cookie.Expires = time.Now().Add(24 * time.Hour)
		c.SetCookie(cookie)
	}
}

func getToken(c echo.Context) string {
	authHeader := c.Request().Header.Get("Authorization")
	if !strings.Contains(authHeader, "Bearer") {
		return ""
	}
	token := strings.Replace(authHeader, "Bearer ", "", -1)
	return token
}

// dummy
func verifyAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}

func verifyToken(c echo.Context, token string) (echo.Context, error) {
	var dummyUserId int64 = 0
	cc := &AuthContext{c, dummyUserId}
	return cc, nil
}
