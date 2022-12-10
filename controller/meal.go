package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/container"
	"github.com/ybkuroki/go-webapp-sample/model/dto"
	"github.com/ybkuroki/go-webapp-sample/service"
)

// MealController is a controller for managing Meals.
type MealController interface {
	GetMeal(c echo.Context) error
	GetMealList(c echo.Context) error
	CreateMeal(c echo.Context) error
}

type MealController struct {
	container container.Container
	service   service.MealService
}

// NewMealController is constructor.
func NewMealController(container container.Container) MealController {
	return &MealController{container: container, service: service.NewMealService(container)}
}

// GetMeal returns one record matched Meal's id.
// @Summary Get a Meal
// @Description Get a Meal
// @Tags Meals
// @Accept  json
// @Produce  json
// @Param Meal_id path int true "Meal ID"
// @Success 200 {object} model.Meal "Success to fetch data."
// @Failure 400 {string} message "Failed to fetch data."
// @Failure 401 {boolean} bool "Failed to the authentication. Returns false."
// @Router /Meals/{Meal_id} [get]
func (controller *MealController) GetMeal(c echo.Context) error {
	Meal, err := controller.service.FindByID(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, Meal)
}

// GetMealList returns the list of matched Meals by searching.
// @Summary Get a Meal list
// @Description Get the list of matched Meals by searching
// @Tags Meals
// @Accept  json
// @Produce  json
// @Param query query string false "Keyword"
// @Param page query int false "Page number"
// @Param size query int false "Item size per page"
// @Success 200 {object} model.Page "Success to fetch a Meal list."
// @Failure 400 {string} message "Failed to fetch data."
// @Failure 401 {boolean} bool "Failed to the authentication. Returns false."
// @Router /Meals [get]
func (controller *MealController) GetMealList(c echo.Context) error {
	Meal, err := controller.service.FindMealsByTitle(c.QueryParam("query"), c.QueryParam("page"), c.QueryParam("size"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, Meal)
}

// CreateMeal create a new Meal by http post.
// @Summary Create a new Meal
// @Description Create a new Meal
// @Tags Meals
// @Accept  json
// @Produce  json
// @Param data body dto.MealDto true "a new Meal data for creating"
// @Success 200 {object} model.Meal "Success to create a new Meal."
// @Failure 400 {string} message "Failed to the registration."
// @Failure 401 {boolean} bool "Failed to the authentication. Returns false."
// @Router /Meals [post]
func (controller *MealController) CreateMeal(c echo.Context) error {
	dto := dto.NewMealDto()
	if err := c.Bind(dto); err != nil {
		return c.JSON(http.StatusBadRequest, dto)
	}
	Meal, result := controller.service.CreateMeal(dto)
	if result != nil {
		return c.JSON(http.StatusBadRequest, result)
	}
	return c.JSON(http.StatusOK, Meal)
}
