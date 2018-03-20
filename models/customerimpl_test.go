package models

import (
	"customer-consumer/helpers/util"
	"log"
	"net/url"
	"testing"

	"golang.org/x/net/context"
)

func TestCreateCustomer(t *testing.T) {
	var customer Customer
	customer.PartnerId = "001"
	customer.Email = "suresh.sidusa@gmail.com"
	customer.FirstName = "Suresh"
	customer.LastName = "Jayapalan"
	customer.Address = "India"
	customer.PhoneNumber = "9990836778"
	ctx := context.Background()
	key, err := CreateCustomer(ctx, "Customers", customer)
	log.Println(err)
	if key.Name != customer.PartnerId {
		t.Errorf("Testing failed", customer.PartnerId)
	}
}

func TestGetCustomer(t *testing.T) {
	kind := util.MustGetenv("DATA_STORE_KIND")
	partner_id := "0001"
	ctx := context.Background()
	customer, err := GetCustomer(ctx, kind, partner_id)
	log.Println(err)
	if customer == nil || customer.PartnerId != "001" {
		t.Errorf("Testing failed", customer)
	}
}

func TestGetAllCustomers(t *testing.T) {
	kind := util.MustGetenv("DATA_STORE_KIND")
	ctx := context.Background()
	v := url.Values{}
	v.Add("partner_id", "001")
	var customers []Customer
	customers, err := GetAllCustomers(kind, ctx, v)
	log.Println(err)
	if customers == nil {
		t.Errorf("Testing failed", customers)
	}
}

func TestGetCustomerDetails(t *testing.T) {
	kind := util.MustGetenv("DATA_STORE_KIND")
	ctx := context.Background()
	v := url.Values{}
	v.Add("partner_id", "001")
	var customers []Customer
	customers, err := GetAllCustomers(kind, ctx, v)
	log.Println(err)
	if customers == nil {
		t.Errorf("Testing failed", customers)
	}
}
