# MX Master Reconnect 🖱️

Automatically fixes Logitech MX Master mouse Bluetooth connection issues on Windows 11 restarts.

## The Problem
MX Master mice often fail to reconnect after Windows 11 restarts due to Bluetooth adapter compatibility issues.

## The Solution
This script automatically removes and re-pairs your MX Master at startup.

## Setup

1. **Download** - Get the `main.py` file
2. **Run at startup** - Add to Windows startup folder (`Win + R` → `shell:startup`) or use Task Scheduler
3. **Restart** - Your mouse should now connect automatically

## Requirements
- Windows 11
- Python 3.x
- Logitech MX Master mouse

## How It Works
At startup, the script finds your MX Master, removes the Bluetooth connection, and re-pairs.