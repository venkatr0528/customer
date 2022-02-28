package service

import (
	"context"

	"github.com/venkat/customer/model"

	repository "github.com/venkat/customer/repository"
)

type customerUsecase struct {
	customerRepo repository.CustomerRepository
}

// NewCustomerUsecase will create new an NewCustomerUsecase object representation of domain.NewCustomerUsecase interface
func NewCustomerUsecase(cr repository.CustomerRepository) CustomerUsecase {
	return &customerUsecase{
		customerRepo: cr,
	}
}

func (a *customerUsecase) Fetch(ctx context.Context) (res []model.Customer, err error) {
	res, err = a.customerRepo.Fetch(ctx)
	return
}

func (a *customerUsecase) GetByID(ctx context.Context, id int64) (res model.Customer, err error) {

	res, err = a.customerRepo.GetByID(ctx, id)
	return
}

func (a *customerUsecase) Update(ctx context.Context, id int64, updatedFileds map[string]interface{}) (err error) {

	return a.customerRepo.Update(ctx, id, updatedFileds)
}

func (a *customerUsecase) Store(ctx context.Context, customer *model.Customer) (err error) {
	err = a.customerRepo.Store(ctx, customer)
	return
}

func (a *customerUsecase) Delete(ctx context.Context, id int64) (err error) {

	existedCustomer, err := a.customerRepo.GetByID(ctx, id)
	if err != nil {
		return
	}
	if existedCustomer == (model.Customer{}) {
		return
	}
	return a.customerRepo.Delete(ctx, id)
}
