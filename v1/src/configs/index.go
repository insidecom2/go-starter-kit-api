package configs

type Config struct {
	DBConfig    DBConfig
	SeverConfig SeverConfig
}

type DBConfig struct {
	Type     string
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type SeverConfig struct {
	Port string
}
