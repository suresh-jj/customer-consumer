package publisher

import (
	"context"
	"customer-consumer/helpers/util"
	"customer-consumer/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"cloud.google.com/go/pubsub"
)

// PublishCustomer responds with 200
func PublishCustomer(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	client, _ := pubsub.NewClient(ctx, util.MustGetenv("GOOGLE_CLOUD_PROJECT"))
	topic := client.Topic(util.MustGetenv("PROCESSED_CUSTOMER_DATA"))
	customerObj := models.Customer{PartnerId: "Chefd001", Email: "test@test.com", FirstName: "test", LastName: "test", Address: "NIT FBD,HRY,IN", PhoneNumber: "9990836778"}
	theJSON, _ := json.Marshal(customerObj)
	var data = string(theJSON)
	result := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(data),
	})
	custId, err1 := result.Get(ctx)
	if err1 != nil {
		log.Printf("%s", err1)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		log.Printf("Published a customer's data; customer_id: %v\n", custId)
		w.WriteHeader(http.StatusOK)
	}
}

// AddCustomer publishes raw customer data
func AddCustomer(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	var customer models.Customer
	err = json.Unmarshal(body, &customer)
	var errormsg = models.CustomerError{}
	w.Header().Add("Content-Type", "application/text; charset=utf-8")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errormsg.Error = err.Error()
		json.NewEncoder(w).Encode(errormsg)
		log.Printf("%s", err)
		return
	}

	//regex validations:
	var validEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	var validPhone = regexp.MustCompile(`^(\([0-9]{3}\) |[0-9]{3}-)[0-9]{3}-[0-9]{4}$`)

	if !validEmail.MatchString(customer.Email) {
		w.WriteHeader(http.StatusBadRequest)
		errormsg.Error = "Email not valid"
		json.NewEncoder(w).Encode(errormsg)
		return
	}
	if !validPhone.MatchString(customer.PhoneNumber) {
		w.WriteHeader(http.StatusBadRequest)
		errormsg.Error = "Phone Number not valid"
		json.NewEncoder(w).Encode(errormsg)
		return
	}

	ctx := context.Background()
	projectID := util.MustGetenv("GOOGLE_CLOUD_PROJECT")
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Printf("Failed to create client: %v", err)
		w.WriteHeader(http.StatusNotFound)
		errormsg.Error = err.Error()
		json.NewEncoder(w).Encode(errormsg)
		return
	}

	t := client.Topic(util.MustGetenv("RAW_CUSTOMER_DATA"))
	output, _ := json.Marshal(customer)
	result := t.Publish(ctx, &pubsub.Message{
		Data: output,
	})
	id, err := result.Get(ctx)
	if err != nil {
		log.Printf("Encountered error publishing: %v", err)
		w.WriteHeader(http.StatusNotFound)
		errormsg.Error = err.Error()
		json.NewEncoder(w).Encode(errormsg)
		return
	}
	log.Printf("Published a customer message in msg ID: %v\n", id)
	var succesmsg = models.CustomerSuccess{}
	succesmsg.Id = id
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(succesmsg)
}
