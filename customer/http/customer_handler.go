package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/venkat/customer/model"
	"github.com/venkat/customer/service"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// CustomerHandler  represent the httphandler for customer
type CustomerHandler struct {
	CUsecase service.CustomerUsecase
}

func NewCustomerHandler(r *mux.Router, us service.CustomerUsecase) {
	handler := &CustomerHandler{
		CUsecase: us,
	}

	r.HandleFunc("/customers", handler.Fetch).Methods("GET")
	r.HandleFunc("/customers", handler.Store).Methods("POST")
	r.HandleFunc("/customers/{id}", handler.Update).Methods("PUT")
	r.HandleFunc("/customers/{id}", handler.GetByID).Methods("GET")
	r.HandleFunc("/customers/{id}", handler.Delete).Methods("DELETE")
}
func (a *CustomerHandler) Fetch(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	res.Header().Set("Content-Type", "application/json")
	result, err := a.CUsecase.Fetch(ctx)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err.Error())
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(result)
}

func (a *CustomerHandler) GetByID(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	idP, err := strconv.Atoi(vars["id"])
	res.Header().Set("Content-Type", "application/json")
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err.Error())
	}

	id := int64(idP)
	ctx := req.Context()

	cust, err := a.CUsecase.GetByID(ctx, id)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err.Error())
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(cust)
}

// Store will store the customer by given request body
func (a *CustomerHandler) Store(res http.ResponseWriter, req *http.Request) {
	var customer model.Customer
	err := json.NewDecoder(req.Body).Decode(&customer)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err.Error())
	}

	ctx := req.Context()
	err = a.CUsecase.Store(ctx, &customer)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err.Error())
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(customer)
}

// Store will update the customer by given request body
func (a *CustomerHandler) Update(res http.ResponseWriter, req *http.Request) {
	var updatedFields map[string]interface{}
	err := json.NewDecoder(req.Body).Decode(&updatedFields)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err.Error())
	}
	vars := mux.Vars(req)
	idP, err := strconv.Atoi(vars["id"])
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err.Error())
	}
	ctx := req.Context()
	err = a.CUsecase.Update(ctx, int64(idP), updatedFields)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err.Error())
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode("Data updated successfully")
}

func (a *CustomerHandler) Delete(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	idP, err := strconv.Atoi(vars["id"])
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err.Error())
	}

	id := int64(idP)
	ctx := req.Context()

	err = a.CUsecase.Delete(ctx, id)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err.Error())
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode("Data deleted successfully")
}
