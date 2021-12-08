package adapters

import (
	"errors"
	"log"

	"github.com/ajaymahar/hexa/domain"
)

type CustomerRepositoryStub struct {
	customers []domain.Customer
}

// NewCustomerRepositoryStub is factory function will return list of customers
func NewCustomerRepositoryStub() *CustomerRepositoryStub {
	custs := []domain.Customer{}
	return &CustomerRepositoryStub{customers: custs}
}

func (crs *CustomerRepositoryStub) FindAll() ([]domain.Customer, error) {
	return crs.customers, nil
}

func (crs *CustomerRepositoryStub) Save(customer domain.Customer) error {
	crs.customers = append(crs.customers, customer)
	log.Printf("Cust Repo Save %+v", crs.customers)
	return nil
}

func (crs *CustomerRepositoryStub) Find(id string) (domain.Customer, error) {

	for _, c := range crs.customers {
		if c.ID == id {
			return c, nil
		}
	}
	return domain.Customer{}, errors.New("not found")
}
