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

func ParseLevel(lvl string) *LoggerLevel {
	switch {
	case lvl == "DEBUG" || lvl == "DBG":
		return DebugLevel
	case lvl == "INFO" || lvl == "INF":
		return InfoLevel
	case lvl == "WARNING" || lvl == "WRN" || lvl == "WARN":
		return WarnLevel
	case lvl == "ERROR" || lvl == "ERR":
		return ErrorLevel
	case lvl == "FATAL" || lvl == "FAT":
		return FatalLevel
	}

	return nil
}
