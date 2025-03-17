package logging

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

// InitLogger initializes the logger with the specified environment (development or production).
func InitLogger(env string) {
    var config zap.Config

    if env == "production" {
        config = zap.NewProductionConfig()
    } else {
        config = zap.NewDevelopmentConfig()
        config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // Colorful logs for development
    }

    // Customize the output paths
    config.OutputPaths = []string{"stdout", "./logs/app.log"}
    config.ErrorOutputPaths = []string{"stderr", "./logs/error.log"}

    // Build the logger
    var err error
    Logger, err = config.Build()
    if err != nil {
        panic("Failed to initialize logger: " + err.Error())
    }

    // Ensure logs are flushed before the application exits
    defer Logger.Sync()
}

// Info logs an info-level message.
func Info(message string, fields ...zap.Field) {
    Logger.Info(message, fields...)
}

// Error logs an error-level message.
func Error(message string, fields ...zap.Field) {
    Logger.Error(message, fields...)
}

// Warn logs a warning-level message.
func Warn(message string, fields ...zap.Field) {
    Logger.Warn(message, fields...)
}

// Debug logs a debug-level message.
func Debug(message string, fields ...zap.Field) {
    Logger.Debug(message, fields...)
}

// Fatal logs a fatal-level message and exits the application.
func Fatal(message string, fields ...zap.Field) {
    Logger.Fatal(message, fields...)
}