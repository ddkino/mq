package cassandra

import (
	"log"

	"github.com/gocql/gocql"
)

func (c *Cassandra) Connect() (*gocql.Session, error) {

	cluster := gocql.NewCluster(c.Host)
	cluster.Port = c.Port
	cluster.Keyspace = c.Keyspace
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: c.Auth.Username,
		Password: c.Auth.Password,
	}
	log.Printf("Try to connect to Cassandra %s %d", c.Host, c.Port)
	session, err := cluster.CreateSession()
	return session, err
}
