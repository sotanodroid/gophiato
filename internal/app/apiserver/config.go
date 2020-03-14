package apiserver

import "github.com/sotanodroid/gophiato/internal/app/store"

// Config holds structure for service config
type Config struct {
	Bindport string `toml:"bind_port"`
	Loglevel string `toml:"log_level"`
	Store    *store.Config
}

// NewConfig return new instance of service config
func NewConfig() *Config {
	return &Config{
		Bindport: "8000",
		Loglevel: "INFO",
		Store:    store.NewConfig(),
	}
}
