package qaframework

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"go-webservices-automation/pkg/config"
)

/*
SendRequest
This is common method to hit specific https request
It accepts SendRequestParams structure with parameters like (Url, Method, Header, Payload)
It returns response status code in int format and response in string format
In case of response error it raises an error
*/
func SendRequest(params SendRequestParams) (int, *http.Response, error) {
	// TODO pass the configs to set these
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    10 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

	payload := strings.NewReader(fmt.Sprintf("%v", params.Payload))
	req, err := http.NewRequest(params.Method, params.Url, payload)
	if err != nil {
		//log.Error("Request failed with error : ", err)
		return 0, nil, nil
	}

	// TODO pass the configs to set these
	// TODO Dynamic generation?
	if params.Header != "" {
		req.Header.Add("Connection", "keep-alive")
		req.Header.Add("content-type", "application/json")
		req.Header.Add("tracer_uid", "35c74d78-6798-44fe-8694-9aac5b9404d6")
		req.Header.Add("user_uid", "5e28b8df-466f-4ebc-8d71-f6f79791a967")
		req.Header.Add("site_uid", "29994380-3ffe-4414-bf92-60de32bd6d7f")
		req.Header.Add("org_uid", "cd3117ac-605a-4a4e-9c30-ad473b95f786")
	}
	// Send Request
	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("\n\n%w\n\n", err)
		return response.StatusCode, response, err
	}

	// Return statusCode and response body
	return response.StatusCode, response, nil
}


func APICallByGET(conf config.Config, ed EndpointData) (APIResponse, error) {
	v := conf.Api.DefaultVersion
	if ed.Version != "" {
		v = ed.Version
	}
	url := fmt.Sprintf(
		"http://%s:%s/%s/%s",
			conf.Api.Host,
			conf.Api.Port,
			v,
			ed.Endpoint,
		)
	if ed.URLParams != "" {
		url = fmt.Sprintf("%s%s", url, ed.URLParams)
	}

	srp := SendRequestParams{Url: url, Method: ed.Method}
	ts := time.Now()
	statusCode, response, err := SendRequest(srp)
	elapsed := time.Since(ts)
	if err != nil {
		return APIResponse{}, err
	}

	resp, err := io.ReadAll(response.Body)
	if err != nil {
		return APIResponse{}, nil
	}

	var responseObject VcsResponse
	if err := json.Unmarshal(resp, &responseObject); err != nil {
		return APIResponse{}, nil
	}

	return APIResponse{
		EndpointData: ed,
		Data:         responseObject,
		StatusCode:   statusCode,
		Timestamp:    int(time.Now().Unix()),
		ResponseTime: elapsed.Milliseconds(),
	}, nil
}


