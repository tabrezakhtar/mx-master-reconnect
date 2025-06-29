# MX Master Reconnect 🖱️

A C# terminal application that automatically fixes Logitech MX Master mouse Bluetooth connection issues on Windows 11 restarts.

## The Problem
MX Master mice often fail to reconnect after Windows 11 restarts due to Bluetooth adapter compatibility issues.

## The Solution
This application automatically restarts the Bluetooth service at startup, which forces your MX Master to reconnect properly.

## Download & Setup

1. **Download** - Go to the [Releases](../../releases) page and download the latest version
2. **Run at startup** - Add the executable to Windows startup folder (`Win + R` → `shell:startup`) or use Task Scheduler
3. **Restart** - Your mouse should now connect automatically after every restart

## Requirements
- Windows 11
- Logitech MX Master mouse

## How It Works
At startup, the application finds your MX Master and restarts the Windows Bluetooth service to force reconnection.