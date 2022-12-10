package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ybkuroki/go-webapp-sample/container"
	"github.com/ybkuroki/go-webapp-sample/controller"

	_ "github.com/ybkuroki/go-webapp-sample/docs" // for using echo-swagger
)

// Init initialize the routing of this application.
func Init(e *echo.Echo, container container.Container) {
	setCORSConfig(e, container)

	setErrorController(e, container)
	setMealController(e, container)
	setFoodController(e, container)
}

func setCORSConfig(e *echo.Echo, container container.Container) {
	if container.GetConfig().Extension.CorsEnabled {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowCredentials: true,
			AllowOrigins:     []string{"*"},
			AllowHeaders: []string{
				echo.HeaderAccessControlAllowHeaders,
				echo.HeaderContentType,
				echo.HeaderContentLength,
				echo.HeaderAcceptEncoding,
			},
			AllowMethods: []string{
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodDelete,
			},
			MaxAge: 86400,
		}))
	}
}

func setErrorController(e *echo.Echo, container container.Container) {
	errorHandler := controller.NewErrorController(container)
	e.HTTPErrorHandler = errorHandler.JSONError
	e.Use(middleware.Recover())
}

func setMealController(e *echo.Echo, container container.Container) {
	Meal := controller.NewMealController(container)
	e.GET(controller.APIMealsID, func(c echo.Context) error { return Meal.GetMeal(c) })
	e.GET(controller.APIMeals, func(c echo.Context) error { return Meal.GetMealList(c) })
	e.POST(controller.APIMeals, func(c echo.Context) error { return Meal.CreateMeal(c) })
	e.PUT(controller.APIMealsID, func(c echo.Context) error { return Meal.UpdateMeal(c) })
	e.DELETE(controller.APIMealsID, func(c echo.Context) error { return Meal.DeleteMeal(c) })
}

func setFoodController(e *echo.Echo, container container.Container) {
	food := controller.NewfoodController(container)
	e.GET(controller.APICategories, func(c echo.Context) error { return food.GetfoodList(c) })
}

func setUserController(e *echo.Echo, container container.Container) {
	user := controller.NewUserController(container)
	e.GET(controller.APIUserLoginStatus, func(c echo.Context) error { return user.GetLoginStatus(c) })
	e.GET(controller.APIUserLoginUser, func(c echo.Context) error { return user.GetLoginUser(c) })

	if container.GetConfig().Extension.SecurityEnabled {
		e.POST(controller.APIUserLogin, func(c echo.Context) error { return user.Login(c) })
		e.POST(controller.APIUserLogout, func(c echo.Context) error { return user.Logout(c) })
	}
}
