package iot

import (
	"net/http"
)

type YaClient struct {
	config     *Config
	httpClient *http.Client
}

func NewIotClient(config *Config) *YaClient {
	return configureNewIotClient(config, &http.Client{})
}

func configureNewIotClient(config *Config, transport *http.Client) *YaClient {
	return &YaClient{
		config:     config,
		httpClient: transport,
	}
}
