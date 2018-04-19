package elastic

var GlobalDev = Server{
	Host: "127.0.0.1",
	Port: 9200,
	Auth: Auth{
		Username: "",
		Password: "",
	},
}

var GlobalProd = Server{
	Host: "127.0.0.1",
	Port: 9200,
	Auth: Auth{
		Username: "",
		Password: "",
	},
}

func New(env string) Server {
	switch env {
	case "production":
		return GlobalProd
	case "dev":
		return GlobalDev
	default:
		return GlobalDev
	}
}
