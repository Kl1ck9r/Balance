package zap

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func InitLogger() (*zap.Logger, error) {
	var logger *zap.Logger
	file, err := os.OpenFile("logger.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("Failed to open file[logger.log]: %v ", err)
	}

	info, err := file.Stat()
	if err != nil {
		return logger, fmt.Errorf("gotten file stat: %v", err)
	}

	if info.IsDir() {
		return logger, fmt.Errorf("%s is directory", info.Name())
	}

	zp := zap.NewProductionEncoderConfig()
	flEconding := zapcore.NewJSONEncoder(zp)

	zpLevel := zap.DebugLevel

	core := zapcore.NewTee(
		zapcore.NewCore(flEconding, zapcore.AddSync(file), zpLevel),
	)

	logger = zap.New(core)

	return logger, nil
}
