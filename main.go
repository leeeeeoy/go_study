package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/leeeeeoy/go_study/db"
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

	e.Logger.Fatal(e.Start(":1323"))

}
