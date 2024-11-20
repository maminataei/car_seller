package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Server
	Postgres
	Redis
}

type Server struct {
	Port    string
	RunMode string
}

type Postgres struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	SSLMode  string
}

type Redis struct {
	Host               string
	Port               string
	Password           string
	Db                 string
	MinIdleConnections int
	PoolSize           int
	PoolTimeout        int
}

func GetConfig() *Config {
	v, err := loadConfig(getConfigPath(os.Getenv("env")), "yaml")
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := parseConfig(v)
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}

func loadConfig(fileName string, fileType string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType(fileType)
	v.SetConfigName(fileName)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return v, nil
}

func parseConfig(v *viper.Viper) (*Config, error) {
	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Printf("Unable to parse config: %v", err)
		return nil, err
	}
	return &cfg, nil
}

func getConfigPath(env string) string {
	if env == "docker" {
		return "config/docker"
	} else if env == "production" {
		return "config/config-production"
	} else {
		return "config/dev"
	}
}
