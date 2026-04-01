package config

import (
	"flag"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type HTTPServer struct {
	Addr string `yaml:"addr"`
}

type Config struct {
	Env        string     `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	Storage    string     `yaml:"storage"`
	HTTPServer HTTPServer `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string

	// 1. Get from ENV
	configPath = os.Getenv("CONFIG_PATH")

	// 2. If not in ENV, get from flag
	if configPath == "" {
		flags := flag.String("config", "", "path to the configuration file")
		flag.Parse()
		configPath = *flags

		if configPath == "" {
			log.Fatal("config path not set")
		}
	}

	// 3. Read file
	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}

	// 4. Parse YAML
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}

	return &cfg
}
