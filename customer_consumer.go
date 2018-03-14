package main

import (
	"customer-consumer/api"
	"customer-consumer/helpers/pkg/publisher"
	"customer-consumer/helpers/pkg/subscriber"
	"customer-consumer/helpers/util"
	"gorilla/mux"
	"net/http"
	//"os"
)

func main() {
	go doHandleRoutes()

	//os.Setenv("GOOGLE_CLOUD_PROJECT", "chefd-dev-xxx")
	//os.Setenv("PUBSUB_EMULATOR_HOST", "localhost:8085")
	//DATASTORE_EMULATOR_HOST=localhost:8081
	//os.Setenv("DATASTORE_EMULATOR_HOST", "localhost:8081")
	//os.Setenv("GOOGLE_CLOUD_PROJECT", "dark-window-194510")
	//os.Setenv("NETSUITE_PRODUCT_BETA", "netsuite-product-beta")
	//os.Setenv("PULL_FOR_NETSUITE_PRODUCT_BETA", "netsuite-product-beta-subscriber")
	//os.Setenv("PRODUCT_PUBLISH_URL", "http://localhost:8080/product-date-publish")
	//os.Setenv("NETSUITE_CONSUMER_KEY", "febf4585d5dab30eac9e6a40643f4b187171432a79901bb98b0ad152d32018bc")
	//os.Setenv("NETSUITE_CONSUMER_SECRET", "8d9eca2dcbf63a9ac1e48c037274307eae791a6d0d9c7c0eb45aca3563d3e1ec")
	//os.Setenv("NETSUITE_REALM", "3972263_SB1")
	//os.Setenv("NETSUITE_TOKEN_KEY", "78e7cbd94e9812e38c2084762dae30873a95a039a63b97d1088e13cea95aff2f")
	//os.Setenv("NETSUITE_TOKEN_SECRET", "95ac952bd9ebda4f4a309cd1ac6b8cc7b3eeb37f5e701c92ec7a4752ecb6bcc6")
	//os.Setenv("NETSUITE_RESTLET_URL", "https://rest.netsuite.com/app/site/hosting/restlet.nl?script=%d&deploy=1")
	//os.Setenv("NETSUITE_MESSAGE_BETA", "netsuite-message-beta")
	//os.Setenv("DATASTORE_EMULATOR_HOST", util.MustGetenv("DATASTORE_EMULATOR_HOST"))
	//os.Setenv("DATA_STORE_KIND", "PRODUCT_KIND")
	//os.Setenv("DATA_STORE_NAME", "PRODUCT")
	//os.Setenv("RAW_PRODUCT_DATA", "raw_product_data")
	//os.Setenv("GET_RAW_PRODUCT_DATA", "get_raw_product_data")
	//os.Setenv("PROCESSED_PRODUCT_DATA", "processed_product_data")
	//os.Setenv("GET_PROCESSED_PRODUCT_DATA", "get_processed_product_data")

	//env variables for customer_consumer:
	//os.Setenv("RAW_CUSTOMER_DATA", "raw_customer_data")
	//os.Setenv("GET_RAW_CUSTOMER_DATA", "get_raw_customer_data")
	//os.Setenv("PROCESSED_CUSTOMER_DATA", "processed_customer_data")
	//os.Setenv("GET_PROCESSED_CUSTOMER_DATA", "get_processed_customer_data")
	//os.Setenv("DATA_STORE_KIND_CUSTOMER", "CUSTOMER_KIND")
	//os.Setenv("DATA_STORE_NAME_CUSTOMER", "CUSTOMER")

	subscriber.SubscribeCustomer()
}

func doHandleRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/", api.HealthCheckFunc)
	router.HandleFunc("/customer-data-publish", publisher.PublishCustomer)
	router.HandleFunc("/customer/{partner_id}", api.GetCustomer).Methods("GET")
	router.HandleFunc("/edit-customer/{partner_id}", api.EditCustomer).Methods("POST")
	router.HandleFunc("/delete-customer/{partner_id}", api.DeleteCustomer).Methods("GET")
	router.HandleFunc("/customers", api.GetAllCustomers).Methods("GET")
	router.HandleFunc("/customer/multifilters/{partner_id}/{email}", api.GetCustomerUsingMultiFilters).Methods("GET")
	http.ListenAndServe(util.Port(), router)
}
