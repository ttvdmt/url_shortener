package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App     AppConfig     `yaml:"app"`
	Storage StorageConfig `yaml:"storage"`
}

type AppConfig struct {
	Host           string `yaml:"host"`
	Port           string `yaml:"port"`
	TTL            string `yaml:"ttl_default"`
	CleanUp_Period string `yaml:"cleanup_period"`
}

type StorageConfig struct {
	SQLite     SQLiteConfig   `yaml:"sqlite"`
	PostgreSQL PostgresConfig `yaml:"postgresql"`
}

type SQLiteConfig struct {
	DBPath string `yaml:"database"`
}

type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

func Load(path string) (Config, error) {
	cfg := Config{}

	file, err := os.Open(path)
	if err != nil {
		return Config{}, fmt.Errorf("cant open config file: %w", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return Config{}, fmt.Errorf("cant decode config: %w", err)
	}

	return cfg, nil
}
