package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/venkat/customer/config"
	customerHttp "github.com/venkat/customer/customer/http"
	repository "github.com/venkat/customer/repository/mysql"
	"github.com/venkat/customer/service"
)

func main() {
	dbConn, err := config.GetDB()

	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	customerRepo := repository.NewMysqlCustomerReposity(dbConn)

	cu := service.NewCustomerUsecase(customerRepo)
	customerHttp.NewCustomerHandler(r, cu)

	log.Fatal(http.ListenAndServe(":8080", r))
}
