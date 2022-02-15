package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/venkat/customer/routes"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/customer/create", routes.CreateCustomer).Methods("POST")
	r.HandleFunc("/customer/{id}", routes.GetCustomer).Methods("GET")
	r.HandleFunc("/customer", routes.GetAllCustomers).Methods("GET")
	r.HandleFunc("/customer/{id}", routes.UpdateCustomer).Methods("PATCH")
	r.HandleFunc("/customer/delete/{id}", routes.DeleteCustomer).Methods("DELETE")
	http.ListenAndServe(":8080", r)
}
