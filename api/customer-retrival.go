package api

import (
	"customer-consumer/helpers/util"
	"customer-consumer/models"
	"encoding/json"
	"fmt"
	"gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"

	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
)

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	params := mux.Vars(r)
	partner_id := params["partner_id"]
	kind := util.MustGetenv("DATA_STORE_KIND_CUSTOMER")
	customer, err := models.GetCustomer(kind, partner_id, ctx)
	w.Header().Set("Content-Type", "application/json")
	if customer == nil {
		if err != nil {
			log.Printf(err.Error())
		}
		var errorObj = "{'Message':'No Customer found'}"
		json.NewEncoder(w).Encode(errorObj)
	} else {
		json.NewEncoder(w).Encode(customer)
	}
}

func EditCustomer(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	var customer models.Customer
	err = json.Unmarshal(body, &customer)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error found: %v", err)
		return
	}

	ctx := context.Background()
	params := mux.Vars(r)
	customer_id := params["partner_id"]
	projectID := util.MustGetenv("GOOGLE_CLOUD_PROJECT")
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Printf("Failed to create client: %v", err)
	}
	kind := util.MustGetenv("DATA_STORE_KIND_CUSTOMER")
	models.EditCustomer(client, kind, customer_id, &customer, ctx)
	json.NewEncoder(w).Encode(customer)
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	params := mux.Vars(r)
	customer_id := params["partner_id"]
	projectID := util.MustGetenv("GOOGLE_CLOUD_PROJECT")
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Printf("Failed to create client: %v", err)
	}
	var customer models.Customer
	kind := util.MustGetenv("DATA_STORE_KIND_CUSTOMER")
	models.DeleteCustomer(client, kind, customer_id, &customer, ctx)
	json.NewEncoder(w).Encode(customer)
}

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var customers []models.Customer
	kind := util.MustGetenv("DATA_STORE_KIND")
	customers, err := models.GetAllCustomers(kind, ctx, r.URL.Query())
	w.Header().Set("Content-Type", "application/json")
	if customers == nil {
		var errorObj = "{'Message':'No Customer found'}"
		if err != nil {
			log.Printf(err.Error())
		}
		json.NewEncoder(w).Encode(errorObj)
	} else {
		json.NewEncoder(w).Encode(customers)
	}
}

func GetCustomerByEmailAndId(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var customers []models.Customer
	kind := util.MustGetenv("DATA_STORE_KIND_CUSTOMER")
	customers, err := models.GetCustomerByEmailAndId(kind, ctx, r.URL.Query())
	w.Header().Set("Content-Type", "application/json")
	if customers == nil {
		var errorObj = "{'Message':'No Customer found'}"
		if err != nil {
			log.Printf(err.Error())
		}
		json.NewEncoder(w).Encode(errorObj)
	} else {
		json.NewEncoder(w).Encode(customers)
	}
}
