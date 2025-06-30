# Logitech MX Master Blurtooth mouse Reconnect 🖱️
I noticed that my MX Master mouse often fails to reconnect to Bluetooth after Windows 11 restarts.
This script simply opens the Windows Bluetooth menu on startup, so you can use your keyboard to remove and re-add (repair) your mouse.

## Download & Setup

1. **Add to Startup** - To run the script at startup:
   - Press `Win + R`, type `shell:startup`, and press Enter to open the Startup folder.
   - Copy `open-bluetooth-settings.cmd` into this folder.
   - The Bluetooth menu will now open automatically each time you log in.
2. **Restart** - After restarting, the Bluetooth menu will appear. Use your keyboard (Tab/Arrow/Enter keys) to remove the MX Master mouse (if present), then add it again to repair the connection.

## Requirements
- Windows 11
- Logitech MX Master mouse

## How It Works
At startup, the script opens the Windows Bluetooth settings. To repair your mouse connection:
1. Use your keyboard to select and remove the MX Master mouse from the Bluetooth device list if it appears.
2. Put your mouse in pairing mode (press the pair button).
3. Use your keyboard (Tab/Arrow/Enter keys) to add a new Bluetooth device and select the MX Master mouse to pair.

This process helps resolve issues where the mouse fails to reconnect after a restart.