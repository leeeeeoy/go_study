package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/leeeeeoy/go_study/db"
	"github.com/leeeeeoy/go_study/dto"
	"github.com/leeeeeoy/go_study/handler"
	"github.com/leeeeeoy/go_study/repository"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/newrelic/go-agent/v3/newrelic"

	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app, _ := newrelic.NewApplication(
		newrelic.ConfigAppName("Amazing Go server"),
		newrelic.ConfigLicense(os.Getenv("DB_PASSWORD")),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)

	client := db.InitDB()
	defer client.Close()

	e := echo.New()
	e.Use(nrecho.Middleware(app))
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
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusOK, map[string]string{
				"code":    "-1",
				"message": "잘못된 파라미터 요청입니다.",
				"err":     fmt.Sprint(err),
			})
		},
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
