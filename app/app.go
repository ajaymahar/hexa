package app

import (
	"log"
	"net/http"

	"github.com/ajaymahar/hexa/domain/adapters"
	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()

	//wiring all the aspects
	repo := adapters.NewCustomerRepositoryStub()
	svc := adapters.NewDefaultCustomerService(repo)
	ch := NewCustomerHandler(svc)

	sr := router.PathPrefix("/api").Subrouter()
	getSub := sr.Methods("GET").Subrouter()
	getSub.HandleFunc("/customers", ch.getAllCustomers)
	getSub.HandleFunc("/customer/{id}", ch.getCustomer)

	postSub := sr.Methods("POST").Subrouter()
	postSub.HandleFunc("/customer", ch.createCustomer)

	log.Println("Starting server on port : 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
