package src

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Conductor struct {
		Bot struct {
			Token    string `yaml:"token"`
			InvitUrl string `yaml:"invitUrl"`
		} `yaml:"bot"`
		Server struct {
			Port string `yaml:"port"`
		} `yaml:"server"`
	} `yaml:"conductor"`
}

func (c *Config) ReadConf() error {
	yamlFile, err := os.ReadFile("./config.yaml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return err
	}
	return nil
}
