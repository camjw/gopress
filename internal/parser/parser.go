package parser

import (
	"encoding/json"
	"io/ioutil"
)

type Testcase struct {
	Testfile string   `json:"testfile"`
	Regexes  []string `json:"regexes"`
}

type GopressConfig struct {
	Directory  string     `json:"directory"`
	Extension  string     `json:"extension"`
	Basebranch string     `json:"basebranch"`
	Tests      []Testcase `json:"tests"`
}

func GetConfig() (GopressConfig, error) {
	configFile, err := ioutil.ReadFile("./gopress.json")
	if err != nil {
		return GopressConfig{}, err
	}
	var config GopressConfig
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		return GopressConfig{}, err
	}
	return config, nil
}
