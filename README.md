# UltiTuner
UltiTuner is a small helper tool to configure functions for Ultimaker S-Line printers, which are not available via the printers' menu.

We all love our Ultimaker printers, but sometimes we want to do things that are not officially supported or designed in another way that doesn't fit your personal workflow. And that's where UltiTuner kicks in, and it tries to fill these gaps and remove some restrictions from the official firmware.

All changes you make with this tool are easily reversible. It doesn't use a different firmware image, and it just modifies some parameters.

## Disclaimer
The software was created and tested with utmost care. But since non-official methods are used to change the firmware, it could happen under certain circumstances that the printer no longer works as usual. Neither I personally nor Ultimaker is responsible for this. **USE THIS TOOL AT YOUR OWN RISK AND RESPONSIBILITY**

## Supported printer models
This tool supports the Ultimaker-S3 and Ultimaker-S5 printers, also called the S-line printers. The much older UM3 is not supported because it uses a slightly different firmware version. But if there is a need, I can look into it. Just raise a feature request.

Also, not all firmware versions are supported. Firmware v6.x and v7.x are fully supported. Older versions may work but are not tested much so far. I had a chance to get some tests with v5.7.2, which seems to work, but older versions are not supported and are excluded from working with this tool.

The tool checks the printer model and the firmware version to avoid any damage or malfunctions. If it detects any unsupported printer, you get a notice that your printer is not supported, and you cannot use this tool.

## Prerequisites
This tool uses ssh to connect to the printer, so you must first enable the "Developer Mode" in the printer menu.

## Available Commands
For more information on each command, see the Example section below.

- `leveling` - Enable or disable the active leveling on the printer
- `cooldown` - Change the safeToTouch temperature
- `restart` - Do a soft restart of the printer
- `reboot` - Do a full reboot of the printer


## Flags
You can get more information for each command when you add the --help (or -h) flag to the end.

A mandatory flag is the `-p` or `--printer-ip` flag, where you have to specify the IP address of your printer.

The `-r` or `restart` flag is used to auto-restart after a command has changed something.

To check the version of UltiTuner run `ultituner --version`

## Command reference
Depending on your operating system, the executable is either ultituner or ultituner.exe. In the following examples, we use ultituner.
### Active Leveling
This command connects to the printer and checks the current active leveling configuration.
You need to restart the printer after any configuration change. You can use the "restart" command or set the "-r" flag to restart it automatically.

`ultituner leveling -p 192.168.0.23`

With the _off_ argument, you can turn off active leveling.

`ultituner leveling off -p 192.168.0.23`

Automatically restart the printer service after disabling active leveling with the "-r" flag at the end.

`ultituner leveling off -p 192.168.0.23 -r`

With the _on_ argument, you can turn it on again.

`ultituner leveling on -p 192.168.0.23`

### Cool Down
The "cooldown" command checks or changes the safeToTouch temperature, which is used during the cool-down phase after the print job.
The value needs to be between 40 and 100 degrees Celsius, and the value is the temperature at which point it is safe to touch the build plate.

If you set the temperature to, i.e., 80, the cool-down phase finishes as soon as the bed temperature reaches 79 degrees Celsius.
The higher the temperature, the faster the cool-down stage will end.

You need to restart the printer after any configuration change. You can use the "restart" command or set the "-r" flag to restart it automatically.

Check the current set temperature

`ultituner cooldown -p 192.168.0.23`

Change the temperature to 80 and restart the printer service

`ultituner cooldown 80 -p 192.168.0.23 -r`

### Restart & Reboot
Just restart the printer. Not the same as power cycling, but it restarts the printer service.

`ultituner restart -p 192.168.0.23`

Or do a complete reboot of the linux system of the printer.

`ultituner reboot -p 192.168.0.23`

## Inside UltiTuner
UltiTuner is not magic. According to the selected command, it modifies some parameters directly in some firmware files. You can do everything with ssh and the vi editor, but it is easier and more failure-proof for most users to have a tool that does these steps, mainly when someone is not used to using Linux.

Additionally, the tool takes care of different firmware versions because some files have changed or parameters are on other locations. Not all firmware versions are supported because I no longer have access to this old firmware. Maybe version 4 or versions older than 5.7.2 would also be working, but these old firmware versions are not supported due to the lack of testing capabilities.

## Upcoming features
- Reset the "clean the build plate" message if you had aborted a print job before it began.
- Enable SSH independent from the developer mode
- ??? open for ideas
