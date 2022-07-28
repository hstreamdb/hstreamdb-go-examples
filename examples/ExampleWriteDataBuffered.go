package examples

import (
	"fmt"
	"github.com/hstreamdb/hstreamdb-go/hstream"
	"github.com/hstreamdb/hstreamdb-go/hstream/Record"
	"log"
	"sync"
)

func ExampleWriteDataBuffered() {
	client, err := hstream.NewHStreamClient(YourHStreamServiceUrl)
	if err != nil {
		log.Fatalf("Creating client error: %s", err)
	}
	defer client.Close()

	producer, err := client.NewBatchProducer("testStream", hstream.WithBatch(10, hstream.DEFAULT_MAX_BATCHRECORDS_SIZE))
	if err != nil {
		log.Fatalf("Creating batch producer error: %s", err)
	}
	defer producer.Stop()

	wg := sync.WaitGroup{}
	wg.Add(1)

	syncStore := sync.Map{}
	key := "test-key-1"

	go func() {
		defer wg.Done()
		result := make([]hstream.AppendResult, 0, 100)
		for i := 0; i < 100; i++ {
			rawRecord, _ := Record.NewHStreamRawRecord("key-1", []byte(fmt.Sprintf("test-value-%s-%d", key, i)))
			r := producer.Append(rawRecord)
			result = append(result, r)
		}
		syncStore.Store(key, result)
	}()

	wg.Wait()
	syncStore.Range(func(key, value interface{}) bool {
		k := key.(string)
		res := value.([]hstream.AppendResult)
		for idx, r := range res {
			resp, err := r.Ready()
			if err != nil {
				log.Printf("write error: %s\n", err.Error())
			}
			log.Printf("[key: %s]: record[%d]=%s\n", k, idx, resp.String())
		}
		return true
	})
}
