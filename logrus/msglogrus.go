package msglogrus

import (
	"runtime"
	"strconv"

	"github.com/sirupsen/logrus"
)

func NewError(log *logrus.Logger) (errmsg msg.I) {
	errmsg = new(msglogError)
	return
}

func newMsgLog(logger *logrus.Logger, depth int) (ml *msgLog) {
	ml = new(msgLog)
	ml.msgLog = logger
	ml.depth = depth
	return
}

type msglog struct {
	log   *logrus.Logger
	depth int
}

type msglogError struct {
	msglog
}

func (ml msgLogError) P(msg string, val ...map[string]interface{}) {
	_, pkgnm, lineno, _ := runtime.Caller(depth)
	lentry := NewEntry(ml.msgLog)
	lentry = lentry.WithField("file", pkgnm+":"+strconv.Itoa(lineno))
	for k, v := range val {
		lentry = lentry.WithFields(k, v)
	}
	lentry.Error(msg)
}
