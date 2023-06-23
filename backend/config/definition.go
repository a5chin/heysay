package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	DB       `yaml:"db"`
	Port     uint   `yaml:"port"`
	Hostname string `yaml:"hostname"`
}

type DB struct {
	Hostname string `yaml:"hostname"`
	Port     uint   `yaml:"port"`
	Name     string `yaml:"name"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Cors struct {
	Allow struct {
		Origin string `yaml:"origin"`
	} `yaml:"allow"`
}

func Load() *Config {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	v.AddConfigPath("./config")

	v.AutomaticEnv()

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := v.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}

	var conf *Config

	err = v.Unmarshal(&conf)
	if err != nil {
		log.Fatalln(err)
	}

	return conf
}
