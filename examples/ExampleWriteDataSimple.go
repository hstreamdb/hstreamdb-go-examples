package examples

import (
	"github.com/hstreamdb/hstreamdb-go/hstream"
	"github.com/hstreamdb/hstreamdb-go/hstream/Record"
	"log"
)

func ExampleWriteDataSimple() {
	client, err := hstream.NewHStreamClient(YourHStreamServiceUrl)
	if err != nil {
		log.Fatalf("Creating client error: %s", err)
	}
	defer client.Close()

	producer, err := client.NewProducer("testStream")
	if err != nil {
		log.Fatalf("Creating producer error: %s", err)
	}

	defer producer.Stop()

	payload := map[string]interface{}{
		"key1": "value1",
		"key2": 123,
	}

	hRecord, err := Record.NewHStreamHRecord("testStream", payload)
	if err != nil {
		log.Fatalf("Creating HRecord error: %s", err)
	}

	for i := 0; i < 500; i++ {
		appendRes := producer.Append(hRecord)
		if resp, err := appendRes.Ready(); err != nil {
			log.Printf("Append error: %s", err)
		} else {
			log.Printf("Append response: %s", resp)
		}
	}
}
