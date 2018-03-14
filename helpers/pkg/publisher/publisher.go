package publisher

import (
	"context"
	"customer-consumer/helpers/util"
	"customer-consumer/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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
		fmt.Printf("Published a customer's data; customer_id: %v\n", custId)
		w.WriteHeader(http.StatusOK)
	}
}
