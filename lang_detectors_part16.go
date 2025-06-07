package main

// 为ObjectiveC添加包检测功能
func (a *App) detectObjectiveCWithPackages() LanguageInfo {
	info := a.detectObjectiveC()

	if info.Installed {
		// 获取已安装的ObjectiveC包
		packages, _ := a.listObjectiveCPackages()
		info.Packages = packages
	}

	return info
}

// 为Bash添加包检测功能
func (a *App) detectBashWithPackages() LanguageInfo {
	info := a.detectBash()

	if info.Installed {
		// 获取已安装的Bash包
		packages, _ := a.listBashPackages()
		info.Packages = packages
	}

	return info
}

// 为PowerShell添加包检测功能
func (a *App) detectPowerShellWithPackages() LanguageInfo {
	info := a.detectPowerShell()

	if info.Installed {
		// 获取已安装的PowerShell包
		psCmd := "powershell"
		if commandExists("pwsh") {
			psCmd = "pwsh"
		}
		packages, _ := a.listPowerShellPackages(psCmd)
		info.Packages = packages
	}

	return info
}

// 为VBA添加包检测功能
func (a *App) detectVBAWithPackages() LanguageInfo {
	info := a.detectVBA()

	if info.Installed {
		// 获取已安装的VBA包
		packages, _ := a.listVBAPackages()
		info.Packages = packages
	}

	return info
}
