package senger

import (
	"time"

	sentry "github.com/getsentry/raven-go"
)

func newSentryClient(dsn string) (*sentry.Client, error) {
	return sentry.New(dsn)
}

func (l *Logger) send(msg string, lvl *LoggerLevel) {
	if l.sentry != nil {
		packet := sentry.NewPacket(msg)

		switch lvl {
		case DebugLevel:
			packet.Level = sentry.DEBUG
		case InfoLevel:
			packet.Level = sentry.INFO
		case WarnLevel:
			packet.Level = sentry.WARNING
		case ErrorLevel:
			packet.Level = sentry.ERROR

			st := sentry.NewStacktrace(2, 2, nil)
			packet.Culprit = st.Culprit()
			packet.Interfaces = append(packet.Interfaces, st)
		case FatalLevel:
			packet.Level = sentry.FATAL

			st := sentry.NewStacktrace(2, 2, nil)
			packet.Culprit = st.Culprit()
			packet.Interfaces = append(packet.Interfaces, st)
		}

		packet.Timestamp = sentry.Timestamp(time.Now())

		if lvl.lvl >= ErrorLevel.lvl {
			_, ch := l.sentry.Capture(packet, nil)
			<-ch
		} else {
			l.sentry.Capture(packet, nil)
		}
	}
}
