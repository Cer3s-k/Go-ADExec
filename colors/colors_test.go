package colors

import (
	"fmt"
	"github.com/gookit/color"
	"testing"
)

func TestColor(t *testing.T) {
	color.Red.Println("message")
	color.Cyan.Println("message")
	color.Gray.Println("message")
	color.Blue.Println("message")
	color.Black.Println("message")
	color.Green.Println("message")
	color.White.Println("message")
	color.Yellow.Println("message")
	color.Magenta.Println("message")
	fmt.Println()
	color.Bold.Println("message")
	color.Normal.Println("message")
	fmt.Println()
	color.LightRed.Println("message")
	color.LightCyan.Println("message")
	color.LightBlue.Println("message")
	color.LightGreen.Println("message")
	color.LightWhite.Println("message")
	color.LightYellow.Println("message")
	color.LightMagenta.Println("message")
	fmt.Println()
	color.HiRed.Println("message")
	color.HiCyan.Println("message")
	color.HiBlue.Println("message")
	color.HiGreen.Println("message")
	color.HiWhite.Println("message")
	color.HiYellow.Println("message")
	color.HiMagenta.Println("message")
	fmt.Println()

	PrintSuccess("this is a success demo 111")
	PrintError("this is a error demo 222")
	PrintSuccessf("this is a format %s demo %s", "success111", "success222")
	PrintErrorf("this is a format %s demo %s", "error111", "error222")

}
