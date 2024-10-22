package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Username     string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DatabaseName string `mapstructure:"dbname"`
}

type JWTConfig struct {
	Secret   string        `mapstructure:"secret"`
	Lifetime time.Duration `mapstructure:"lifetime"`
}

type S3Config struct {
	AccessKey string `mapstructure:"access_key"`
	SecretKey string `mapstructure:"secret_key"`
	Endpoint  string `mapstructure:"endpoint"`
	Bucket    string `mapstructure:"bucket"`
	Region    string `mapstructure:"region"`
}

type Config struct {
	Database DatabaseConfig `mapstructure:"database"`
	JWT      JWTConfig      `mapstructure:"jwt"`
	S3       S3Config       `mapstructure:"s3"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("toml")

	var config Config

	if err := viper.ReadInConfig(); err != nil {
		return &config, fmt.Errorf("error reading config file: %v", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return &config, fmt.Errorf("unable to decode into struct: %v", err)
	}

	return &config, nil
}

func (config *Config) GetDatabaseURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.DatabaseName)
}

func (config *Config) GetS3Config() *S3Config {
	return &config.S3
}

func (config *Config) GetSecretKey() string {
	return config.JWT.Secret
}
