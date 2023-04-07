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

// Setup setup
func Setup(sectionName string) (QAFramework, error) {
	// Debugger Setup
	qaf.Log, _ = logger.GetDevelopmentLogger("qa-vcs-service-dtn","1", false)
	qaf.Log.Infow("test setup", "section", sectionName)

	// Log Configs
	c, err := config.GetConfig("../../qa-vcs-service-dtnp.config.json")
	if err != nil {
		return QAFramework{}, err
	}
	qaf.Config = c

	return qaf, nil
}
