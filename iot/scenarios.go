package iot

import (
	"encoding/json"
	"errors"
	"fmt"
)

func (ya *YaClient) TriggerScenario(scenarioID string) (*IotResponseStatus, error) {
	var errMethodName = "[triggerScenario]"

	if scenarioID == "" {
		return nil, ErrorReturn(errMethodName, errors.New(" scenarioID not provided"))
	}

	body, err := ya.post(fmt.Sprintf(ya.config.IotUrl+"/v1.0/scenarios/"+scenarioID+"/actions"), nil)
	if err != nil {
		return nil, ErrorReturn(errMethodName, err)
	}

	var result IotResponseStatus
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, ErrorReturn(errMethodName, err)
	}

	return &result, nil
}
