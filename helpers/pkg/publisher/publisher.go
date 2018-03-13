package publisher

import (
	"context"
	"log"
	"net/http"
	"encoding/json"
	"customer-consumer/helpers/util"
	"customer-consumer/models"
	"cloud.google.com/go/pubsub"
	"fmt"
	"io/ioutil"
)

// PublishCustomer responds with 200
func PublishCustomer(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	client, _ := pubsub.NewClient(ctx, util.MustGetenv("GOOGLE_CLOUD_PROJECT"))
	topic := client.Topic(util.MustGetenv("PROCESSED_CUSTOMER_DATA"))
	customerObj := models.Customer{PartnerId: "Chefd001",Email:"vikas.taank@sidusa.com",FirstName:"Vikas",LastName:"Taank",Address:"NIT FBD,HRY,IN",PhoneNumber:"9990836778"}
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
		fmt.Printf("Published a customer's data; customer_id: %v\n", custId)
		w.WriteHeader(http.StatusOK)
	}
}


// InsertCustomerData publishes raw customer data from netsuite to pub/sub.
func InsertCustomerData(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		var customer models.Customer
		err = json.Unmarshal(body, &customer)
	
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Error found: %v", err)
			return
		}
	
		ctx := context.Background()
		projectID := util.MustGetenv("GOOGLE_CLOUD_PROJECT")
		client, err := pubsub.NewClient(ctx, projectID)
		if err != nil {
			log.Printf("Failed to create client: %v", err)
		}
	
		t := client.Topic(util.MustGetenv("RAW_CUSTOMER_DATA"))
		output, _ := json.Marshal(customer)
		result := t.Publish(ctx, &pubsub.Message{
			Data: output,
		})
		id, err := result.Get(ctx)
		if err != nil {
			log.Printf("Encountered error publishing: %v", err)
		}
		log.Printf("Published a customer message in msg ID: %v\n", id)
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/text; charset=utf-8")
	
	}