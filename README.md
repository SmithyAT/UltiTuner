# UltiTuner - the Swiss army knife for Ultimaker printers
UltiTuner is a helper tool to configure functions for network-enabled Ultimaker printers (S3, S5), which are not available via the printers' menu.

We all love our Ultimaker printers, but sometimes we want to do things that are not officially supported or designed in another way that doesn't fit your personal workflow. And that's where UltiTuner kicks in, to fill these gaps and remove some restrictions from the official firmware.

All changes you make with this tool are easily reversible. It doesn't use a different firmware image, it just modifies some files or parameters.

## Disclaimer
The software was created and tested with the utmost care. But if something goes wrong, neither Ultimaker nor I am responsible if the printer no longer works as usual or breaks.

**USE THIS TOOL AT YOUR OWN RISK AND RESPONSIBILITY**

## Supported printer models
This tool supports the Ultimaker-S3 and Ultimaker-S5 printers, also called the S-line printers. The much older UM3 is currently not supported because it uses a slightly different firmware version and not all functions are available or needed for the UM3. 

Not all firmware versions are supported. Firmware v6.x and v7.x are fully supported. Older versions may work but are not tested much so far. I had a chance to get some tests with v5.7.2, which seems to work.

Each command has a function included, to check if the printer and firmware is compatible with that action to avoid any malfunctions. If it detects any incompatible printer, you get a notice that your printer is not supported, and you cannot use this command.

## Prerequisites
This tool uses ssh to connect to the printer, so you must first enable the "Developer Mode" in the printer menu.

## Installation
UltiTuner is just a CLI application, there is no need to install it. Just open a terminal or cmd window and run the application.

## Available Commands
For more information on each command, see the command reference section below with examples.

- `leveling` - Enable or disable the active leveling behavior
- `safetotouch` - Change the safeToTouch temperature
- `cooldown` - Enable or disable the post-print cool-down stage
- `clearabort` - Clear the waiting state after aborting a print job
- `sshserver`- Enable or disable the ssh daemon on the printer
- `restart` - Restart the griffin printer service
- `reboot` - Reboot the printer

## Flags
### -p
A mandatory flag is the `-p` or `--printer-ip` flag, where you have to specify the IP address of your printer.
The IP address is shown in the display of the printer after you have enabled the developer mode.

### -r
You must power cycle the printer to take effect any configuration changes.
But you can use the `-r` or `--restart` flag with the command to restart it automatically. Only the griffin printer service will be restarted, which is a bit faster than a full reboot or power cycle.

### -h
To get more information for each command you can add the `-h` or `-help` flag.

## Command reference & examples
Depending on your operating system, the executable is either ultituner or ultituner.exe. In the following examples, we use ultituner.

Note that you must power cycle the printer for most of these changes to take effect or use the "-r" flag to restart automatically.

### Active Leveling
With the `leveling` command, you can check, enable or disable the active leveling behavior of your printer.

Note that you must do a manual leveling from the printer menu before you start your first print job!

#### Check the current behavior
`ultituner leveling -p 192.168.0.23`

#### Disable active leveling 
`ultituner leveling off -p 192.168.0.23`

#### Enable active leveling and auto-restart the printer
`ultituner leveling on -p 192.168.0.23 -r`


### SafeToTouch Temperature
The `safetotouch` command checks or changes the safeToTouch temperature, which is used during the post-print cool-down stage. The value needs to be between 40 and 200 degrees Celsius, and the value is the temperature at which point it is safe to touch the build plate.

If you set the temperature to, i.e. 80, the cool-down stage finishes as soon as the bed temperature reaches 79 degrees Celsius. The higher the temperature, the faster the cool-down stage will end.

#### Check the current set temperature
`ultituner cooldown -p 192.168.0.23`

#### Change the temperature to 80 and restart the printer service
`ultituner cooldown 80 -p 192.168.0.23 -r`


### Cooldown 
The `cooldown` command is a simplified variant of the safetotouch command.
With `on` the safeToTouch temperature is set to the default of 60, and with `off` the temperature is set to 200, meaning that the cool-down stage at the end of a print job immediately ends.

Without an argument `[on|off]` you get the current set temperature, which is identical to the `safetotouch` command.

#### Disable the cool-down stage and restart the printer service
`ultituner cooldown off -p 192.168.0.23 -r`

#### Enable the cool-down stage
`ultituner cooldown on -p 192.168.0.23`


### Clear Abort, Awaiting Clean-Up
When you abort a print job in an early stage, so during the heat-up phase, before the print actually starts, the printer is asking you after the abort if you want to retry the job or not. In the DigitalFactory you see a message like "Aborted, Awaiting Clean-Up". You cannot start another print job until someone presses the button on the printer's display.

With the `clearabort` command, you can clear the message on the printer's display, and the printer is ready again to accept new print jobs.

Notice that it is not working when the printer has already started to print.

#### Clear the message
`ultituner clearabort -p 192.168.0.23`


### SSH Server
The `sshserver` command permanently enables the ssh daemon on the printer, independent of developer mode. It can be very helpful if you work with the firmware because the developer mode is just a flag, and the firmware starts the ssh daemon during startup.

If you have an error in a file and the startup routine throws an exception, the printer is bricked because you can no longer ssh into the printer to fix the problem. The ssh daemon is always available when enabling it permanently at system start, regardless of the startup routines of the printer firmware.

Remember that disabling the developer mode won't disable the ssh daemon anymore. If you want to disable ssh at all, you have to use UltiTuner again.

#### Enable the SSH daemon
`ultituner sshserver on -p 192.168.0.23`

#### Disable the SSH daemon
`ultituner sshserver off -p 192.168.0.23`


### Restart
Is used to restart the griffin printer service. It is enough to restart the printer service to take effect configuration changes instead of a full reboot.

`ultituner restart -p 192.168.0.23`


### Reboot
Is used to reboot the linux system of the printer.

`ultituner reboot -p 192.168.0.23`


## Update notification
There are ongoing improvements and bug fixes to this tool. Therefore an update checker is implemented, checking for new releases on GitHub once per day.

The tool creates an additional file (ultituner_upd-cache.json) along with the ultituner executable to avoid an update check on every run. If you delete the file accidentally, it will be automatically created again at the next start.


## Upcoming features
- [ ] UM3 support
- [ ] Reset the message if you have aborted a print job before it has started to print.
- [ ] LED switch
- [ ] I am open for ideas
