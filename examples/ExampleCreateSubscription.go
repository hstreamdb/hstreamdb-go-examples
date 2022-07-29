package examples

import (
	"github.com/hstreamdb/hstreamdb-go/hstream"
	"log"
)

func ExampleCreateSubscription() error {
	client, err := hstream.NewHStreamClient(YourHStreamServiceUrl)
	if err != nil {
		log.Fatalf("Creating client error: %s", err)
	}
	defer client.Close()

	streamName := "testStream"
	subId0 := "SubscriptionId0"
	subId1 := "SubscriptionId1"

	// Create a new subscription with ack timeout = 600s, max unAcked records num set to 10000 by default
	if err := client.CreateSubscription(subId0, streamName, func(sub *hstream.Subscription) { sub.AckTimeoutSeconds = 600 }); err != nil {
		log.Fatalf("Creating subscription error: %s", err)
	}

	if err := client.CreateSubscription(subId1, streamName, func(sub *hstream.Subscription) { sub.AckTimeoutSeconds = 600; sub.MaxUnackedRecords = 5000 }); err != nil {
		log.Fatalf("Creating subscription error: %s", err)
	}

	return nil
}
