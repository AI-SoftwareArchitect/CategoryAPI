package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Environment string          `mapstructure:"environment"`
	Cassandra   CassandraConfig `mapstructure:"cassandra"`
	Server      ServerConfig    `mapstructure:"server"`
}

type CassandraConfig struct {
	Hosts    []string `mapstructure:"hosts"`
	Keyspace string   `mapstructure:"keyspace"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

var AppConfig Config

// LoadConfig reads configuration from file or environment variables.
func LoadConfig() {
	// Set the file name and path
	viper.SetConfigName("config") // config.yaml
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs") // ./configs/config.yaml
	viper.AddConfigPath(".")         // Fallback to current directory

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// Unmarshal the config into AppConfig
	if err := viper.Unmarshal(&AppConfig); err != nil {
		panic(fmt.Errorf("unable to decode into struct: %w", err))
	}

	// Optionally, read environment variables
	viper.AutomaticEnv()
}
