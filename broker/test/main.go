package main

import (
	"log"

	"github.com/nats-io/go-nats"

	"fmt"

	broker "github.com/user/mq/broker"
)

func main() {
	b := broker.NewBrokerNats()
	fmt.Print(b)
	n, err := b.ConnectCluster([]broker.NatsType{b})
	log.Println(n)
	if err != nil {
		fmt.Errorf("signup.sendmail #SendRequest %s", err)
		panic(err)
	} else {
		log.Println("Connected OK ")
	}
}

func ok() {

	_, err := nats.Connect("nats://dede:123456@localhost:4222")

	if err != nil {
		fmt.Errorf("signup.sendmail #SendRequest %s", err)
		panic(err)
	} else {
		log.Println("Connected OK ")
	}

}
