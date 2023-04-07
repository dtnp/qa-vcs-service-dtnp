// Package logger for logger functions
package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// GetProductionLogger initialises production environment logger
func GetProductionLogger(appName string, appVersion string) (*zap.SugaredLogger, error) {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.DisableStacktrace = true
	config.InitialFields = map[string]interface{}{
		"service": appName,
		"version": appVersion,
	}
	lb, err := config.Build()
	if err != nil {
		return nil, err
	}
	return lb.Sugar(), nil
}

// GetDevelopmentLogger initialises development environment logger
func GetDevelopmentLogger(appName string, appVersion string, jsonEncoding bool) (*zap.SugaredLogger, error) {
	// Fetch app version if not found
	if appVersion == "" {
		// Ignore errors, this is optional
		version, _ := os.ReadFile("../../VERSION")
		if version != nil {
			appVersion = strings.TrimSpace(string(version))
		}
	}

	config := zap.NewDevelopmentConfig()
	if jsonEncoding {
		config.Encoding = "json"
	}
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.OutputPaths = []string{"stdout"}
	config.DisableStacktrace = true
	config.InitialFields = map[string]interface{}{
		"service": appName,
		"version": appVersion,
	}
	lb, err := config.Build()
	if err != nil {
		return nil, err
	}
	log := lb.Sugar()

	return log, nil
}
