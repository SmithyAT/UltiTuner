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
	Long: `UltiTuner - Tune the Active Leveling

The "leveling" command is used to check, enable or disable the active leveling of your printer. 
It is important that you do a manual leveling from the printer menu, after you have disabled the active leveling and before you start your first print job.

Note, that the printer will restart after you send the command to enable or disable the active leveling.

UltiTuner uses ssh to connect to the printer, so you need to enable the "Developer Mode" in the printer menu before. 
`,
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
			sshCmd(client, "sed -i 's/self.__probing_mode = ProbeMode.NEVER/self.__probing_mode = ProbeMode.DETAILED/g' /usr/share/griffin/griffin/printer/procedures/pre_and_post_print/auto_bed_level_adjust/alignZAxisProcedure.py && systemctl restart griffin.printer")
			fmt.Println("Active Leveling turned ON")
		} else if len(args) == 1 && args[0] == "off" {
			sshCmd(client, "sed -i 's/self.__probing_mode = ProbeMode.DETAILED/self.__probing_mode = ProbeMode.NEVER/g' /usr/share/griffin/griffin/printer/procedures/pre_and_post_print/auto_bed_level_adjust/alignZAxisProcedure.py && systemctl restart griffin.printer")
			fmt.Println("Active Leveling turned OFF")
			fmt.Println("Important: Do a manual leveling from the menu before you start your first print job!")
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
