package qaframework

import (
	"testing"

	"github.com/dailymotion/allure-go"

	"go-webservices-automation/pkg/config"
)

func RunEndpointFunction(t *testing.T, config config.Config, desc string, f func()) {
	if config.Allure.Enabled {
		allure.Test(t, allure.Action(func() {
			allure.Step(
				allure.Description(desc),
				allure.Action(f),
			)
		}))
	} else {
		f()
	}
}