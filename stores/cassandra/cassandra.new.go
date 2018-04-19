package cassandra

var GlobalDev = Cassandra{
	Host:     "127.0.0.1",
	Port:     9042,
	Keyspace: "cha_global",
	Auth: Auth{
		Username: "cassandra",
		Password: "cassandra",
	},
}

var GlobalProd = Cassandra{
	Host:     "127.0.0.1",
	Port:     9042,
	Keyspace: "cha_global",
	Auth: Auth{
		Username: "cassandra",
		Password: "cassandra",
	},
}

func NewStoreCassandra() Cassandra {
	return GlobalDev
}

func New(env string) Cassandra {
	switch env {
	case "production":
		return GlobalProd
	case "dev":
		return GlobalDev
	default:
		return GlobalDev
	}
}
