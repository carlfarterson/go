package main

import (
	"io/ioutil"
	"log"

	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Host      string `yaml:"Host"`
	AccessKey string `yaml:"AccessKey"`
	SecretKey string `yaml:"SecretKey"`
}

func (c *Config) getConfig() *Config {

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func main() {

	var config Config
	config.getConfig()

	client := new(client.CommonClient).Init(config.Host)
	resp, err := client.GetTimestamp()

	if err != nil {
		applogger.Error("Get timestamp error: %s", err)
	} else {
		applogger.Info("Get timestamp: %d", resp)
	}
}
