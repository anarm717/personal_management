package controllers

import (
	"net/http"
	"personal_management/api/middlewares"
)

func (s *Server) initializeRoutes() {

	sh := http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("./swaggerui/")))
	s.Router.PathPrefix("/swaggerui/").Handler(sh)

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Category routes
	s.Router.HandleFunc("/categories", middlewares.SetMiddlewareJSON(s.GetCategories)).Methods("GET")
	s.Router.HandleFunc("/categories/{id}", middlewares.SetMiddlewareJSON(s.GetCategory)).Methods("GET")
	s.Router.HandleFunc("/categories", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.CreateCategory))).Methods("POST")

	//employees
	s.Router.HandleFunc("/employees/{PageNumber}/{RowCount}", middlewares.SetMiddlewareJSON(s.GetEmployees)).Methods("GET")
}
