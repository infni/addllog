package log

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type DefaultLogger struct {
	Logger GoLoggerInterface
}

func NewFileLogger(logger GoLoggerInterface) *DefaultLogger {
	r := new(DefaultLogger)

	r.Logger = logger

	return r
}

func (logger *DefaultLogger) Log(message string, action string, addlInfo AddlInfo) {
	logger.logMessage(message, "INFO", action, addlInfo)
}

func (logger *DefaultLogger) LogCritical(message string, action string, addlInfo AddlInfo) {
	logger.logMessage(message, "CRITICAL", action, addlInfo)
}

func (logger *DefaultLogger) logMessage(message string, logLevel string, action string, addlInfo AddlInfo) {

	if len(action) == 0 {
		action = "logging"
	}

	data := map[string]interface{}{
		Action:    action,
		LogLevel:  logLevel,
		Message:   strings.Replace(strings.Replace(message, "\"", "\\\"", -1), "\n", "\\n", -1),
		Timestamp: time.Now().Format(time.RFC3339Nano),
	}
	if len(addlInfo) != 0 {
		data[AdditionalInfo] = addlInfo.ToJson()
	}

	if jsonMsg, jsonErr := json.Marshal(data); jsonErr != nil {
		data["message"] = fmt.Errorf("error serilizing json : %w", jsonErr).Error()
		if jsonMsg2, jsonErr2 := json.Marshal(data); jsonErr2 != nil {
			fmt.Print(fmt.Errorf("Failed to LOG!! error serilizing json : %w\nOriginal message: %s", jsonErr2, message).Error())
		} else {
			logger.Logger.Print(string(jsonMsg2))
		}
	} else {
		logger.Logger.Print(string(jsonMsg))
	}
}
