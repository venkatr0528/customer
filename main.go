package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/venkat/customer/config"
	customerHttp "github.com/venkat/customer/customer/http"
	repository "github.com/venkat/customer/repository/mysql"
	"github.com/venkat/customer/usecase"
)

func main() {
	dbConn, err := config.GetDB()

	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	customerRepo := repository.NewMysqlCustomerReposity(dbConn)

	cu := usecase.NewCustomerUsecase(customerRepo)
	customerHttp.NewCustomerHandler(e, cu)

	log.Fatal(e.Start(":8080"))
}
