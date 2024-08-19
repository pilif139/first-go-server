package main

import (
	"crud_server/db"
	handlers "crud_server/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	db, err := db.GetDB()
	if err != nil {
		e.Logger.Fatal(err)
	}
	db.AutoMigrate(&handlers.User{})

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	var uh = handlers.UserHandler{DB: db}
	e.GET("/users", uh.GetAllUsers)
	e.POST("/users", uh.CreateUser)
	e.GET("/users/:id", uh.GetUser)
	e.PUT("/users/:id", uh.UpdateUser)
	e.DELETE("/users/:id", uh.DeleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}
