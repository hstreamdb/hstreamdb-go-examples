package examples

import (
	"github.com/hstreamdb/hstreamdb-go/hstream"
	"log"
)

func ExampleDeleteStream() {
	client, err := hstream.NewHStreamClient(YourHStreamServiceUrl)
	if err != nil {
		log.Fatalf("Creating client error: %s", err)
	}
	defer client.Close()

	if err := client.DeleteStream("testStream"); err != nil {
		log.Fatalf("Deleting stream error: %s", err)
	}
}
