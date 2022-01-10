package cmd

import (
	"fmt"
	"github.com/melbahja/goph"
	"github.com/spf13/cobra"
	"strings"
)

var clearabortCmd = &cobra.Command{
	Use:                   "clearabort -p <ip_address of the printer>",
	Args:                  cobra.MatchAll(cobra.ExactArgs(0)),
	DisableFlagsInUseLine: true,
	Short:                 "Clear the waiting state after aborting a print job",
	Long: `
When you abort a print job in an early stage, so during the heat-up phase, 
before the print actually starts, the printer is asking you after the abort 
if you want to retry the job or not. In the DigitalFactory you see a message
like "Aborted, Awaiting Clean-Up". You cannot start another print job until 
someone presses the button on the printer's display.

With that command, you can clear the message on the printer's display, 
and the printer is ready again to accept new print jobs.

Notice that it is not working when the printer has already started to print.`,

	Run: func(cmd *cobra.Command, args []string) {

		// Connect to the printer
		fmt.Print("Connecting to the printer " + printerIP + ".....")
		client := sshConnect()
		defer func(client *goph.Client) {
			err := client.Close()
			if err != nil {
				fmt.Println("ERROR: Something went wrong - unable to close the ssh connection")
			}
		}(client)
		fmt.Println("done, connected")

		// Print the properties of the printer
		printPrinterProperties(client)

		// Reboot the printer
		fmt.Print("Clearing the abort message.....")
		result := sshCmd(client, "dbus-send --system --dest=nl.ultimaker.printer --type=method_call --print-reply /nl/ultimaker/printer nl.ultimaker.messageProcedure string:\"REPRINT_AFTER_ABORT\" string:\"ABORT\"")
		if strings.Contains(result, "boolean true") {
			fmt.Println("done, printer is available again")
		} else {
			fmt.Println("failed, maybe the printer is in a different state?")
		}

	},
}

func init() {
	rootCmd.AddCommand(clearabortCmd)

}
