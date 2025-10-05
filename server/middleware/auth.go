package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	expire, err := token.Claims.GetExpirationTime()
	if err != nil {
		return fmt.Errorf("Could not get expiration date")
	}
	if expire.Before(time.Now()) {
		return fmt.Errorf("Token is expired please login again")
	}

	return nil
}

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, "Missing authorization header")
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		err := verifyToken(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Invalid Token")
		}

		return next(c)
	}
}
