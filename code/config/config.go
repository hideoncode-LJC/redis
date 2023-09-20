package config

import (
	"log"
	"github.com/spf13/viper"
)

type RedisConfig struct {
	IP string
	Port string
	Password string
	DB int
}

func GetRdb() *RedisConfig {
	viper.SetConfigFile("/code/go-redis/config/config.yaml")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalln(err)
	}

	return &RedisConfig{
		IP: viper.GetString("redis.ip"),
		Port: viper.GetString("redis.port"),
		Password: viper.GetString("redis.password"),
		DB: viper.GetInt("redis.db"),
	}
}