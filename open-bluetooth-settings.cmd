@echo off
echo Opening Windows Bluetooth settings...
powershell -command "Start-Process ms-settings:bluetooth; Start-Sleep -Seconds 1; Add-Type -AssemblyName System.Windows.Forms; [System.Windows.Forms.SendKeys]::SendWait('{TAB}');"
echo Windows Bluetooth settings opened. Use your keyboard to add or reconnect your MX Master mouse.
pause
