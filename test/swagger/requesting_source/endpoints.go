
package requesting_source

import (
	"go-webservices-automation/pkg/qaframework"
)

// EndpointData are the individual endpoints for this specific folder
func EndpointData() map[string]qaframework.EndpointData {
	return map[string]qaframework.EndpointData {

		"TestGETRequestingsourceAliasAlias400": {
			Method:     	  "GET",
			Endpoint:   	  "/requesting-source/alias/{alias}",
			Version:    	  "v1",
			URLParams:  	  "",
			StatusCode:       400,
			MaxExecutionTime: 300,
		},
"TestGETRequestingsourceAliasAlias404": {
		Method:     	  "GET",
		Endpoint:   	  "/requesting-source/alias/{alias}",
		Version:    	  "v1",
		URLParams:  	  "",
		StatusCode:       404,
		MaxExecutionTime: 300,
	},
		"TestGETRequestingsourceAliasAlias200": {
			Method:     	  "GET",
			Endpoint:   	  "/requesting-source/alias/{alias}",
			Version:    	  "v1",
			URLParams:  	  "",
			StatusCode:       200,
			MaxExecutionTime: 300,
		},
		"TestGETRequestingsourceAliasAlias200v1": {
			Method:     	  "GET",
			Endpoint:   	  "/requesting-source/alias/{alias}",
			Version:    	  "v1",
			URLParams:  	  "",
			StatusCode:       200,
			MaxExecutionTime: 300,
		},
		"TestGETRequestingsourceAliasAlias200v2": {
			Method:     	  "GET",
			Endpoint:   	  "/requesting-source/alias/{alias}",
			Version:    	  "v1",
			URLParams:  	  "",
			StatusCode:       200,
			MaxExecutionTime: 300,
		},
		"TestGETRequestingsourceAliasAlias200v3": {
			Method:     	  "GET",
			Endpoint:   	  "/requesting-source/alias/{alias}",
			Version:    	  "v1",
		},

"TestGETRequestingsource200": {
		Method:     	  "GET",
		Endpoint:   	  "/requesting-source",
		Version:    	  "v1",
		URLParams:  	  "",
		StatusCode:       200,
		MaxExecutionTime: 300,
	},
"TestPOSTRequestingsource201": {
		Method:     	  "POST",
		Endpoint:   	  "/requesting-source",
		Version:    	  "v1",
		URLParams:  	  "",
		StatusCode:       201,
		MaxExecutionTime: 300,
	},
"TestPOSTRequestingsource400": {
		Method:     	  "POST",
		Endpoint:   	  "/requesting-source",
		Version:    	  "v1",
		URLParams:  	  "",
		StatusCode:       400,
		MaxExecutionTime: 300,
	},
"TestGETRequestingsourceId200": {
		Method:     	  "GET",
		Endpoint:   	  "/requesting-source/{id}",
		Version:    	  "v1",
		URLParams:  	  "",
		StatusCode:       200,
		MaxExecutionTime: 300,
	},
"TestGETRequestingsourceId400": {
		Method:     	  "GET",
		Endpoint:   	  "/requesting-source/{id}",
		Version:    	  "v1",
		URLParams:  	  "",
		StatusCode:       400,
		MaxExecutionTime: 300,
	},
"TestGETRequestingsourceId404": {
		Method:     	  "GET",
		Endpoint:   	  "/requesting-source/{id}",
		Version:    	  "v1",
		URLParams:  	  "",
		StatusCode:       404,
		MaxExecutionTime: 300,
	},
"TestDELETERequestingsourceId200": {
		Method:     	  "DELETE",
		Endpoint:   	  "/requesting-source/{id}",
		Version:    	  "v1",
		URLParams:  	  "",
		StatusCode:       200,
		MaxExecutionTime: 300,
	},
"TestDELETERequestingsourceId400": {
		Method:     	  "DELETE",
		Endpoint:   	  "/requesting-source/{id}",
		Version:    	  "v1",
		URLParams:  	  "",
		StatusCode:       400,
		MaxExecutionTime: 300,
	},
"TestDELETERequestingsourceId404": {
		Method:     	  "DELETE",
		Endpoint:   	  "/requesting-source/{id}",
		Version:    	  "v1",
		URLParams:  	  "",
		StatusCode:       404,
		MaxExecutionTime: 300,
	},
"TestPATCHRequestingsourceId404": {
		Method:     	  "PATCH",
		Endpoint:   	  "/requesting-source/{id}",
		Version:    	  "v1",
		URLParams:  	  "",
		StatusCode:       404,
		MaxExecutionTime: 300,
	},
"TestPATCHRequestingsourceId200": {
		Method:     	  "PATCH",
		Endpoint:   	  "/requesting-source/{id}",
		Version:    	  "v1",
		URLParams:  	  "",
		StatusCode:       200,
		MaxExecutionTime: 300,
	},
"TestPATCHRequestingsourceId400": {
		Method:     	  "PATCH",
		Endpoint:   	  "/requesting-source/{id}",
		Version:    	  "v1",
		URLParams:  	  "",
		StatusCode:       400,
		MaxExecutionTime: 300,
	},
"TestPATCHRequestingsourceUndeleteId200": {
		Method:     	  "PATCH",
		Endpoint:   	  "/requesting-source/undelete/{id}",
		Version:    	  "v1",
		URLParams:  	  "",
		StatusCode:       200,
		MaxExecutionTime: 300,
	},
"TestPATCHRequestingsourceUndeleteId400": {
		Method:     	  "PATCH",
		Endpoint:   	  "/requesting-source/undelete/{id}",
		Version:    	  "v1",
		URLParams:  	  "",
		StatusCode:       400,
		MaxExecutionTime: 300,
	},
"TestPATCHRequestingsourceUndeleteId404": {
		Method:     	  "PATCH",
		Endpoint:   	  "/requesting-source/undelete/{id}",
		Version:    	  "v1",
		URLParams:  	  "",
		StatusCode:       404,
		MaxExecutionTime: 300,
	},
	}
}
