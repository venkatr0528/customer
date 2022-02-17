package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/venkat/customer/config"
	"github.com/venkat/customer/model"
	"gorm.io/gorm"
)

type App struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (a *App) Init() {
	db, err := config.GetDB()
	if err != nil {
		log.Fatal("Issue with database connection:", err.Error())
	}
	a.DB = db
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}
func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/customer/create", a.CreateCustomer).Methods("POST")
	a.Router.HandleFunc("/customer/{id}", a.GetCustomer).Methods("GET")
	a.Router.HandleFunc("/customer", a.GetAllCustomers).Methods("GET")
	a.Router.HandleFunc("/customer/{id}", a.UpdateCustomer).Methods("PATCH")
	a.Router.HandleFunc("/customer/delete/{id}", a.DeleteCustomer).Methods("DELETE")
}
func (a *App) Run(address string) {
	log.Fatal(http.ListenAndServe(address, a.Router))
}
func (a *App) CreateCustomer(res http.ResponseWriter, req *http.Request) {
	var customer model.Customer
	var response model.Response
	err := json.NewDecoder(req.Body).Decode(&customer)
	res.Header().Set("Content-Type", "application/json")
	response.Status = 200
	if err != nil {
		response.Error = true
		response.Data = err.Error()
		json.NewEncoder(res).Encode(response)
	}
	err = customer.CreateCustomer(a.DB)
	if err != nil {
		response.Error = true
		response.Data = err.Error()
	} else {
		response.Data = "Data created successfully"

	}
	json.NewEncoder(res).Encode(response)

}

func (a *App) GetAllCustomers(res http.ResponseWriter, req *http.Request) {
	var response model.Response
	customers, err := model.GetAllCustomer()
	res.Header().Set("Content-Type", "application/json")
	response.Status = 200
	if err != nil {
		response.Error = true
		response.Data = err.Error()
	} else {
		response.Data = customers
		response.Error = false
	}
	json.NewEncoder(res).Encode(response)
}
func (a *App) GetCustomer(res http.ResponseWriter, req *http.Request) {
	var response model.Response
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	customer := model.Customer{ID: id}

	if err := customer.GetCustomer(a.DB); err != nil {
		response.Error = true
		response.Data = err.Error()
	} else {
		response.Data = customer
		response.Error = false
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(response)

}
func (a *App) UpdateCustomer(res http.ResponseWriter, req *http.Request) {
	var response model.Response
	var updatedFields map[string]interface{}
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	err := json.NewDecoder(req.Body).Decode(&updatedFields)
	customer := model.Customer{ID: id}
	err = customer.UpdateCustomer(a.DB, updatedFields)
	response.Status = 200
	response.Error = false
	if err != nil {
		response.Error = true
		response.Data = err.Error()
	} else {
		response.Data = "data updated successfully"

	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(response)

}
func (a *App) DeleteCustomer(res http.ResponseWriter, req *http.Request) {
	var response model.Response
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	customer := model.Customer{ID: id}
	err := customer.DeleteCustomer(a.DB)
	response.Status = 200
	response.Error = false
	if err != nil {
		response.Error = true
		response.Data = err.Error()
	} else {
		response.Data = "data deleted successfully"

	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(response)

}
