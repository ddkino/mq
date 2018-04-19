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
		if _, err := dat["id"]; !err {
			log.Print("missing id")
			return
		}
		ctx := context.Background()
		elasticConnection.Delete().
			Index(elasticStore.IndexByCha.Name).
			Type(elasticStore.IndexByCha.Type).
			Id(dat["id"].(string)).
			Do(ctx)
		// >>>>>>> end code
	})
}
