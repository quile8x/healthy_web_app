package service

import (
	"errors"

	"github.com/ybkuroki/go-webapp-sample/container"
	"github.com/ybkuroki/go-webapp-sample/model"
	"github.com/ybkuroki/go-webapp-sample/model/dto"
	"github.com/ybkuroki/go-webapp-sample/repository"
	"github.com/ybkuroki/go-webapp-sample/util"
)

// MealService is a service for managing meals.
type MealService interface {
	FindByID(id string) (*model.Meal, error)
	FindAllMeals() (*[]model.Meal, error)
	FindAllMealsByPage(page string, size string) (*model.Page, error)
	FindMealsByName(meal_name string, page string, size string) (*model.Page, error)
	CreateMeal(dto *dto.MealDto) (*model.Meal, map[string]string)
}

type mealService struct {
	container container.Container
}

// NewMealService is constructor.
func NewMealService(container container.Container) MealService {
	return &mealService{container: container}
}

// FindByID returns one record matched meal's id.
func (m *mealService) FindByID(id string) (*model.Meal, error) {
	if !util.IsNumeric(id) {
		return nil, errors.New("failed to fetch data")
	}

	rep := m.container.GetRepository()
	meal := model.Meal{}
	var result *model.Meal
	var err error
	if result, err = meal.FindByID(rep, util.ConvertToUint(id)).Take(); err != nil {
		return nil, err
	}
	return result, nil
}

// FindAllMeals returns the list of all meals.
func (m *mealService) FindAllMeals() (*[]model.Meal, error) {
	rep := m.container.GetRepository()
	meal := model.Meal{}
	result, err := meal.FindAll(rep)
	if err != nil {
		m.container.GetLogger().GetZapLogger().Errorf(err.Error())
		return nil, err
	}
	return result, nil
}

// FindAllMealsByPage returns the page object of all meals.
func (m *mealService) FindAllMealsByPage(page string, size string) (*model.Page, error) {
	rep := m.container.GetRepository()
	meal := model.Meal{}
	result, err := meal.FindAllByPage(rep, page, size)
	if err != nil {
		m.container.GetLogger().GetZapLogger().Errorf(err.Error())
		return nil, err
	}
	return result, nil
}

// FindMealsByTitle returns the page object of meals matched given meal title.
func (m *mealService) FindMealsByName(meal_name string, page string, size string) (*model.Page, error) {
	rep := m.container.GetRepository()
	meal := model.Meal{}
	result, err := meal.FindByName(rep, meal_name, page, size)
	if err != nil {
		m.container.GetLogger().GetZapLogger().Errorf(err.Error())
		return nil, err
	}
	return result, nil
}

// CreateMeal register the given meal data.
func (m *mealService) CreateMeal(dto *dto.MealDto) (*model.Meal, map[string]string) {
	if errors := dto.Validate(); errors != nil {
		return nil, errors
	}

	rep := m.container.GetRepository()
	var result *model.Meal
	var err error

	if trerr := rep.Transaction(func(txrep repository.Repository) error {
		result, err = txCreateMeal(txrep, dto)
		return err
	}); trerr != nil {
		m.container.GetLogger().GetZapLogger().Errorf(trerr.Error())
		return nil, map[string]string{"error": "Failed to the registration"}
	}
	return result, nil
}

func txCreateMeal(txrep repository.Repository, dto *dto.MealDto) (*model.Meal, error) {
	var result *model.Meal
	var err error
	meal := dto.Create()

	category := model.Category{}
	if meal.Category, err = category.FindByID(txrep, dto.CategoryID).Take(); err != nil {
		return nil, err
	}

	format := model.Format{}
	if meal.Format, err = format.FindByID(txrep, dto.FormatID).Take(); err != nil {
		return nil, err
	}

	if result, err = meal.Create(txrep); err != nil {
		return nil, err
	}

	return result, nil
}
