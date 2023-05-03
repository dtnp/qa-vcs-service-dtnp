package qaframework

import (
	"go-webservices-automation/pkg/config"
	"go-webservices-automation/pkg/logger"
	"go.uber.org/zap"
)

type QAFramework struct {
	Config   config.Config
	Log      *zap.SugaredLogger
}

var qaf QAFramework

// Setup sets up the overall sections in the test folder
func Setup(sectionName string) (QAFramework, error) {
	// Debugger Setup
	// TODO eww, don't hard code anything in here
	qaf.Log, _ = logger.GetDevelopmentLogger("qa-vcs-service-dtn","1", false)
	qaf.Log.Infow("test setup", "section", sectionName)

	// Log Configs
	// TODO eww, don't hard code anything in here
	c, err := config.GetConfig("../../../qa-vcs-service-dtnp.config.json")
	if err != nil {
		return QAFramework{}, err
	}
	qaf.Config = c

	return qaf, nil
}
