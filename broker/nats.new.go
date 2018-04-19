package broker

var n = NatsType{
	Host: "localhost",
	Port: 4222,
	Auth: Auth{
		Username: "dede",
		Password: "123456",
	},
}

var docker = NatsType{
	Host: "nats-1",
	Port: 4222,
	Auth: Auth{
		Username: "dede",
		Password: "123456",
	},
}

func NewBrokerNats(env string) NatsType {
	switch env {
	case "production":
		return docker
	case "devlopment":
		return n
	default:
		return n
	}
}
