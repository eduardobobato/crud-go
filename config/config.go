package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Config : config interface
type Config interface {
	Read() (string, string, string)
}

type conf struct {
	ServerURI  string
	Database   string
	Collection string
}

// NewConfig : Return a new config
func NewConfig() Config {
	return &conf{}
}

func (c *conf) Read() (string, string, string) {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
	return c.ServerURI, c.Database, c.Collection
}
