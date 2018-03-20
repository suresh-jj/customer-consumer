package main

import (
	"customer-consumer/api"
	"customer-consumer/helpers/pkg/publisher"
	"customer-consumer/helpers/pkg/subscriber"
	"customer-consumer/helpers/util"
	"gorilla/mux"
	"net/http"
)

func main() {
	go doHandleRoutes()

	subscriber.SubscribeCustomer()
}

func doHandleRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/", api.HealthCheckFunc)
	router.HandleFunc("/customer-data-publish", publisher.PublishCustomer)
	router.HandleFunc("/add-customer", publisher.AddCustomer).Methods("POST")
	router.HandleFunc("/customer/{partner_id}", api.GetCustomer).Methods("GET")
	router.HandleFunc("/edit-customer/{partner_id}", api.EditCustomer).Methods("POST")
	router.HandleFunc("/delete-customer/{partner_id}", api.DeleteCustomer).Methods("GET")
	router.HandleFunc("/customers", api.GetAllCustomers).Methods("GET")
	router.HandleFunc("/customer/multifilters/{partner_id}/{email}", api.GetCustomerUsingMultiFilters).Methods("GET")
	http.ListenAndServe(util.Port(), router)
}
