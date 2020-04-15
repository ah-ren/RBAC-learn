package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type MetaConfig struct {
	RunMode     string `yaml:"run-mode"`
	Port        int    `yaml:"port"`
	LogPath     string `yaml:"log-path"`
	DefPageSize int32  `yaml:"def-page-size"`
	MaxPageSize int32  `yaml:"max-page-size"`
}

type JwtConfig struct {
	Secret        string `yaml:"secret"`
	Issuer        string `yaml:"issuer"`
	Expire        int64  `yaml:"expire"`
	RefreshExpire int64  `yaml:"refresh-expire"`
}

type CasbinConfig struct {
	ConfigPath string `yaml:"config-path"`
}

type MySqlConfig struct {
	Host     string `yaml:"host"`
	Port     int32  `yaml:"port"`
	Name     string `yaml:"name"`
	Charset  string `yaml:"charset"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	IsLog    bool   `yaml:"log"`
}

type ServerConfig struct {
	MetaConfig   *MetaConfig   `yaml:"meta"`
	JwtConfig    *JwtConfig    `yaml:"jwt"`
	CasbinConfig *CasbinConfig `yaml:"casbin"`
	MySqlConfig  *MySqlConfig  `yaml:"mysql"`
}

func Load(path string) (*ServerConfig, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	config := &ServerConfig{}
	err = yaml.Unmarshal(f, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
