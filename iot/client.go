package iot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func (ya *YaClient) get(url string) ([]byte, error) {
	return ya.do(url, "GET", nil)
}

func (ya *YaClient) post(url string, body interface{}) ([]byte, error) {
	bBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return ya.do(url, "POST", bBody)
}

func (ya *YaClient) do(url, method string, reqBody []byte) ([]byte, error) {

	var bearer = "Bearer " + ya.config.Token

	var reqReader io.Reader = nil
	if reqBody != nil {
		reqReader = bytes.NewBuffer(reqBody)
	}

	req, err := http.NewRequest(method, url, reqReader)
	if err != nil {
		return nil, fmt.Errorf("request build err: %s", err)
	}

	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")

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
