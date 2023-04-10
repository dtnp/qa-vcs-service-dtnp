package requesting_tech

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"go-webservices-automation/pkg/config"
	"go-webservices-automation/pkg/qaframework"
	"go.uber.org/zap"
)

type TestSuite struct {
	config   config.Config
	Log      *zap.SugaredLogger
	SectionName string
}

var TS TestSuite
var SectionName = "requesting_tech"

// TestMain sets up things before and after tests
func TestMain(m *testing.M) {
	qas, err := qaframework.Setup(SectionName)
	if err != nil {
		// Better than panic?
		fmt.Printf("error for %s: %s", SectionName, err)
		os.Exit(0)
	}

	TS.SectionName = SectionName
	TS.Log = qas.Log
	TS.config = qas.Config

	// Run Tests
	success := m.Run()
	qas.Log.Infow("test completion", "section", SectionName)
	//ts.teardown()
	os.Exit(success)
}

func GetEndpointData(method string, name string) (qaframework.EndpointData, error) {
	me := strings.ToLower(fmt.Sprintf("%s %s", method, name))

	e := endpoints()
	ed, ok := e[me]
	if !ok {
		return qaframework.EndpointData{}, fmt.Errorf("no endpoint data: %s", me)
	}
	return ed, nil
}
