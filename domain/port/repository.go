package port

import "github.com/ajaymahar/hexa/domain"

// secondary port
type CustomerRepository interface {
	FindAll() ([]domain.Customer, error)
	Save(domain.Customer) error
	Find(string) (domain.Customer, error)
}
