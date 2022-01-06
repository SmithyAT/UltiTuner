package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var printerIP string

var rootCmd = &cobra.Command{
	Version: "20220106",
	Use:     "ultituner",
	Short:   "UltiTuner, the small helper tool",
	Long: `UltiTuner - Written by Christian Schmied aka Smithy
UltiTuner is a small helper tool, to configure and manage printer configurations which are not available via the printers menu. 

Disclaimer: This tool uses the ssh connection (developer mode) to connect to your printer and modifies directly some files of the firmware.
Use this tool at your own risk!`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&printerIP, "printer-ip", "p", "", "The ip address of your printer i.e. -p 192.168.0.23")
	_ = rootCmd.MarkPersistentFlagRequired("printer-ip")
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
