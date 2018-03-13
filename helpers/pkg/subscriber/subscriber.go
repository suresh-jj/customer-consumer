package subscriber

import (
	"context"
	"customer-consumer/helpers/util"
	"customer-consumer/models"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"cloud.google.com/go/pubsub"
)

// SubscribeCustomer listens for new customer messages and save to log (and database)
func SubscribeCustomer() {
	ctx := context.Background()
	projectID := util.MustGetenv("GOOGLE_CLOUD_PROJECT")
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Printf("Failed to create client: %v", err)
	}

	var mu sync.Mutex
	sub := client.Subscription(util.MustGetenv("GET_RAW_CUSTOMER_DATA"))
	cctx, cancel := context.WithCancel(ctx)
	err = sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		mu.Lock()
		defer mu.Unlock()
		log.Printf(fmt.Sprintf("{\"message\":%s}", string(msg.Data)))
		kind := util.MustGetenv("DATA_STORE_KIND_CUSTOMER")
		var customer models.Customer
		err = json.Unmarshal(msg.Data, &customer)
		if err != nil {
			log.Printf("Failed to insert customer data: %v", err)
		}
		key, err := models.InsertCustomer(kind, customer, ctx)
		if key != nil {
			log.Printf(key.Name + " published")
		}
		if err != nil {
			log.Printf(err.Error())
		}
		msg.Ack()
	})
	if err == nil {
		cancel()
		return
	}
	log.Printf("SubscribeCustomer() exited with error: %v", err)
}
