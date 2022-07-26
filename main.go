package main

import (
	"github.com/hstreamdb/hstreamdb-go-examples/examples"
	"log"
	"reflect"
	"runtime"
)

func main() {
	xs := []func(){
		examples.ExampleCreateStream,
		examples.ExampleListStreams,

		examples.ExampleWriteDataSimple,
		examples.ExampleWriteDataBuffered,
		examples.ExampleWriteDataWithKey,

		examples.ExampleCreateSubscription,
		examples.ExampleListSubscriptions,

		examples.ExampleConsumeDataSimple,
		examples.ExampleConsumeDataShared,

		examples.ExampleDeleteSubscription,
		examples.ExampleDeleteStream}

	for _, x := range xs {
		runFuncWithLog(x)
	}

}

func getFuncName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func runFuncWithLog(f func()) {
	funcName := getFuncName(f)
	log.Printf("start %s", funcName)
	f()
	log.Printf("end %s", funcName)
}
