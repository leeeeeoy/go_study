package main

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/leeeeeoy/go_study/db"
	"github.com/leeeeeoy/go_study/dto"
	"github.com/leeeeeoy/go_study/handler"
	"github.com/leeeeeoy/go_study/repository"
	_ "github.com/lib/pq"
)

func main() {

	client := db.InitDB()
	defer client.Close()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	ur := repository.NewUserRepository(client)
	uh := handler.NewUserHandler(ur)
	uh.InitUserHandler(e)

	adminGroup := e.Group("/admin")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(dto.JwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}

	adminGroup.Use(echojwt.WithConfig(config))

	adminGroup.GET("", func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*dto.JwtCustomClaims)
		name := claims.Name
		return c.String(http.StatusOK, "안녕하세요! "+name)
	})

	e.Logger.Fatal(e.Start(":1323"))

}
