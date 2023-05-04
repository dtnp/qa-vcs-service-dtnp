package qaframework

import (
	"fmt"
	"testing"

	"github.com/dailymotion/allure-go"

	"go-webservices-automation/pkg/config"
)

// RunEndpointFunction we might or might not want to use Allure
//  This functions allows us to choose which we want to do
func RunEndpointFunction(t *testing.T, config config.Config, desc string, f func()) {
	// Is allure enabled in config?
	if config.Allure.Enabled {
		allure.Test(t, allure.Action(func() {
			allure.Step(
				allure.Description(desc),
				allure.Action(f),
			)
		}))
	} else {
		// If NOT enabled, then just call the function itself
		f()
	}
}

// RunEndpointFunctionv2 we might or might not want to use Allure
//  This functions allows us to choose which we want to do
func RunEndpointFunctionv2(t *testing.T, config config.Config, epd EndpointData,  f func(data EndpointData)) {
	description := fmt.Sprintf(
		"Verify %s (%d): version: %s, endpoint '%s'",
		epd.Method,
		epd.StatusCode,
		epd.Version,
		epd.Endpoint,
	)

	// Is allure enabled in config?
	if config.Allure.Enabled {
		allure.Test(t, allure.Action(func() {
			allure.Step(
				allure.Description(description),
				allure.Action(func() {f(epd)}),
			)
		}))
	} else {
		// If NOT enabled, then just call the function itself
		f(epd)
	}
}