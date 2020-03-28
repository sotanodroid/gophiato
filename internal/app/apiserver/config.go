package apiserver

// Config holds structure for service config
type Config struct {
	Bindport    string `toml:"bind_port"`
	Loglevel    string `toml:"log_level"`
	DataBaseURL string `toml:"database_url"`
	SessionKey  string `toml:"session_key"`
}

// NewConfig return new instance of service config
func NewConfig() *Config {
	return &Config{
		Bindport: "8000",
		Loglevel: "INFO",
	}
}
