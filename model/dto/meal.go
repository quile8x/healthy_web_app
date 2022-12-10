package dto

import (
	"encoding/json"
	"time"

	"github.com/ybkuroki/go-webapp-sample/model"
	"gopkg.in/go-playground/validator.v9"
)

const (
	required string = "required"
)

const (
	ValidationErrMessageMealName string = "Please enter the name with 3 to 50 characters."
	ValidationErrMessageDefault string = "This field is required."
)

// MealDto defines a data transfer object for Meal.
type MealDto struct {
	meal_name string    `validate:"required" json:"meal_name"`
	user_id   string    `validate:"required" `json:"food_id"`
	food_id   uint      `validate:"required" `json:"food_id"`
	meal_at   time.Time `validate:"required" `json:"meal_at"`
}

// NewMealDto is constructor.
func NewMealDto() *MealDto {
	return &MealDto{}
}

// Create creates a Meal model from this DTO.
func (m *MealDto) Create() *model.Meal {
	return model.NewMeal(m.user_id, m.food_id, m.meal_name, m.meal_at)
}

// Validate performs validation check for the each item.
func (b *MealDto) Validate() map[string]string {
	return validateDto(b)
}

func validateDto(b interface{}) map[string]string {
	err := validator.New().Struct(b)
	if err == nil {
		return nil
	}

	errors := err.(validator.ValidationErrors)
	if len(errors) == 0 {
		return nil
	}

	return createErrorMessages(errors)
}

func createErrorMessages(errors validator.ValidationErrors) map[string]string {
	result := make(map[string]string)
	for i := range errors {
		switch errors[i].StructField() {
		case "meal_name":
			switch errors[i].Tag() {
			case required:
				result["meal_name"] = ValidationErrMessageMealName
			}
		case "user_id":
			switch errors[i].Tag() {
			case required:
				result["user_id"] = ValidationErrMessageDefault
			}
		}
	}
	return result
}

// ToString is return string of object
func (u *MealDto) ToString() (string, error) {
	bytes, err := json.Marshal(u)
	return string(bytes), err
}
