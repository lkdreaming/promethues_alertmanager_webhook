package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"text/template"
)

type Server struct {
	Port int `yaml:"port"`
}

var DingTalkTpl *template.Template

type Config struct {
	Server Server `yaml:"server"`
	Chat   Chat   `yaml:"chat"`
}

var AppConf Config

type Chat struct {
	DingTalk DingTalk `yaml:"dingTalk"`
}

type DingTalk struct {
	AccessToken string `yaml:"access_token"`
	Secret      string `yaml:"secret"`
	Template    string `yaml:"template"`
}

func BootStrapInit() error {
	content, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(content, &AppConf)
	if err != nil {
		return err
	}
	DingTalkTpl, err = template.New("dingTalk").Parse(AppConf.Chat.DingTalk.Template)
	return nil
}
