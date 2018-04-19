package handler

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gocql/gocql"
	"github.com/user/mq/stores/cassandra/queries"

	"github.com/nats-io/go-nats"
)

func Handler(msg *nats.Msg, ENV *string, natsConnection *nats.Conn, cassandraSession *gocql.Session) {

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

}
