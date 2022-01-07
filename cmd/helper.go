package cmd

import (
	"fmt"
	"github.com/melbahja/goph"
	"os"
	"strconv"
	"strings"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func sshConnect() *goph.Client {
	client, err := goph.NewUnknown("root", printerIP, goph.Password("ultimaker"))
	if err != nil {
		fmt.Println("ERROR: Unable to connect to your printer! Check if the developer mode is enabled and if the ip address it correct.")
		os.Exit(1)
	}
	return client
}

func sshCmd(c *goph.Client, cmd string) string {
	out, err := c.Run(cmd)
	if err != nil {
		fmt.Println("ERROR: Something went wrong - unable to complete the action.")
		os.Exit(1)
	}
	result := string(out)
	result = strings.TrimSuffix(result, "\n")
	return result
}

func getPrinterProperties(c *goph.Client) string {
	result := sshCmd(c, "grep 'machine_name' /var/lib/griffin/system_preferences.json | cut -d':' -f2 | sed 's/[\"|,| ]//g' && cat /etc/ultimaker_firmware | sed 's/article_numbers/model/g'")
	return result
}

func getVersion(c *goph.Client) (model string, version int) {
	result := sshCmd(c, "echo $(grep article_numbers /etc/ultimaker_firmware | sed 's/article_numbers: //g'):$(grep version /etc/ultimaker_firmware | sed 's/version: //g' | cut -c1)")
	details := strings.Split(result, ":")
	tmpVersion, _ := strconv.Atoi(details[1])
	return details[0], tmpVersion
}

func checkVersion(c *goph.Client, requiredVersion int) bool {
	unsupportedPrinters := []string{"9066", "9511"}

	model, printerVersion := getVersion(c)
	if contains(unsupportedPrinters, model) || printerVersion < requiredVersion {
		return false
	} else {
		return true
	}
}
