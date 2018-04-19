package handler

import (
	"github.com/user/mq/email"

	"encoding/json"

	"github.com/gocql/gocql"

	"github.com/nats-io/go-nats"
)

func SubHandler(topic string, Q string, ENV string, natsConnection *nats.Conn, cassandraSession *gocql.Session) {
	natsConnection.QueueSubscribe(topic, Q, func(msg *nats.Msg) {

		var dat map[string]interface{}
		json.Unmarshal(msg.Data, &dat)

		// <<<<<<< insert your code here
		// do not exit code on error

		// todo ;
		/**
		var templates *template.Template
		templates = template.Must(template.ParseFiles("edit.html", "view.html"))
		*/

		myEmail := email.New()
		myEmail.Init(dat["from"].(string),
			[]string{dat["email"].(string)},
			"signup welcome",
			"",
		)
		err := myEmail.FactoryEmail("signup", map[string]string{
			"Name": dat["username"].(string),
			"URL":  dat["url"].(string),
		})

		if err != nil {
			panic(err)
		}
		err = myEmail.SendEmail(ENV)
		if err != nil {
			panic(err)
		}
		// >>>>>>> end code
	})
}
