package signup

import (
	"github.com/user/mq/crypto"

	"fmt"
	"log"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/nats-io/go-nats"
	error "github.com/user/mq/models/error"
	user "github.com/user/mq/models/user"
)

func Email(natsConnection *nats.Conn) {

	subject := "event.signup.welcome"
	token := "234567890,051cb46b-724d-4591-8770-1bbe6f060d08"
	const SECRET_KEY = "dede"
	PRIVATE_KEY := fmt.Sprintf("%s %d %d %s",
		SECRET_KEY,
		time.Now().Day(),
		time.Now().Hour(),
		token)
	PUBLIC_KEY := crypto.Sha256(PRIVATE_KEY)
	URL := fmt.Sprintf("https://mysite.fr?key=%s&id=%s", PUBLIC_KEY, token)
	reply, _ := natsConnection.Request(subject, []byte(fmt.Sprintf(`{
		"url": "%s", 
		"username": "dede.dede", 
		"email": "david.ddkino@gmail.com", 
		"from": "ddkino.dd@gmail.com"}`, URL)), time.Second*3)
	log.Println("Published message on subject " + subject)
	_ = reply
}

func Test(natsConnection *nats.Conn) {

	subject := "store.user.test"
	reply, _ := natsConnection.Request(subject, []byte(`{
		"username": "dede.dede", 
		"email": "david.ddkino@gmail.com", 
		"password": "123456"}`), 1*time.Second)
	log.Println("Published message on subject " + subject)
	model := &user.Model{}
	error := &error.Type{}
	err := proto.Unmarshal(reply.Data, model)
	if err != nil {
		log.Print("unmarshaling error: ", err)
	}
	log.Printf("Received [%v] {%s}: '%v'\n", reply.Subject, string(reply.Data), error.GetCode())
}
