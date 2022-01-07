package cmd

import (
	"fmt"
	"github.com/melbahja/goph"
	"github.com/spf13/cobra"
	"os"
)

var restartCmd = &cobra.Command{
	Use:                   "restart -p <ip_address of the printer>",
	Args:                  cobra.MatchAll(cobra.ExactArgs(0)),
	DisableFlagsInUseLine: true,
	Short:                 "Do a restart of the printer service",
	Long: `UltiTuner

The "restart" command is used to restart the griffin printer service.`,
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

		fmt.Print("Restarting griffin printer service...")
		sshCmd(client, "systemctl restart griffin.printer")
		fmt.Println("DONE")
		fmt.Println()

	},
}

func init() {
	rootCmd.AddCommand(restartCmd)

}
