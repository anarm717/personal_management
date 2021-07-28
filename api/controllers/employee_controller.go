package controllers

import (
	"fmt"
	"net/http"
	"personal_management/api/models"
	"personal_management/api/responses"
	"strconv"
	"github.com/gorilla/mux"
)

//type PageSettings struct {
//	PageNumber int
//	RowCount int
//}
//swagger:parameters ListEmployees
type EmployeeRequest struct {
	// in: path
	// required: true
	PageNumber int
	// in: path
	// required: true
	RowCount int
}
// swagger:route GET /employees/{PageNumber}/{RowCount} employees ListEmployees
// responses:
// 200:employeeResponse
func (server *Server) GetEmployees(w http.ResponseWriter, r *http.Request) {
	employee := models.Employee{}
	vars := mux.Vars(r)
	pageNumber, err := strconv.ParseUint(vars["PageNumber"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	fmt.Println(pageNumber)
	rowCount, err := strconv.ParseUint(vars["RowCount"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	fmt.Println(rowCount)
	employees, err := employee.GetAllEmployees(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, employees)
}
