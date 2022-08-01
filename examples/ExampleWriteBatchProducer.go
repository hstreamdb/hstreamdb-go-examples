package examples

import (
	"fmt"
	"github.com/hstreamdb/hstreamdb-go/hstream"
	"github.com/hstreamdb/hstreamdb-go/hstream/Record"
	"log"
	"sync"
)

func ExampleWriteBatchProducer() error {
	client, err := hstream.NewHStreamClient(YourHStreamServiceUrl)
	if err != nil {
		log.Fatalf("Creating client error: %s", err)
	}
	defer client.Close()

	producer, err := client.NewBatchProducer("testStream", hstream.WithBatch(10, 500))
	if err != nil {
		log.Fatalf("Creating producer error: %s", err)
	}
	defer producer.Stop()

	keys := []string{"test-key1", "test-key2", "test-key3", "test-key4", "test-key5"}
	rids := sync.Map{}
	wg := sync.WaitGroup{}
	wg.Add(5)

	for _, key := range keys {
		go func(key string) {
			result := make([]hstream.AppendResult, 0, 100)
			for i := 0; i < 100; i++ {
				rawRecord, _ := Record.NewHStreamRawRecord(key, []byte(fmt.Sprintf("test_value---%s-%d", key, i)))
				r := producer.Append(rawRecord)
				result = append(result, r)
			}
			rids.Store(key, result)
			wg.Done()
		}(key)
	}

	wg.Wait()
	rids.Range(func(key, value interface{}) bool {
		k := key.(string)
		res := value.([]hstream.AppendResult)
		for i := 0; i < 100; i++ {
			resp, err := res[i].Ready()
			if err != nil {
				log.Printf("write error: %s\n", err.Error())
			}
			log.Printf("[key: %s]: record[%d]=%s\n", k, i, resp.String())
		}
		return true
	})

	return nil
}
