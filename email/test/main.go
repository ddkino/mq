package main

import (
	"github.com/user/mq/email"
)

func main() {

	myEmail := email.New()
	myEmail.Init("ddkino.dd@gmail.com",
		[]string{"david.ddkino@gmail.com"},
		"signup subject",
		"",
	)
	err := myEmail.FactoryEmail("signup", map[string]string{
		"Name": "dede",
		"URL":  "https://google.fr",
	})

	if err != nil {
		panic(err)
	}
	err = myEmail.SendEmail("development")
	if err != nil {
		panic(err)
	}
}
