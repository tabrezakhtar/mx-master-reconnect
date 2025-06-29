# MX Master Reconnect 🖱️
I noticed that my MX Master mouse often fail to reconnect to Bluetooth after Windows 11 restarts.
This script automatically removes and re-pairs the mouse at startup.

## Setup

1. **Download** - Get the `main.py` file
2. **Run at startup** - Add to Windows startup folder (`Win + R` → `shell:startup`) or use Task Scheduler
3. **Restart** - Your mouse should now connect automatically

## Requirements
- Windows 11
- Python 3.x
- Logitech MX Master mouse

## How It Works
At startup, the script finds the MX Master mouse, removes the Bluetooth connection, and re-pairs.