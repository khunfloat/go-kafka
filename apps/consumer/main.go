package main

import (
	"fmt"
	"log"

	"github.com/khunfloat/go-kafka/config"
	"github.com/khunfloat/go-kafka/pkg/utils"
)

func main() {
	cfg := config.KafkaConnCfg{
		Url:   "localhost:9092",
		Topic: "shop",
	}
	conn := utils.KafkaConn(cfg)

	for {
		message, err := conn.ReadMessage(10e3)
		if err != nil {
			break
		}
		fmt.Println(string(message.Value))
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}
}
