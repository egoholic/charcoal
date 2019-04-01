package config

import (
	"fmt"
	"time"
)

func ServiceName() string {
	return "auth"
}

func ServiceInstanceName() string {
	return "1"
}

func ServiceFullName() string {
	return fmt.Sprintf("%s-%s", ServiceName(), ServiceInstanceName())
}

func MongoDBConnectionString() string {
	return "mongodb://localhost:27017"
}

func MongoDBDatabaseName() string {
	return ServiceName()
}

func HTTPServicePort() string {
	return ":8080"
}

func HTTPTimeOut() time.Duration {
	return 5 * time.Second
}
