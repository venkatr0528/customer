package model

import "github.com/venkat/customer/config"

type Customer struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	MobileNumber string `json:"mobile_number"`
	EmailId      string `json:"email_id"`
	Address      string `json:"address"`
}
type Response struct {
	Status int         `json:"status"`
	Error  bool        `json:"error"`
	Data   interface{} `json:"data"`
}

func CreateCustomer(customer Customer) (bool, error) {
	db, err := config.GetDB()
	if err != nil {
		return true, err
	}
	res := db.Create(&customer)
	if res.Error != nil {
		return true, err
	}
	return false, nil
}

func UpdateCustomer(id int, customer Customer) (bool, error) {
	db, err := config.GetDB()
	if err != nil {
		return true, err
	}
	res := db.Create(&customer)
	if res.Error != nil {
		return true, err
	}
	return false, nil

}

func DeleteCustomer(id int) (bool, error) {
	var customer Customer
	db, err := config.GetDB()
	if err != nil {
		return true, err
	}
	res := db.Delete(&customer, id)
	if res.Error != nil {
		return true, err
	}
	return false, nil

}

func GetCustomer(id int) (Customer, error) {
	var customer Customer
	db, err := config.GetDB()
	if err != nil {
		return customer, err
	}
	res := db.First(&customer, id)
	if res.Error != nil {
		return customer, err
	}
	return customer, nil
}

func GetAllCustomer() ([]Customer, error) {
	var customers []Customer
	db, err := config.GetDB()
	if err != nil {
		return customers, err
	}
	res := db.Find(&customers)
	if res.Error != nil {
		return customers, err
	}
	return customers, nil
}
