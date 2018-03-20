package api

import (
	"customer-consumer/helpers/util"
	"customer-consumer/models"
	"encoding/json"
	"gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"

	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
)

//GetCustomer gives a single customer
func GetCustomer(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	params := mux.Vars(r)
	partnerID := params["customer_id"]
	kind := util.MustGetenv("DATA_STORE_KIND_CUSTOMER")
	customer, err := models.GetCustomer(ctx, kind, partnerID)
	w.Header().Set("Content-Type", "application/json")
	if customer == nil {
		if err != nil {
			log.Printf(err.Error())
		}
		var errormsg = models.CustomerError{}
		errormsg.Error = "No Customer found"
		json.NewEncoder(w).Encode(errormsg)
	} else {
		json.NewEncoder(w).Encode(customer)
	}
}

//EditCustomer edits and updates a particular customer
func EditCustomer(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	var customer models.Customer
	err = json.Unmarshal(body, &customer)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var errormsg = models.CustomerError{}
		errormsg.Error = err.Error()
		json.NewEncoder(w).Encode(errormsg)
		log.Printf(err.Error())
		return
	}

	ctx := context.Background()
	params := mux.Vars(r)
	partnerID := params["customer_id"]
	projectID := util.MustGetenv("GOOGLE_CLOUD_PROJECT")
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var errormsg = models.CustomerError{}
		errormsg.Error = err.Error()
		json.NewEncoder(w).Encode(errormsg)
		log.Printf(err.Error())
		return
	}
	kind := util.MustGetenv("DATA_STORE_KIND_CUSTOMER")
	models.EditCustomer(ctx, client, kind, partnerID, &customer)
	json.NewEncoder(w).Encode(customer)
}

//DeleteCustomer deletes a particular customer
func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	params := mux.Vars(r)
	partnerID := params["customer_id"]
	projectID := util.MustGetenv("GOOGLE_CLOUD_PROJECT")
	client, err := datastore.NewClient(ctx, projectID)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var errormsg = models.CustomerError{}
		errormsg.Error = err.Error()
		json.NewEncoder(w).Encode(errormsg)
		log.Printf(err.Error())
		return
	}
	var customer models.Customer
	kind := util.MustGetenv("DATA_STORE_KIND_CUSTOMER")
	models.DeleteCustomer(ctx, client, kind, partnerID, &customer)
	json.NewEncoder(w).Encode(customer)
}

//GetAllCustomers gives all cutomers
func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var customers []models.Customer
	kind := util.MustGetenv("DATA_STORE_KIND_CUSTOMER")
	params := mux.Vars(r)
	var customerfilter models.CustomerFilter
	customerfilter.PartnerId = params["partner_id"]
	customers, err := models.GetAllCustomers(ctx, kind, customerfilter)
	w.Header().Set("Content-Type", "application/json")
	if customers == nil {
		if err != nil {
			log.Printf(err.Error())
		}
		w.WriteHeader(http.StatusBadRequest)
		var errormsg = models.CustomerError{}
		errormsg.Error = "No Customer found"
		json.NewEncoder(w).Encode(errormsg)
	} else {
		json.NewEncoder(w).Encode(customers)
	}
}

//GetCustomerUsingMultiFilters gives customer data with specific search filters
func GetCustomerUsingMultiFilters(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	params := mux.Vars(r)
	var customers []models.Customer
	kind := util.MustGetenv("DATA_STORE_KIND_CUSTOMER")

	var customerfilter models.CustomerFilter
	customerfilter.PartnerId = params["partner_id"]
	customerfilter.Email = params["email"]

	customers, err := models.GetCustomerUsingMultiFilters(ctx, kind, customerfilter)
	w.Header().Set("Content-Type", "application/json")
	if customers == nil {
		if err != nil {
			log.Printf(err.Error())
		}
		var errormsg = models.CustomerError{}
		errormsg.Error = "No Customer found"
		json.NewEncoder(w).Encode(errormsg)
	} else {
		json.NewEncoder(w).Encode(customers)
	}
}
