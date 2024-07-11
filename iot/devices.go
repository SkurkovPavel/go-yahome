package iot

import (
	"encoding/json"
	"errors"
	"fmt"
)

func (ya *YaClient) GetDevice(deviceID string) (*IotDeviceResponse, error) {
	var errMethodName = "[getDevices]"

	if deviceID == "" {
		return nil, ErrorReturn(errMethodName, errors.New("deviceID not provided"))
	}

	body, err := ya.get(fmt.Sprintf(ya.config.IotUrl + "/v1.0/devices/" + deviceID))
	if err != nil {
		return nil, ErrorReturn(errMethodName, err)
	}

	var result IotDeviceResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("%s error while unmarshal the response body: %s", errMethodName, err.Error())
	}

	return &result, nil
}

func (ya *YaClient) QueryDevices(devices *[]Device) (*IotResponseStatus, error) {
	var errMethodName = "[unlinkAccount]"

	body, err := ya.post(fmt.Sprintf(ya.config.IotUrl+"/v1.0/user/devices/query"), devices)
	if err != nil {
		return nil, ErrorReturn(errMethodName, err)
	}

	var result IotResponseStatus
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("%s error while unmarshal the response body: %s", errMethodName, err.Error())
	}

	return &result, nil
}

func (ya *YaClient) SetActionToDevice(device Device) (*[]Device, error) {
	var errMethodName = "[setActionToDevice]"

	if err := device.Validate(); err != nil {
		return nil, ErrorReturn(errMethodName, err)
	}

	body, err := ya.post(fmt.Sprintf(ya.config.IotUrl+"/v1.0/devices/actions"), nil)
	if err != nil {
		return nil, ErrorReturn(errMethodName, err)
	}

	var result IotInfoResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("%s error while unmarshal the response body: %s", errMethodName, err.Error())
	}

	return &result.Devices, nil
}
