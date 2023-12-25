package main

import (
	"Go-ADExec/cmd"
	"github.com/gookit/color"
)

func main() {
	var logo string = `
     ______            ___    ____  ______              
    / ____/___        /   |  / __ \/ ____/  _____  _____
   / / __/ __ \______/ /| | / / / / __/ | |/_/ _ \/ ___/
  / /_/ / /_/ /_____/ ___ |/ /_/ / /____>  </  __/ /__  
  \____/\____/     /_/  |_/_____/_____/_/|_|\___/\___/
`
	color.Cyan.Print(logo)
	cmd.Execute()
}
