package colors

import (
	"fmt"
	"github.com/gookit/color"
)

// SuccessPrintln success output when program execution
func SuccessPrintln(args ...interface{}) {

	//fmt.Printf("%s %s\n", blue("[+]"), green(args...))
	color.HEXStyle("#00e500").Printf("[+] ")
	color.HEXStyle("#70f3ff").Printf("%s\n", args...)
}

// SuccessPrintf success output when program execution
func SuccessPrintf(format string, args ...interface{}) {
	formats := fmt.Sprintf(format, args...)
	color.HEXStyle("#00e500").Printf("[+] ")
	color.HEXStyle("#70f3ff").Printf("%s", formats)
}

// ErrorPrintln error output when program execution
func ErrorPrintln(args ...interface{}) {
	color.HEXStyle("#ff0097").Printf("[-] ")
	color.HEXStyle("#ff2121").Printf("%s\n", args...)
}

// ErrorPrintf error output when program execution
func ErrorPrintf(format string, args ...interface{}) {
	formats := fmt.Sprintf(format, args...)
	color.HEXStyle("#ff0097").Printf("[-] ")
	color.HEXStyle("#ff2121").Printf("%s", formats)
}

// InfoPrintln information output when program execution
func InfoPrintln(format string, args ...interface{}) {
	formats := fmt.Sprintf(format, args...)
	color.HEXStyle("#3eede7").Printf("[*] ")
	color.HEXStyle("#3de1ad").Printf("%s\n", formats)
}

// InfoPrintf information output when program execution
func InfoPrintf(format string, args ...interface{}) {
	formats := fmt.Sprintf(format, args...)
	color.HEXStyle("#3eede7").Printf("[*] ")
	color.HEXStyle("#3de1ad").Printf("%s", formats)
}
