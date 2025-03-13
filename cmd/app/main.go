package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"simple_api/internal/database"
	"simple_api/internal/handlers"
	"simple_api/internal/taskService"
	"simple_api/internal/userService"
	"simple_api/internal/web/tasks"
	"simple_api/internal/web/users"
)

func main() {
	database.InitDB()

	tasksRepo := taskService.NewTaskRepository(database.DB)
	tasksService := taskService.NewService(tasksRepo)
	tasksHandler := handlers.NewHandler(tasksService)

	usersRepo := userService.NewUserRepository(database.DB)
	usersService := userService.NewUserService(usersRepo)
	usersHandler := handlers.NewUserHandler(usersService)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	tasksStrictHandler := tasks.NewStrictHandler(tasksHandler, nil)
	tasks.RegisterHandlers(e, tasksStrictHandler)

	usersStrictHandler := users.NewStrictHandler(usersHandler, nil)
	users.RegisterHandlers(e, usersStrictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
