package service

import (
	"github.com/ybkuroki/go-webapp-sample/container"
	"github.com/ybkuroki/go-webapp-sample/model"
)

// FoodService is a service for managing master data such as format and food.
type FoodService interface {
	FindAllFoods() *[]model.Food
}

type foodService struct {
	container container.Container
}

// NewFoodService is constructor.
func NewFoodService(container container.Container) FoodService {
	return &foodService{container: container}
}

// FindAllFoods returns the list of all foods.
func (m *foodService) FindAllFoods() *[]model.Food {
	rep := m.container.GetRepository()
	food := model.Food{}
	result, err := food.FindAll(rep)
	if err != nil {
		m.container.GetLogger().GetZapLogger().Errorf(err.Error())
		return nil
	}
	return result
}
