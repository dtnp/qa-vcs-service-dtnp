/*
Package utils implements different utility methods
*/

package openapi

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const MaxExecutionTimeDefault = 300

func ReadSwaggerDocument(swaggerFilePath string) {
	var result map[string]interface{}
	jsonFile, _ := os.Open(swaggerFilePath)
	byteValue, _ := io.ReadAll(jsonFile)
	err := json.Unmarshal(byteValue, &result)
	if err != nil {
		fmt.Println(err)
	}
	endpointPaths := result["paths"].(map[string]interface{})
	testData := []string{}

	for endpointPath, value := range endpointPaths {
		if !strings.Contains(endpointPath, "requesting-source") {
			continue
		}

		filePath := CreateTestFile(endpointPath)
		for method, methodResp := range value.(map[string]interface{}) {

			//pp, _ := json.MarshalIndent(methodResp, "", " ")
			//fmt.Printf("!!Pretty Print:\n%s\n", string(pp))
			//break


			methodResponse := methodResp.(map[string]interface{})
			headers := make(map[string]interface{})
			pathParam := make(map[string]interface{})
			queryParam := make(map[string]interface{})
			payload := make(map[string]interface{})
			response := make(map[string]interface{})
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
					} else if testData["in"] == "body" {
						schemaPath := testData["schema"].(map[string]interface{})["$ref"].(string)
						payload = ReadPayloadProperties(swaggerFilePath, make(map[string]interface{}), schemaPath)
					}
				}
			}
			if _, present := methodResponse["requestBody"]; present {
				requestBody := methodResponse["requestBody"].(map[string]interface{})["content"].(map[string]interface{})["application/json"].(map[string]interface{})["schema"].(map[string]interface{})["$ref"].(string)
				payload = ReadPayloadProperties(swaggerFilePath, make(map[string]interface{}), requestBody)
			}
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
				td := AddTestData(endpointPath, method, statusCode, payload, pathParam, queryParam, headers, response)
				//testData = fmt.Sprintf("%s\n%s", testData, td)
				//fmt.Printf("\n\nTest Data: %s", td)
				testData = append(testData, td)

				CreateTestCase(filePath, endpointPath, method, statusCode)
			}
		}
	}
	CreateEndpointsFile("requesting_source", testData)
}

func CreateEndpointsFile(endpoint string, data []string) {
	endpointDirPath := fmt.Sprintf("test/swagger/%s", endpoint)
	endpointFilePath := fmt.Sprintf("%s/endpoints.go", endpointDirPath)

	td := ""
	for _, d := range data {
		td = fmt.Sprintf("%s\n%s", td, d)
	}

	fmt.Printf("TEST DATA: %v", td)

	if !FileExists(endpointFilePath) {
		f, err := os.Create(endpointFilePath)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		template := `
package %s

import (
	"go-webservices-automation/pkg/qaframework"
)

// endpoints are the individual endpoints for this specific folder
func endpoints() map[string]qaframework.EndpointData {
	return map[string]qaframework.EndpointData {
		%s
	}
}
`
		n, err := f.WriteString(fmt.Sprintf(template, endpoint, td))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("wrote %d bytes\n", n)
		// TODO : Check returned error
		_ = f.Sync()
	}
	return
}

func CreateTestFile(endpoint string) string {
	fileName := strings.Split(endpoint, "/")[1]
	fileName = strings.ReplaceAll(fileName, "-", "_")
	swaggerDirPath := fmt.Sprintf("test/swagger/%s", fileName)
	testFilePath := fmt.Sprintf("%s/%s_test.go", swaggerDirPath, fileName)

	if !FileExists(swaggerDirPath) {
		if err := os.MkdirAll(swaggerDirPath, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}
	if !FileExists(testFilePath) {
		f, err := os.Create(testFilePath)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
//		n, err := f.WriteString(`
//package test
//
//import (
//	"go-webservices-automation/config"
//	"go-webservices-automation/utils"
//	"testing"
//
//	"github.com/dailymotion/allure-go"
//	"github.com/stretchr/testify/require"
//)
//		`)
		n, err := f.WriteString(NewTestTemplate(fileName))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("wrote %d bytes\n", n)
		// TODO : Check returned error
		_ = f.Sync()
	}
	return swaggerDirPath
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func CreateTestCase(filePath string, endpoint string, method string, statusCode string) {
	replacer := strings.NewReplacer("{", "", "}", "", "-", "", "_", "")
	filteredEndpoint := replacer.Replace(endpoint)
	endpointPathList := strings.Split(filteredEndpoint, "/")
	for i := range endpointPathList {
		endpointPathList[i] = strings.Title(endpointPathList[i])
	}
	testCaseName := fmt.Sprintf("Test%s%s%s", strings.ToUpper(method), strings.Join(endpointPathList, ""), statusCode)
	//description := fmt.Sprintf(`Test case to verify %s status code on '%s' endpoint with %s request`, statusCode, endpoint, strings.ToUpper(method))

	methodFile := fmt.Sprintf("%s/%s_test.go", filePath, strings.ToLower(method))

	if !FileExists(methodFile) {
		f, err := os.Create(methodFile)
		if err != nil {
			log.Fatal(err)
		}
		fileName := strings.Split(endpoint, "/")[1]
		fileName = strings.ReplaceAll(fileName, "-", "_")
		testCaseTemplate := `package %s

import (
	"fmt"
	"testing"

	"go-webservices-automation/pkg/qaframework"

	"github.com/stretchr/testify/require"
)

`
		testCase := fmt.Sprintf(testCaseTemplate, fileName)

		_, err = f.WriteString(testCase)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
		f.Close()
	}

	// Read the contents of the file into a string
	fileBytes, err := ioutil.ReadFile(methodFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fileContent := string(fileBytes)

	// Check if the search string is present in the file
	if !strings.Contains(fileContent, testCaseName) {
//		testCaseTemplate := `
//func %s(t *testing.T) {
//	allure.Test(t, allure.Action(func() {
//	allure.Step(allure.Description("%s"),
//	allure.Action(func() {
//		config.GenerateLog(utils.GetCurrentFuncName())
//		require := require.New(utils.WrapT(t))
//		requestParams:= utils.GetTestData("%s")
//		statusCode, _ := utils.SendRequest(requestParams)
//		require.Equal(statusCode, %s, "Status not matching")
//	}))
//}))
//}
//`

		testCase := TemplateGet(testCaseName, method)
		//testCase := fmt.Sprintf(testCaseTemplate, testCaseName, description, testCaseName, statusCode)
		// Open the file in append mode
		file, err := os.OpenFile(methodFile, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()
		// Write the text to the file
		_, err = file.WriteString(testCase)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
}

func TemplateGet(name string, method string) string {
	output := `
func #FUNCTION_NAME#(t *testing.T) {
	method := "#METHOD#"
	name := "#FUNCTION_NAME#"
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

		req.Equal(ed.StatusCode, res.StatusCode, "Status code mismatch")
		req.True(res.Data.Success)
		req.NotEmpty(res.Data.Data)
		req.GreaterOrEqual(res.Data.Timestamp, res.Timestamp)
		req.LessOrEqual(res.ResponseTime , int64(ed.MaxExecutionTime))
	})
}
`
	output = strings.ReplaceAll(output, "#FUNCTION_NAME#", name)
	output = strings.Replace(output, "#METHOD#", method, 1)
	return output
}

func AddTestData(endpoint string, method string, statusCode string, payload map[string]interface{}, pathParams map[string]interface{}, queryParams map[string]interface{}, headers map[string]interface{}, response map[string]interface{}) string {
	replacer := strings.NewReplacer("{", "", "}", "", "-", "", "_", "")
	filteredEndpoint := replacer.Replace(endpoint)
	endpointPathList := strings.Split(filteredEndpoint, "/")
	for i := range endpointPathList {
		endpointPathList[i] = strings.Title(endpointPathList[i])
	}
	testCaseName := fmt.Sprintf("Test%s%s%s", strings.ToUpper(method), strings.Join(endpointPathList, ""), statusCode)
	//testDataFilePath := config.FindRootDir() + "/data/testData.json"
//	testDataFilePath := "/data/testData.json"
//
//	if !FileExists(testDataFilePath) {
//		f, err := os.Create(testDataFilePath)
//		if err != nil {
//			log.Fatal(err)
//		}
//		defer f.Close()
//		n, err := f.WriteString(`{
//
//}`)
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Printf("wrote %d bytes\n", n)
//		err = f.Sync()
//		if err != nil {
//			log.Fatal(err)
//		}
//	}
//
//	// Convert the map to a JSON string
//	payloadjsonData, err := json.Marshal(payload)
//	if err != nil {
//		panic(err)
//	}
//	pathParamsjsonData, err := json.Marshal(pathParams)
//	if err != nil {
//		panic(err)
//	}
//	headersjsonData, err := json.Marshal(headers)
//	if err != nil {
//		panic(err)
//	}
//	queryParamsjsonData, err := json.Marshal(queryParams)
//	if err != nil {
//		panic(err)
//	}
//	responsejsonData, err := json.Marshal(response)
//	if err != nil {
//		panic(err)
//	}
//
//	// Read the contents of the file into a string
//	fileBytes, err := ioutil.ReadFile(testDataFilePath)
//	if err != nil {
//		fmt.Println("Error reading file:", err)
//		return
//	}
//	fileContent := string(fileBytes)
	// Check if the search string is present in the file
	//if !strings.Contains(fileContent, testCaseName) {
	//	testDataTemplate := `
	//	"%s": {
	//		"method": "%s",
	//		"endpoint": "%s",
	//		"payload": %v,
	//		"PathParams": %v,
	//		"header": %v,
	//		"queryParams": %v,
	//		"response": %v
	//	}
	//	`

	testDataTemplate := `"%s": {
		Method:     	  "%s",
		Endpoint:   	  "%s",
		Version:    	  "%s",
		URLParams:  	  "",
		StatusCode:       %d,
		MaxExecutionTime: %d,
	},`
// 	URLParams:  	  #URLParams#,

	intStatusCode, err := strconv.Atoi(statusCode)
	if err != nil {
		fmt.Printf("string conversion: %s", err)
		return ""
	}

	testData := fmt.Sprintf(testDataTemplate, testCaseName, strings.ToUpper(method), endpoint, "v1", intStatusCode, MaxExecutionTimeDefault)
	//
	//	file, err := os.OpenFile(testDataFilePath, os.O_RDWR, 0644)
	//	if err != nil {
	//		fmt.Println(err)
	//		return ""
	//	}
	//	defer file.Close()
	//
	//	// read the contents of the file into a buffer
	//	scanner := bufio.NewScanner(file)
	//	var buffer []string
	//	for scanner.Scan() {
	//		buffer = append(buffer, scanner.Text())
	//	}
	//
	//	// modify the buffer by updating the desired line
	//	if len(buffer) > 3 {
	//		testData = ", " + testData
	//	}
	//	buffer[len(buffer)-1] = testData + "}"
	//
	//	// write the modified buffer back to the file
	//	var prettyJSON bytes.Buffer
	//	error := json.Indent(&prettyJSON, []byte(fmt.Sprintf("%s\n", strings.Join(buffer, "\n"))), "", "\t")
	//	if error != nil {
	//		log.Println("JSON parse error: ", error)
	//		return ""
	//	}
	//	err = ioutil.WriteFile(testDataFilePath, prettyJSON.Bytes(), 0644)
	//	if err != nil {
	//		fmt.Println(err)
	//		return ""
	//	}
	////}

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
package template_section

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
func GetEndpointData(method string, name string) (qaframework.EndpointData, error) {
	// me := strings.ToLower(fmt.Sprintf("%s %s", method, name))

	e := endpoints()
	ed, ok := e[name]
	if !ok {
		return qaframework.EndpointData{}, fmt.Errorf("no endpoint data: %s", name)
	}
	return ed, nil
}`

	return strings.ReplaceAll(output, "template_section", section)
}