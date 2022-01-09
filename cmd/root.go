package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var printerIP string

var rootCmd = &cobra.Command{
	Use:   "ultituner",
	Short: "UltiTuner - the Swiss army knife for Ultimaker printers",
	Long: `
UltiTuner is a helper tool to configure functions for network-enabled 
Ultimaker printers (S3, S5).

This tool uses ssh to connect to the printer, so you must first enable 
the "Developer Mode" in the printer menu.

Check the available commands below. 
You get more help for each command when you add --help to the command line.
More information can be found on GitHub https://github.com/SmithyAT/UltiTuner

DISCLAIMER: 
The software was created and tested with the utmost care. But if something goes wrong, 
neither Ultimaker nor I am responsible if the printer no longer works as usual or breaks.
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
