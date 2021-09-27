package config

import (
	"encoding/json"
	"os"
)

const FileName = "../../../config/config.json"

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
