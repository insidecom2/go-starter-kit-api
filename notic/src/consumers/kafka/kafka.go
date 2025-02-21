package consumers

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

func ConnectKafka() (sarama.Consumer, error) {

	urlConnect := []string{os.Getenv("KAFKA_URL")}
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.AutoCommit.Enable = true

	return sarama.NewConsumer(urlConnect, config)

}

func ConsumerMessage(topic string) error {
	msgCount := 0
	worker, err := ConnectKafka()
	if err != nil {
		log.Printf("error :%v", worker)
	}
	log.Printf("connected: %v", "Kafka")

	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Printf("error :%v", consumer)
	}
	log.Printf("consumer start !!!")

	signCha := make(chan os.Signal, 1)
	signal.Notify(signCha, syscall.SIGINT, syscall.SIGTERM)

	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				log.Printf("%v", err)
			case msg := <-consumer.Messages():
				msgCount++
				log.Printf(" message: %d %v", msgCount, string(msg.Value))
			case <-signCha:
				log.Println("Interrupt detection")
				doneCh <- struct{}{}
			}
		}
	}()
	<-doneCh

	log.Println(" message done")
	if err := worker.Close(); err != nil {
		log.Printf(" error: %v", err)
	}

	return nil
}
