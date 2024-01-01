package colors

import (
	"fmt"
	"github.com/gookit/color"
)

// SuccessPrintln success output when program execution
func SuccessPrintln(args ...interface{}) {
	formats := fmt.Sprint(args...)
	color.HEXStyle("#00e500").Printf("[+] ")
	color.HEXStyle("#70f3ff").Printf("%s\n", formats)
}

// SuccessPrintf success output when program execution
func SuccessPrintf(format string, args ...interface{}) {
	formats := fmt.Sprintf(format, args...)
	color.HEXStyle("#00e500").Printf("[+] ")
	color.HEXStyle("#70f3ff").Printf("%s", formats)
}

// ErrorPrintln error output when program execution
func ErrorPrintln(args ...interface{}) {
	formats := fmt.Sprint(args...)
	color.HEXStyle("#ff0097").Printf("[-] ")
	color.HEXStyle("#ff2121").Printf("%s\n", formats)
}

// ErrorPrintf error output when program execution
func ErrorPrintf(format string, args ...interface{}) {
	formats := fmt.Sprintf(format, args...)
	color.HEXStyle("#ff0097").Printf("[-] ")
	color.HEXStyle("#ff2121").Printf("%s", formats)
}

// InfoPrintln information output when program execution
func InfoPrintln(args ...interface{}) {
	formats := fmt.Sprint(args...)
	color.HEXStyle("#3eede7").Printf("[*] ")
	color.HEXStyle("#3de1ad").Printf("%s\n", formats)
}

// InfoPrintf information output when program execution
func InfoPrintf(format string, args ...interface{}) {
	formats := fmt.Sprintf(format, args...)
	color.HEXStyle("#3eede7").Printf("[*] ")
	color.HEXStyle("#3de1ad").Printf("%s", formats)
}

// NormalPrintln success output when program execution
func NormalPrintln(args ...interface{}) {
	formats := fmt.Sprint(args...)
	color.HEXStyle("#70f3ff").Printf("%s\n", formats)
}

// NormalPrintf success output when program execution
func NormalPrintf(format string, args ...interface{}) {
	formats := fmt.Sprintf(format, args...)
	color.HEXStyle("#70f3ff").Printf("%s", formats)
}
