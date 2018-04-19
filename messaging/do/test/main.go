package main

import (
	"flag"
	"fmt"

	"github.com/user/mq/services/do"
)

func main() {
	//-------------- FLAGS
	ENV := flag.String("env", "development", "environnement")
	flag.Parse()
	fmt.Println(*ENV)

	natsConnection, cassandraSession := do.PoolConnect(ENV)
	fmt.Printf("natsConnection %v\n", natsConnection.IsConnected())
	fmt.Printf("cassandraSession %v\n", cassandraSession)
	// _, err := natsConnection.Subscribe("topic", func(msg *nats.Msg) {
	// })
	// if err != nil {
	// 	panic(err)
	// }
}
