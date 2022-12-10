package model

import (
	"github.com/moznion/go-optional"
	"github.com/ybkuroki/go-webapp-sample/repository"
)

// Food defines struct of Food data.
type Food struct {
	food_id     uint    `gorm:"primary_key" json:"id"`
	food_name   string  `validate:"required" json:"food_name"`
	calo_amount float64 `validate:"required" json:"calo_amount"`
}

// TableName returns the table name of Food struct and it is used by gorm.
func (Food) TableName() string {
	return "foods"
}

// NewFood is constructor
func NewFood(food_name string) *Food {
	return &Food{food_name: food_name}
}

// Exist returns true if a given Food exits.
func (f *Food) Exist(rep repository.Repository, food_id uint) (bool, error) {
	var count int64
	if err := rep.Where("food_id = ?", food_id).Count(&count).Error; err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

// FindByID returns a Food full matched given Food's ID.
func (f *Food) FindByID(rep repository.Repository, food_id uint) optional.Option[*Food] {
	var Food Food
	if err := rep.Where("id = ?", foods).First(&Food).Error; err != nil {
		return optional.None[*Food]()
	}
	return optional.Some(&Food)
}

// FindAll returns all categories of the Food table.
func (f *Food) FindAll(rep repository.Repository) (*[]Food, error) {
	var categories []Food
	if err := rep.Find(&foods).Error; err != nil {
		return nil, err
	}
	return &categories, nil
}

// Create persists this Food data.
func (f *Food) Create(rep repository.Repository) (*Food, error) {
	if err := rep.Create(f).Error; err != nil {
		return nil, err
	}
	return f, nil
}

// ToString is return string of object
func (f *Food) ToString() string {
	return toString(f)
}
