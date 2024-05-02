package logger

import "fmt"

// Simple Logger for log messages in usecase layer
type Logger interface {
	Log(message string) error
}

type SimpleStdLogger struct{}

func (l *SimpleStdLogger) Log(message string) error {
	_, err := fmt.Println(message)
	return err
}
