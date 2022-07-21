package main

import "github.com/hstreamdb/hstreamdb-go-examples/examples"

func main() {
	examples.ExampleCreateStream()
	examples.ExampleListStreams()

	examples.ExampleWriteDataSimple()
	examples.ExampleWriteDataBuffered()
	examples.ExampleWriteDataWithKey()

	examples.ExampleCreateSubscription()
	examples.ExampleListSubscriptions()

	examples.ExampleConsumeDataSimple()
	examples.ExampleConsumeDataShared()

	examples.ExampleDeleteSubscription()
	examples.ExampleDeleteStream()
}
