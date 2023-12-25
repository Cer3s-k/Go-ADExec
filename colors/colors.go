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

	fmt.Printf("%s %s\n", blue("[+]"), green(args...))
}

// PrintSuccessf implementation of Printf colors output when program execution success
func PrintSuccessf(format string, args ...interface{}) {
	formats := fmt.Sprintf(format, args...)
	fmt.Printf("%s %s\n", blue("[+]"), green(formats))
}

// PrintError implementation of Println colors output when program execution error
func PrintError(args ...interface{}) {
	fmt.Printf("%s %s\n", red("[-]"), magenta(args...))
}

// PrintErrorf implementation of Printf colors output when program execution error
func PrintErrorf(format string, args ...interface{}) {
	formats := fmt.Sprintf(format, args...)
	fmt.Printf("%s %s\n", red("[-]"), magenta(formats))
}
