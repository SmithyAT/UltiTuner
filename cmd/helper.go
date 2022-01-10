package cmd

import (
	"fmt"
	"github.com/melbahja/goph"
	"os"
	"strconv"
	"strings"
)

// Connect and return the client
func sshConnect() *goph.Client {
	client, err := goph.NewUnknown("root", printerIP, goph.Password("ultimaker"))
	if err != nil {
		fmt.Println("error, unable to connect to printer")
		fmt.Println("Check if the developer mode is enabled and if the ip address it correct.")
		os.Exit(1)
	}
	return client
}

// Exec remote commands and return the result
func sshCmd(c *goph.Client, cmd string) string {
	out, err := c.Run(cmd)
	if err != nil {
		fmt.Println()
		fmt.Println("ERROR: Something went wrong - unable to complete the action.")
		fmt.Println(err)
		os.Exit(1)
	}
	result := string(out)
	result = strings.TrimSuffix(result, "\n")
	return result
}

// Fetch the printer name, model and firmware version and output the result
func printPrinterProperties(c *goph.Client) {
	result := sshCmd(c, "grep 'machine_name' /var/lib/griffin/system_preferences.json | cut -d':' -f2 | sed 's/[\"|,| ]//g' && cat /etc/ultimaker_firmware | sed 's/article_numbers/model/g'")
	fmt.Println("--------------------")
	fmt.Println(result)
	fmt.Println("--------------------")
}

// The ssh command must return a string true if compatible and false if not (...&& echo true || echo false)
func checkCompatibility(c *goph.Client, cmd string) {
	fmt.Print("Checking printer/firmware compatibility.....")
	printerVersion := sshCmd(c, "grep article_numbers /etc/ultimaker_firmware | sed 's/article_numbers: //g'")

	result := sshCmd(c, cmd)
	supported, _ := strconv.ParseBool(result)
	if supported && printerVersion != "9066" && printerVersion != "9511" {
		fmt.Println("done, compatible")
	} else {
		fmt.Println("done, but not compatible")
		os.Exit(1)
	}
}

// Restart griffin service
func restartGriffin(c *goph.Client) {
	fmt.Print("Restarting griffin printer service.....")
	sshCmd(c, "systemctl restart griffin.printer")
	fmt.Println("done, wait a moment until the printer is ready again")
}
