package qaframework

import (
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