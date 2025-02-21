package producer

import (
	"log"
	"os"

	"github.com/IBM/sarama"
)

func connectKafka(url []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	return sarama.NewSyncProducer(url, config)

}

func SendMessage(topic string, key string, message []byte) error {
	url := []string{os.Getenv("KAFKA_URL")}
	producer, err := connectKafka(url)
	if err != nil {
		log.Printf("error : %v", err)
		return err
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Key:   sarama.StringEncoder(key),
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}
	log.Printf("Send message topic(%s) partition(%d) offset(%d)", topic, partition, offset)
	return nil
}
