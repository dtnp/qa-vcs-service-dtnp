package requesting_source

import (
	"fmt"
	"os"
	"testing"
	"time"

	"go-webservices-automation/pkg/config"
	"go-webservices-automation/pkg/qaframework"
	"go.uber.org/zap"
)

type TestSuite struct {
	config       config.Config
	Log          *zap.SugaredLogger
	SectionName  string
	//EndpointData map[string]qaframework.EndpointData
}

var TS TestSuite
var SectionName = "requesting_source"

// TestMain sets up things before and after tests
func TestMain(m *testing.M) {
	qas, err := qaframework.Setup(SectionName)
	if err != nil {
		// Better than panic?
		fmt.Printf("error for %s: %s", SectionName, err)
		os.Exit(0)
	}

	TS.SectionName = SectionName
	//TS.EndpointData = EndpointData()
	TS.Log = qas.Log
	TS.config = qas.Config

	time.Sleep(1*time.Second)

	// Run Tests
	success := m.Run()
	qas.Log.Infow("test completion", "section", SectionName)
	//ts.teardown()
	os.Exit(success)
}

// GetEndpointData reads the basic details from the endpoints file in
//   this specific folder
func GetEndpointData(method string, name string) (qaframework.EndpointData, error) {
	// me := strings.ToLower(fmt.Sprintf("%s %s", method, name))

	e := EndpointData()
	ed, ok := e[name]
	if !ok {
		return qaframework.EndpointData{}, fmt.Errorf("no endpoint data: %s", name)
	}
	return ed, nil
}

// GetEndpointData reads the basic details from the endpoints file in
//   this specific folder
func GetEndpointDatav2(endpointName string) (qaframework.EndpointData, error) {
	e := EndpointData()
	ed, ok := e[endpointName]
	if !ok {
		return qaframework.EndpointData{}, fmt.Errorf("no endpoint data: %s", endpointName)
	}
	return ed, nil
}

// GetEndpointData reads the basic details from the endpoints file in
//   this specific folder
func RunTest(t *testing.T, endpointName string, f func(data qaframework.EndpointData)) {
	ed, err := GetEndpointDatav2(endpointName)
	if err != nil {
		t.Error(err)
		return
	}

	qaframework.RunEndpointFunctionv2(t, TS.config, ed, f)
}
