package do

import (
	"fmt"
	"log"
	"sync"

	"github.com/gocql/gocql"
	"github.com/nats-io/go-nats"
	"github.com/olivere/elastic"
	"github.com/user/mq/broker"
	"github.com/user/mq/stores/cassandra"
	elasticStore "github.com/user/mq/stores/elastic"
)

func PoolConnect(ENV string) (*nats.Conn, *gocql.Session) {
	//-------------- CONNECTION async + waitgroup
	var natsConnection *nats.Conn
	var cassandraSession *gocql.Session
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		natsBroker := broker.NewBrokerNats(ENV)
		natsConn, err := natsBroker.Connect()
		if err != nil {
			log.Println("natsBroker err")
			panic(err)
		}
		natsConnection = natsConn
		fmt.Printf("natsConn %v", natsConn.IsConnected())
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		cassandraStore := cassandra.New(ENV)
		session, err := cassandraStore.Connect()
		if err != nil {
			log.Println("cassandraStore err")
			panic(err)
		}
		cassandraSession = session
		wg.Done()
	}()
	wg.Wait()
	return natsConnection, cassandraSession
}

func NatsConnect(ENV string) *nats.Conn {
	var natsConnection *nats.Conn
	natsBroker := broker.NewBrokerNats(ENV)
	natsConn, err := natsBroker.Connect()
	if err != nil {
		log.Print("natsBroker err")
		panic(err)
	}
	natsConnection = natsConn
	fmt.Printf("natsConn %v", natsConn.IsConnected())
	return natsConnection
}

func CassandraConnect(ENV string) *gocql.Session {
	var cassandraSession *gocql.Session
	cassandraStore := cassandra.New(ENV)
	session, err := cassandraStore.Connect()
	if err != nil {
		log.Print("cassandraStore err")
		panic(err)
	}
	cassandraSession = session
	return cassandraSession
}

func ElasticConnect(ENV string) *elastic.Client {
	var elasticClient *elastic.Client
	elasticStore := elasticStore.New(ENV)
	elasticClient, err := elasticStore.Connect()
	if err != nil {
		log.Print("elasticConnection err")
		panic(err)
	}
	return elasticClient
}
