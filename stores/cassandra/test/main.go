package main

import (
	"fmt"

	"github.com/gocql/gocql"
	"github.com/user/mq/stores/cassandra"
)

func main() {
	//ok()
	c := cassandra.NewStoreCassandra()
	session, err := c.Connect()
	if err != nil {
		panic(err)
	} else {
		fmt.Println(session)
	}
}

func ok() {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Port = 9042
	cluster.Keyspace = "cha_global"
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: "cha_admin",
		Password: "123456",
	}
	session, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	} else {
		fmt.Println(session)
	}
}
