package iot

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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

func (ya *YaClient) do(url, method string) ([]byte, error) {

	var bearer = "Bearer " + ya.config.Token

	req, err := http.NewRequest(method, url, nil)

	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	resp, err := ya.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s", resp.Status)
	}

	defer func(Body io.ReadCloser) {
		if err = Body.Close(); err != nil {
			panic(err)
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, fmt.Errorf("error while reading the response bytes: %s", err.Error())
	}

	return body, nil
}

func (ya *YaClient) GetInfo() (*IotInfoResponse, error) {
	var errMethodName = "[getInfo]"
	var url = ya.config.BaseUrl + "/v1.0/user/info"

	body, err := ya.do(url, "GET")
	if err != nil {
		return nil, err
	}

	var result IotInfoResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("%s error while unmarshal the response body: %s", errMethodName, err.Error())
	}

	return &result, nil
}

func (ya *YaClient) GetDevice(deviceID string) (*IotDeviceResponse, error) {
	var errMethodName = "[getDevices]"
	var url = fmt.Sprintf(ya.config.BaseUrl + "/v1.0/devices/" + deviceID)

	body, err := ya.do(url, "GET")
	if err != nil {
		return nil, fmt.Errorf("%s error: %s", errMethodName, err.Error())
	}

	var result IotDeviceResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("%s error while unmarshal the response body: %s", errMethodName, err.Error())
	}

	return &result, nil
}

func (ya *YaClient) GetGroup(groupID string) (*IotGroupResponse, error) {
	var errMethodName = "[getDevices]"
	var url = fmt.Sprintf(ya.config.BaseUrl + "/v1.0/groups/" + groupID)

	body, err := ya.do(url, "GET")
	if err != nil {
		return nil, fmt.Errorf("%s error: %s", errMethodName, err.Error())
	}

	var result IotGroupResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("%s error while unmarshal the response body: %s", errMethodName, err.Error())
	}

	return &result, nil
}

func (ya *YaClient) TriggerScenario(scenarioID string) (*IotResponseStatus, error) {
	var errMethodName = "[triggerScenario]"
	var url = fmt.Sprintf(ya.config.BaseUrl + "/v1.0/scenarios/" + scenarioID + "/actions")

	body, err := ya.do(url, "POST")
	if err != nil {
		return nil, fmt.Errorf("%s error: %s", errMethodName, err.Error())
	}

	var result IotResponseStatus
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("%s error while unmarshal the response body: %s", errMethodName, err.Error())
	}

	return &result, nil
}
