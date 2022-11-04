package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sync"
)

var (
	conf       *config
	onceConfig sync.Once
)

type config struct {
	ServerParam struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	DbParams struct {
		Port     string `yaml:"port"`
		Host     string `yaml:"host"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Db       string `yaml:"db"`
	} `yaml:"db_params"`
	CacheDbParams struct {
		Port     string `yaml:"port"`
		Host     string `yaml:"host"`
		Password string `yaml:"password"`
		Db       string `yaml:"db"`
	} `yaml:"cache_db_params"`
}

func GetConfig() *config {
	onceConfig.Do(func() {
		conf = &config{}
		data, err := getFile()
		if err != nil {
			panic(err)
		}

		if err := yaml.Unmarshal(data, conf); err != nil {
			panic(err)
		}
	})

	return conf
}

func getFile() ([]byte, error) {
	return ioutil.ReadFile("./config/config.yml")
}
