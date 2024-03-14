package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	SystemLog *zap.Logger // SystemLog is the logger for system logs.
)

func init() {
	// Create a new production encoder configuration.
	config := zap.NewProductionEncoderConfig()
	// Set the time encoder to ISO8601.
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	// Open the info log file with necessary permissions.
	infoFile, _ := os.OpenFile("./log/file"+"/info-log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// Open the error log file with necessary permissions.
	errorFile, _ := os.OpenFile("./log/file"+"/error-log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	// Create a new core for the logger that writes to both the info and error log files.
	treeCore := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.AddSync(infoFile), zapcore.InfoLevel),
		zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.AddSync(errorFile), zapcore.ErrorLevel),
	)

	// Initialize the logger with the created core and add the caller option.
	SystemLog = zap.New(treeCore, zap.AddCaller())
}
