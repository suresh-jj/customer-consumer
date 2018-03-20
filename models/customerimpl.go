package models

import (
	"customer-consumer/helpers/util"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"

	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
)

//CreateCustomer creates new customer
func CreateCustomer(ctx context.Context, kind string, customer Customer) (*datastore.Key, error) {
	name := customer.Id
	taskKey := datastore.NameKey(kind, name, nil)
	client, cerr := DataStoreClient()
	if cerr != nil {
		return nil, cerr
	}

	//regex validations:
	var validEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	var validPhone = regexp.MustCompile(`^(\([0-9]{3}\) |[0-9]{3}-)[0-9]{3}-[0-9]{4}$`)

	if !validEmail.MatchString(customer.Email) {
		return nil, errors.New("Email not valid")
	}
	if !validPhone.MatchString(customer.PhoneNumber) {
		return nil, errors.New("Phone Number must contain 10 digits only")
	}

	//check if customer is already existing
	var isExisting bool
	isExisting, err := isExistingCustomer(ctx, kind, customer)
	log.Printf("isExistingCustomer:: %t", isExisting)
	if isExisting {
		if err != nil {
			log.Printf(err.Error())
		}
		log.Printf("Failed to insert as the customer is already existing in storage!")
		return nil, errors.New("Customer already exists")
	}

	key, err := client.Put(ctx, taskKey, &customer)
	if err != nil {
		log.Printf(err.Error())
		log.Printf("Failed to insert the details of the customer into data-storage, due to: %v", err)
		return nil, err
	}
	return key, err
}

//GetCustomer gives a single customer
func GetCustomer(ctx context.Context, kind string, name string) (*Customer, error) {
	var customer Customer
	taskKey := datastore.NameKey(kind, name, nil)
	client, err := DataStoreClient()
	err = client.Get(ctx, taskKey, &customer)
	if err != nil {
		if strings.HasPrefix(err.Error(), `datastore: no such entity`) {
			err = fmt.Errorf(`Customer '%v' not found`, name)
			return nil, err
		}
	}
	return &customer, err
}

//EditCustomer edits and updates a particular customer
func EditCustomer(ctx context.Context, client *datastore.Client, kind string, name string, dst interface{}) (*Customer, error) {
	var customer Customer
	taskKey := datastore.NameKey(kind, name, nil)
	client, err := DataStoreClient()
	_, err = client.Put(ctx, taskKey, dst)
	if err != nil {
		log.Printf("Failed to update the customer details from data-storage, due to: %v", err)
		if strings.HasPrefix(err.Error(), `datastore: no such entity`) {
			err = fmt.Errorf(`Customer '%v' not found`, name)
			return nil, err
		}
	}
	return &customer, err
}

//DeleteCustomer deletes a particular customer
func DeleteCustomer(ctx context.Context, client *datastore.Client, kind string, name string, dst interface{}) (*Customer, error) {
	var customer Customer
	taskKey := datastore.NameKey(kind, name, nil)
	client, err := DataStoreClient()
	err = client.Delete(ctx, taskKey)
	if err != nil {
		log.Printf("Failed to delete the customer details from data-storage, due to: %v", err)
		if strings.HasPrefix(err.Error(), `datastore: no such entity`) {
			err = fmt.Errorf(`Customer '%v' not found`, name)
			return nil, err
		}
	}
	return &customer, err
}

//GetAllCustomers gives all cutomers
func GetAllCustomers(ctx context.Context, kind string, customerfilter CustomerFilter) ([]Customer, error) {
	client, err := DataStoreClient()
	var customers []Customer
	q := datastore.NewQuery(kind).Order("FirstName")

	if customerfilter.PartnerId != "" {
		q = q.Filter("PartnerId=", customerfilter.PartnerId)
	}

	_, err = client.GetAll(ctx, q, &customers)
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	fmt.Println("dd", customers)
	return customers, err
}

//GetCustomerUsingMultiFilters gives customer data with specific search filters
func GetCustomerUsingMultiFilters(ctx context.Context, kind string, customerfilter CustomerFilter) ([]Customer, error) {
	client, err := DataStoreClient()
	var customers []Customer
	q := datastore.NewQuery(kind)

	if customerfilter.PartnerId != "" {
		q = q.Filter("PartnerId=", customerfilter.PartnerId)
		log.Printf("search filter - PartnerId : %s", customerfilter.PartnerId)
	}
	if customerfilter.Email != "" {
		q = q.Filter("Email=", customerfilter.Email)
		log.Printf("search filter - Email : %s", customerfilter.Email)
	}

	_, err = client.GetAll(ctx, q, &customers)
	if err != nil {
		fmt.Println(err)
		log.Printf(err.Error())
		return nil, err
	}
	return customers, err
}

//isExistingCustomer validates if the Customer is already existing in system
func isExistingCustomer(ctx context.Context, kind string, customer Customer) (bool, error) {
	var isExisting bool
	client, err := DataStoreClient()
	var customers []Customer
	q := datastore.NewQuery(kind)

	if customer.PartnerId != "" {
		q = q.Filter("PartnerId=", customer.PartnerId)
	}

	if customer.Email != "" {
		q = q.Filter("Email=", customer.Email)
	}

	_, err = client.GetAll(ctx, q, &customers)
	if err != nil {
		log.Printf(err.Error())
		return false, err
	}
	if customers == nil {
		isExisting = false
	} else {
		isExisting = true
	}
	return isExisting, err
}

//DataStoreClient returns the datastore client object
func DataStoreClient() (*datastore.Client, error) {
	ctx := context.Background()
	projectID := util.MustGetenv("GOOGLE_CLOUD_PROJECT")
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Printf(err.Error())
	}
	return client, err
}
