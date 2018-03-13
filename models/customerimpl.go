package models

import (
	"customer-consumer/helpers/util"
	"fmt"
	"log"
	"strings"

	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
)

const (
	invalidData = `error: invalid User data`
)

func InsertCustomer(kind string, customer Customer, ctx context.Context) (*datastore.Key, error) {
	name := customer.PartnerId
	taskKey := datastore.NameKey(kind, name, nil)
	client, cerr := DataStoreClient()
	if cerr != nil {
		return nil, cerr
	}

	key, err := client.Put(ctx, taskKey, &customer)
	if err != nil {
		log.Printf(err.Error())
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

func GetCustomerByEmailAndId(kind string, ctx context.Context, paramMap map[string][]string) ([]Customer, error) {
	client, err := DataStoreClient()
	var customers []Customer
	partner_id := paramMap["partner_id"]
	email_id := paramMap["email"]
	var partnerId string
	var email string
	q := datastore.NewQuery(kind)

	if partner_id != nil {
		partnerId = partner_id[0]
		q = q.Filter("PartnerID=", partnerId)
		log.Printf("search filter - PartnerID : %s", partnerId)
	}
	if email_id != nil {
		email = email_id[0]
		q = q.Filter("Email=", email)
		log.Printf("search filter - Email : %s", email)
	}

	//q = q.Filter("Email=", "vikas.sidusa@gmail.com")
	//q = q.Filter("PartnerID=", "001")

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
