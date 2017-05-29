package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/Unknwon/goconfig"
)

const (
	// defaultSection Default section name
	defaultSection = "TestEnv"
)

const configPath = "/src/brush/core/config/config.ini"

var (
	conf *Config
)

func GetConfig() *Config {
	if conf == nil {
		absPath, err := filepath.Abs(os.Getenv("GOPATH") + configPath)
		if err != nil {
			log.Fatal(err)
		}

		conf = NewConfig(absPath)
	}

	return conf
}

type Config struct {
	configFile *goconfig.ConfigFile
	section    string
}

func NewConfig(configDir string) *Config {
	var err error
	conf := &Config{}
	conf.section = defaultSection
	conf.configFile, err = goconfig.LoadConfigFile(configDir)
	if err != nil {
		log.Fatal(err)
	}
	return conf
}

func (c *Config) GetValue(key string) string {
	value, err := c.configFile.GetValue(c.section, key)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

func (c *Config) GetInt64(key string) int64 {
	value, err := c.configFile.Int64(c.section, key)
	if err != nil {
		log.Fatal(err)
	}
	return value
}
