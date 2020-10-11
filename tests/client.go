package test

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func Request(
	baseUrl, 
	method string, 
	inputBody io.Reader, 
	headers map[string]string,
	uris ...string) (*http.Response, map[string]interface{}, error) {

	url := baseUrl
	for _, uri := range uris {
		url = fmt.Sprintf("%s%s", url, uri)
	}

	client := &http.Client{}

	req, err := http.NewRequest(method, url, inputBody)
	if err != nil {
		return nil, nil, err
	}

	if len(headers) > 0 {
		for headerName, headerValue := range headers {
			req.Header.Add(headerName, headerValue)
		}
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}

	defer res.Body.Close()
	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	var respM interface{}
	if err := json.Unmarshal(respBody, &respM); err != nil {
		return nil, nil, err
	}

	var resClone http.Response
	resClone = *res
	if resClone.StatusCode != 200 {
		return &resClone, nil, errors.New("status not 201")
	}
	return &resClone, respM.(map[string]interface{}), nil
}

func Health(baseUrl string) (*http.Response, map[string]interface{}, error) {
	return Request(baseUrl, "GET", nil, nil, "/health")
}

