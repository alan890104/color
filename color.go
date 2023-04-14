package color

import (
	"fmt"
)

const (
	_Black   = "\033[30m"
	_Red     = "\033[31m"
	_Green   = "\033[32m"
	_Yellow  = "\033[33m"
	_Blue    = "\033[34m"
	_Magenta = "\033[35m"
	_Cyan    = "\033[36m"
	_White   = "\033[37m"
	_Reset   = "\033[0m"
)

// Colorize 會將給定的字串加上顏色
func Colorize(color, text string) string {
	return color + text + _Reset
}

// Colorizef 會將給定的字串加上顏色
func Colorizef(color, format string, a ...interface{}) string {
	return color + fmt.Sprintf(format, a...) + _Reset
}

// Blackf 會將給定的字串加上黑色
func Blackf(format string, a ...interface{}) string {
	return Colorizef(_Black, format, a...)
}

// Redf 會將給定的字串加上紅色
func Redf(format string, a ...interface{}) string {
	return Colorizef(_Red, format, a...)
}

// Greenf 會將給定的字串加上綠色
func Greenf(format string, a ...interface{}) string {
	return Colorizef(_Green, format, a...)
}

// Yellowf 會將給定的字串加上黃色
func Yellowf(format string, a ...interface{}) string {
	return Colorizef(_Yellow, format, a...)
}

// Bluef 會將給定的字串加上藍色
func Bluef(format string, a ...interface{}) string {
	return Colorizef(_Blue, format, a...)
}

// Magentaf 會將給定的字串加上紫色
func Magentaf(format string, a ...interface{}) string {
	return Colorizef(_Magenta, format, a...)
}

// Cyanf 會將給定的字串加上青色
func Cyanf(format string, a ...interface{}) string {
	return Colorizef(_Cyan, format, a...)
}

// Whitef 會將給定的字串加上白色
func Whitef(format string, a ...interface{}) string {
	return Colorizef(_White, format, a...)
}

// Black 會將給定的字串加上黑色
func Black(text string) string {
	return Colorize(_Black, text)
}

// Red 會將給定的字串加上紅色
func Red(text string) string {
	return Colorize(_Red, text)
}

// Green 會將給定的字串加上綠色
func Green(text string) string {
	return Colorize(_Green, text)
}

// Yellow 會將給定的字串加上黃色
func Yellow(text string) string {
	return Colorize(_Yellow, text)
}

// Blue 會將給定的字串加上藍色
func Blue(text string) string {
	return Colorize(_Blue, text)
}

// Magenta 會將給定的字串加上紫色
func Magenta(text string) string {
	return Colorize(_Magenta, text)
}

// Cyan 會將給定的字串加上青色
func Cyan(text string) string {
	return Colorize(_Cyan, text)
}

// White 會將給定的字串加上白色
func White(text string) string {
	return Colorize(_White, text)
}

// Success 會將給定的字串加上綠色
func Success(text string) string {
	return Colorize(_Green, text)
}

// Error 會將給定的字串加上紅色
func Error(text string) string {
	return Colorize(_Red, text)
}

// Warning 會將給定的字串加上黃色
func Warning(text string) string {
	return Colorize(_Yellow, text)
}

// Info 會將給定的字串加上藍色
func Info(text string) string {
	return Colorize(_Blue, text)
}

// Debug 會將給定的字串加上青色
func Debug(text string) string {
	return Colorize(_Cyan, text)
}

// Fatal 會將給定的字串加上紫色
func Fatal(text string) string {
	return Colorize(_Magenta, text)
}

// Successf 會將給定的字串加上綠色
func Successf(format string, a ...interface{}) string {
	return Colorizef(_Green, format, a...)
}

// Errorf 會將給定的字串加上紅色
func Errorf(format string, a ...interface{}) string {
	return Colorizef(_Red, format, a...)
}

// Warningf 會將給定的字串加上黃色
func Warningf(format string, a ...interface{}) string {
	return Colorizef(_Yellow, format, a...)
}

// Infof 會將給定的字串加上藍色
func Infof(format string, a ...interface{}) string {
	return Colorizef(_Blue, format, a...)
}

// Debugf 會將給定的字串加上青色
func Debugf(format string, a ...interface{}) string {
	return Colorizef(_Cyan, format, a...)
}

// Fatalf 會將給定的字串加上紫色
func Fatalf(format string, a ...interface{}) string {
	return Colorizef(_Magenta, format, a...)
}
