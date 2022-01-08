package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var cooldownCmd = &cobra.Command{
	Use:                   "cooldown [on | off] -p <ip_address of the printer>",
	Args:                  cobra.MatchAll(cobra.MaximumNArgs(1)),
	DisableFlagsInUseLine: true,
	Short:                 "Enable or disable the cool-down stage",
	Long: `UltiTuner - Tune the cool-down stage

The "cooldown" command is an easy variant of the safetotouch command.
With "on" the safeToTouch temperature is set to the default of 60, and with "off" the temperature is set to 200.  
Without an argument [on|off], you get the current set temperature, and it is then identical to the "safetotouch" command.

Note that you need to restart the printer after any configuration change. You can use the "restart" command or set the "-r" flag to restart it automatically.`,

	Run: func(cmd *cobra.Command, args []string) {
		restartFlag, _ := cmd.Flags().GetBool("restart")

		// Enable the cool-down stage, safeToTouch = 60
		if len(args) == 1 && args[0] == "on" {
			if restartFlag {
				safetotouchCmd.SetArgs([]string{
					fmt.Sprintf("-r"),
				})
			}
			safetotouchCmd.Run(cmd, []string{"60"})

			// Disable the cool-down stage, safeToTouch = 200
		} else if len(args) == 1 && args[0] == "off" {
			if restartFlag {
				safetotouchCmd.SetArgs([]string{
					fmt.Sprintf("-r"),
				})
			}
			safetotouchCmd.Run(cmd, []string{"200"})

			// Fetch configured safeToTouch temperature
		} else {
			safetotouchCmd.Run(cmd, []string{})
		}

	},
}

func init() {
	rootCmd.AddCommand(cooldownCmd)
	cooldownCmd.Flags().BoolP("restart", "r", false, "Restart the printer service after the change")
}
