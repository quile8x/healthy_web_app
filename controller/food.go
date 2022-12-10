package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/container"
	"github.com/ybkuroki/go-webapp-sample/service"
)

// FoodController is a controller for managing Food data.
type FoodController interface {
	GetFoodList(c echo.Context) error
}

type FoodController struct {
	container container.Container
	service   service.FoodService
}

// NewFoodController is constructor.
func NewFoodController(container container.Container) FoodController {
	return &FoodController{container: container, service: service.NewFoodService(container)}
}

// GetFoodList returns the list of all foods.
// @Summary Get a Food list
// @Description Get a Food list
// @Tags Food
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Food "Success to fetch a Food list."
// @Failure 401 {string} false "Failed to the authentication."
// @Router /foods [get]
func (controller *FoodController) GetFoodList(c echo.Context) error {
	return c.JSON(http.StatusOK, controller.service.FindAllFood())
}
