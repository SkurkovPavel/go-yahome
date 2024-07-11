package iot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_New(t *testing.T) {

	act := NewIotClient(&Config{})
	exp := configureNewIotClient(&Config{}, &http.Client{})

	assert.Equal(t, exp, act)
}

func TestService_Info(t *testing.T) {

	jsonStringIfo := `{"status": "ok","request_id": "retest","rooms": [],"groups": [],"devices": [],"scenarios": [],"households": []}`

	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, jsonStringIfo) }))
	defer ts.Close()

	cl := configureNewIotClient(&Config{IotUrl: ts.URL}, ts.Client())

	t.Run("TEST DO", func(t *testing.T) {
		body, err := cl.do(cl.config.IotUrl, "GET", nil)
		require.NoError(t, err)
		assert.Equal(t, []byte(jsonStringIfo+"\n"), body)
	})

	t.Run("TEST GetInfo", func(t *testing.T) {
		res, err := cl.GetInfo()
		require.NoError(t, err)

		var testDataInfo IotInfoResponse
		_ = json.Unmarshal([]byte(jsonStringIfo), &testDataInfo)

		assert.Equal(t, &testDataInfo, res)
	})

}

func TestService_Device(t *testing.T) {

	jsonTestString := `{"status":"ok","request_id":"test","id":"test","name":"Лампочка 2","aliases":[],"type":"devices.types.light","external_id":"test","skill_id":"Test","state":"online","groups":["test"],"room":"test","capabilities":[{"retrievable":true,"type":"devices.capabilities.color_setting","parameters":{"color_model":"hsv","temperature_k":{"min":2700,"max":6500}},"state":{"instance":"temperature_k","value":5600},"last_updated":1637944661.8522422}],"properties":[]}`

	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, jsonTestString) }))
	defer ts.Close()

	cl := configureNewIotClient(&Config{IotUrl: ts.URL}, ts.Client())

	t.Run("TEST GetDevice", func(t *testing.T) {

		res, err := cl.GetDevice("test")
		require.NoError(t, err)

		var testData IotDeviceResponse
		_ = json.Unmarshal([]byte(jsonTestString), &testData)

		assert.Equal(t, &testData, res)
	})

}

func TestService_Group(t *testing.T) {

	jsonTestString := `{"status":"ok","request_id":"test","id":"test","name":"Люстра","aliases":[],"type":"devices.types.light","state":"split","capabilities":[{"retrievable":true,"type":"devices.capabilities.color_setting","parameters":{"color_model":"hsv","temperature_k":{"min":2700,"max":6500}},"state":{"instance":"temperature_k","value":5600}}],"devices":[{"id":"test","name":"Лампочка 2","type":"devices.types.light"}]}`

	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, jsonTestString) }))
	defer ts.Close()

	cl := configureNewIotClient(&Config{IotUrl: ts.URL}, ts.Client())

	t.Run("TEST GetGroup", func(t *testing.T) {

		res, err := cl.GetGroup("test")
		require.NoError(t, err)

		var testData IotGroupResponse
		_ = json.Unmarshal([]byte(jsonTestString), &testData)

		assert.Equal(t, &testData, res)
	})

}

func TestService_TriggerScenario(t *testing.T) {

	jsonTestString := `{"status":"ok","request_id":"test"}`

	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, jsonTestString) }))
	defer ts.Close()

	cl := configureNewIotClient(&Config{IotUrl: ts.URL}, ts.Client())

	t.Run("TEST TriggerScenario", func(t *testing.T) {

		res, err := cl.TriggerScenario("test")
		require.NoError(t, err)

		var testData IotResponseStatus
		_ = json.Unmarshal([]byte(jsonTestString), &testData)

		assert.Equal(t, &testData, res)
	})

}

func TestService_Errors(t *testing.T) {

	jsonTestString := `{"status":1,"request_id":"test"}`

	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, jsonTestString) }))
	defer ts.Close()

	cl := configureNewIotClient(&Config{IotUrl: ts.URL}, ts.Client())

	t.Run("TEST GetInfo ERROR", func(t *testing.T) {
		_, err := cl.GetInfo()
		require.Error(t, err)
	})
	t.Run("TEST TriggerScenario ERROR", func(t *testing.T) {
		_, err := cl.GetGroup("test")
		require.Error(t, err)
	})
	t.Run("TEST TriggerScenario EMPTY", func(t *testing.T) {
		_, err := cl.GetGroup("")
		require.Error(t, err)
	})
	t.Run("TEST GetDevice ERROR", func(t *testing.T) {
		_, err := cl.GetDevice("test")
		require.Error(t, err)
	})
	t.Run("TEST GetDevice EMPTY", func(t *testing.T) {
		_, err := cl.GetDevice("")
		require.Error(t, err)
	})
	t.Run("TEST TriggerScenario ERROR", func(t *testing.T) {
		_, err := cl.TriggerScenario("test")
		require.Error(t, err)
	})
	t.Run("TEST TriggerScenario EMPTY", func(t *testing.T) {
		_, err := cl.TriggerScenario("")
		require.Error(t, err)
	})
}

func TestService_Unauthorised(t *testing.T) {
	t.Parallel()
	cl := NewIotClient(NewConfig())

	t.Run("TEST GetInfo 401", func(t *testing.T) {
		_, err := cl.get(cl.config.IotUrl + "/v1.0/user/info")
		require.Error(t, err)
		assert.Equal(t, err.Error(), "401 Unauthorized")
	})
	t.Run("TEST GetInfo 401", func(t *testing.T) {
		_, err := cl.post(cl.config.IotUrl+"/v1.0/user/unlink", nil)
		require.Error(t, err)
		assert.Equal(t, err.Error(), "404 Not Found")
	})
}
