package logger

import (
	"fmt"
	"router-gostashlg-template/entities/app"

	"github.com/randyardiansyah25/gostashlg"
)

var Logger gostashlg.LoggerEngine

func PrintLogf(format string, message ...interface{}) {
	PrintLog(fmt.Sprintf(format, message...))
}

func PrintLog(message ...interface{}) {
	printOut(gostashlg.LOG, "sys_log", fmt.Sprint(message...))
}

func PrintWarnf(format string, message ...interface{}) {
	PrintWarn(fmt.Sprintf(format, message...))
}

func PrintWarn(message ...interface{}) {
	printOut(gostashlg.WARN, "sys_warn", fmt.Sprint(message...))
}

func PrintErrorf(format string, message ...interface{}) {
	PrintError(fmt.Sprintf(format, message...))
}

func PrintError(message ...interface{}) {
	printOut(gostashlg.ERROR, "sys_err", fmt.Sprint(message...))
}

func PrintInfof(format string, message ...interface{}) {
	PrintInfo(fmt.Sprintf(format, message...))
}

func PrintInfo(message ...interface{}) {
	printOut(gostashlg.INFO, "sys_info", fmt.Sprint(message...))
}

func printOut(level gostashlg.Level, event string, data interface{}) {
	field := gostashlg.NewFields().
		SetIdentifierName(app.Identifier).
		SetLevel(level).
		SetEvent(event).
		SetMessage("SYSTEM_OUT").
		SetData(data).
		Get()
	Logger.Write(field)
}

func PrintErrorWithData(event string, message string, data ...interface{}) {
	printWithData(gostashlg.ERROR, event, message, data...)
}

func PrintWarnWithData(event string, message string, data ...interface{}) {
	printWithData(gostashlg.WARN, event, message, data...)
}

func PrintInfoWithData(event string, message string, data ...interface{}) {
	printWithData(gostashlg.INFO, event, message, data...)
}

func printWithData(level gostashlg.Level, event string, message string, data ...interface{}) {
	field := gostashlg.NewFields().
		SetIdentifierName(app.Identifier).
		SetLevel(level).
		SetEvent(event).
		SetMessage(message).
		SetData(fmt.Sprint(data...)).
		Get()
	Logger.Write(field)
}

func PrintField(field gostashlg.Fields) {
	Logger.Write(field)
}
