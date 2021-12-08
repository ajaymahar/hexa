package adapters

import (
	"github.com/ajaymahar/hexa/domain"
	"github.com/ajaymahar/hexa/domain/port"
)

type DefaultCustomerService struct {
	repo port.CustomerRepository
}

// NewDefaultCustomerService is factory function will inject repository dependency
func NewDefaultCustomerService(repo port.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repo}
}

func (cs DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return cs.repo.FindAll()
}

func (cs DefaultCustomerService) CreateCustomer(customer domain.Customer) error {
	return cs.repo.Save(customer)
}

func (cs DefaultCustomerService) GetCustomer(id string) (domain.Customer, error) {
	return cs.repo.Find(id)
}
