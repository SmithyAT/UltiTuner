package cmd

import (
	"fmt"
	"github.com/melbahja/goph"
	"github.com/spf13/cobra"
)

var webcamCmd = &cobra.Command{
	Use:                   "webcam -p <ip_address of the printer>",
	Args:                  cobra.MatchAll(cobra.MaximumNArgs(1)),
	DisableFlagsInUseLine: true,
	Short:                 "Enable additional webcams",
	Long: `
The "webcam" command modifies the internal firewall and opens addtional ports 8081 to 8083,
to be able to connect and access additional webcams. 

Note that you must power cycle the printer for this change to take effect 
or use the "-r" flag to restart automatically.`,

	Run: func(cmd *cobra.Command, args []string) {
		restartFlag, _ := cmd.Flags().GetBool("restart")

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

		// Check if this command is compatible with the connected printer. Must return a string true if compatible and false if not (...&& echo true || echo false)
		checkCompatibility(client, "egrep -q 'add rule ip filter INPUT ip protocol 6 ct state new tcp dport 8080 accept' /usr/share/griffin/griffin/network/firewall/nftables_firewall_off.conf && echo true || echo false")

		// Change the firewall rule
		fmt.Print("Changing firewall rules.....")
		sshCmd(client, "sed -i 's/add rule ip filter INPUT ip protocol 6 ct state new tcp dport 8080 accept/add rule ip filter INPUT ip protocol 6 ct state new tcp dport { 8080-8083 } accept/' /usr/share/griffin/griffin/network/firewall/nftables_firewall_off.conf")
		fmt.Println("done")
		if restartFlag {
			restartGriffin(client)
		}

	},
}

func init() {
	rootCmd.AddCommand(webcamCmd)
	webcamCmd.Flags().BoolP("restart", "r", false, "Restart the printer service after the change")
}
