package requesting_source

import (
	"go-webservices-automation/pkg/qaframework"
)

// endpoints are the individual endpoints for this specific folder
func endpoints() map[string]qaframework.EndpointData {
	return map[string]qaframework.EndpointData {
		"get requestingsourcequery": {
			Method: "GET",
			Endpoint: "requesting-source",
			Version:   "v1",
			URLParams: "",
		},
		"get requestingsourcequerybyalias": {
			Method: "GET",
			Endpoint: "requesting-source/alias/{alias}",
			Version:   "v1",
			URLParams: "",
		},
		"get requestingsourcequerybyid": {
			Method: "GET",
			Endpoint: "requesting-source/{id}",
			Version:   "v1",
			URLParams: "",
		},
	}
}

