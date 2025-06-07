package main

// 为Elixir添加包检测功能
func (a *App) detectElixirWithPackages() LanguageInfo {
	info := a.detectElixir()

	if info.Installed {
		// 获取已安装的Elixir包
		packages, _ := a.listElixirPackages()
		info.Packages = packages
	}

	return info
}

// 为F#添加包检测功能
func (a *App) detectFSharpWithPackages() LanguageInfo {
	info := a.detectFSharp()

	if info.Installed {
		// 获取已安装的F#包
		packages, _ := a.listFSharpPackages()
		info.Packages = packages
	}

	return info
}

// 为Prolog添加包检测功能
func (a *App) detectPrologWithPackages() LanguageInfo {
	info := a.detectProlog()

	if info.Installed {
		// 获取已安装的Prolog包
		packages, _ := a.listPrologPackages()
		info.Packages = packages
	}

	return info
}

// 为Lisp添加包检测功能
func (a *App) detectLispWithPackages() LanguageInfo {
	info := a.detectLisp()

	if info.Installed {
		// 获取已安装的Lisp包
		packages, _ := a.listLispPackages()
		info.Packages = packages
	}

	return info
}
