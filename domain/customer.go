package domain

import "context"

type Customer struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	MobileNumber string `json:"mobile_number"`
	EmailId      string `json:"email_id"`
	Address      string `json:"address"`
}

// CustomerUsecase represent the customer's usecases
type CustomerUsecase interface {
	Fetch(context.Context) ([]Customer, error)
	GetByID(context.Context, int64) (Customer, error)
	Update(context.Context, int64, map[string]interface{}) error
	Store(context.Context, *Customer) error
	Delete(context.Context, int64) error
}

// CustomerRepository represent the customer's repository contract
type CustomerRepository interface {
	Fetch(context.Context) ([]Customer, error)
	GetByID(context.Context, int64) (Customer, error)
	Update(context.Context, int64, map[string]interface{}) error
	Store(context.Context, *Customer) error
	Delete(context.Context, int64) error
}
