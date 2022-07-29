package examples

import (
	"github.com/hstreamdb/hstreamdb-go/hstream"
	"github.com/hstreamdb/hstreamdb-go/hstream/Record"
	"log"
)

func ExampleConsumeDataSimple() {
	client, err := hstream.NewHStreamClient(YourHStreamServiceUrl)
	if err != nil {
		log.Fatalf("Creating client error: %s", err)
	}
	defer client.Close()

	subId0 := "SubscriptionId0"
	consumer := client.NewConsumer("consumer-0", subId0)
	defer consumer.Stop()

	dataChan := consumer.StartFetch()
	fetchedRecords := make([]Record.RecordId, 0, 100)
	for recordMsg := range dataChan {
		receivedRecords, err := recordMsg.Result, recordMsg.Err
		if err != nil {
			log.Printf("Stream fetching error: %s", err)
			break
		}

		for _, record := range receivedRecords {
			recordId := record.GetRecordId()
			log.Printf("Receive %s record: record id = %s, payload = %s",
				record.GetRecordType(), record.GetRecordId(), record.GetPayload())
			fetchedRecords = append(fetchedRecords, recordId)
			record.Ack()
		}

		if len(fetchedRecords) == 100 {
			log.Println("Stream fetching stopped")
			break
		}
	}
}
