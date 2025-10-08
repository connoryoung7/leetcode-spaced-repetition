package internal

import (
	"sync"

	"go.uber.org/zap"
)

type Logger struct {
	Logger *zap.Logger
}

var (
	instance *zap.Logger
	once     sync.Once
)

// func GetInstance() *Logger {
// 	once.Do(func() {
// 		// Create a production logger (JSON format)
// 		zapLogger, err := zap.NewProduction()
// 		if err != nil {
// 			panic("failed to initialize logger: " + err.Error())
// 		}

// 		instance = &Logger{Logger: zapLogger}
// 	})
// 	return instance
// }
