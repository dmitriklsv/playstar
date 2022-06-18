package configs

import "github.com/ilyakaznacheev/cleanenv"

type Configs struct {
	Server struct {
		Addr string `yaml:"addr"`
	} `yaml:"server"`

	Client struct {
		TimeoutSeconds int `yaml:"timeout"`
	} `yaml:"client"`

	RabbitMQ struct {
		Addr string `yaml:"addr"`
	} `yaml:"rabbitmq"`
}

func GetConfigs() (*Configs, error) {
	configs := &Configs{}
	return configs, cleanenv.ReadConfig("config.yml", configs)
}
