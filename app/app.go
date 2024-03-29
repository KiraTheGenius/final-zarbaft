package app

import (
	"bankai/controllers/usercontroller"
	"bankai/repository/userRepository"
	"bankai/services/userService"

	"bankai/middlewares"

	"github.com/labstack/echo"
)

type App struct {
	E *echo.Echo
}

func NewApp() *App {
	e := echo.New()
	routing(e)
	return &App{
		E: e,
	}
}

func (a *App) Start(addr string) error {
	a.E.Logger.Fatal(a.E.Start(addr))
	return nil
}

func routing(e *echo.Echo) {
	userRepo := userRepository.NewGormUserRepository()
	UserService := userService.NewUserService(userRepo)
	UserController := usercontroller.UserController{UserService: UserService}
	// public routing
	e.POST("/signup", UserController.Signup)
	e.POST("/login", UserController.Login)
	e.POST("/token", UserController.GetToken)
	e.GET("/:locale/slider/:tag", UserController.Slider)
	e.GET("/:locale/collection/:slug", UserController.GetCollection)

	// protected routing
	e.GET("/now", UserController.GetTime, middlewares.IsLoggedIn, middlewares.IsAdmin)
}
