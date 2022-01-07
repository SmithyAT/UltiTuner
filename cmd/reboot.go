package cmd

import (
	"fmt"
	"github.com/melbahja/goph"
	"github.com/spf13/cobra"
	"os"
)

var rebootCmd = &cobra.Command{
	Use:                   "reboot -p <ip_address of the printer>",
	Args:                  cobra.MatchAll(cobra.ExactArgs(0)),
	DisableFlagsInUseLine: true,
	Short:                 "Do a full reboot of the printer",
	Long: `UltiTuner

The "reboot" command is used to reboot the linux system of the printer.`,
	Run: func(cmd *cobra.Command, args []string) {

		client := sshConnect()
		defer func(client *goph.Client) {
			err := client.Close()
			if err != nil {
				fmt.Println("ERROR: Something went wrong - unable to complete the action")
			}
		}(client)
		fmt.Println(getPrinterProperties(client))
		if checkVersion(client, 5) {
			fmt.Println("supported by ultituner: YES")
			fmt.Println()
		} else {
			fmt.Println("supported by ultituner: NO, your printer or firmware is not supported")
			os.Exit(1)
		}

		fmt.Print("Rebooting...")
		sshCmd(client, "shutdown -r now")
		fmt.Println("DONE")
		fmt.Println("Printer is now starting up again.")
		fmt.Println()

	},
}

func init() {
	rootCmd.AddCommand(rebootCmd)

}
