package examples

import (
	"fmt"
	"github.com/hstreamdb/hstreamdb-go/hstream"
	"log"
	"sync"
)

func ExampleWriteDataWithKey() {
	client, err := hstream.NewHStreamClient(YourHStreamServiceUrl)
	if err != nil {
		log.Fatalf("Creating client error: %s", err)
	}
	defer client.Close()

	producer, err := client.NewBatchProducer("testStream", hstream.WithBatch(10, 150))
	defer producer.Stop()

	keys := []string{"test-key1", "test-key2", "test-key3"}
	rids := sync.Map{}
	wg := sync.WaitGroup{}
	wg.Add(3)
	for _, key := range keys {
		go func(key string) {
			result := make([]hstream.AppendResult, 0, 100)
			for i := 0; i < 100; i++ {
				rawRecord, _ := hstream.NewHStreamRawRecord("key-1", []byte(fmt.Sprintf("test-value-%s-%d", key, i)))
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
