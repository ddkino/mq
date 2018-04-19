package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/gocql/gocql"
	"github.com/nats-io/go-nats"
	"github.com/user/mq/broker"
	"github.com/user/mq/stores/cassandra"
)

func main() {
	//-------------- CONNECTION async + waitgroup
	var wg sync.WaitGroup
	var natsConnection nats.Conn
	var ENV *string
	tmp := "dev"
	ENV = &tmp
	wg.Add(1)
	go func() {
		natsBroker := broker.NewBrokerNats(*ENV)
		natsConn, err := natsBroker.Connect()
		if err != nil {
			log.Print("natsBroker err")
			panic(err)
		}
		natsConnection = *natsConn
		fmt.Printf("natsConn %v\n", natsConn.IsConnected())
		wg.Done()
	}()

	var cassandraSession gocql.Session
	wg.Add(1)
	go func(session *gocql.Session) {
		cassandraStore := cassandra.New(*ENV)
		session, err := cassandraStore.Connect()
		if err != nil {
			log.Print("cassandraStore err")
			panic(err)
		}
		wg.Done()
	}(&cassandraSession)
	// _ = natsConnection
	// _ = cassandraSession

	wg.Wait()
	log.Println("output:")
	fmt.Printf("natsConnection %v\n", natsConnection.IsConnected())
	fmt.Printf("cassandraSession %v", cassandraSession)
}
