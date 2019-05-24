package parser

import (
	"encoding/json"
	"fmt"
	"io"
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

func GetConfig(r io.Reader) (Config, error) {
	var config Config
	err := json.NewDecoder(r).Decode(&config)
	if err != nil {
		fmt.Println(err)
		return Config{}, err
	}
	config.cleanDirectory()
	return config, nil
}

func (c *Config) GetFilePath(t Testcase) string {
	return c.Directory + t.Testfile + c.Extension
}

func (c *Config) cleanDirectory() {
	trailing := c.Directory[len(c.Directory)-1:]
	if trailing != "/" {
		c.Directory = c.Directory + "/"
	}
}
