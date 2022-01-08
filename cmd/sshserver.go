package cmd

import (
	"fmt"
	"github.com/melbahja/goph"
	"github.com/spf13/cobra"
)

var sshserverCmd = &cobra.Command{
	Use:                   "sshserver [on | off] -p <ip_address of the printer>",
	Args:                  cobra.MatchAll(cobra.OnlyValidArgs, cobra.MaximumNArgs(1)),
	ValidArgs:             []string{"on", "off"},
	DisableFlagsInUseLine: true,
	Short:                 "Enable or disable the ssh daemon on the printer",
	Long: `UltiTuner - Enable/Disable the SSH Daemon

The "sshserver" command permanently enables the ssh daemon on the printer, independent of developer mode.
It can be very helpful if you work with the firmware because the developer mode is just a flag, and the firmware starts the ssh daemon during startup.
If you have an error in a file and the startup routine throws an exception, the printer is bricked because you can no longer shh into the printer and fix the problem.
The ssh daemon is always available when enabling it permanently as the system default, regardless of the startup routines of the printer firmware.

Remember that disabling the developer mode won't disable the ssh daemon anymore. If you want to turn it off again, you have to use UltiTuner again.`,

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

		// Check which ssh daemon is installed - depends on the firmware
		sshDaemon := sshCmd(client, "systemctl list-units | grep -q dropbear && echo dropbear || echo ssh")

		// Enable ssh daemon
		if len(args) == 1 && args[0] == "on" {
			fmt.Print("Enabling " + sshDaemon + " daemon.....")
			sshCmd(client, "systemctl enable "+sshDaemon)
			fmt.Println("done")

			//Disable ssh daemon
		} else if len(args) == 1 && args[0] == "off" {
			fmt.Print("Disabling " + sshDaemon + " daemon.....")
			sshCmd(client, "systemctl disable "+sshDaemon)
			fmt.Println("done")
			fmt.Println("Remember also to disable the developer mode if you don't want someone to have SSH access.")

			// Fetch the status of the ssh daemon
		} else {
			fmt.Print("Checking if " + sshDaemon + " is configured as a system service.....")
			if sshDaemon == "ssh" {
				result := sshCmd(client, "systemctl is-enabled ssh")
				if result == "enabled" {
					fmt.Println("done, enabled")
				} else {
					fmt.Println("done, disabled")
				}
			} else {
				result := sshCmd(client, "test -f /etc/rc5.d/S01dropbear && echo enabled || echo disabled")
				fmt.Println("done, " + result)
			}
		}

		fmt.Println()

	},
}

func init() {
	rootCmd.AddCommand(sshserverCmd)
}
