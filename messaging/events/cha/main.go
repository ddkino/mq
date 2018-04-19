package main

import (
	"fmt"
	"log"

	// "github.com/nats-io/go-nats"
	"github.com/user/mq/chassis"
	create "github.com/user/mq/messaging/events/cha/create"
	delete "github.com/user/mq/messaging/events/cha/delete"

	"runtime"

	"github.com/user/mq/messaging/do"
)

/**
go run main.go -topic=welcome
todo flag for queue Q1 Q2 ....
*/
func main() {
	signature := chassis.Signature{
		Namespace:   "event.cha",
		Description: "",
		Version:     "0.1",
	}
	do.Flags(&signature)
	fmt.Printf("%+v", signature)

	//-------------- CONNECTION NATS
	// natsConnection, _ := do.PoolConnect(signature.Env)
	natsConnection := do.NatsConnect(signature.Env)
	elasticConnection := do.ElasticConnect(signature.Env)

	//-------------- SUBSCRIPTION
	log.Println("TOPIC = " + signature.Topic)
	switch signature.Topic {
	case "create":
		// step  #1
		_, err := create.SubHandler(signature.TopicName(), signature.TopicName()+"Q1", signature.Env, natsConnection, elasticConnection)
		if err != nil {
			log.Print(err)
		}
	case "delete":
		// step  #1
		_, err := delete.SubHandler(signature.TopicName(), signature.TopicName()+"Q1", signature.Env, natsConnection, elasticConnection)
		if err != nil {
			log.Print(err)
		}
	default:
		// all ideal for testing
	}

	// Keep the connection alive
	runtime.Goexit()
}
