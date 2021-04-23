package config

import (
	"encoding/json"
	"io/ioutil"
	"sync"
)

type Config struct {
	Application string `json:"application"`
	Description string `json:"description"`
	Version     string `json:"version"`
	DevMode     bool   `json:"dev_mode"`

	Postgres struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Dbname   string `json:"dbname"`
	} `json:"postgres"`

	Nexmo struct {
		Key    string `json:"key"`
		Secret string `json:"secret"`
	} `json:"nexmo"`

	Email struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     int    `json:"port"`
	} `json:"email"`

	Sentry struct {
		Dsn string `json:"dsn"`
	} `json:"sentry"`
}

var instance *Config

var once sync.Once

func GetInstance() *Config {
	once.Do(func() {
		file, _ := ioutil.ReadFile("./config/config.json")
		var config *Config
		_ = json.Unmarshal(file, config)
		instance = config
	})

	return instance
}
