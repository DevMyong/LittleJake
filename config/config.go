package config

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const FileName = "./config.json"

type Config struct {
	Token  string              `json:"token"`
	Prefix string              `json:"prefix"`
	Usages map[string][]string `json:"usages"`
}

type FlagPattern map[string][]string
type ArgPattern []string

var Token string

func Init() {
	loadEnv()
}

func init() {
	cfg, err := ParseConfigFromJSONFile(FileName)
	if err != nil {
		panic(err)
	}
}

type Config struct {
	Token  string              `json:"token"`
	Prefix string              `json:"prefix"`
	Usages map[string][]string `json:"usages"`
}
type FlagPattern map[string][]string
type ArgPattern []string

func ParseConfigFromJSONFile(fileName string) (c *Config, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer f.Close()

	c = new(Config)
	err = json.NewDecoder(f).Decode(c)

	return
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Failed to load .env file")
	}
	fmt.Println("Success to load .env file")
}
