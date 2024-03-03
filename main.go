package main

import (
	"go.uber.org/zap"
)

func main() {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	storage := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}

	logger.Info("Server started")
	if _, ok := storage["key1"]; ok {
		logger.Info("Key1 found")
	}
	if _, ok := storage["key4"]; !ok {
		logger.Error("Key4 not found")
	}
}
