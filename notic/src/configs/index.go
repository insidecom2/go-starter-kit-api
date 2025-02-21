package configs

type Configs struct {
	KafkaConfig KafkaConfig
}

type KafkaConfig struct {
	Broker string
}
