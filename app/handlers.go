package app

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"

	"github.com/ajaymahar/hexa/domain"
	"github.com/ajaymahar/hexa/domain/port"
	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	svc port.CustomerService
}

// NewCustomerHandler is factory method injected with CustomerService
func NewCustomerHandler(svc port.CustomerService) CustomerHandler {
	return CustomerHandler{svc: svc}
}

func (ch CustomerHandler) getAllCustomers(rw http.ResponseWriter, r *http.Request) {

	customers, err := ch.svc.GetAllCustomers()
	if err != nil {

		log.Println(err)
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		rw.Header().Set("Content-Type", "application/xml")
		if err := xml.NewEncoder(rw).Encode(customers); err != nil {

			log.Println(err)
			http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	} else {
		rw.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(rw).Encode(customers); err != nil {

			log.Println(err)
			http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
}

func (ch CustomerHandler) createCustomer(rw http.ResponseWriter, r *http.Request) {
	var customer domain.Customer

	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {

		log.Println(err)
		http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	if err := ch.svc.CreateCustomer(customer); err != nil {

		log.Println(err)
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(rw).Encode(customer); err != nil {

		log.Println(err)
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (ch CustomerHandler) getCustomer(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	customer, err := ch.svc.GetCustomer(id)
	if err != nil {
		log.Println(err)
		if err.Error() == "not found" {

			http.Error(rw, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(rw).Encode(customer); err != nil {
		log.Println(err)
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
