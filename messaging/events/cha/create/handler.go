package handler

import (
	"context"
	"log"

	"encoding/json"

	"github.com/nats-io/go-nats"
	"github.com/olivere/elastic"
	elasticStore "github.com/user/mq/stores/elastic"
)

func SubHandler(topic string, Q string, ENV string, natsConnection *nats.Conn, elasticConnection *elastic.Client) (*nats.Subscription, error) {
	return natsConnection.QueueSubscribe(topic, Q, func(msg *nats.Msg) {

		log.Print(topic)
		var dat map[string]interface{}
		json.Unmarshal(msg.Data, &dat)

		// <<<<<<< insert your code here
		// do not exit code on error
		log.Print(dat)
		if _, err := dat["id"]; !err {
			log.Println("id missing")
			return
		}
		if _, err := dat["a_id"]; !err {
			log.Println("user missing")
			return
		}
		ctx := context.Background()
		cha := elasticStore.Cha{
			C:       dat["c"].(string),
			H:       dat["h"].(string),
			C_alt:   dat["c_alt"].(string),
			H_alt:   dat["h_alt"].(string),
			Lang:    dat["lang"].(string),
			Created: dat["datetime_created_at"].(string),
			User: elasticStore.User{
				Id: dat["a_id"].(string),
			},
		}
		_, err := elasticConnection.Index().
			Index(elasticStore.IndexByCha.Name).
			Type(elasticStore.IndexByCha.Type).
			Id(dat["id"].(string)).
			BodyJson(cha).
			Do(ctx)
		if err != nil {
			// Handle error
			log.Println(err)
			return
		}
		// >>>>>>> end code
	})
}
