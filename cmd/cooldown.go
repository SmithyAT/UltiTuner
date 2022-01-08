package cmd

import (
	"fmt"
	"github.com/melbahja/goph"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

var cooldownCmd = &cobra.Command{
	Use:                   "cooldown [temperature] -p <ip_address of the printer>",
	Args:                  cobra.MatchAll(cobra.MaximumNArgs(1)),
	DisableFlagsInUseLine: true,
	Short:                 "Check or change the safeToTouch temperature",
	Long: `UltiTuner - Tune the safeToTouch Temperature

The "cooldown" command checks or changes the safeToTouch temperature, which is used during the cool-down phase after the print job. 
The value needs to be between 40 and 100 degrees Celsius, and the value is the temperature at which point it is safe to touch the build plate. 

If you set the temperature to, i.e., 80, the cool-down phase finishes as soon as the bed temperature reaches 79 degrees Celsius. 
The higher the temperature, the faster the cool-down stage will end. 

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
		checkCompatibility(client, "egrep -q 'property var safeToTouch: Okuda.SystemStateProxy.printBed.temperature < Okuda.Defines.safeToTouchTemperature' /usr/share/okuda/components/progress/PrintProgress.qml && echo true || echo false")

		// Check if the argument is an integer and within range
		var newTemp string
		if len(args) > 0 {
			tmpTemp, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("ERROR: An integer value is needed for the safeToTouch temperature")
				os.Exit(1)
			}
			if tmpTemp < 40 || tmpTemp > 100 {
				fmt.Println("ERROR: The safeToTouch temperature must be a value between 40 and 100")
				os.Exit(1)
			}
			newTemp = args[0]
		}

		// Check which file has to be modified - depends on the firmware
		checkPrinterVariant := sshCmd(client, "grep -q 'property var safeToTouch: Okuda.SystemStateProxy.printBed.temperature < Okuda.Defines.safeToTouchTemperature' /usr/share/okuda/components/progress/PrintProgress.qml && echo 'new_version' || echo 'old_version'")

		// Change safeToTouch temperature
		if len(args) > 0 {
			if checkPrinterVariant == "new_version" {
				sshCmd(client, "sed -i 's/property int safeToTouchTemperature: .*$/property int safeToTouchTemperature: 100 \\/\\/ [degC] see safe_human_touchable_temperature in um3.json/' /usr/share/okuda/Okuda/QmlPlugin/Okuda/Defines.qml")
			} else {
				sshCmd(client, "sed -i 's/property var safeToTouch: Okuda.SystemStateProxy.printBed.temperature <.*$/property var safeToTouch: Okuda.SystemStateProxy.printBed.temperature < "+args[0]+"/' /usr/share/okuda/components/progress/PrintProgress.qml")
			}
			fmt.Println("Changed safeToTouch temperature to : " + newTemp)
			if restartFlag {
				restartGriffin(client)
			}

			// Fetch configured safeToTouch temperature
		} else {
			var result string
			if checkPrinterVariant == "new_version" {
				result = sshCmd(client, "grep 'property int safeToTouchTemperature:' /usr/share/okuda/Okuda/QmlPlugin/Okuda/Defines.qml | cut -d\":\" -f 2 | sed 's/ //g' | cut -d\"/\" -f 1")
			} else {
				result = sshCmd(client, "grep 'property var safeToTouch: Okuda.SystemStateProxy.printBed.temperature <' /usr/share/okuda/components/progress/PrintProgress.qml | cut -d\"<\" -f 2 | sed 's/ //g'")
			}
			fmt.Println("Current configured safeToTouch temperature : " + result)
		}

		fmt.Println()

	},
}

func init() {
	rootCmd.AddCommand(cooldownCmd)
	cooldownCmd.Flags().BoolP("restart", "r", false, "Restart the printer service after the change")
}
