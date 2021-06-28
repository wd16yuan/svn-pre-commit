package log

import (
	"os"
	"path"
	"time"

	nested "github.com/antonfisher/nested-logrus-formatter"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

var homeDir = "logs"
var fileName = "pre_commit"
var log = logrus.New()

func init() {
	initDir(homeDir)
	initConfig(path.Join(homeDir, fileName))
}

func initDir(home string) {
	info, err := os.Stat(home)
	if err == nil {
		if info.IsDir() {
			return
		}
	}
	os.Mkdir(home, os.ModeDir)
}

func initConfig(logPath string) {
	writer, _ := rotatelogs.New(
		logPath+".%Y%m%d",
		//rotatelogs.WithLinkName(logPath),
		rotatelogs.WithRotationCount(2),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	log.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		TimestampFormat: time.RFC3339,
	})
	log.SetOutput(writer)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Warning(args ...interface{}) {
	log.Warning(args...)
}
