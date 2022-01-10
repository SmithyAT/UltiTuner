package cmd

import (
	"fmt"
	"github.com/melbahja/goph"
	"github.com/spf13/cobra"
	"strings"
)

var clearbedCmd = &cobra.Command{
	Use:                   "clearbed -p <ip_address of the printer>",
	Args:                  cobra.MatchAll(cobra.ExactArgs(0)),
	DisableFlagsInUseLine: true,
	Short:                 "Clear the print bed state",
	Long: `
When you abort a print job in an early stage, so during the heat-up phase, 
before the print actually starts, the printer is asking you after the abort 
if you want to retry the job or not. In the DigitalFactory you see a message
like "Aborted, Awaiting Clean-Up". You cannot start another print job until 
someone presses the button on the printer's display.

Another case is when your print job has finished, you removed the object,
but finally forgot to click the button that you have cleared the print bed.

With that command, you can clear the message on the printer's display, 
and the printer is ready again to accept new print jobs.`,

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

		// Clear the printer state
		fmt.Print("Clearing the print bed state.....")
		result := sshCmd(client, "dbus-send --system --dest=nl.ultimaker.printer --type=method_call --print-reply /nl/ultimaker/printer nl.ultimaker.messageProcedure string:\"REPRINT_AFTER_ABORT\" string:\"ABORT\"")
		if strings.Contains(result, "boolean true") {
			fmt.Println("done, printer is available again")
		} else {
			result := sshCmd(client, "dbus-send --system --dest=nl.ultimaker.printer --type=method_call --print-reply /nl/ultimaker/printer nl.ultimaker.messageProcedure string:\"PRINT\" string:\"PRINTER_CLEANED\"")
			if strings.Contains(result, "boolean true") {
				fmt.Println("done, printer is available again")
			} else {
				fmt.Println("failed, maybe the printer is in a state, which could not be handled")
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(clearbedCmd)

}
