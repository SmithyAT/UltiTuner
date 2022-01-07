# UltiTuner
UltiTuner is a small helper tool to configure functions for Ultimaker S-Line printers, which are not available via the printers menu.

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

- `leveling` - Enable or disable the active leveling on the printer.


## Flags
You can get more information for each command when you add the --help (or -h) flag to the end.

A mandatory flag is the `-p` or `--printer-ip` flag, where you have to specify the IP address of your printer.

To check the version of UltiTuner run `ultituner --version`

## Example
Depending on your operating system, the executable is either ultituner or ultituner.exe. In the following examples, we use ultituner.
### Active Leveling
This command connects to the printer and checks the current active leveling configuration. This command is good to test if your printer is supported.

`ultituner leveling -p 192.168.0.23`

With the _off_ argument, you can turn off active leveling. Your printer will do a soft restart after the command.

`ultituner leveling off -p 192.168.0.23`

With the _on_ argument, you can turn it on again. Your printer will do a soft restart after the command.#

`ultituner leveling on -p 192.168.0.23`

## Inside UltiTuner
UltiTuner is not magic. According to the selected command, it modifies some parameters directly in some firmware files. You can do everything with ssh and the vi editor, but it is easier and more failure-proof for most users to have a tool that does these steps, mainly when someone is not used to using Linux.

Additionally, the tool takes care of different firmware versions because some files have changed or parameters are on other locations. Not all firmware versions are supported because I no longer have access to this old firmware. Maybe version 4 or versions older than 5.7.2 would also be working, but these old firmware versions are not supported due to the lack of testing capabilities.
