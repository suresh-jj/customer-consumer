package main

import (
	"customer-consumer/api"
	"customer-consumer/helpers/util"
	"customer-consumer/services"
	"gorilla/mux"
	"net/http"
)

func main() {
	go doHandleRoutes()

	services.SubscribeCustomer()
}

func doHandleRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/", api.HealthCheckFunc)
	router.HandleFunc("/customer-data-publish", services.PublishCustomer)
	router.HandleFunc("/customers", services.AddCustomer).Methods("POST")
	router.HandleFunc("/customers/{customer_id}", api.GetCustomer).Methods("GET")
	router.HandleFunc("/customers/{customer_id}", api.EditCustomer).Methods("PUT")
	router.HandleFunc("/customers/{customer_id}", api.DeleteCustomer).Methods("DELETE")
	router.HandleFunc("/customers", api.GetAllCustomers).Methods("GET")
	router.HandleFunc("/customers/multifilters/{partner_id}/{email}", api.GetCustomerUsingMultiFilters).Methods("GET")
	http.ListenAndServe(util.Port(), router)
}
