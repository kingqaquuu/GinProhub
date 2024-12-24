package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Database struct {
		Username  string
		Password  string
		Host      string
		Port      string
		Name      string
		Charset   string
		ParseTime bool
		Loc       string
	}
	Server struct {
		ip   string
		port string
	}
	Redis struct {
		ip   string
		port string
		DB   int
	}
}

func LoadConfig() (*Config, error) {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %v", err)
	}
	var cfg Config
	if err := yaml.Unmarshal(yamlFile, &cfg); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %v", err)
	}
	return &cfg, nil
}
