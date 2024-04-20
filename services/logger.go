package services

// Simple Logger for log messages in usecase layer
type Logger interface {
	Log(message string) error
}
