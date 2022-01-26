package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"path/filepath"
)

var zapLogger *zap.SugaredLogger
var logFileDesc *os.File

var logLevels = map[int]zapcore.Level {
	10: zap.DebugLevel,
	20: zap.InfoLevel,
	30: zap.WarnLevel,
	40: zap.ErrorLevel,
	50: zap.FatalLevel,
}

func InitLogger(filename string, logLevel int) {
	if zapLogger != nil {
		return
	}

	err := os.MkdirAll(filepath.Dir(filename), 0750)
	if err != nil {
		log.Fatal(err)
	}

	logFileDesc, err = os.OpenFile(filepath.Clean(filename), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		log.Fatal(err)
	}

	logFile := &lumberjack.Logger{
		Filename: filename,
		MaxSize: 1,
		MaxBackups: 3,
		MaxAge: 28,
	}

	sync := zapcore.WriteSyncer(zapcore.AddSync(logFile))

	encConfig := zap.NewProductionEncoderConfig()
	encConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encConfig)

	core := zapcore.NewCore(encoder, sync, logLevels[logLevel])

	logger := zap.New(core, zap.AddCaller())

	zapLogger = logger.Sugar()
}

func Info(args ...interface{}) {
	zapLogger.Info(args)
}

func Debug(args ...interface{}) {
	zapLogger.Debug(args)
}

func Warn(args ...interface{}) {
	zapLogger.Warn(args)
}

func Error(args ...interface{}) {
	zapLogger.Error(args)
}

func Fatal(args ...interface{}) {
	zapLogger.Fatal(args)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}