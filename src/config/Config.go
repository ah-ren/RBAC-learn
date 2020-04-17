package config

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type MetaConfig struct {
	RunMode     string `json:"run-mode"      yaml:"run-mode"`
	Port        int    `json:"port"          yaml:"port"`
	LogPath     string `json:"log-path"      yaml:"log-path"`
	DefPageSize int32  `json:"def-page-size" yaml:"def-page-size"`
	MaxPageSize int32  `json:"max-page-size" yaml:"max-page-size"`
}

type MySqlConfig struct {
	Host     string `json:"host"     yaml:"host"`
	Port     int32  `json:"port"     yaml:"port"`
	Name     string `json:"name"     yaml:"name"`
	Charset  string `json:"charset"  yaml:"charset"`
	User     string `json:"user"     yaml:"user"`
	Password string `json:"password" yaml:"password"`
	IsLog    bool   `json:"log"      yaml:"log"`
}

type RedisConfig struct {
	ConnType string `json:"conn-type" yaml:"conn-type"`
	Host     string `json:"host"      yaml:"host"`
	Port     int32  `json:"port"      yaml:"port"`
	Db       int32  `json:"db"        yaml:"db"`
	Password string `json:"password"  yaml:"password"`

	ConnectTimeout int32 `yaml:"connect-timeout"`
	ReadTimeout    int32 `yaml:"read-timeout"`
	WriteTimeout   int32 `yaml:"write-timeout"`
}

type JwtConfig struct {
	Secret        string            `json:"secret"         yaml:"secret"`
	Issuer        string            `json:"issuer"         yaml:"issuer"`
	Expire        int64             `json:"expire"         yaml:"expire"`
	RefreshExpire int64             `json:"refresh-expire" yaml:"refresh-expire"`
	RedisFmt      string            `json:"redis-fmt"      yaml:"redis-fmt"`
	FakeToken     map[string]uint32 `json:"fake"           yaml:"fake"`
}

type CasbinConfig struct {
	ConfigPath string `json:"config-path" yaml:"config-path"`
}

type ServerConfig struct {
	MetaConfig   *MetaConfig   `json:"meta"   yaml:"meta"`
	MySqlConfig  *MySqlConfig  `json:"mysql"  yaml:"mysql"`
	RedisConfig  *RedisConfig  `json:"redis"  yaml:"redis"`
	JwtConfig    *JwtConfig    `json:"jwt"    yaml:"jwt"`
	CasbinConfig *CasbinConfig `json:"casbin" yaml:"casbin"`
}

func Load(path string) (*ServerConfig, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	config := &ServerConfig{}

	ext := filepath.Ext(path)
	switch ext {
	case ".yml":
		fallthrough
	case ".yaml":
		err = yaml.Unmarshal(f, config)
	case ".json":
		err = json.Unmarshal(f, config)
	default:
		return nil, fmt.Errorf("Expected a yaml or json file, got a \"%s\" file\n", ext)
	}

	if err != nil {
		return nil, err
	}
	return config, nil
}
