package cmd

import (
	"fmt"
	"github.com/melbahja/goph"
	"github.com/spf13/cobra"
	"strings"
)

var levelingCmd = &cobra.Command{
	Use:                   "leveling [on | off] -p <ip_address of the printer>",
	Args:                  cobra.MatchAll(cobra.OnlyValidArgs, cobra.MaximumNArgs(1)),
	ValidArgs:             []string{"on", "off"},
	DisableFlagsInUseLine: true,
	Short:                 "Enable or disable the active leveling on the printer",
	Long: `UltiTuner - Tune the Active Leveling

The "leveling" command is used to check, enable or disable the active leveling of your printer. 
You must do manual leveling from the printer menu after you have disabled the active leveling and before you start your first print job.

Note that you need to restart the printer after any configuration change. You can use the "restart" command or set the "-r" flag to restart it automatically.`,

	Run: func(cmd *cobra.Command, args []string) {
		restartFlag, _ := cmd.Flags().GetBool("restart")

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

		// Check if this command is compatible with the connected printer. Must return a string true if compatible and false if not (...&& echo true || echo false)
		checkCompatibility(client, "egrep -q 'self.__probing_mode = ProbeMode.(DETAILED|NEVER)' /usr/share/griffin/griffin/printer/procedures/pre_and_post_print/auto_bed_level_adjust/alignZAxisProcedure.py && echo true || echo false")

		// Enable Active Leveling
		if len(args) == 1 && args[0] == "on" {
			fmt.Print("Enabling active leveling.....")
			sshCmd(client, "sed -i 's/self.__probing_mode = ProbeMode.NEVER/self.__probing_mode = ProbeMode.DETAILED/g' /usr/share/griffin/griffin/printer/procedures/pre_and_post_print/auto_bed_level_adjust/alignZAxisProcedure.py")
			fmt.Println("done, turned ON")
			if restartFlag {
				restartGriffin(client)
			}

			//Disable Active Leveling
		} else if len(args) == 1 && args[0] == "off" {
			fmt.Print("Disabling active leveling.....")
			sshCmd(client, "sed -i 's/self.__probing_mode = ProbeMode.DETAILED/self.__probing_mode = ProbeMode.NEVER/g' /usr/share/griffin/griffin/printer/procedures/pre_and_post_print/auto_bed_level_adjust/alignZAxisProcedure.py")
			fmt.Println("done, turned OFF")
			fmt.Println("IMPORTANT: Do a manual leveling from the menu before you start your first print job!")
			if restartFlag {
				restartGriffin(client)
			}

			// Fetch Active Leveling status
		} else {
			fmt.Print("Checking the status of active leveling.....")
			result := sshCmd(client, "grep 'self.__probing_mode = ProbeMode.' /usr/share/griffin/griffin/printer/procedures/pre_and_post_print/auto_bed_level_adjust/alignZAxisProcedure.py")
			if strings.Contains(result, "DETAILED") {
				fmt.Println("done, currently ENABLED")
			} else {
				fmt.Println("done, currently DISABLED")
			}
		}

		fmt.Println()

	},
}

func init() {
	rootCmd.AddCommand(levelingCmd)
	levelingCmd.Flags().BoolP("restart", "r", false, "Restart the printer service after the change")
}
