package model

import (
	"github.com/venkat/customer/config"
	"gorm.io/gorm"
)

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

func (c *Customer) CreateCustomer(db *gorm.DB) error {
	res := db.Create(&c)
	return res.Error
}

func (c *Customer) UpdateCustomer(db *gorm.DB, updatedFileds map[string]interface{}) error {

	res := db.Model(&c).Updates(updatedFileds)
	return res.Error

}

func (c *Customer) DeleteCustomer(db *gorm.DB) error {

	res := db.Delete(&c)
	return res.Error

}

func (c *Customer) GetCustomer(db *gorm.DB) error {
	res := db.Find(&c)
	return res.Error
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
