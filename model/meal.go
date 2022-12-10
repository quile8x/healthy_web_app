package model

import (
	"database/sql"
	"errors"
	"math"
	"time"

	"github.com/moznion/go-optional"
	"github.com/ybkuroki/go-webapp-sample/repository"
	"github.com/ybkuroki/go-webapp-sample/util"
	"gorm.io/gorm"
)

// Meal defines struct of Meal data.
type Meal struct {
	meal_id   uint      `gorm:"primary_key" json:"id"`
	meal_name string    `json:"meal_name"`
	user_id   uint      `json:"user_id"`
	food_id   uint      `json:"food_id"`
	meal_at   time.Time `json:"meal_at"`
}

// RecordMeal defines struct represents the record of the database.
type RecordMeal struct {
	meal_id   uint
	meal_name string
	user_id   uint
	food_id   uint
	meal_at   time.Time
}

const (
	selectMeal = "select m.meal_id as meal_id, m.meal_name as meal_name, m.meal_at as meal_at, " +
		"f.food_id as food_id, f.food_name as food_name" +
		"from meals m inner join foods f on f.food_id = m.food_id"
	findByID   = " where m.meal_id = ?"
	findByName = " where meal_name like ? "
)

// TableName returns the table name of Meal struct and it is used by gorm.
func (Meal) TableName() string {
	return "Meal"
}

// NewMeal is constructor
func NewMeal(meal_name string, user_id uint, food_id uint, meal_at time.Time) *Meal {
	return &Meal{meal_name: meal_name, user_id: user_id, food_id: food_id, meal_at: meal_at}
}

// FindByID returns a Meal full matched given Meal's ID.
func (m *Meal) FindByID(rep repository.Repository, id uint) optional.Option[*Meal] {
	var rec RecordMeal
	args := []interface{}{id}

	createRaw(rep, selectMeal+findByID, "", "", args).Scan(&rec)
	return convertToMeal(&rec)
}

// FindAll returns all Meals of the Meal table.
func (m *Meal) FindAll(rep repository.Repository) (*[]Meal, error) {
	var Meals []Meal
	var err error

	if Meals, err = findRows(rep, selectMeal, "", "", []interface{}{}); err != nil {
		return nil, err
	}
	return &Meals, nil
}

// FindAllByPage returns the page object of all Meals.
func (m *Meal) FindAllByPage(rep repository.Repository, page string, size string) (*Page, error) {
	var Meals []Meal
	var err error

	if Meals, err = findRows(rep, selectMeal, page, size, []interface{}{}); err != nil {
		return nil, err
	}
	p := createPage(&Meals, page, size)
	return p, nil
}

// FindByName returns the page object of Meals partially matched given Meal title.
func (b *Meal) FindByName(rep repository.Repository, name string, page string, size string) (*Page, error) {
	var Meals []Meal
	var err error
	args := []interface{}{"%" + name + "%"}

	if Meals, err = findRows(rep, selectMeal+findByName, page, size, args); err != nil {
		return nil, err
	}
	p := createPage(&Meals, page, size)
	return p, nil
}

func findRows(rep repository.Repository, sqlquery string, page string, size string, args []interface{}) ([]Meal, error) {
	var Meals []Meal

	var rec RecordMeal
	var rows *sql.Rows
	var err error

	if rows, err = createRaw(rep, sqlquery, page, size, args).Rows(); err != nil {
		return nil, err
	}
	for rows.Next() {
		if err = rep.ScanRows(rows, &rec); err != nil {
			return nil, err
		}

		opt := convertToMeal(&rec)
		if opt.IsNone() {
			return nil, errors.New("failed to fetch data")
		}
		Meal, _ := opt.Take()
		Meals = append(Meals, *Meal)
	}
	return Meals, nil
}

func createRaw(rep repository.Repository, sql string, pageNum string, pageSize string, args []interface{}) *gorm.DB {
	if util.IsNumeric(pageNum) && util.IsNumeric(pageSize) {
		page := util.ConvertToInt(pageNum)
		size := util.ConvertToInt(pageSize)
		args = append(args, size)
		args = append(args, page*size)
		sql += " limit ? offset ? "
	}
	if len(args) > 0 {
		return rep.Raw(sql, args...)
	}
	return rep.Raw(sql)
}

func createPage(Meals *[]Meal, page string, size string) *Page {
	p := NewPage()
	p.Page = util.ConvertToInt(page)
	p.Size = util.ConvertToInt(size)
	p.NumberOfElements = p.Size
	p.TotalElements = len(*Meals)
	if p.TotalPages = int(math.Ceil(float64(p.TotalElements) / float64(p.Size))); p.TotalPages < 0 {
		p.TotalPages = 0
	}
	p.Content = Meals

	return p
}

// Save persists this Meal data.
func (m *Meal) Save(rep repository.Repository) (*Meal, error) {
	if err := rep.Save(m).Error; err != nil {
		return nil, err
	}
	return b, nil
}

// Create persists this Meal data.
func (b *Meal) Create(rep repository.Repository) (*Meal, error) {
	if err := rep.Select("user_id", "food_id", "meal_name").Create(b).Error; err != nil {
		return nil, err
	}
	return b, nil
}

func convertToMeal(rec *RecordMeal) optional.Option[*Meal] {
	if rec.meal_id == 0 {
		return optional.None[*Meal]()
	}
	f := &Food{food_id: rec.food_id, food_name: rec.food_name}
	return optional.Some(
		&Meal{meal_id: rec.meal_id, meal_name: rec.meal_name, user_id: rec.user_id, food_id: rec.food_id, Food: f})
}

// ToString is return string of object
func (b *Meal) ToString() string {
	return toString(b)
}
