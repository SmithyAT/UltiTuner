package cmd

import (
	"fmt"
	"github.com/melbahja/goph"
	"github.com/spf13/cobra"
)

var rebootCmd = &cobra.Command{
	Use:                   "reboot -p <ip_address of the printer>",
	Args:                  cobra.MatchAll(cobra.ExactArgs(0)),
	DisableFlagsInUseLine: true,
	Short:                 "Do a full reboot of the printer",
	Long: `UltiTuner - Reboot

The "reboot" command is used to reboot the linux system of the printer.`,

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

		// Reboot the printer
		fmt.Print("Rebooting the printer.....")
		sshCmd(client, "shutdown -r now")
		fmt.Println("done, printer is now starting up again")
		fmt.Println()

	},
}

func init() {
	rootCmd.AddCommand(rebootCmd)

}
