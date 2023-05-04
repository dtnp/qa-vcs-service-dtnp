package requesting_source

import (
	"fmt"
	"strings"
	"testing"

	"go-webservices-automation/pkg/qaframework"

	"github.com/stretchr/testify/require"
)


func TestGETRequestingsourceAliasAlias400(t *testing.T) {
	method := "get"
	name := "TestGETRequestingsourceAliasAlias400"
	desc := fmt.Sprintf("%s method for %s", method, name)
	// This endpoint can pass in multiple tests, comma separated
	tests := []string{"postmanz"}

	qaframework.RunEndpointFunction(t, TS.config, desc, func() {
		req := require.New(t)
		ed, err := GetEndpointData(method, name)
		if err != nil {
			t.Error(err)
			return
		}

		e := ed.Endpoint
		for _, ep := range tests {
			ed.Endpoint = strings.Replace(e, "{alias}", ep, 1)
			res, err := qaframework.APICallByGET(TS.config, ed)
			if err != nil {
				return
			}

			t.Logf("endpoint: %s %s", ed.Method, ed.Endpoint)
			req.Equal(ed.StatusCode, res.StatusCode, "Status code mismatch")
			req.True(res.Data.Success)
			req.NotEmpty(res.Data.Data)
			req.GreaterOrEqual(res.Data.Timestamp, res.Timestamp)
			req.LessOrEqual(res.ResponseTime, int64(ed.MaxExecutionTime))
		}
	})
}

func TestGETRequestingsourceAliasAlias404(t *testing.T) {
	method := "get"
	name := "TestGETRequestingsourceAliasAlias404"
	desc := fmt.Sprintf("%s method for %s", method, name)
	// This endpoint can pass in multiple tests, comma separated
	tests := []string{"porkchops"}

	qaframework.RunEndpointFunction(t, TS.config, desc, func() {
		req := require.New(t)
		ed, err := GetEndpointData(method, name)
		if err != nil {
			t.Error(err)
			return
		}

		e := ed.Endpoint
		for _, ep := range tests {
			ed.Endpoint = strings.Replace(e, "{alias}", ep, 1)
			res, err := qaframework.APICallByGET(TS.config, ed)
			if err != nil {
				return
			}

			t.Logf("endpoint: %s %s", ed.Method, ed.Endpoint)
			req.Equal(ed.StatusCode, res.StatusCode, "Status code mismatch")
			req.False(res.Data.Success)
			req.Empty(res.Data.Data)
			req.GreaterOrEqual(res.Data.Timestamp, res.Timestamp)
			req.LessOrEqual(res.ResponseTime, int64(ed.MaxExecutionTime))
		}
	})
}

func TestGETRequestingsourceAliasAlias200v1(t *testing.T) {
	method := "get"
	name := "TestGETRequestingsourceAliasAlias200v1"
	desc := fmt.Sprintf("%s method for %s", method, name)
	// This endpoint can pass in multiple tests, comma separated
	tests := []string{"postman"}

	qaframework.RunEndpointFunction(t, TS.config, desc, func() {
		req := require.New(t)
		ed, err := GetEndpointData(method, name)
		if err != nil {
			t.Error(err)
			return
		}

		e := ed.Endpoint
		for _, ep := range tests {
			ed.Endpoint = strings.Replace(e, "{alias}", ep, 1)
			res, err := qaframework.APICallByGET(TS.config, ed)
			if err != nil {
				return
			}

			t.Logf("endpoint: %s %s", ed.Method, ed.Endpoint)
			req.Equal(ed.StatusCode, res.StatusCode, "Status code mismatch")
			req.True(res.Data.Success)
			req.NotEmpty(res.Data.Data)
			req.GreaterOrEqual(res.Data.Timestamp, res.Timestamp)
			req.LessOrEqual(res.ResponseTime, int64(ed.MaxExecutionTime))
		}
	})
}

func TestGETRequestingsourceAliasAlias200v2(t *testing.T) {
	ed := qaframework.EndpointData{
		Method:     	  "GET",
		Endpoint:   	  "/requesting-source/alias/{alias}",
		Version:    	  "v1",
		//URLParams:  	  "",
		StatusCode:       200,
		MaxExecutionTime: 300,
	}
	// This endpoint can pass in multiple tests, comma separated
	tests := []string{"postman"}
	desc := fmt.Sprintf("%s method for %s", ed.Method, ed.Endpoint)

	qaframework.RunEndpointFunction(t, TS.config, desc, func() {
		req := require.New(t)
		for _, ep := range tests {
			ed.Endpoint = strings.Replace(ed.Endpoint, "{alias}", ep, 1)
			res, err := qaframework.APICallByGET(TS.config, ed)
			if err != nil {
				return
			}

			t.Logf("endpoint: %s %s", ed.Method, ed.Endpoint)
			req.Equal(ed.StatusCode, res.StatusCode, "Status code mismatch")
			req.True(res.Data.Success)
			req.NotEmpty(res.Data.Data)
			req.GreaterOrEqual(res.Data.Timestamp, res.Timestamp)
			req.LessOrEqual(res.ResponseTime, int64(ed.MaxExecutionTime))
		}
	})
}

func TestGETRequestingsourceAliasAlias200v3(t *testing.T) {
	RunTest(t, "TestGETRequestingsourceAliasAlias200v3", func(e qaframework.EndpointData) {
		req := require.New(t)
		res, err := qaframework.APICallByGETDynamic(TS.config, e, map[string]string{"{alias}": "postman"})
		if err != nil {
			t.Error(err)
			return
		}

		t.Logf("method: %s, url: %s", e.Method, res.URL)
		req.Equal(200, res.StatusCode, "Status code mismatch")
		req.True(res.Data.Success)
		req.NotEmpty(res.Data.Data)
		req.GreaterOrEqual(res.Data.Timestamp, res.Timestamp)
		req.LessOrEqual(res.ResponseTime, int64(300))
	})
}

func TestGETRequestingsource200(t *testing.T) {
	method := "get"
	name := "TestGETRequestingsource200"
	desc := fmt.Sprintf("%s method for %s", method, name)

	qaframework.RunEndpointFunction(t, TS.config, desc, func() {
		req := require.New(t)
		ed, err := GetEndpointData(method, name)
		if err != nil {
			t.Error(err)
			return
		}

		res, err := qaframework.APICallByGET(TS.config, ed)
		if err != nil {
			t.Error(err)
			return
		}

		req.Equal(ed.StatusCode, res.StatusCode, "Status code mismatch")
		req.True(res.Data.Success)
		req.NotEmpty(res.Data.Data)
		req.GreaterOrEqual(res.Data.Timestamp, res.Timestamp)
		req.LessOrEqual(res.ResponseTime , int64(ed.MaxExecutionTime))
	})
}

func TestGETRequestingsourceId200(t *testing.T) {
	method := "get"
	name := "TestGETRequestingsourceId200"
	desc := fmt.Sprintf("%s method for %s", method, name)
	// This endpoint can pass in multiple tests, comma separated
	tests := []string{""}

	qaframework.RunEndpointFunction(t, TS.config, desc, func() {
		req := require.New(t)
		ed, err := GetEndpointData(method, name)
		if err != nil {
			t.Error(err)
			return
		}

		e := ed.Endpoint
		for _, ep := range tests {
			ed.Endpoint = strings.Replace(e, "{id}", ep, 1)
			res, err := qaframework.APICallByGET(TS.config, ed)
			if err != nil {
				return
			}

			t.Logf("endpoint: %s %s", ed.Method, ed.Endpoint)
			req.Equal(ed.StatusCode, res.StatusCode, "Status code mismatch")
			req.True(res.Data.Success)
			req.NotEmpty(res.Data.Data)
			req.GreaterOrEqual(res.Data.Timestamp, res.Timestamp)
			req.LessOrEqual(res.ResponseTime, int64(ed.MaxExecutionTime))
		}
	})
}

func TestGETRequestingsourceId400(t *testing.T) {
	method := "get"
	name := "TestGETRequestingsourceId400"
	desc := fmt.Sprintf("%s method for %s", method, name)
	// This endpoint can pass in multiple tests, comma separated
	tests := []string{""}

	qaframework.RunEndpointFunction(t, TS.config, desc, func() {
		req := require.New(t)
		ed, err := GetEndpointData(method, name)
		if err != nil {
			t.Error(err)
			return
		}

		e := ed.Endpoint
		for _, ep := range tests {
			ed.Endpoint = strings.Replace(e, "{id}", ep, 1)
			res, err := qaframework.APICallByGET(TS.config, ed)
			if err != nil {
				return
			}

			t.Logf("endpoint: %s %s", ed.Method, ed.Endpoint)
			req.Equal(ed.StatusCode, res.StatusCode, "Status code mismatch")
			req.True(res.Data.Success)
			req.NotEmpty(res.Data.Data)
			req.GreaterOrEqual(res.Data.Timestamp, res.Timestamp)
			req.LessOrEqual(res.ResponseTime, int64(ed.MaxExecutionTime))
		}
	})
}

func TestGETRequestingsourceId404(t *testing.T) {
	method := "get"
	name := "TestGETRequestingsourceId404"
	desc := fmt.Sprintf("%s method for %s", method, name)
	// This endpoint can pass in multiple tests, comma separated
	tests := []string{""}

	qaframework.RunEndpointFunction(t, TS.config, desc, func() {
		req := require.New(t)
		ed, err := GetEndpointData(method, name)
		if err != nil {
			t.Error(err)
			return
		}

		e := ed.Endpoint
		for _, ep := range tests {
			ed.Endpoint = strings.Replace(e, "{id}", ep, 1)
			res, err := qaframework.APICallByGET(TS.config, ed)
			if err != nil {
				return
			}

			t.Logf("endpoint: %s %s", ed.Method, ed.Endpoint)
			req.Equal(ed.StatusCode, res.StatusCode, "Status code mismatch")
			req.True(res.Data.Success)
			req.NotEmpty(res.Data.Data)
			req.GreaterOrEqual(res.Data.Timestamp, res.Timestamp)
			req.LessOrEqual(res.ResponseTime, int64(ed.MaxExecutionTime))
		}
	})
}
