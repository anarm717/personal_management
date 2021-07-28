package models

import (
	"gorm.io/gorm"
	"personal_management/api/util"
	"time"
)

type Employee struct {
	ID int64
	Surname string
	Name string
	Patronymic string
	Birthdate time.Time
}
// swagger:response employeeResponse
type employeeResponse struct {
	// in: body
	Body []Employee
}

func (p *Employee) GetAllEmployees(db *gorm.DB,rowCount,pageNumber int) (*[]Employee, error) {
	var err error
	employees := []Employee{}
	err = db.Raw(util.Sql_getEmploees,rowCount,pageNumber).Find(&employees).Offset(0).Limit(5).Error
	if err != nil {
		return &[]Employee{}, err
	}
	return &employees, nil
}
