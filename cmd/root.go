package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var printerIP string

var rootCmd = &cobra.Command{
	Version: "0.5",
	Use:     "ultituner",
	Short:   "UltiTuner, the small helper tool",
	Long: `UltiTuner - Written by Christian Schmied aka Smithy (https://github.com/SmithyAT/UltiTuner)

UltiTuner is a small helper tool to configure functions for Ultimaker S-Line printers, which are not available via the printers menu.
This tool uses ssh to connect to the printer, so you need to first enable the "Developer Mode" in the printer menu.

Check the available commands below. You get more help for each command when you add --help to the command line.
More information can be found on GitHub https://github.com/SmithyAT/UltiTuner

DISCLAIMER: 
The software was created and tested with utmost care. But since non-official methods are used to change the firmware, 
it could happen under certain circumstances that the printer no longer works as usual. Neither I personally nor Ultimaker is responsible for this. 
**USE THIS TOOL AT YOUR OWN RISK AND RESPONSIBILITY**`,
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
