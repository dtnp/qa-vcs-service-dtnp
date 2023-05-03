package qaframework

import "time"

type EndpointData struct {
	Method           string `json:"method"`
	Endpoint         string `json:"endpoint"`
	Version          string `json:"version"`
	URLParams        string `json:"appendAPIendpoint"`
	StatusCode       int    `json:"status_code"`
	MaxExecutionTime int    `json:"max_execution_time"`
}

type APIDetails struct {
	Method   string
	Endpoint string
	Filename string
	URL      string
}

type APIResponse struct {
	EndpointData EndpointData
	Data         VcsResponse
	StatusCode   int
	Timestamp    int
	ResponseTime int64
}

type SendRequestParams struct {
	Url     string
	Method  string
	Header  string
	Payload map[string]interface{}
}

type VcsResponse struct {
	Success   bool `json:"success"`
	Timestamp int  `json:"timestamp"`
	Data      []struct {
		ID          string    `json:"id"`
		Alias       string    `json:"alias"`
		Label       string    `json:"label"`
		Description string    `json:"description"`
		CreatedOn   time.Time `json:"created_on"`
		UpdatedOn   time.Time `json:"updated_on"`
	} `json:"data"`
}

type VCSResults struct {
	Success   bool `json:"success"`
	Timestamp int  `json:"timestamp"`
	Data      []struct {
		ID      string `json:"id"`
		VcsID   string `json:"vcs_id"`
		Message string `json:"message"`
		Results struct {
			Msg string `json:"msg"`
		} `json:"results"`
		CreatedOn time.Time `json:"created_on"`
	} `json:"data"`
}
