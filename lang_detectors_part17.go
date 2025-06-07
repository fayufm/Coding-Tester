package main

// 为SQL添加包检测功能
func (a *App) detectSQLWithPackages() LanguageInfo {
	info := a.detectSQL()

	if info.Installed {
		// 获取已安装的SQL包
		packages, _ := a.listSQLPackages()
		info.Packages = packages
	}

	return info
}

// 为HTML添加包检测功能
func (a *App) detectHTMLWithPackages() LanguageInfo {
	info := a.detectHTML()

	if info.Installed {
		// 获取已安装的HTML包
		packages, _ := a.listHTMLPackages()
		info.Packages = packages
	}

	return info
}
