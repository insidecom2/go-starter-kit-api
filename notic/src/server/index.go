package server

import (
	"go/consumer/src/configs"
	consumers "go/consumer/src/consumers/kafka"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type SeverStruct struct {
	configs *configs.Configs
}

func Init() {
	config := configs.Configs{}
	server := &SeverStruct{
		configs: &config,
	}
	server.loadEnv()
}

func (s *SeverStruct) loadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Cannot load .env : %v", err)
	}

	s.configs.KafkaConfig.Broker = os.Getenv("KAFKA_URL")
	consumers.ConsumerMessage("notic")
}
