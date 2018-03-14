package models

import (
	"customer-consumer/helpers/util"
	"fmt"
	"log"
	"strings"

	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
)

func CreateCustomer(kind string, customer Customer, ctx context.Context) (*datastore.Key, error) {
	name := customer.PartnerId
	taskKey := datastore.NameKey(kind, name, nil)
	client, cerr := DataStoreClient()
	if cerr != nil {
		return nil, cerr
	}

	key, err := client.Put(ctx, taskKey, &customer)
	if err != nil {
		log.Printf(err.Error())
		log.Printf("Failed to insert the details of the customer into data-storage, due to: %v", err)
		return nil, err
	}
	log.Printf(key.Name)
	return key, err
}

func GetCustomer(kind string, name string, ctx context.Context) (*Customer, error) {
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

func EditCustomer(client *datastore.Client, kind string, name string, dst interface{}, ctx context.Context) (*Customer, error) {
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

func DeleteCustomer(client *datastore.Client, kind string, name string, dst interface{}, ctx context.Context) (*Customer, error) {
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

func GetAllCustomers(kind string, ctx context.Context, paramMap map[string][]string) ([]Customer, error) {
	client, err := DataStoreClient()
	var customers []Customer
	partner_id := paramMap["partner_id"]
	var partnerId string
	q := datastore.NewQuery(kind)

	if partner_id != nil {
		partnerId = partner_id[0]
		q = q.Filter("PartnerID=", partnerId)
	}

	_, err = client.GetAll(ctx, q, &customers)
	if err != nil {
		fmt.Println(err)
		log.Printf(err.Error())
		return nil, err
	}
	return customers, err
}

func GetCustomerUsingMultiFilters(kind string, ctx context.Context, paramMap map[string][]string) ([]Customer, error) {
	client, err := DataStoreClient()
	var customers []Customer
	q := datastore.NewQuery(kind)

	partner_id := paramMap["partner_id"]
	email := paramMap["email"]

	if partner_id != nil {
		q = q.Filter("PartnerId=", partner_id)
		log.Printf("search filter - PartnerId : %s", partner_id)
	}
	if email != nil {
		q = q.Filter("Email=", email)
		log.Printf("search filter - Email : %s", email)
	}

	_, err = client.GetAll(ctx, q, &customers)
	if err != nil {
		fmt.Println(err)
		log.Printf(err.Error())
		return nil, err
	}
	return customers, err
}

func DataStoreClient() (*datastore.Client, error) {
	ctx := context.Background()
	projectID := util.MustGetenv("GOOGLE_CLOUD_PROJECT")
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Printf(err.Error())
	}
	return client, err
}
