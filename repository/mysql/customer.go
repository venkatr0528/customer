package mysql

import (
	"context"

	"github.com/venkat/customer/model"
	"github.com/venkat/customer/repository"
	"gorm.io/gorm"
)

type mysqlCustomerRepository struct {
	Conn *gorm.DB
}

func NewMysqlCustomerReposity(conn *gorm.DB) repository.CustomerRepository {
	return &mysqlCustomerRepository{conn}
}

func (m *mysqlCustomerRepository) Fetch(ctx context.Context) (res []model.Customer, err error) {

	query := m.Conn.Find(&res)
	err = query.Error
	return
}
func (m *mysqlCustomerRepository) GetByID(ctx context.Context, id int64) (res model.Customer, err error) {
	query := m.Conn.Find(&res, id)
	err = query.Error

	return
}

func (m *mysqlCustomerRepository) Store(ctx context.Context, customer *model.Customer) (err error) {
	err = m.Conn.Create(&customer).Error
	return
}

func (m *mysqlCustomerRepository) Delete(ctx context.Context, id int64) (err error) {
	err = m.Conn.Delete(model.Customer{}, id).Error
	return
}
func (m *mysqlCustomerRepository) Update(ctx context.Context, id int64, updatedFields map[string]interface{}) (err error) {
	err = m.Conn.Model(model.Customer{}).Where("id=?", id).Updates(updatedFields).Error
	return
}
