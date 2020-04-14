package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type MetaConfig struct {
	RunMode string `yaml:"run-mode"`
	Port    int    `yaml:"port"`
	LogPath string `yaml:"log-path"`
}

type JwtConfig struct {
	Secret        string `yaml:"secret"`
	Issuer        string `yaml:"issuer"`
	Expire        int64  `yaml:"expire"`
	RefreshExpire int64  `yaml:"refresh-expire"`
}

type Config struct {
	MetaConfig *MetaConfig `yaml:"meta"`
	JwtConfig  *JwtConfig  `yaml:"jwt"`
}

func Load(path string) (*Config, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	err = yaml.Unmarshal(f, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
