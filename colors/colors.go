package colors

import (
	"fmt"
	"github.com/gookit/color"
)

var red = color.FgRed.Render
var green = color.FgGreen.Render
var blue = color.FgBlue.Render
var magenta = color.Magenta.Render

// PrintSuccess implementation of Println colors output when program execution success
func PrintSuccess(args ...interface{}) {

	//fmt.Printf("%s %s\n", blue("[+]"), green(args...))
	color.HEXStyle("#00e500").Printf("[+] ")
	color.HEXStyle("#70f3ff").Printf("%s\n", args...)
}

// PrintSuccessf implementation of Printf colors output when program execution success
func PrintSuccessf(format string, args ...interface{}) {
	formats := fmt.Sprintf(format, args...)
	color.HEXStyle("#00e500").Printf("[+] ")
	color.HEXStyle("#70f3ff").Printf("%s\n", formats)
}

// PrintError implementation of Println colors output when program execution error
func PrintError(args ...interface{}) {
	color.HEXStyle("#ff0097").Printf("[-] ")
	color.HEXStyle("#ff2121").Printf("%s\n", args...)
}

// PrintErrorf implementation of Printf colors output when program execution error
func PrintErrorf(format string, args ...interface{}) {
	formats := fmt.Sprintf(format, args...)
	color.HEXStyle("#ff0097").Printf("[-] ")
	color.HEXStyle("#ff2121").Printf("%s\n", formats)
}
