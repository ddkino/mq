package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/user/mq/broker"
	"github.com/user/mq/chassis"
	"github.com/user/mq/stores/cassandra"
	"github.com/user/mq/stores/cassandra/queries"

	"encoding/json"
	"runtime"

	"github.com/nats-io/go-nats"
)

func main() {
	signature := chassis.Signature{
		Namespace:   "store.user",
		Topic:       "byid",
		Description: "",
		Version:     "0.1",
	}
	SIGNATURE := fmt.Sprintf("%+v", signature)
	//-------------- FLAGS
	ENV := flag.String("env", "development", "environnement")
	flag.Parse()
	args := flag.Args()
	fmt.Println(*ENV)
	if len(args) > 0 {
		if args[0] == "h" || args[0] == "help" {
			fmt.Println(*ENV)
			fmt.Println(SIGNATURE)
			os.Exit(3)
		}
	}

	//-------------- CONNECTION
	natsBroker := broker.NewBrokerNats(*ENV)
	natsConnection, err := natsBroker.Connect()
	if err != nil {
		panic(err)
	}
	cassandraStore := cassandra.New(*ENV)
	cassandraSession, err := cassandraStore.Connect()
	if err != nil {
		panic(err)
	}
	//-------------- SUBSCRIPTION
	var topic = fmt.Sprintf("%s.%s", signature.Namespace, signature.Topic)
	natsConnection.Subscribe(topic, func(msg *nats.Msg) {
		var dat map[string]interface{}
		json.Unmarshal(msg.Data, &dat)
		// <<<<<<< insert your code here
		// do not exit code on error

		// ~~~~~~~~~~~~ query example
		var error = false
		var out_params  string
		if err := queries.<queryname>(
			cassandraSession,
			dat["input_params"].(string)).Scan(&out_params); err != nil {
			if *ENV != "development" {
				log.Printf("UserByLoginPassword, error=%s", err.Error())
			}
			error = true
			natsConnection.Publish(msg.Reply, []byte(fmt.Sprintf(`{error: "%s"}`, err.Error())))
		}
		// ~~~~~~~~~~~~ end query

		// ~~~~~~~~~~~~ reply
		if ! error {
			if *ENV != "development" {
				fmt.Printf("%s", out_params)
			}
			natsConnection.Publish(msg.Reply, []byte(`{msg: <out_params>}`))
		}
		// >>>>>>> end code
	})
	// Keep the connection alive
	runtime.Goexit()
}
