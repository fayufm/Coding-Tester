package main

// 为CoffeeScript添加包检测功能
func (a *App) detectCoffeeScriptWithPackages() LanguageInfo {
	info := a.detectCoffeeScript()

	if info.Installed {
		// 获取已安装的CoffeeScript包
		packages, _ := a.listCoffeeScriptPackages()
		info.Packages = packages
	}

	return info
}

// 为J语言添加包检测功能
func (a *App) detectJWithPackages() LanguageInfo {
	info := a.detectJ()

	if info.Installed {
		// 获取已安装的J语言包
		packages, _ := a.listJPackages()
		info.Packages = packages
	}

	return info
}

// 为Vala添加包检测功能
func (a *App) detectValaWithPackages() LanguageInfo {
	info := a.detectVala()

	if info.Installed {
		// 获取已安装的Vala包
		packages, _ := a.listValaPackages()
		info.Packages = packages
	}

	return info
}

// 为Clojure添加包检测功能
func (a *App) detectClojureWithPackages() LanguageInfo {
	info := a.detectClojure()

	if info.Installed {
		// 获取已安装的Clojure包
		packages, _ := a.listClojurePackages()
		info.Packages = packages
	}

	return info
}
