package broker

import (
	"fmt"
	"log"

	"github.com/nats-io/go-nats"
)

func (n *NatsType) Connect() (*nats.Conn, error) {

	params := fmt.Sprintf("nats://%s:%s@%s:%d",
		n.Auth.Username,
		n.Auth.Password,
		n.Host,
		n.Port)

	log.Println("Try to connect to broker NATS with params: " + params)
	natsConnection, err := nats.Connect(params)
	if !natsConnection.IsConnected() {
		fmt.Printf("error connection")
	}
	return natsConnection, err
}

func (n *NatsType) ConnectCluster(paramsArr []NatsType) (*nats.Conn, error) {

	var servers string
	for _, n := range paramsArr {
		servers = fmt.Sprintf("nats://%s:%s@%s:%d",
			n.Auth.Username,
			n.Auth.Password,
			n.Host,
			n.Port)
	}

	log.Println("Try to connect to servers: " + servers)
	natsConnection, err := nats.Connect(servers)
	if !natsConnection.IsConnected() {
		fmt.Printf("error connection")
	}
	return natsConnection, err
}
