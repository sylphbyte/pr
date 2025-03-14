package pr

import (
	"encoding/json"
	"fmt"
	C "github.com/fatih/color"
	"os"
	"strings"
	"time"
)

var (
	enabled bool
	filter  = C.New(C.BgYellow, C.Bold)
)

func Enabled() bool {
	return enabled
}

func Disable() {
	enabled = false
}

func Enable() {
	enabled = true
}

func PrintBlack(format string, a ...interface{}) {
	base(false, "black", format, a...)
}

func PrintRed(format string, a ...interface{}) {
	base(false, "red", format, a...)
}

func PrintGreen(format string, a ...interface{}) {
	base(false, "green", format, a...)
}

func PrintYellow(format string, a ...interface{}) {
	base(false, "yellow", format, a...)
}

func PrintBlue(format string, a ...interface{}) {
	base(false, "blue", format, a...)
}

func PrintMagenta(format string, a ...interface{}) {
	base(false, "magenta", format, a...)
}

func PrintCyan(format string, a ...interface{}) {
	base(false, "cyan", format, a...)
}

func PrintWhite(format string, a ...interface{}) {
	base(false, "white", format, a...)
}

func SprintBlack(format string, a ...interface{}) string {
	return baseString("black", format, a...)
}

func SprintRed(format string, a ...interface{}) string {
	return baseString("red", format, a...)
}

func SprintGreen(format string, a ...interface{}) string {
	return baseString("green", format, a...)
}

func SprintYellow(format string, a ...interface{}) string {
	return baseString("yellow", format, a...)
}

func SprintBlue(format string, a ...interface{}) string {
	return baseString("blue", format, a...)
}

func SprintMagenta(format string, a ...interface{}) string {
	return baseString("magenta", format, a...)
}

func SprintCyan(format string, a ...interface{}) string {
	return baseString("cyan", format, a...)
}

func SprintWhite(format string, a ...interface{}) string {
	return baseString("white", format, a...)
}

func base(force bool, kind string, format string, a ...interface{}) {
	if !force && !enabled {
		return
	}

	switch kind {
	case "black":
		C.Black(format, a...)
		break
	case "red":
		C.Red(format, a...)
		break
	case "green":
		C.Green(format, a...)
		break
	case "yellow":
		C.Yellow(format, a...)
		break
	case "blue":
		C.Blue(format, a...)
		break
	case "magenta":
		C.Magenta(format, a...)
		break
	case "cyan":
		C.Cyan(format, a...)
		break
	case "white":
		C.White(format, a...)
		break
	}

	C.Unset()
}

func Black(format string, a ...interface{}) {
	base(false, "black", format, a...)
}

func Red(format string, a ...interface{}) {
	base(false, "red", format, a...)
}

func Green(format string, a ...interface{}) {
	base(false, "green", format, a...)
}

func Yellow(format string, a ...interface{}) {
	base(false, "yellow", format, a...)
}

func Blue(format string, a ...interface{}) {
	base(false, "blue", format, a...)
}

func Magenta(format string, a ...interface{}) {
	base(false, "magenta", format, a...)
}

func Cyan(format string, a ...interface{}) {
	base(false, "cyan", format, a...)
}

func White(format string, a ...interface{}) {
	base(false, "white", format, a...)
}

func ForceBlack(format string, a ...interface{}) {
	base(true, "black", format, a...)
}

func ForceRed(format string, a ...interface{}) {
	base(true, "red", format, a...)
}

func ForceGreen(format string, a ...interface{}) {
	base(true, "green", format, a...)
}

func ForceYellow(format string, a ...interface{}) {
	base(true, "yellow", format, a...)
}

func ForceBlue(format string, a ...interface{}) {
	base(true, "blue", format, a...)
}

func ForceMagenta(format string, a ...interface{}) {
	base(true, "magenta", format, a...)
}

func ForceCyan(format string, a ...interface{}) {
	base(true, "cyan", format, a...)
}

func ForceWhite(format string, a ...interface{}) {
	base(true, "white", format, a...)
}

func SBlack(format string, a ...interface{}) string {
	return baseString("black", format, a...)
}

func SRed(format string, a ...interface{}) string {
	return baseString("red", format, a...)
}

func SGreen(format string, a ...interface{}) string {
	return baseString("green", format, a...)
}

func SYellow(format string, a ...interface{}) string {
	return baseString("yellow", format, a...)
}

func SBlue(format string, a ...interface{}) string {
	return baseString("blue", format, a...)
}

func SMagenta(format string, a ...interface{}) string {
	return baseString("magenta", format, a...)
}

func SCyan(format string, a ...interface{}) string {
	return baseString("cyan", format, a...)
}

func SWhite(format string, a ...interface{}) string {
	return baseString("white", format, a...)
}

func baseString(kind string, format string, a ...interface{}) string {
	if !enabled {
		return ""
	}

	switch kind {
	case "black":
		return C.BlackString(format, a...)
	case "red":
		return C.RedString(format, a...)
	case "green":
		return C.GreenString(format, a...)
	case "yellow":
		return C.YellowString(format, a...)
	case "blue":
		return C.BlueString(format, a...)
	case "magenta":
		return C.MagentaString(format, a...)
	case "cyan":
		return C.CyanString(format, a...)
	case "white":
		return C.WhiteString(format, a...)
	}

	return ""
}

func Panic(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	panic(s)
}
func RedPanic(format string, a ...interface{}) {
	s := SRed(format, a...)
	panic(s)
}

// 系统类: Cyan
// 错误类: Red
// 提示类: Magenta
// 强提示: Yellow
// 成功: Green

func prefix(mark, format string) string {
	ts := time.Now().Format("2006-01-02 15:04:05.000")
	mark = fmt.Sprintf("[%s]", mark)
	return fmt.Sprintf("%23s %9s: %s", ts, mark, format)

}

func Mark(format string, a ...interface{}) {
	format = prefix("mark", format)
	ForceRed(format, a...)
	os.Exit(0)
}

func System(format string, a ...interface{}) {
	format = prefix("system", format)
	ForceCyan(format, a...)
}

func Error(format string, a ...interface{}) {
	format = prefix("error", format)
	ForceRed(format, a...)
}

func Notice(format string, a ...interface{}) {
	format = prefix("notice", format)
	Magenta(format, a...)
}

func Warning(format string, a ...interface{}) {
	format = prefix("warning", format)
	Yellow(format, a...)
}

func Info(format string, a ...interface{}) {
	format = prefix("info", format)
	Green(format, a...)
}

func Json(obj any) {
	if !enabled {
		return
	}
	js, _ := json.Marshal(obj)
	base(false, "red", string(js))
}

func JsonExit(obj any) {
	if !enabled {
		return
	}
	js, _ := json.Marshal(obj)
	base(false, "red", string(js))
	os.Exit(0)
}

func Filter(format string, a ...interface{}) {
	if !enabled {
		return
	}

	if strings.HasSuffix(format, "\n") {
		format = strings.TrimSuffix(format, "\n")
	}

	filter = C.New(C.BgYellow, C.Bold)
	format = prefix("filter", format)
	_, _ = filter.Printf(format, a...)
	fmt.Println()
}

func Request(format string, a ...interface{}) {
	if !enabled {
		return
	}

	if strings.HasSuffix(format, "\n") {
		format = strings.TrimSuffix(format, "\n")
	}

	filter = C.New(C.BgCyan, C.Bold)
	format = prefix("request", format)
	_, _ = filter.Printf(format, a...)
	fmt.Println()
}

func SRequest(format string, a ...interface{}) string {
	if !enabled {
		return ""
	}

	if strings.HasSuffix(format, "\n") {
		format = strings.TrimSuffix(format, "\n")
	}

	filter = C.New(C.BgCyan, C.Bold)
	format = prefix("request", format)
	return filter.Sprintf(format, a...)
}

func RequestError(format string, a ...interface{}) {
	if !enabled {
		return
	}

	if strings.HasSuffix(format, "\n") {
		format = strings.TrimSuffix(format, "\n")
	}

	filter = C.New(C.BgRed, C.Bold)
	format = prefix("error", format)
	_, _ = filter.Printf(format, a...)
	fmt.Println()
}
