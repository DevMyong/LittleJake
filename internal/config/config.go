package config

import (
	"Jake/internal/logger"
	"encoding/json"
	"github.com/joho/godotenv"
	"os"
)

var C *config

func init() {
	logger.L.Info("Start to set config...")
	C = _new()
	C.Init()
}

type config struct {
	Token         string `json:"-"`
	CommandPrefix string `json:"prefix"`
}

func _new() *config {
	c := new(config)

	return c
}

func (c *config) Init() {
	c.SetSecrets()
	c.SetVariables()
}

func (c *config) SetSecrets() {
	if err := godotenv.Load(); err != nil {
		logger.L.Fatal("Failed to load .env file. config doesn't exists.")
		return
	}

	c.Token = os.Getenv("TOKEN")
	//todo: add DB url, port...

	logger.L.Info("Success to load config")
}

func (c *config) SetVariables() {
	f, err := os.Open("config.json")
	if err != nil {
		logger.L.Fatal(err)
		return
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(c)
	if err != nil {
		logger.L.Fatal(err)
		return
	}

	//todo: add invokes...
}
