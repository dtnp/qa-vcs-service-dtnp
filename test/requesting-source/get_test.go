package requesting_source

import (
	"fmt"
	"testing"

	"go-webservices-automation/pkg/qaframework"

	"github.com/stretchr/testify/require"
)

func TestGetRequestingSourceQuery(t *testing.T) {
	method := "GET"
	name := "RequestingSourceQuery"
	desc := fmt.Sprintf("%s method for %s", method, name)

	qaframework.RunEndpointFunction(t, TS.config, desc, func() {
		req := require.New(t)
		ed, err := GetEndpointData(method, name)
		if err != nil {
			t.Error(err)
		}

		res, err := qaframework.APICallByGET(TS.config, ed)
		if err != nil {
			t.Error(err)
		}

		req.Equal(200, res.StatusCode, "Status code mismatch")
		req.True(res.Data.Success)
		req.NotEmpty(res.Data.Data)
		req.GreaterOrEqual(res.Data.Timestamp, res.Timestamp)
		req.LessOrEqual(res.ResponseTime, int64(300))

		// Hardcoded as of April 04, 2023 based on total values in DB at the time
		req.GreaterOrEqual(len(res.Data.Data), 5)
	})
}