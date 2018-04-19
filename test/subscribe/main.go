package main

// Import packages
import (
	"fmt"
	"runtime"

	"github.com/nats-io/go-nats"
)

func main() {
	natsConnection, _ := nats.Connect("nats://dede:123456@localhost:4222")
	natsConnection.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	runtime.Goexit()
}
