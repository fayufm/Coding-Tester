package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

// 检测ABAP
func (a *App) detectABAP() LanguageInfo {
	info := LanguageInfo{
		Name:            "ABAP",
		Installed:       false,
		DownloadURL:     "https://developers.sap.com/trials-downloads.html",
		InstallTutorial: "https://developers.sap.com/tutorials/abap-environment-trial-onboarding.html",
		PackageManager:  "ABAP Package Manager",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "ABAP SDK for Google Cloud",
				Description: "连接SAP系统和Google Cloud的SDK",
				Version:     "latest",
			},
			{
				Name:        "ABAP Git",
				Description: "ABAP的Git客户端",
				Version:     "latest",
			},
			{
				Name:        "ABAP RESTful Application Programming Model",
				Description: "用于构建企业服务和应用程序的框架",
				Version:     "latest",
			},
		},
	}

	// ABAP通常只在SAP环境中运行，无法通过命令行检测
	return info
}

// 检测ActionScript
func (a *App) detectActionScript() LanguageInfo {
	info := LanguageInfo{
		Name:            "ActionScript",
		Installed:       false,
		DownloadURL:     "https://www.adobe.com/products/animate.html",
		InstallTutorial: "https://helpx.adobe.com/animate/using/creating-publishing-html5-canvas-document.html",
		PackageManager:  "N/A",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "Adobe Animate",
				Description: "用于创建交互式动画的工具",
				Version:     "latest",
			},
			{
				Name:        "Apache Flex",
				Description: "用于构建移动和Web应用程序的SDK",
				Version:     "4.16.1",
			},
			{
				Name:        "FlashDevelop",
				Description: "ActionScript的开源IDE",
				Version:     "5.3.3",
			},
		},
	}

	// 检测Adobe Animate或Flash
	if commandExists("animate") {
		info.Installed = true
		info.Version = "Adobe Animate (版本无法自动检测)"
	}

	return info
}

// 检测APL
func (a *App) detectAPL() LanguageInfo {
	info := LanguageInfo{
		Name:            "APL",
		Installed:       false,
		DownloadURL:     "https://www.dyalog.com/download-zone.htm",
		InstallTutorial: "https://www.dyalog.com/uploads/documents/MasteringDyalogAPL.pdf",
		PackageManager:  "N/A",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "Dyalog APL",
				Description: "APL的现代实现",
				Version:     "18.2",
			},
			{
				Name:        "GNU APL",
				Description: "APL的自由实现",
				Version:     "1.8",
			},
			{
				Name:        "NARS2000",
				Description: "APL的开源实现",
				Version:     "latest",
			},
		},
	}

	// 检测Dyalog APL
	if commandExists("dyalog") {
		info.Installed = true
		info.Version = "Dyalog APL (版本无法自动检测)"
		return info
	}

	// 检测GNU APL
	if commandExists("apl") {
		output, err := executeCommandWithTimeout("apl", "--version")
		if err == nil {
			info.Installed = true
			info.Version = "GNU APL: " + output
		}
	}

	return info
}

// 检测Ballerina
func (a *App) detectBallerina() LanguageInfo {
	info := LanguageInfo{
		Name:            "Ballerina",
		Installed:       false,
		DownloadURL:     "https://ballerina.io/downloads/",
		InstallTutorial: "https://ballerina.io/learn/getting-started/",
		PackageManager:  "Ballerina Central",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "ballerina/http",
				Description: "HTTP客户端和服务器实现",
				Version:     "latest",
			},
			{
				Name:        "ballerina/io",
				Description: "I/O API",
				Version:     "latest",
			},
			{
				Name:        "ballerina/jwt",
				Description: "JWT验证和生成",
				Version:     "latest",
			},
			{
				Name:        "ballerina/mysql",
				Description: "MySQL客户端",
				Version:     "latest",
			},
		},
	}

	if !commandExists("bal") {
		return info
	}

	output, err := executeCommandWithTimeout("bal", "version")
	if err == nil {
		info.Installed = true
		info.Version = output
	}

	return info
}

// 检测BASIC
func (a *App) detectBASIC() LanguageInfo {
	info := LanguageInfo{
		Name:            "BASIC",
		Installed:       false,
		DownloadURL:     "https://www.freebasic.net/",
		InstallTutorial: "https://www.freebasic.net/wiki/DocToc",
		PackageManager:  "N/A",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "FreeBASIC",
				Description: "开源的BASIC编译器",
				Version:     "1.10.0",
			},
			{
				Name:        "QB64",
				Description: "QuickBASIC的现代版本",
				Version:     "latest",
			},
			{
				Name:        "SmallBASIC",
				Description: "结构化BASIC方言",
				Version:     "12.24",
			},
		},
	}

	// 检测FreeBASIC
	if commandExists("fbc") {
		output, err := executeCommandWithTimeout("fbc", "-v")
		if err == nil {
			info.Installed = true
			info.Version = "FreeBASIC: " + output
			return info
		}
	}

	// 检测QB64
	if commandExists("qb64") {
		info.Installed = true
		info.Version = "QB64 (版本无法自动检测)"
		return info
	}

	return info
}

// 检测Boo
func (a *App) detectBoo() LanguageInfo {
	info := LanguageInfo{
		Name:            "Boo",
		Installed:       false,
		DownloadURL:     "https://github.com/boo-lang/boo",
		InstallTutorial: "https://github.com/boo-lang/boo/wiki/Getting-Started",
		PackageManager:  "NuGet",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "Boo.Lang",
				Description: "Boo语言运行时库",
				Version:     "2.0.9999.3",
			},
			{
				Name:        "Boo.Lang.Compiler",
				Description: "Boo编译器",
				Version:     "2.0.9999.3",
			},
			{
				Name:        "Boo.Lang.Parser",
				Description: "Boo解析器",
				Version:     "2.0.9999.3",
			},
		},
	}

	// 检测Boo编译器
	if commandExists("booc") {
		output, err := executeCommandWithTimeout("booc", "-version")
		if err == nil {
			info.Installed = true
			info.Version = output
		}
	}

	return info
}

// 检测Ceylon
func (a *App) detectCeylon() LanguageInfo {
	info := LanguageInfo{
		Name:            "Ceylon",
		Installed:       false,
		DownloadURL:     "https://ceylon-lang.org/download/",
		InstallTutorial: "https://ceylon-lang.org/documentation/1.3/tour/",
		PackageManager:  "Ceylon Herd",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "ceylon.collection",
				Description: "集合API",
				Version:     "1.3.3",
			},
			{
				Name:        "ceylon.http",
				Description: "HTTP客户端和服务器API",
				Version:     "1.3.3",
			},
			{
				Name:        "ceylon.json",
				Description: "JSON解析和生成",
				Version:     "1.3.3",
			},
		},
	}

	if !commandExists("ceylon") {
		return info
	}

	output, err := executeCommandWithTimeout("ceylon", "--version")
	if err == nil {
		info.Installed = true
		info.Version = output
	}

	return info
}

// 检测CoffeeScript
func (a *App) detectCoffeeScript() LanguageInfo {
	info := LanguageInfo{
		Name:            "CoffeeScript",
		Installed:       false,
		DownloadURL:     "https://coffeescript.org/",
		InstallTutorial: "https://coffeescript.org/#installation",
		PackageManager:  "npm",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "coffeescript",
				Description: "CoffeeScript编译器",
				Version:     "2.7.0",
			},
			{
				Name:        "coffee-script-redux",
				Description: "CoffeeScript编译器的重写版本",
				Version:     "2.0.0-beta8",
			},
			{
				Name:        "coffeelint",
				Description: "CoffeeScript的代码质量工具",
				Version:     "2.1.0",
			},
		},
	}

	if !commandExists("coffee") {
		return info
	}

	output, err := executeCommandWithTimeout("coffee", "--version")
	if err == nil {
		info.Installed = true
		info.Version = output

		// 获取已安装的CoffeeScript相关包
		packages, _ := a.listCoffeeScriptPackages()
		info.Packages = packages
	}

	return info
}

// 列出已安装的CoffeeScript相关包
func (a *App) listCoffeeScriptPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 检查npm是否安装
	if !commandExists("npm") {
		return packages, nil
	}

	// 检查全局安装的CoffeeScript
	output, err := executeCommandWithTimeout("npm", "list", "-g", "coffeescript")
	if err == nil && !strings.Contains(output, "empty") {
		// 提取版本信息
		lines := strings.Split(output, "\n")
		for _, line := range lines {
			if strings.Contains(line, "coffeescript@") {
				parts := strings.Split(line, "@")
				version := parts[len(parts)-1]

				packages = append(packages, PackageInfo{
					Name:        "coffeescript",
					Description: "CoffeeScript编译器",
					Version:     version,
					Installed:   true,
				})

				break
			}
		}
	}

	// 检查相关的CoffeeScript工具
	coffeeTools := []string{
		"coffee-script", "coffeelint", "coffee-script-redux", "coffee-react",
		"coffee-react-transform", "cson", "coffee-loader", "coffeescript-compiler",
	}

	for _, tool := range coffeeTools {
		output, err := executeCommandWithTimeout("npm", "list", "-g", tool)
		if err == nil && !strings.Contains(output, "empty") && strings.Contains(output, tool+"@") {
			// 提取版本信息
			lines := strings.Split(output, "\n")
			for _, line := range lines {
				if strings.Contains(line, tool+"@") {
					parts := strings.Split(line, "@")
					version := parts[len(parts)-1]

					var description string
					switch tool {
					case "coffee-script":
						description = "CoffeeScript编译器(旧版)"
					case "coffeelint":
						description = "CoffeeScript代码质量工具"
					case "coffee-script-redux":
						description = "CoffeeScript编译器的重写版本"
					case "coffee-react":
						description = "CoffeeScript的React JSX支持"
					case "coffee-react-transform":
						description = "CoffeeScript的React转换工具"
					case "cson":
						description = "CoffeeScript对象表示法解析器"
					case "coffee-loader":
						description = "Webpack的CoffeeScript加载器"
					case "coffeescript-compiler":
						description = "CoffeeScript编译器包装器"
					default:
						description = "CoffeeScript相关工具"
					}

					packages = append(packages, PackageInfo{
						Name:        tool,
						Description: description,
						Version:     version,
						Installed:   true,
					})

					break
				}
			}
		}
	}

	// 检查本地项目中的包
	// 尝试在当前目录或常见的项目目录中查找package.json
	projectDirs := []string{
		".", "./coffee", "./src",
	}

	for _, dir := range projectDirs {
		packageJsonPath := filepath.Join(dir, "package.json")
		if _, err := os.Stat(packageJsonPath); err == nil {
			content, err := os.ReadFile(packageJsonPath)
			if err != nil {
				continue
			}

			// 解析JSON内容
			var packageJson struct {
				Dependencies    map[string]string `json:"dependencies"`
				DevDependencies map[string]string `json:"devDependencies"`
			}

			if err := json.Unmarshal(content, &packageJson); err != nil {
				continue
			}

			// 检查依赖中的CoffeeScript相关包
			for _, deps := range []map[string]string{packageJson.Dependencies, packageJson.DevDependencies} {
				for pkg, version := range deps {
					if pkg == "coffeescript" || pkg == "coffee-script" ||
						strings.Contains(pkg, "coffee") || strings.Contains(pkg, "cson") {
						packages = append(packages, PackageInfo{
							Name:      pkg,
							Version:   version,
							Installed: true,
						})

						// 限制数量
						if len(packages) >= 20 {
							return packages, nil
						}
					}
				}
			}
		}
	}

	return packages, nil
}

// 检测Elm
func (a *App) detectElm() LanguageInfo {
	info := LanguageInfo{
		Name:            "Elm",
		Installed:       false,
		DownloadURL:     "https://guide.elm-lang.org/install/elm.html",
		InstallTutorial: "https://guide.elm-lang.org/",
		PackageManager:  "elm-package",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "elm/core",
				Description: "Elm的核心库",
				Version:     "1.0.5",
			},
			{
				Name:        "elm/browser",
				Description: "控制浏览器的库",
				Version:     "1.0.2",
			},
			{
				Name:        "elm/html",
				Description: "HTML库",
				Version:     "1.0.0",
			},
			{
				Name:        "elm/json",
				Description: "JSON编码和解码",
				Version:     "1.1.3",
			},
			{
				Name:        "elm/http",
				Description: "HTTP请求",
				Version:     "2.0.0",
			},
		},
	}

	if !commandExists("elm") {
		return info
	}

	output, err := executeCommandWithTimeout("elm", "--version")
	if err == nil {
		info.Installed = true
		info.Version = output
	}

	return info
}

// 检测Hack
func (a *App) detectHack() LanguageInfo {
	info := LanguageInfo{
		Name:            "Hack",
		Installed:       false,
		DownloadURL:     "https://hacklang.org/",
		InstallTutorial: "https://docs.hhvm.com/hack/getting-started/getting-started",
		PackageManager:  "Composer",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "hhvm/hhvm-autoload",
				Description: "Hack的自动加载器",
				Version:     "3.3.0",
			},
			{
				Name:        "hhvm/hsl",
				Description: "Hack标准库",
				Version:     "4.108.1",
			},
			{
				Name:        "facebook/fbexpect",
				Description: "Hack的单元测试库",
				Version:     "2.10.0",
			},
		},
	}

	if !commandExists("hhvm") {
		return info
	}

	output, err := executeCommandWithTimeout("hhvm", "--version")
	if err == nil {
		info.Installed = true
		info.Version = output
	}

	return info
}

// 检测J
func (a *App) detectJ() LanguageInfo {
	info := LanguageInfo{
		Name:            "J",
		Installed:       false,
		DownloadURL:     "https://www.jsoftware.com/",
		InstallTutorial: "https://code.jsoftware.com/wiki/System/Installation",
		PackageManager:  "pacman",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "math/misc",
				Description: "数学杂项",
				Version:     "latest",
			},
			{
				Name:        "graphics/plot",
				Description: "绘图库",
				Version:     "latest",
			},
			{
				Name:        "tables/csv",
				Description: "CSV文件处理",
				Version:     "latest",
			},
		},
	}

	if !commandExists("jconsole") {
		return info
	}

	// 检测J版本
	output, err := executeCommandWithTimeout("jconsole", "-js", "JVERSION")
	if err == nil && output != "" {
		info.Installed = true
		info.Version = strings.TrimSpace(output)

		// 列出J包
		packages, _ := a.listJPackages()
		info.Packages = packages
	}

	return info
}

// 列出J包
func (a *App) listJPackages() ([]PackageInfo, error) {
	packages := []PackageInfo{}

	// 使用J的pacman包管理器列出已安装的包
	output, err := executeCommandWithTimeout("jconsole", "-js", "showinstalled''")
	if err != nil || output == "" {
		return packages, err
	}

	lines := strings.Split(output, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || !strings.Contains(line, " ") {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) >= 2 {
			name := parts[0]
			version := "latest"

			// 尝试提取版本信息
			if len(parts) >= 3 {
				version = parts[1]
			}

			packages = append(packages, PackageInfo{
				Name:      name,
				Version:   version,
				Installed: true,
			})
		}
	}

	return packages, nil
}

// 检测Jython
func (a *App) detectJython() LanguageInfo {
	info := LanguageInfo{
		Name:            "Jython",
		Installed:       false,
		DownloadURL:     "https://www.jython.org/download.html",
		InstallTutorial: "https://www.jython.org/installation.html",
		PackageManager:  "pip/easy_install",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "django",
				Description: "Python Web框架",
				Version:     "1.11.29",
			},
			{
				Name:        "requests",
				Description: "HTTP库",
				Version:     "2.25.1",
			},
			{
				Name:        "numpy",
				Description: "科学计算库",
				Version:     "1.16.6",
			},
		},
	}

	if !commandExists("jython") {
		return info
	}

	output, err := executeCommandWithTimeout("jython", "--version")
	if err == nil {
		info.Installed = true
		info.Version = output
	}

	return info
}

// 检测LOLCODE
func (a *App) detectLOLCODE() LanguageInfo {
	info := LanguageInfo{
		Name:            "LOLCODE",
		Installed:       false,
		DownloadURL:     "https://github.com/justinmeza/lci",
		InstallTutorial: "https://github.com/justinmeza/lci/blob/master/README.md",
		PackageManager:  "N/A",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "lci",
				Description: "LOLCODE解释器",
				Version:     "0.11.2",
			},
		},
	}

	if !commandExists("lci") {
		return info
	}

	info.Installed = true
	info.Version = "LOLCODE (版本无法自动检测)"

	return info
}

// 检测PureScript
func (a *App) detectPureScript() LanguageInfo {
	info := LanguageInfo{
		Name:            "PureScript",
		Installed:       false,
		DownloadURL:     "https://www.purescript.org/",
		InstallTutorial: "https://github.com/purescript/documentation/blob/master/guides/Getting-Started.md",
		PackageManager:  "spago",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "purescript-prelude",
				Description: "PureScript的基本函数和类型",
				Version:     "6.0.1",
			},
			{
				Name:        "purescript-effect",
				Description: "副作用处理",
				Version:     "4.0.0",
			},
			{
				Name:        "purescript-console",
				Description: "控制台输出",
				Version:     "6.0.0",
			},
			{
				Name:        "purescript-aff",
				Description: "异步效果",
				Version:     "7.1.0",
			},
			{
				Name:        "purescript-halogen",
				Description: "UI库",
				Version:     "7.0.0",
			},
		},
	}

	if !commandExists("purs") {
		return info
	}

	output, err := executeCommandWithTimeout("purs", "--version")
	if err == nil {
		info.Installed = true
		info.Version = output
	}

	return info
}

// 检测Q#
func (a *App) detectQSharp() LanguageInfo {
	info := LanguageInfo{
		Name:            "Q#",
		Installed:       false,
		DownloadURL:     "https://docs.microsoft.com/en-us/quantum/",
		InstallTutorial: "https://docs.microsoft.com/en-us/quantum/quickstarts/install-command-line",
		PackageManager:  "NuGet",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "Microsoft.Quantum.Standard",
				Description: "Q#标准库",
				Version:     "0.28.302812",
			},
			{
				Name:        "Microsoft.Quantum.Development.Kit",
				Description: "Q#开发工具包",
				Version:     "0.28.302812",
			},
			{
				Name:        "Microsoft.Quantum.Numerics",
				Description: "数值计算库",
				Version:     "0.28.302812",
			},
		},
	}

	if !commandExists("dotnet") {
		return info
	}

	// 检查是否安装了Q#项目模板
	output, err := executeCommandWithTimeout("dotnet", "new", "--list")
	if err == nil && strings.Contains(output, "Q#") {
		info.Installed = true
		info.Version = "Q# (通过.NET SDK安装)"
	}

	return info
}

// 检测Red
func (a *App) detectRed() LanguageInfo {
	info := LanguageInfo{
		Name:            "Red",
		Installed:       false,
		DownloadURL:     "https://www.red-lang.org/p/download.html",
		InstallTutorial: "https://github.com/red/red/wiki/Getting-started",
		PackageManager:  "Red Package Manager",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "redbin",
				Description: "Red二进制格式",
				Version:     "latest",
			},
			{
				Name:        "view",
				Description: "GUI系统",
				Version:     "latest",
			},
			{
				Name:        "parse",
				Description: "解析方言",
				Version:     "latest",
			},
		},
	}

	if !commandExists("red") {
		return info
	}

	info.Installed = true
	info.Version = "Red (版本无法自动检测)"

	return info
}

// 检测ReScript
func (a *App) detectReScript() LanguageInfo {
	info := LanguageInfo{
		Name:            "ReScript",
		Installed:       false,
		DownloadURL:     "https://rescript-lang.org/docs/manual/latest/installation",
		InstallTutorial: "https://rescript-lang.org/docs/manual/latest/installation",
		PackageManager:  "npm",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "rescript",
				Description: "ReScript编译器和标准库",
				Version:     "10.1.4",
			},
			{
				Name:        "@rescript/react",
				Description: "React绑定",
				Version:     "0.11.0",
			},
			{
				Name:        "@rescript/core",
				Description: "ReScript核心库",
				Version:     "0.5.0",
			},
		},
	}

	if !commandExists("rescript") {
		return info
	}

	output, err := executeCommandWithTimeout("rescript", "--version")
	if err == nil {
		info.Installed = true
		info.Version = output

		// 获取已安装的ReScript相关包
		packages, _ := a.listReScriptPackages()
		info.Packages = packages
	}

	return info
}

// 列出已安装的ReScript相关包
func (a *App) listReScriptPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 检查npm是否安装
	if !commandExists("npm") {
		return packages, nil
	}

	// 检查全局安装的ReScript
	output, err := executeCommandWithTimeout("npm", "list", "-g", "rescript")
	if err == nil && !strings.Contains(output, "empty") {
		// 提取版本信息
		lines := strings.Split(output, "\n")
		for _, line := range lines {
			if strings.Contains(line, "rescript@") {
				parts := strings.Split(line, "@")
				version := parts[len(parts)-1]

				packages = append(packages, PackageInfo{
					Name:        "rescript",
					Description: "ReScript编译器和标准库",
					Version:     version,
					Installed:   true,
				})

				break
			}
		}
	}

	// 检查相关的ReScript工具
	rescriptTools := []string{
		"@rescript/react", "@rescript/core", "bs-platform", "gentype",
		"@glennsl/bs-jest", "bs-fetch", "bs-json", "rescript-compiler",
	}

	for _, tool := range rescriptTools {
		output, err := executeCommandWithTimeout("npm", "list", "-g", tool)
		if err == nil && !strings.Contains(output, "empty") && strings.Contains(output, tool+"@") {
			// 提取版本信息
			lines := strings.Split(output, "\n")
			for _, line := range lines {
				if strings.Contains(line, tool+"@") {
					parts := strings.Split(line, "@")
					version := parts[len(parts)-1]

					var description string
					switch tool {
					case "@rescript/react":
						description = "ReScript的React绑定"
					case "@rescript/core":
						description = "ReScript核心库"
					case "bs-platform":
						description = "BuckleScript平台(ReScript旧名)"
					case "gentype":
						description = "ReScript到TypeScript/Flow类型转换工具"
					case "@glennsl/bs-jest":
						description = "ReScript的Jest绑定"
					case "bs-fetch":
						description = "ReScript的Fetch API绑定"
					case "bs-json":
						description = "ReScript的JSON解析库"
					case "rescript-compiler":
						description = "ReScript编译器"
					default:
						description = "ReScript相关工具"
					}

					packages = append(packages, PackageInfo{
						Name:        tool,
						Description: description,
						Version:     version,
						Installed:   true,
					})

					break
				}
			}
		}
	}

	// 检查本地项目中的包
	projectDirs := []string{
		".", "./src", "./lib/bs",
	}

	for _, dir := range projectDirs {
		// 检查bsconfig.json文件
		bsconfigPath := filepath.Join(dir, "bsconfig.json")
		if _, err := os.Stat(bsconfigPath); err == nil {
			content, err := os.ReadFile(bsconfigPath)
			if err != nil {
				continue
			}

			// 解析JSON内容
			var bsconfig struct {
				Dependencies   map[string]string `json:"dependencies"`
				BsDependencies []string          `json:"bs-dependencies"`
			}

			if err := json.Unmarshal(content, &bsconfig); err != nil {
				continue
			}

			// 添加依赖
			for pkg, version := range bsconfig.Dependencies {
				packages = append(packages, PackageInfo{
					Name:      pkg,
					Version:   version,
					Installed: true,
				})

				// 限制数量
				if len(packages) >= 20 {
					return packages, nil
				}
			}

			// 添加bs-dependencies
			for _, pkg := range bsconfig.BsDependencies {
				// 检查是否已经添加
				alreadyAdded := false
				for _, existingPkg := range packages {
					if existingPkg.Name == pkg {
						alreadyAdded = true
						break
					}
				}

				if !alreadyAdded {
					packages = append(packages, PackageInfo{
						Name:      pkg,
						Version:   "installed",
						Installed: true,
					})

					// 限制数量
					if len(packages) >= 20 {
						return packages, nil
					}
				}
			}
		}

		// 检查package.json
		packageJsonPath := filepath.Join(dir, "package.json")
		if _, err := os.Stat(packageJsonPath); err == nil {
			content, err := os.ReadFile(packageJsonPath)
			if err != nil {
				continue
			}

			// 解析JSON内容
			var packageJson struct {
				Dependencies    map[string]string `json:"dependencies"`
				DevDependencies map[string]string `json:"devDependencies"`
			}

			if err := json.Unmarshal(content, &packageJson); err != nil {
				continue
			}

			// 检查依赖中的ReScript相关包
			for _, deps := range []map[string]string{packageJson.Dependencies, packageJson.DevDependencies} {
				for pkg, version := range deps {
					if pkg == "rescript" || pkg == "bs-platform" ||
						strings.Contains(pkg, "rescript") || strings.Contains(pkg, "bs-") {
						// 检查是否已经添加
						alreadyAdded := false
						for _, existingPkg := range packages {
							if existingPkg.Name == pkg {
								alreadyAdded = true
								break
							}
						}

						if !alreadyAdded {
							packages = append(packages, PackageInfo{
								Name:      pkg,
								Version:   version,
								Installed: true,
							})

							// 限制数量
							if len(packages) >= 20 {
								return packages, nil
							}
						}
					}
				}
			}
		}
	}

	return packages, nil
}

// 检测Scratch
func (a *App) detectScratch() LanguageInfo {
	info := LanguageInfo{
		Name:            "Scratch",
		Installed:       false,
		DownloadURL:     "https://scratch.mit.edu/download",
		InstallTutorial: "https://scratch.mit.edu/download",
		PackageManager:  "N/A",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "Scratch Desktop",
				Description: "Scratch离线编辑器",
				Version:     "3.29.1",
			},
			{
				Name:        "ScratchJr",
				Description: "适合年幼儿童的Scratch",
				Version:     "latest",
			},
		},
	}

	// Scratch主要是基于Web的，没有命令行工具
	// 这里检查是否安装了Scratch Desktop
	if commandExists("scratch-desktop") {
		info.Installed = true
		info.Version = "Scratch Desktop (版本无法自动检测)"
	}

	return info
}

// 检测Vala
func (a *App) detectVala() LanguageInfo {
	info := LanguageInfo{
		Name:            "Vala",
		Installed:       false,
		DownloadURL:     "https://wiki.gnome.org/Projects/Vala",
		InstallTutorial: "https://wiki.gnome.org/Projects/Vala/Documentation",
		PackageManager:  "Meson",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "glib-2.0",
				Description: "GLib库",
				Version:     "2.76.3",
			},
			{
				Name:        "gtk+-3.0",
				Description: "GTK库",
				Version:     "3.24.38",
			},
			{
				Name:        "json-glib-1.0",
				Description: "JSON库",
				Version:     "1.6.6",
			},
			{
				Name:        "libsoup-2.4",
				Description: "HTTP客户端/服务器库",
				Version:     "2.74.3",
			},
		},
	}

	if !commandExists("valac") {
		return info
	}

	output, err := executeCommandWithTimeout("valac", "--version")
	if err == nil {
		info.Installed = true
		info.Version = output

		// 获取已安装的Vala包
		packages, _ := a.listValaPackages()
		info.Packages = packages
	}

	return info
}

// 列出已安装的Vala包
func (a *App) listValaPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 检查pkg-config是否安装
	if !commandExists("pkg-config") {
		return packages, nil
	}

	// 获取已安装的Vala库
	valaLibs := []string{
		"glib-2.0", "gobject-2.0", "gtk+-3.0", "gtk4", "json-glib-1.0",
		"libsoup-2.4", "libsoup-3.0", "gee-0.8", "gio-2.0", "gdk-3.0", "cairo",
		"pango", "sqlite3", "libxml-2.0", "gstreamer-1.0", "clutter-1.0",
	}

	for _, lib := range valaLibs {
		output, err := executeCommandWithTimeout("pkg-config", "--modversion", lib)
		if err == nil {
			packages = append(packages, PackageInfo{
				Name:      lib,
				Version:   strings.TrimSpace(output),
				Installed: true,
			})

			// 限制数量
			if len(packages) >= 20 {
				break
			}
		}
	}

	return packages, nil
}

// 检测XSLT
func (a *App) detectXSLT() LanguageInfo {
	info := LanguageInfo{
		Name:            "XSLT",
		Installed:       false,
		DownloadURL:     "http://xmlsoft.org/XSLT/",
		InstallTutorial: "http://xmlsoft.org/XSLT/tutorial/libxslttutorial.html",
		PackageManager:  "N/A",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "libxslt",
				Description: "XSLT库",
				Version:     "1.1.37",
			},
			{
				Name:        "Saxon-HE",
				Description: "XSLT和XQuery处理器",
				Version:     "11.4",
			},
		},
	}

	// 检查xsltproc
	if commandExists("xsltproc") {
		output, err := executeCommandWithTimeout("xsltproc", "--version")
		if err == nil {
			info.Installed = true
			info.Version = "xsltproc: " + output
			return info
		}
	}

	// 检查Saxon
	if commandExists("saxon") {
		output, err := executeCommandWithTimeout("saxon", "--version")
		if err == nil {
			info.Installed = true
			info.Version = "Saxon: " + output
			return info
		}
	}

	return info
}
