package main

import (
	"flag"

	"go-webservices-automation/pkg/openapi"
)

func main() {
	var swaggerFile string
	flag.StringVar(&swaggerFile, "swaggerFile", "/data/lighthouse_swagger.json", "Path for swagger file")
	flag.Parse()

	openapi.ReadSwaggerDocument(swaggerFile)
}
