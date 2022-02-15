package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/venkat/customer/model"
)

func CreateCustomer(res http.ResponseWriter, req *http.Request) {
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
	reqStatus, err := model.CreateCustomer(customer)
	response.Error = reqStatus
	if err != nil {
		response.Data = err.Error()
	} else {
		response.Data = "Data created successfully"

	}
	json.NewEncoder(res).Encode(response)

}

func GetAllCustomers(res http.ResponseWriter, req *http.Request) {
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
func GetCustomer(res http.ResponseWriter, req *http.Request) {
	var response model.Response
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	customer, err := model.GetCustomer(id)
	if err != nil {
		response.Error = true
		response.Data = err.Error()
	} else {
		response.Data = customer
		response.Error = false
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(response)

}
func UpdateCustomer(res http.ResponseWriter, req *http.Request) {

}
func DeleteCustomer(res http.ResponseWriter, req *http.Request) {
	var response model.Response
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	reqStatus, err := model.DeleteCustomer(id)
	response.Status = 200
	response.Error = reqStatus
	if err != nil {

		response.Data = err.Error()
	} else {
		response.Data = "data deleted successfully"

	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(response)

}
