package parser

import (
	"encoding/json"
	"io/ioutil"
)

type Testcase struct {
	Testfile string   `json:"testfile"`
	Regexes  []string `json:"regexes"`
}

type Config struct {
	Directory  string     `json:"directory"`
	Extension  string     `json:"extension"`
	Basebranch string     `json:"basebranch"`
	Tests      []Testcase `json:"tests"`
}

func GetConfig() (Config, error) {
	configFile, err := ioutil.ReadFile("./gopress.json")
	if err != nil {
		return Config{}, err
	}
	var config Config
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

func (c *Config) GetFilePath(t Testcase) string {
	return c.Directory + t.Testfile + c.Extension
}
