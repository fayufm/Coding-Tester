@echo off
echo 正在创建桌面快捷方式...

REM 使用PowerShell创建带自定义图标的快捷方式
powershell -Command "$WshShell = New-Object -ComObject WScript.Shell; $Shortcut = $WshShell.CreateShortcut([Environment]::GetFolderPath('Desktop') + '\Networ Tester.lnk'); $Shortcut.TargetPath = '%~dp0build\bin\networ-tester.exe'; $Shortcut.IconLocation = '%~dp0icon.ico'; $Shortcut.WorkingDirectory = '%~dp0build\bin'; $Shortcut.Save()"

echo 桌面快捷方式创建完成！
pause 