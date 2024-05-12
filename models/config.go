package models

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	S3 struct {
		Bucket    string `yaml:"bucket"`
		Region    string `yaml:"region"`
		SecretKey string `yaml:"secret_key"`
		AccessKey string `yaml:"access_key"`
	}

	Dump struct {
		Command string `yaml:"command"`
	}

	SMTP struct {
		User              string `yaml:"user"`
		Password          string `yaml:"password"`
		Host              string `yaml:"host"`
		Secure            bool   `yaml:"secure"`
		Port              int    `yaml:"port"`
		Email             string `yaml:"email"`
		NotificationEmail string `yaml:"notification_email"`
	}
}

func LoadConfig() *Config {
	config := &Config{}
	configPath := os.Getenv("MYSQL_BACKUP_CONFIG_FILE")
	if configPath == "" {
		userConfigDir, err := os.UserConfigDir()

		if err != nil {
			panic("Error loading config file: " + err.Error())
		}

		configPath = userConfigDir + "/mysql-to-s3-backup/config.yaml"
	}

	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		panic("Error loading config file: " + err.Error())
	}

	err = yaml.Unmarshal(yamlFile, config)

	if err != nil {
		panic("Error loading config file: " + err.Error())
	}

	return config
}
