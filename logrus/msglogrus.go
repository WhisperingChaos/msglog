package msglogrus

import (
	msg "msgloggit"
	"runtime"
	"strconv"

	"github.com/sirupsen/logrus"
)

func NewError(log *logrus.Logger) msg.I {
	emsg := new(msgLogError)
	emsg.msgLog.setMsgLog(log, 0)
	return emsg
}

type msgLog struct {
	logger *logrus.Logger
	depth  int
}

func (ml msgLog) setMsgLog(logger *logrus.Logger, depth int) {
	ml.logger = logger
	ml.depth = depth
}

type msgLogError struct {
	msgLog
}

func (ml msgLogError) P(msg string, val ...map[string]interface{}) {
	_, pkgnm, lineno, _ := runtime.Caller(ml.depth + 3)
	lentry := logrus.NewEntry(ml.logger)
	lentry = lentry.WithField("file", pkgnm+":"+strconv.Itoa(lineno))
	for _, mp := range val {
		for k, v := range mp {
			lentry = lentry.WithField(k, v)
		}
	}
	lentry.Error(msg)
}
