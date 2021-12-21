package logger

//TODO: Redo all of this

type ErrorLogger interface {
	Error(msg string)
	Errorf(format string, args ...interface{})

	// CheckErrorWithMessage check error with formatted message and log messages with stack. And then return true if error is nil and return false if error is not nil.
	CheckErrorWithMessage(err error, format string, args ...interface{}) bool
	// CheckErrorNoStackWithMessage check error with formatted message and return true if error is nil and return false if error is not nil.
	CheckErrorNoStackWithMessage(err error, format string, args ...interface{}) bool
	// CheckError check error and return true if error is nil and return false if error is not nil.
	CheckError(err error) bool
	// CheckErrorNoStack CheckError check error and return true if error is nil and return false if error is not nil.
	CheckErrorNoStack(err error) bool
}

// Logger implementation interface
type Logger interface {
	ErrorLogger
	Debug(msg string)
	Debugf(format string, args ...interface{})
	Info(msg string)
	Infof(format string, args ...interface{})
	Warn(msg string)
	Warnf(format string, args ...interface{})
	Fatal(msg string)
	Fatalf(format string, args ...interface{})
	Panic(msg string)
	Panicf(format string, args ...interface{})

}

// Fields for the logger
type Fields map[string]interface{}
