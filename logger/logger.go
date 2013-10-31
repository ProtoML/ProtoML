package logger

type ProtoMLLogger interface {
	SetPrefix(string)
    Print(args ...interface{})
    Printf(format string, args ...interface{})
    Fatal(args ...interface{})
    Fatalf(format string, args ...interface{})
}


var Debug bool = false
var Logger ProtoMLLogger

func setPrefix(logName string) {
	if Logger != nil {
		Logger.SetPrefix("[" + logName + "] ")
	}
}

func logOut(logName string, tag string, format string, v ...interface{}) {
	if Logger != nil {
		setPrefix(logName)
		Logger.Printf("{" + tag + "}\t" + format + "\n", v...)
	}
}

func logFatal(logName string, tag string, err error, format string, v ...interface{}) {
	if Logger != nil {
		setPrefix(logName)
		ev := append(v,err)
		Logger.Fatalf("{" + tag + "}\t" + format + "\n\terror: %v\n", ev...)
	}
}

func LogInfo(tag string, format string, v ...interface{}) {
	logOut("INFO",tag,format,v...)
}

func LogDebug(tag string, format string, v ...interface{}) {
	if Debug {
		logOut("DEBUG",tag,format,v...)
	}
}

func LogFatal(tag string, err error, format string, v ...interface{}) {
	logFatal("FATAL",tag,err,format,v...)
}
