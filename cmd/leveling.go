package cmd

import (
	"fmt"
	"github.com/melbahja/goph"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var levelingCmd = &cobra.Command{
	Use:                   "leveling [on | off] -p <ip_address of the printer>",
	Args:                  cobra.MatchAll(cobra.OnlyValidArgs, cobra.MaximumNArgs(1)),
	ValidArgs:             []string{"on", "off"},
	DisableFlagsInUseLine: true,
	Short:                 "Enable or disable the active leveling on the printer",
	Long: `Enable or disable the active leveling on the printer.

You need to enable the "Developer Mode" on your printer, that ultituner is able to connect to your printer.

When you call the command leveling without the on or off argument, the current status of active leveling is shown.`,
	Run: func(cmd *cobra.Command, args []string) {

		client := sshConnect()
		defer func(client *goph.Client) {
			err := client.Close()
			if err != nil {
				fmt.Println("ERROR: Something went wrong - unable to complete the action.")
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

		if len(args) == 1 && args[0] == "on" {
			sshCmd(client, "sed -i 's/self.__probing_mode = ProbeMode.NEVER/self.__probing_mode = ProbeMode.DETAILED/g' /usr/share/griffin/griffin/printer/procedures/pre_and_post_print/auto_bed_level_adjust/alignZAxisProcedure.py")
			fmt.Println("Active Leveling turned ON")
		} else if len(args) == 1 && args[0] == "off" {
			sshCmd(client, "sed -i 's/self.__probing_mode = ProbeMode.DETAILED/self.__probing_mode = ProbeMode.NEVER/g' /usr/share/griffin/griffin/printer/procedures/pre_and_post_print/auto_bed_level_adjust/alignZAxisProcedure.py")
			fmt.Println("Active Leveling turned OFF")
		} else {
			result := sshCmd(client, "grep 'self.__probing_mode = ProbeMode.' /usr/share/griffin/griffin/printer/procedures/pre_and_post_print/auto_bed_level_adjust/alignZAxisProcedure.py")
			if strings.Contains(result, "DETAILED") {
				fmt.Println("Active Leveling is currently ENABLED")
			} else {
				fmt.Println("Active Leveling is currently DISABLED")
			}
		}
		fmt.Println()

	},
}

func init() {
	rootCmd.AddCommand(levelingCmd)

}
