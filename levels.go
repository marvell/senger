package senger

type LoggerLevel struct {
	msg   string
	lvl   uint8
	color []byte
}

var (
	DebugLevel = &LoggerLevel{"[DBG]", 1, bGreen}
	InfoLevel  = &LoggerLevel{"[INF]", 2, bCyan}
	WarnLevel  = &LoggerLevel{"[WRN]", 3, bYellow}
	ErrorLevel = &LoggerLevel{"[ERR]", 4, bRed}
	FatalLevel = &LoggerLevel{"[FAT]", 5, bWhiteOnRed}
)

func (lvl *LoggerLevel) String() string {
	return withColor(lvl.color, lvl.msg)
}
