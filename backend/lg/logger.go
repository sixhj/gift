package lg

import (
	"go.uber.org/zap"
)

var (
	MyLog *zap.SugaredLogger
)

func Info(args ...interface{}) {
	MyLog.Info(args)
}
