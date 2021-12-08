package port

import "github.com/ajaymahar/hexa/domain"

// primary port
type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	CreateCustomer(domain.Customer) error
	GetCustomer(string) (domain.Customer, error)
}
