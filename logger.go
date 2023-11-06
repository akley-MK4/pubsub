package pubsub

import "sync/atomic"

type ILogger interface {
	All(v ...interface{})
	AllF(format string, v ...interface{})
	Debug(v ...interface{})
	DebugF(format string, v ...interface{})
	Info(v ...interface{})
	InfoF(format string, v ...interface{})
	Warning(v ...interface{})
	WarningF(format string, v ...interface{})
	Error(v ...interface{})
	ErrorF(format string, v ...interface{})
}

type exampleLogger struct{}

func (t *exampleLogger) All(v ...interface{}) {

}

func (t *exampleLogger) AllF(format string, v ...interface{}) {

}

func (t *exampleLogger) Debug(v ...interface{}) {

}

func (t *exampleLogger) DebugF(format string, v ...interface{}) {

}

func (t *exampleLogger) Info(v ...interface{}) {

}

func (t *exampleLogger) InfoF(format string, v ...interface{}) {

}

func (t *exampleLogger) Warning(v ...interface{}) {

}

func (t *exampleLogger) WarningF(format string, v ...interface{}) {

}

func (t *exampleLogger) Error(v ...interface{}) {

}

func (t *exampleLogger) ErrorF(format string, v ...interface{}) {

}

var (
	exaLoggerInst ILogger = &exampleLogger{}
	extLoggerInst ILogger = nil
	loggerSetTag  int32
)

func getLoggerInstance() ILogger {
	if extLoggerInst != nil {
		return extLoggerInst
	}
	return exaLoggerInst
}

func SetLoggerInstance(logger ILogger) bool {
	if !atomic.CompareAndSwapInt32(&loggerSetTag, 0, 1) {
		return false
	}

	extLoggerInst = logger
	return true
}
