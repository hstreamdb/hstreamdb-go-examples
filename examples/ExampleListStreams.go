package examples

import (
	"fmt"
	"github.com/hstreamdb/hstreamdb-go/hstream"
	"log"
)

func ExampleListStreams() {
	client, err := hstream.NewHStreamClient(YourHStreamServiceUrl)
	if err != nil {
		log.Fatalf("Creating client error: %s", err)
	}
	defer client.Close()

	streams, err := client.ListStreams()
	if err != nil {
		log.Fatalf("Listing streams error: %s", err)
	}

	fmt.Println(streams)
}
