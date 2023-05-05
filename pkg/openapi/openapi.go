/*
Package utils implements different utility methods
*/

package openapi

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Endpoint stores information about a given endpoint
type Endpoint struct {
	Endpoint     string
	PackageName  string
	TestCaseName string
	Method       string
	StatusCode   int
	isDynamic    bool
}

// ParseSwaggerDocument opens, reads, parses, and generates information
//   from the swagger doc supplied.
func ParseSwaggerDocument(swaggerFilePath string) error {
	// Open the swagger docs and make sure
	var result map[string]interface{}
	jsonFile, _ := os.Open(swaggerFilePath)
	byteValue, _ := io.ReadAll(jsonFile)
	err := json.Unmarshal(byteValue, &result)
	if err != nil {
		return err
	}

	// Large array of every possible path we find
	endpointPaths := result["paths"].(map[string]interface{})

	for endpointPath, value := range endpointPaths {
		// This is an overall array for all test data that we extract
		var endpointData []string

		// Empty endpoint object to hold useful data
		e := Endpoint{
			Endpoint: endpointPath,
			PackageName: GetPackageName(endpointPath),
		}
		if err = SetupTestPackage(e.PackageName); err != nil {
			return err
		}

		for method, methodResp := range value.(map[string]interface{}) {
			// Save the method
			e.Method = method

			// Each loops sets up a new set of endpoint details
			//   this is how golang sets up empty "dictionary" placeholders
			methodResponse := methodResp.(map[string]interface{})
			headers := make(map[string]interface{})
			pathParam := make(map[string]interface{})
			queryParam := make(map[string]interface{})
			//payload := make(map[string]interface{})
			response := make(map[string]interface{})

			// Check for parameters
			if _, present := methodResponse["parameters"]; present {
				for _, param := range methodResponse["parameters"].([]interface{}) {
					testData := param.(map[string]interface{})
					if testData["in"] == "header" {
						// TODO: Need to add condition whether parameter is required or not
						headers[testData["name"].(string)] = "Update valid values before running test"
					} else if testData["in"] == "path" {
						pathParam[testData["name"].(string)] = "Update valid values before running test"
					} else if testData["in"] == "query" {
						queryParam[testData["name"].(string)] = "Update valid values before running test"
					//} else if testData["in"] == "body" {
						//schemaPath := testData["schema"].(map[string]interface{})["$ref"].(string)
						//payload = ReadPayloadProperties(swaggerFilePath, make(map[string]interface{}), schemaPath)
					}
				}
			}
			//if _, present := methodResponse["requestBody"]; present {
			//	requestBody := methodResponse["requestBody"].(map[string]interface{})["content"].(map[string]interface{})["application/json"].(map[string]interface{})["schema"].(map[string]interface{})["$ref"].(string)
			//	payload = ReadPayloadProperties(swaggerFilePath, make(map[string]interface{}), requestBody)
			//}
			for statusCode, respBody := range methodResponse["responses"].(map[string]interface{}) {
				if statusCode == "500" {
					continue
				}
				for key, value := range respBody.(map[string]interface{}) {
					if key == "schema" {
						if _, present := value.(map[string]interface{})["$ref"]; present {
							response["resp"] = ReadPayloadProperties(swaggerFilePath, make(map[string]interface{}), value.(map[string]interface{})["$ref"].(string))
						} else {
							response[key] = value.(map[string]interface{})["type"].(string)
						}
					} else if key == "content" {
						for keys, value := range value.(map[string]interface{})["application/json"].(map[string]interface{})["schema"].(map[string]interface{})["properties"].(map[string]interface{}) {
							if _, present := value.(map[string]interface{})["$ref"]; present {
								response["resp"] = ReadPayloadProperties(swaggerFilePath, make(map[string]interface{}), value.(map[string]interface{})["$ref"].(string))
							} else {
								response[keys] = value.(map[string]interface{})["type"].(string)
							}
						}
					} else {
						response[key] = value.(string)
					}
				}

				e.TestCaseName = CreateTestCaseName(e, statusCode)
				e.isDynamic = strings.Contains(endpointPath, "{")
				e.StatusCode,_ = strconv.Atoi(statusCode)

				// Returns the data needed for the endpoints.go file
				ed := GetEndpointData(e)
				endpointData = append(endpointData, ed)

				if err = CreateTestCase(e); err != nil {
					return err
				}

				//fmt.Printf("%s (%s): %d\n", e.Method, e.Endpoint, e.StatusCode)

			}
		}



		if err = CreateEndpointsFile(e, endpointData); err != nil {
			return err
		}
	}
	return nil
}

func CreateTestCaseName(e Endpoint, statusCode string) string {
	replacer := strings.NewReplacer("{", "", "}", "", "-", "", "_", "")
	filteredEndpoint := replacer.Replace(e.Endpoint)
	endpointPathList := strings.Split(filteredEndpoint, "/")
	for i := range endpointPathList {
		endpointPathList[i] = strings.Title(endpointPathList[i])
	}
	return fmt.Sprintf("Test%s%s%s", strings.ToUpper(e.Method), strings.Join(endpointPathList, ""), statusCode)
}

// CreateTestCase creates a different file for each method
//	and populates it with tests
//  methods include: get, post, delete, patch, etc
func CreateTestCase(e Endpoint) error {

	// create the path to the method files
	methodFile := fmt.Sprintf(
		"test/swagger/%s/%s_test.go",
		e.PackageName,
		strings.ToLower(e.Method))

	// If the file didn't exist, create it with correct permissions
	if err := CreateMethodFile(e, methodFile); err != nil {
		return fmt.Errorf("create method file: %w", err)
	}

	// Read the contents of the file into a string
	//fileBytes, err := ioutil.ReadFile(methodFile) -- old

	// Open the file in append mode
	file, err := os.OpenFile(methodFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("opening method file: %w", err)
	}
	fileBytes, err := ioutil.ReadFile(methodFile)
	if err != nil {
		return fmt.Errorf("reading method file: %w", err)
	}
	fileContent := string(fileBytes)

	// Check if the search string is present in the file
	if strings.Contains(fileContent, e.TestCaseName) {
		return nil
	}

	// Some tests have dynamic aspects that we pass in like {id} for ID's
	var testCase string
	if e.isDynamic {
		testCase = TemplateGetDynamic(e)
	} else {
		testCase = TemplateGet(e)
	}

	// Write the text to the file
	if _, err = file.WriteString(testCase); err != nil {
		return fmt.Errorf("writing file: %w", err)
	}

	return nil
}

func CreateMethodFile(e Endpoint, methodFile string) error {
	if !FileExists(methodFile) {
		f, err := os.Create(methodFile)
		if err != nil {
			return err
		}
		testCaseTemplate := `package %s

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go-webservices-automation/pkg/qaframework"
)

`
		_, err = f.WriteString(fmt.Sprintf(testCaseTemplate, e.PackageName))
		if err != nil {
			return fmt.Errorf("writing to file: %w", err)
		}
	}
	return nil
}


// GetPackageName split an endpoint into pieces and returns the name of
//   the package.
func GetPackageName(endpointPath string) string {
	pieces := strings.Split(endpointPath, "/")
	if len(pieces) == 1 {
		return "/"
	}

	return strings.ReplaceAll(pieces[1], "-", "_")
}


// SetupTestPackage creates the folders and initial files needed
func SetupTestPackage(packageName string) error {
	swaggerDirPath := fmt.Sprintf("test/swagger/%s", packageName)
	testFile := fmt.Sprintf("%s/%s_test.go", swaggerDirPath, packageName)
	endpointFile := fmt.Sprintf("%s/endpoints.go", swaggerDirPath)

	// Check if the folder already exists.  If not, create it
	if !FileExists(swaggerDirPath) {
		if err := os.MkdirAll(swaggerDirPath, os.ModePerm); err != nil {
			return err
		}
	}
	// Creates the main test runner file
	if !FileExists(testFile) {
		f, err := os.Create(testFile)
		if err != nil {
			return err
		}
		tpl := NewTestTemplate(packageName)
		fmt.Printf("%s\n--------------------\n\n", tpl)
		n, err := f.WriteString(tpl)
		if err != nil {
			return err
		}
		fmt.Printf("created %s: %d bytes\n", testFile, n)
	}

	// Creates the endpoints file for all the individual endpoints
	if !FileExists(endpointFile) {
		_, err := os.Create(testFile)
		if err != nil {
			return err
		}
	}
	return nil
}

// FileExists checks if a file and/or folder exists
//    TODO: Fix this so that we split file and folder checks apart
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}


func CreateEndpointsFile(e Endpoint, endpointData []string) error {
	endpointFilePath := fmt.Sprintf("test/swagger/%s/endpoints.go", e.PackageName)

	// Build a big string of all test case data
	var builder strings.Builder
	for _, d := range endpointData {
		builder.WriteString(d)
	}

	if !FileExists(endpointFilePath) {
		f, err := os.Create(endpointFilePath)
		if err != nil {
			return fmt.Errorf("create endpoints file: %w", err)
		}
		template := `
package %s

import (
	"go-webservices-automation/pkg/qaframework"
)

// endpoints are the individual endpoints for this specific folder
func endpoints() map[string]qaframework.EndpointData {
	return map[string]qaframework.EndpointData {
		%s
		// ENDPOINTDATA
	}
}
`
		_, err = f.WriteString(fmt.Sprintf(template, e.PackageName, builder.String()))
		if err != nil {
			return fmt.Errorf("write endpoints file: %w", err)
		}
		return nil
	}

	fileBytes, err := ioutil.ReadFile(endpointFilePath)
	if err != nil {
		return fmt.Errorf("reading method file: %w", err)
	}
	fileContent := string(fileBytes)

	fileContent = strings.Replace(fileContent, "// ENDPOINTDATA", builder.String(), 1)
	f, err := os.OpenFile(endpointFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return fmt.Errorf("overwriting endpoint file: %w", err)
	}
	_, err = f.WriteString(fileContent)
	if err != nil {
		return fmt.Errorf("overwriting endpoint file: %w", err)
	}

	return nil
}



func TemplateGet(e Endpoint) string {
	output := `
func #FUNCTION_NAME#(t *testing.T) {
	RunTest(t, "#FUNCTION_NAME#", func(e qaframework.EndpointData) {
		req := require.New(t)
		res, err := qaframework.APICallByGETDynamic(TS.config, e, nil)
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
`
	output = strings.ReplaceAll(output, "#FUNCTION_NAME#", e.TestCaseName)
	return output
}

func TemplateGetDynamic(e Endpoint) string {
	var re = regexp.MustCompile(`{[a-zA-Z0-9]+}`)
	match := re.FindStringSubmatch(e.Endpoint)
	dynamic := match[len(match)-1]

	output := `
func #FUNCTION_NAME#(t *testing.T) {
	// This endpoint can pass in multiple tests, comma separated
	tests := []string{""}

	RunTest(t, "#FUNCTION_NAME#", func(e qaframework.EndpointData) {
		req := require.New(t)

		for _, tval := range tests {
			res, err := qaframework.APICallByGETDynamic(TS.config, e, map[string]string{"#DYNAMIC#": tval})
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
		}
	})
}
`
	output = strings.ReplaceAll(output, "#FUNCTION_NAME#", e.TestCaseName)
	output = strings.Replace(output, "#DYNAMIC#", dynamic, 1)

	return output
}

func GetEndpointData(e Endpoint) string {
	testDataTemplate := `
		"%s": {
			Method:     	  "%s",
			Endpoint:   	  "%s",
			Version:    	  "%s",
			URLParams:  	  "",
		},`
	// TODO: change the hardcoded v1 here
	testData := fmt.Sprintf(testDataTemplate, e.TestCaseName, strings.ToUpper(e.Method), e.Endpoint, "v1")

	return testData
}

func ReadPayloadProperties(swaggerFilePath string, payload map[string]interface{}, location string) map[string]interface{} {
	var result map[string]interface{}
	jsonFile, _ := os.Open(swaggerFilePath)
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err := json.Unmarshal([]byte(byteValue), &result)
	if err != nil {
		panic(err)
	}
	payloadPath := strings.Split(location, "/")

	var payloadschemas map[string]interface{}
	payloadschemas = result

	for _, val := range payloadPath[1:] {
		payloadschemas = payloadschemas[val].(map[string]interface{})
	}

	for key, value := range payloadschemas["properties"].(map[string]interface{}) {
		payload[key] = value.(map[string]interface{})["type"].(string)
	}
	return payload
}

func NewTestTemplate(section string) string {
	output := `
package #PACKAGE_NAME#

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
	config   config.Config
	Log      *zap.SugaredLogger
	SectionName string
}

var TS TestSuite
var SectionName = "template_section"

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

	time.Sleep(1*time.Second)

	// Run Tests
	success := m.Run()
	qas.Log.Infow("test completion", "section", SectionName)
	//ts.teardown()
	os.Exit(success)
}

// GetEndpointData reads the basic details from the endpoints file in
//   this specific folder
func GetEndpointData(name string) (qaframework.EndpointData, error) {
	// me := strings.ToLower(fmt.Sprintf("%s %s", method, name))

	e := endpoints()
	ed, ok := e[name]
	if !ok {
		return qaframework.EndpointData{}, fmt.Errorf("no endpoint data: %s", name)
	}
	return ed, nil
}

// GetEndpointData reads the basic details from the endpoints file in
//   this specific folder
func RunTest(t *testing.T, endpointName string, f func(data qaframework.EndpointData)) {
	ed, err := GetEndpointData(endpointName)
	if err != nil {
		t.Error(err)
		return
	}

	qaframework.RunEndpointFunction(t, TS.config, ed, f)
}
`
	o := strings.ReplaceAll(output, "#PACKAGE_NAME#", section)
	fmt.Println(o)
	return o
}