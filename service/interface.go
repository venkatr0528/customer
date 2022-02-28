package service

import (
	"context"

	"github.com/venkat/customer/model"
)

// CustomerUsecase represent the customer's usecases
type CustomerUsecase interface {
	Fetch(context.Context) ([]model.Customer, error)
	GetByID(context.Context, int64) (model.Customer, error)
	Update(context.Context, int64, map[string]interface{}) error
	Store(context.Context, *model.Customer) error
	Delete(context.Context, int64) error
}
