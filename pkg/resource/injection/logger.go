package injection

import "github.com/sirupsen/logrus"

type Logger struct {
	*logrus.Logger
}

func NewLogger() Logger {
	l := logrus.New()
	logFormat := &logrus.TextFormatter{
		ForceColors:     false,
		DisableColors:   false,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	}

	l.SetFormatter(logFormat)

	return Logger{
		Logger: l,
	}
}
