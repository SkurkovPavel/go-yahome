package iot

type Config struct {
	Token   string `toml:"token"`
	BaseUrl string `toml:"base_url"`
}

func NewConfig() *Config {
	return &Config{
		BaseUrl: "https://api.iot.yandex.net",
	}
}
