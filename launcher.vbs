Set WshShell = CreateObject("WScript.Shell")
strPath = WshShell.CurrentDirectory
strApp = strPath & "\build\bin\networ-tester.exe"
WshShell.Run Chr(34) & strApp & Chr(34), 1, false 