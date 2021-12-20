package logger

import "fmt"

type Logger struct {
	level string
}

func New(level string) *Logger {
	return &Logger{
		level: level,
	}
}

func (l *Logger) stdout(level string, message string) {
	fmt.Printf(
		"[%s]: %s\n",
		level,
		message,
	)
}



func (l *Logger) Debug(message string) {
	l.stdout("DEBUG", message)
}
