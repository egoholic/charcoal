package config

type MongoDBCreds struct {
	ConnectionString func() string
}

var MongoDB = MongoDBCreds{func() string { return "mongodb://127.0.0.1:27017" }}
