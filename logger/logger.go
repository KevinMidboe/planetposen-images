package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

// Event stores messages to log later, from our standard interface
type Event struct {
	id      int
	message string
}

// ApplicationLogger enforces specific log message formats
type ApplicationLogger struct {
	*logrus.Logger
}

// NewLogger initializes the standard logger
func InitLogger() *ApplicationLogger {
	var baseLogger = logrus.New()
	var applicationLogger = &ApplicationLogger{baseLogger}

	f, _ := os.OpenFile("/var/log/planetposen_logs/planetposen-images.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	applicationLogger.SetOutput(f)
	applicationLogger.Formatter = &logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyMsg:  "message",
			logrus.FieldKeyTime: "@timestamp",
		},
	}

	return applicationLogger
}

// Declare variables to store log messages as new Events
var (
	infoMessage            = Event{1, "%s"}
	invalidArgMessage      = Event{2, "Invalid arg: %s"}
	invalidArgValueMessage = Event{3, "Invalid value for argument: %s: %v"}
	missingArgMessage      = Event{4, "Missing arg: %s"}
)

func (l *ApplicationLogger) Info(argumentName string) {
	l.WithFields(logrus.Fields{
		"application": "planetposen-images",
	}).Infof(infoMessage.message, argumentName)
}

func (l *ApplicationLogger) InfoWithPath(message string, path string) {
	l.WithFields(logrus.Fields{
		"application": "planetposen-images",
		"path":        path,
	}).Info(message)
}

func (l *ApplicationLogger) InfoWithFilename(message string, filename string) {
	l.WithFields(logrus.Fields{
		"application": "planetposen-images",
		"filename":    filename,
	}).Info(message)
}

func (l *ApplicationLogger) UploadSuccessMessage(path string, url string) {
	l.WithFields(logrus.Fields{
		"application": "planetposen-images",
		"path":        path,
		"url":         url,
	}).Info("Uploaded and available")
}

func (l *ApplicationLogger) Error(message string, error error) {
	l.WithFields(logrus.Fields{
		"application": "planetposen-images",
		"err":         error,
	}).Error(message)
}

func (l *ApplicationLogger) Fatal(error error) {
	l.WithFields(logrus.Fields{
		"application": "planetposen-images",
		"err":         error,
	}).Fatal("Fatal application exception!")
}

func (l *ApplicationLogger) GoogleApiError(message string, error error) {
	l.WithFields(logrus.Fields{
		"application": "planetposen-images",
		"go_package":  "cloud.google.com/go/storage",
		"err":         error,
	}).Error(message)
}

// InvalidArg is a standard error message
func (l *ApplicationLogger) InvalidArg(argumentName string) {
	l.Errorf(invalidArgMessage.message, argumentName)
}

// InvalidArgValue is a standard error message
func (l *ApplicationLogger) InvalidArgValue(argumentName string, argumentValue string) {
	l.Errorf(invalidArgValueMessage.message, argumentName, argumentValue)
}

// MissingArg is a standard error message
func (l *ApplicationLogger) MissingArg(argumentName string) {
	l.Errorf(missingArgMessage.message, argumentName)
}
