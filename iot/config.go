package iot

type Config struct {
	Token  string `toml:"token"`
	IotUrl string `toml:"base_url"`
}

func NewConfig() *Config {
	return &Config{
		IotUrl: "https://api.iot.yandex.net",
	}
}
