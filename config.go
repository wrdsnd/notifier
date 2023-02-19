package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Telegram struct {
	BotToken string `yaml:"botToken"`
	ChatID   string `yaml:"chatId"`
}

type Config struct {
	Telegram    Telegram
	Directories []string
	Filter      string
}

func ReadConfig() (*Config, error) {
	file, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	var config Config

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
