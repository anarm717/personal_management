package models

import (
	"errors"
	"gorm.io/gorm"
	"html"
	"strings"
)

type Category struct {
	ID   uint64 `gorm:"primary_key;auto_increment" json:"id";`
	Name string `gorm:"size:255;not null;unique" json:"name"`
}
type CategoryName struct{
	Name string `json: "name"`
}
//swagger:parameters Category
type CategoruRequest struct {
	// in: body
	// required: true
	CategoryName CategoryName
}
// swagger:response categoryResponse
type categoryResponse struct {
	// in: body
	Body []Category
}
func (Category) TableName() string {
	return "work_with_staff.category"
}

func (p *Category) Prepare() {
	p.ID = 0
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
}

func (p *Category) Validate() error {

	if p.Name == "" {
		return errors.New("Required Name")
	}
	return nil
}

func (p *Category) GetAllCategory(db *gorm.DB) (*[]Category, error) {
	var err error
	categories := []Category{}
	err = db.Debug().Model(&Category{}).Find(&categories).Error
	if err != nil {
		return &[]Category{}, err
	}
	return &categories, nil
}

func (p *Category) FindCategoryByID(db *gorm.DB, pid uint64) (*Category, error) {
	var err error
	err = db.Debug().Model(&Category{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Category{}, err
	}
	return p, nil
}

func (p *Category) SaveCategory(db *gorm.DB) (*Category, error) {
	var err error
	err = db.Debug().Model(&Category{}).Create(&p).Error
	if err != nil {
		return &Category{}, err
	}
	return p, nil
}
