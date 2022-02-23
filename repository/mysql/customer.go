package repository

import (
	"context"

	"github.com/venkat/customer/domain"
	"gorm.io/gorm"
)

type mysqlCustomerRepository struct {
	Conn *gorm.DB
}

func NewMysqlCustomerReposity(conn *gorm.DB) domain.CustomerRepository {
	return &mysqlCustomerRepository{conn}
}

func (m *mysqlCustomerRepository) Fetch(ctx context.Context) (res []domain.Customer, err error) {

	query := m.Conn.Find(&res)
	err = query.Error
	return
}
func (m *mysqlCustomerRepository) GetByID(ctx context.Context, id int64) (res domain.Customer, err error) {
	query := m.Conn.Find(&res, id)
	err = query.Error

	return
}

func (m *mysqlCustomerRepository) Store(ctx context.Context, customer *domain.Customer) (err error) {
	err = m.Conn.Create(&customer).Error
	return
}

func (m *mysqlCustomerRepository) Delete(ctx context.Context, id int64) (err error) {
	err = m.Conn.Delete(domain.Customer{}, id).Error
	return
}
func (m *mysqlCustomerRepository) Update(ctx context.Context, id int64, updatedFields map[string]interface{}) (err error) {
	err = m.Conn.Model(domain.Customer{}).Where("id=?", id).Updates(updatedFields).Error
	return
}
