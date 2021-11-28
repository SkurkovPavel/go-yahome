package iot

import (
	"encoding/json"
	"errors"
	"fmt"
)

func (ya *YaClient) GetGroup(groupID string) (*IotGroupResponse, error) {
	var errMethodName = "[getDevices]"

	if groupID == "" {
		return nil, ErrorReturn(errMethodName, errors.New(" groupID not provided"))
	}

	body, err := ya.get(fmt.Sprintf(ya.config.IotUrl + "/v1.0/groups/" + groupID))
	if err != nil {
		return nil, ErrorReturn(errMethodName, err)
	}

	var result IotGroupResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, ErrorReturn(errMethodName, err)
	}

	return &result, nil
}

func (ya *YaClient) SetActionToGroup(group Group) (*IotGroupResponse, error) {
	var errMethodName = "[setActionToGroup]"

	if err := group.Validate(); err != nil {
		return nil, ErrorReturn(errMethodName, err)
	}

	body, err := ya.post(fmt.Sprintf(ya.config.IotUrl+"/v1.0/groups/"+group.Id+"/actions"), group)
	if err != nil {
		return nil, ErrorReturn(errMethodName, err)
	}

	var result IotGroupResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, ErrorReturn(errMethodName, err)
	}

	return &result, nil
}
