package server

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/HanmaDevin/image-processing/database"
	"github.com/HanmaDevin/image-processing/server/middleware"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var (
	jwtSecret = os.Getenv("JWT_SECRET")
)

func NewServer() *echo.Echo {
	e := echo.New()

	e.POST("/register", signup)
	e.POST("/login", login)

	safe := e.Group("/images")
	safe.Use(middleware.AuthMiddleware)
	safe.POST("", upload)
	safe.POST("/:id/transform", transform)
	safe.GET("/:id", retrieve)

	return e
}

func StartServer(e *echo.Echo) {
	e.Logger.Fatal(e.Start(":8000"))
}

func listImages(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "Not implemented yet")
}

func retrieve(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "Not implemented yet")
}

func transform(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "Not implemented yet")
}

func upload(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "Not implemented yet")
}

func login(c echo.Context) error {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.Bind(&creds); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user, err := database.LoginUser(creds.Username, creds.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 2).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Could not generate token")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": fmt.Sprintf("Welcome, %s!", user.Username),
		"token":   tokenString,
	})
}

func signup(c echo.Context) error {
	var user database.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := database.RegisterUser(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	output := fmt.Sprintf("%s registered successfully", user.Username)
	return c.JSON(http.StatusCreated, output)
}
