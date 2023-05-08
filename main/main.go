package main

import (
	"flag"
	"fmt"
	"os"

	"go-webservices-automation/pkg/openapi"
)

func main() {
	var swaggerFile string
	flag.StringVar(&swaggerFile, "swaggerFile", "/data/lighthouse_swagger.json", "Path for swagger file")
	flag.Parse()

	if err := openapi.ParseSwaggerDocument(swaggerFile); err != nil {
		fmt.Printf("\nOpenAPI/Swagger import failed: %s\n\n", err)
		os.Exit(1)
	}

}
