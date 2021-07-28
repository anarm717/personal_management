package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"personal_management/api/auth"
	"personal_management/api/models"
	"personal_management/api/responses"
	"personal_management/api/utils/formaterror"
	"strconv"
)
// swagger:route GET /categories categories ListCategories
// responses:
// 200:categoryResponse
func (server *Server) GetCategories(w http.ResponseWriter, r *http.Request) {
	category := models.Category{}
	categories, err := category.GetAllCategory(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, categories)
}

func (server *Server) GetCategory(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	category := models.Category{}

	categoryReceived, err := category.FindCategoryByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, categoryReceived)
}
// swagger:route POST /categories categories Category
// security:
//   api_key:
// responses:
// 200:categoryResponse
func (server *Server) CreateCategory(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	category := models.Category{}
	err = json.Unmarshal(body, &category)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	category.Prepare()
	err = category.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	_, err = auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	categoryCreated, err := category.SaveCategory(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, categoryCreated.ID))
	responses.JSON(w, http.StatusCreated, categoryCreated)
}
