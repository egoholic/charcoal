package config

type MongoDBCreds struct {
	ConnectionString func() string
}

var MongoDB = MongoDBCreds{func() string { return "mongodb://localhost:27017" }}
