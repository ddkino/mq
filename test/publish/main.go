package main

// Import packages
import (
	"flag"
	"log"

	"github.com/user/mq/test/publish/signup"

	"github.com/nats-io/go-nats"
	"github.com/user/mq/test/publish/user"
)

func usage() {
	log.Fatalf("Usage: run test --topic=<topic> \n")
}

func main() {
	flag.Usage = usage
	flag.String("env", "development", "environnement")
	TOPIC := flag.String("topic", "full path", "topic.user.login")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		usage()
	}
	natsConnection, _ := nats.Connect("nats://dede:123456@localhost:4222")

	log.Println("Connected to " + nats.DefaultURL)
	log.Println("TOPIC = " + *TOPIC)
	switch *TOPIC {
	case "store.user.login":
		user.Login(natsConnection)
	case "store.user.byid":
		user.Byid(natsConnection)
	case "event.signup.email":
		signup.Email(natsConnection)
	}
	defer natsConnection.Close()

}
