package main

// 为C/C++添加包检测功能
func (a *App) detectCppWithPackages() LanguageInfo {
	info := a.detectCpp()

	if info.Installed {
		// 获取已安装的C/C++包
		packages, _ := a.listCppPackages()
		info.Packages = packages
	}

	return info
}

// 为Dart添加包检测功能
func (a *App) detectDartWithPackages() LanguageInfo {
	info := a.detectDart()

	if info.Installed {
		// 获取已安装的Dart包
		packages, _ := a.listDartPackages()
		info.Packages = packages
	}

	return info
}

// 为TypeScript添加包检测功能
func (a *App) detectTypeScriptWithPackages() LanguageInfo {
	info := a.detectTypeScript()

	if info.Installed {
		// 获取已安装的TypeScript包
		packages, _ := a.listTypeScriptPackages()
		info.Packages = packages
	}

	return info
}

// 为Lua添加包检测功能
func (a *App) detectLuaWithPackages() LanguageInfo {
	info := a.detectLua()

	if info.Installed {
		// 获取已安装的Lua包
		packages, _ := a.listLuaPackages()
		info.Packages = packages
	}

	return info
}

// 为R添加包检测功能
func (a *App) detectRWithPackages() LanguageInfo {
	info := a.detectR()

	if info.Installed {
		// 获取已安装的R包
		packages, _ := a.listRPackages()
		info.Packages = packages
	}

	return info
}

// 为Scala添加包检测功能
func (a *App) detectScalaWithPackages() LanguageInfo {
	info := a.detectScala()

	if info.Installed {
		// 获取已安装的Scala包
		packages, _ := a.listScalaPackages()
		info.Packages = packages
	}

	return info
}

// 为Haskell添加包检测功能
func (a *App) detectHaskellWithPackages() LanguageInfo {
	info := a.detectHaskell()

	if info.Installed {
		// 获取已安装的Haskell包
		packages, _ := a.listHaskellPackages()
		info.Packages = packages
	}

	return info
}
