package main

import "fmt"

type DataService interface {
	Connect() string
}

type ClickHouse struct {
}

func (ch *ClickHouse) Connect() string {
	return "Connected to ClickHouse"
}

type Redis struct {
}

func (rd *Redis) Connect() string {
	return "Connected to Redis"
}

// DataService Factory
func getService(serviceType string) DataService {
	switch serviceType {
	case "clickhouse":
		return &ClickHouse{}
	case "redis":
		return &Redis{}
	default:
		return nil
	}
}

func main() {
	clickhouse := getService("clickhouse")
	fmt.Println(clickhouse.Connect())

	redis := getService("redis")
	fmt.Println(redis.Connect())
}
