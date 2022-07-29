package examples

import (
	"github.com/hstreamdb/hstreamdb-go/hstream"
	"github.com/hstreamdb/hstreamdb-go/hstream/Record"
	"log"
	"sync"
)

func ExampleConsumeDataShared() error {
	client, err := hstream.NewHStreamClient(YourHStreamServiceUrl)
	if err != nil {
		log.Fatalf("Creating client error: %s", err)
	}
	defer client.Close()

	subId1 := "SubscriptionId1"

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		consumer := client.NewConsumer("consumer-1", subId1)
		defer wg.Done()
		defer consumer.Stop()

		dataChan := consumer.StartFetch()
		fetchedRecords := make([]Record.RecordId, 0, 100)
		for recordMsg := range dataChan {
			receivedRecords, err := recordMsg.Result, recordMsg.Err
			if err != nil {
				log.Printf("[consumer-1]: Stream fetching error: %s", err)
				continue
			}

			for _, record := range receivedRecords {
				recordId := record.GetRecordId()
				log.Printf("[consumer-1]: Receive %s record: record id = %s, payload = %s",
					record.GetRecordType(), record.GetRecordId(), record.GetPayload())
				fetchedRecords = append(fetchedRecords, recordId)
				record.Ack()
			}

			if len(fetchedRecords) == 100 {
				log.Println("[consumer-1]: Stream fetching stopped")
				break
			}
		}
	}()

	go func() {
		consumer := client.NewConsumer("consumer-2", subId1)
		defer wg.Done()
		defer consumer.Stop()

		dataChan := consumer.StartFetch()
		fetchedRecords := make([]Record.RecordId, 0, 100)
		for recordMsg := range dataChan {
			receivedRecords, err := recordMsg.Result, recordMsg.Err
			if err != nil {
				log.Printf("[consumer-2]: Stream fetching error: %s", err)
				continue
			}

			for _, record := range receivedRecords {
				recordId := record.GetRecordId()
				log.Printf("[consumer-2]: Receive %s record: record id = %s, payload = %s",
					record.GetRecordType(), record.GetRecordId(), record.GetPayload())
				fetchedRecords = append(fetchedRecords, recordId)
				record.Ack()
			}

			if len(fetchedRecords) == 100 {
				log.Println("[consumer-2]: Stream fetching stopped")
				break
			}
		}
	}()

	wg.Wait()

	return nil
}
