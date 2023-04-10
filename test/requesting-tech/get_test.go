package requesting_tech

import (
	"fmt"
	"strings"
	"testing"

	"go-webservices-automation/pkg/qaframework"

	"github.com/stretchr/testify/require"
)

func TestGetRequestingTechQuery(t *testing.T) {
	method := "GET"
	name := "RequestingTechQuery"
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

		req.Equal(200, res.StatusCode, "Status code mismatch")
		req.True(res.Data.Success)
		req.NotEmpty(res.Data.Data)
		req.GreaterOrEqual(res.Data.Timestamp, res.Timestamp)
		req.LessOrEqual(res.ResponseTime , int64(300))

		// Hardcoded as of April 04, 2023 based on total values in DB at the time
		req.GreaterOrEqual(len(res.Data.Data), 5)
	})
}

func TestGetRequestingTechQueryByAlias(t *testing.T) {
	method := "GET"
	name := "RequestingTechQueryByAlias"
	desc := fmt.Sprintf("%s method for %s", method, name)
	tests := []string{"api-rest", "graphql", "manual-import"}

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
			req.Equal(200, res.StatusCode, "Status code mismatch")
			req.True(res.Data.Success)
			req.NotEmpty(res.Data.Data)
			req.GreaterOrEqual(res.Data.Timestamp, res.Timestamp)
			req.LessOrEqual(res.ResponseTime, int64(300))
		}
	})
}

func TestGetRequestingTechQueryByID(t *testing.T) {
	method := "GET"
	name := "RequestingTechQueryByID"
	desc := fmt.Sprintf("%s method for %s", method, name)
	tests := []string{"1", "3", "5"}

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
				t.Error(err)
				return
			}

			t.Logf("endpoint: %s %s", ed.Method, ed.Endpoint)
			req.Equal(200, res.StatusCode, "Status code mismatch")
			req.True(res.Data.Success)
			req.NotEmpty(res.Data.Data)
			req.GreaterOrEqual(res.Data.Timestamp, res.Timestamp)
			req.LessOrEqual(res.ResponseTime, int64(300))
		}
	})
}