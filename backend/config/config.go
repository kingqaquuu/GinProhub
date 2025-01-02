package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
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
	configPath:="../config/config.yaml"
	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %v", err)
	}
	var cfg Config
	if err := yaml.Unmarshal(yamlFile, &cfg); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %v", err)
	}
	return &cfg, nil
}

func (cfg *Config) DSN() string {
    return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
        cfg.Database.Username,
        cfg.Database.Password,
        cfg.Database.Host,
        cfg.Database.Port,
        cfg.Database.Name,
        cfg.Database.Charset,
        //cfg.Database.ParseTime,
        cfg.Database.Loc)
}
