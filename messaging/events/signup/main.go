package main

import (
	"fmt"
	"log"

	// "github.com/nats-io/go-nats"
	"github.com/user/mq/chassis"
	welcome "github.com/user/mq/messaging/events/signup/welcome"

	"runtime"

	"github.com/user/mq/messaging/do"
)

/**
go run main.go -topic=welcome
todo flag for queue Q1 Q2 ....
*/
func main() {
	signature := chassis.Signature{
		Namespace:   "event.signup",
		Description: "",
		Version:     "0.1",
	}
	do.Flags(&signature)
	fmt.Printf("%+v", signature)

	//-------------- CONNECTION
	natsConnection, _ := do.PoolConnect(signature.Env)
	// natsConnection := do.NatsConnect(signature.Env)

	//-------------- SUBSCRIPTION
	log.Println("TOPIC = " + signature.Topic)
	switch signature.Topic {
	case "welcome":
		// step  #1
		_, err := welcome.SubHandler(signature.TopicName(), signature.TopicName()+"Q1", signature.Env, natsConnection)
		if err != nil {
			log.Print(err)
		}
	case "validation":
		// step  #2
		// welcome.SubHandler(signature.TopicName(), signature.Env, natsConnection, cassandraSession)
	default:
		// all ideal for testing
	}
	// Keep the connection alive
	runtime.Goexit()
}
