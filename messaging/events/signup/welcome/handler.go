package handler

import (
	"log"

	"github.com/user/mq/email"

	"encoding/json"

	"github.com/nats-io/go-nats"
)

func SubHandler(topic string, Q string, ENV string, natsConnection *nats.Conn) (*nats.Subscription, error) {
	return natsConnection.QueueSubscribe(topic, Q, func(msg *nats.Msg) {

		log.Print(topic)
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
		}, ENV)

		if err != nil {
			log.Print("FactoryEmail err")
			log.Println(err)
			return
		}
		err = myEmail.SendEmail(ENV)
		if err != nil {
			log.Print("sendmail err")
			log.Println(err)
			return
		}
		// >>>>>>> end code
	})
}
