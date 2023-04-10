package requesting_tech

import (
	"go-webservices-automation/pkg/qaframework"
)

// endpoints are the individual endpoints for this specific folder
func endpoints() map[string]qaframework.EndpointData {
	return map[string]qaframework.EndpointData {
		"get requestingtechquery": {
			Method: "GET",
			Endpoint: "requesting-tech",
			Version:   "v1",
			URLParams: "",
		},
		"get requestingtechquerybyalias": {
			Method: "GET",
			Endpoint: "requesting-tech/alias/{alias}",
			Version:   "v1",
			URLParams: "",
		},
		"get requestingtechquerybyid": {
			Method: "GET",
			Endpoint: "requesting-tech/{id}",
			Version:   "v1",
			URLParams: "",
		},
	}
}

