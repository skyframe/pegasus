package log

import (
	"go.uber.org/zap"
	nLog "log"
	"os"
	"path/filepath"
)

type Config struct {
	Filename        string
	LogLevel        Level
	SlackLevel      Level
	ApplicationName string
}

var sugaredLogger *zap.SugaredLogger

func init() {
	config := &Config{
		ApplicationName: "Pegasus",
		Filename:        "./logs/pegasus.log",
		LogLevel:        DebugLevel,
		SlackLevel:      ErrorLevel,
	}

	outputPaths := []string{"stdout"}
	if config.Filename != "" {
		if _, err := os.Stat(config.Filename); os.IsNotExist(err) {
			err := os.MkdirAll(filepath.Dir(config.Filename), os.ModePerm)
			if err != nil {
				nLog.Fatal("Could not create logs dir: ", err)
			}
			f, err := os.Create(config.Filename)
			if err != nil {
				nLog.Fatal("Could not create log file "+config.Filename+": ", err)
			}
			f.Close()
		}

		outputPaths = append(outputPaths, config.Filename)
	}

	zapConfig := zap.NewDevelopmentConfig()
	zapConfig.OutputPaths = outputPaths
	zapConfig.Level = config.LogLevel.ToZap()

	zapper, err := zapConfig.Build(zap.AddCallerSkip(1))
	if err != nil {
		nLog.Fatal("Could not instantiate zap logger: ", err)
	}

	defer zapper.Sync()
	sugaredLogger = zapper.Sugar()
}
func Error(args ...interface{}) {
	sugaredLogger.Error()
}

func Info(args ...interface{}) {

}
