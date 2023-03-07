package configs

import "github.com/ilyakaznacheev/cleanenv"

type Configs struct {
	Server struct {
		Addr string `yaml:"addr"`
	} `yaml:"server"`

}

func GetConfigs() (*Configs, error) {
	configs := &Configs{}
	return configs, cleanenv.ReadConfig("config.yml", configs)
}
