package examples

import (
	"log"

	"github.com/hstreamdb/hstreamdb-go/hstream"
)

func ExampleCreateStream() error {
	client, err := hstream.NewHStreamClient(YourHStreamServiceUrl)
	if err != nil {
		log.Fatalf("Creating client error: %s", err)
	}
	defer client.Close()

	// Create a new stream with 1 replica, set the data retention to 1800s.
	err = client.CreateStream("testStream",
		hstream.WithReplicationFactor(1),
		hstream.EnableBacklog(1800))
	if err != nil {
		log.Fatalf("Creating stream error: %s", err)
	}

	return nil
}
