package iot

import (
	"encoding/json"
	"fmt"
)

// UnlinkAccount Use carefully. If the user has disabled accounts,
// the user's token is canceled regardless of the correct response to the request received.
func (ya *YaClient) UnlinkAccount() (*IotResponseStatus, error) {
	var errMethodName = "[unlinkAccount]"

	body, err := ya.post(fmt.Sprintf(ya.config.IotUrl+"/v1.0/user/unlink"), nil)
	if err != nil {
		return nil, ErrorReturn(errMethodName, err)
	}

	var result IotResponseStatus
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("%s error while unmarshal the response body: %s", errMethodName, err.Error())
	}

	return &result, nil
}

func (ya *YaClient) GetInfo() (*IotInfoResponse, error) {
	var errMethodName = "[getInfo]"

	body, err := ya.get(ya.config.IotUrl + "/v1.0/user/info")
	if err != nil {
		return nil, err
	}

	var result IotInfoResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("%s error while unmarshal the response body: %s", errMethodName, err.Error())
	}

	return &result, nil
}
