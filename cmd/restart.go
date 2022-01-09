package cmd

import (
	"fmt"
	"github.com/melbahja/goph"
	"github.com/spf13/cobra"
)

var restartCmd = &cobra.Command{
	Use:                   "restart -p <ip_address of the printer>",
	Args:                  cobra.MatchAll(cobra.ExactArgs(0)),
	DisableFlagsInUseLine: true,
	Short:                 "Restart the griffin printer service",
	Long: `
The "restart" command is used to restart the griffin printer service.
It is enough to restart the printer service to take effect configuration 
changes instead of a full reboot.`,

	Run: func(cmd *cobra.Command, args []string) {

		// Connect to the printer
		fmt.Print("Connecting to the printer " + printerIP + ".....")
		client := sshConnect()
		defer func(client *goph.Client) {
			err := client.Close()
			if err != nil {
				fmt.Println("ERROR: Something went wrong - unable to complete the action")
			}
		}(client)
		fmt.Println("done, connected")

		// Print the properties of the printer
		printPrinterProperties(client)

		// Restart the griffin service
		restartGriffin(client)
		fmt.Println()

	},
}

func init() {
	rootCmd.AddCommand(restartCmd)

}
