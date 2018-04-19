package user

import (
	"log"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/nats-io/go-nats"
	error "github.com/user/mq/models/error"
	user "github.com/user/mq/models/user"
)

func Signup(natsConnection *nats.Conn) {

	subject := "signup.sendmail"
	reply, _ := natsConnection.Request(subject, []byte(`{"url": "token1234567890,051cb46b-724d-4591-8770-1bbe6f060d08", 
		"username": "dede.dede", 
		"email": "david.ddkino@gmail.com", 
		"from": "ddkino.dd@gmail.com"}`), 1*time.Second)
	log.Println("Published message on subject " + subject)
	model := &user.Model{}
	error := &error.Type{}
	err := proto.Unmarshal(reply.Data, model)
	if err != nil {
		log.Print("unmarshaling error: ", err)
	}
	log.Printf("Received [%v] {%s}: '%v'\n", reply.Subject, string(reply.Data), error.GetCode())
}

func Login(natsConnection *nats.Conn) {

	subject := "store.user.login"
	reply, _ := natsConnection.Request(subject, []byte(`{
		"username": "dede.dede", 
		"email": "david.ddkino@gmail.com", 
		"password": "123456"}`), 1*time.Second)
	log.Println("Published message on subject " + subject)
	// model := &user.Model{}
	error := &error.Type{}
	err := proto.Unmarshal(reply.Data, error)
	if err != nil {
		log.Print("unmarshaling error: ", err)
	}
	log.Printf("Received [%v] error :%s / errocode'%v'\n", reply.Subject, error.GetMsg(), error.GetCode())
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

func Byid(natsConnection *nats.Conn) {
	subject := "store.user.byid"
	reply, _ := natsConnection.Request(subject, []byte(`{"user_id": "051cb46b-724d-4591-8770-1bbe6f060d08"}`), 1*time.Second)
	_ = reply
	log.Println("Published message on subject " + subject)
	model := &user.Model{}
	error := &error.Type{}
	err := proto.Unmarshal(reply.Data, model)
	if err != nil {
		log.Print("unmarshaling error: ", err)
	}
	log.Printf("Received [%v] {%s}: '%v'\n", reply.Subject, string(reply.Data), error.GetCode())
}
