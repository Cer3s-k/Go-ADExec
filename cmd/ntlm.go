package cmd

import (
	"Go-ADExec/colors"
	"github.com/spf13/cobra"
)

// ntlm module
var ntlmCmd = &cobra.Command{
	Use:   "ntlm",
	Short: "ntlm the ActiveDirectory configuration",
	Long:  "ntlm the ActiveDirectory configuration",
	Run: func(cmd *cobra.Command, args []string) {
		colors.ErrorPrintln("Go-ADExec ntlm -h")
	},
}
