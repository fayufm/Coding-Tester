package main

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// 为ABAP添加包检测功能
func (a *App) detectABAPWithPackages() LanguageInfo {
	info := a.detectABAP()

	if info.Installed {
		// 获取已安装的ABAP包
		packages, _ := a.listABAPPackages()
		info.Packages = packages
	}

	return info
}

// 列出ABAP包
func (a *App) listABAPPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// ABAP通常在SAP系统中运行，尝试检查一些本地SAP GUI配置
	sapPaths := []string{
		"C:\\Program Files\\SAP\\FrontEnd\\SAPgui",
		"C:\\Program Files (x86)\\SAP\\FrontEnd\\SAPgui",
		"/usr/sap",
	}

	for _, path := range sapPaths {
		if _, err := os.Stat(path); err == nil {
			// 如果找到SAP目录，尝试扫描已安装的组件
			filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
				if err != nil {
					return nil
				}

				// 查找可能的ABAP组件目录
				if info.IsDir() && (strings.Contains(strings.ToLower(info.Name()), "addon") ||
					strings.Contains(strings.ToLower(info.Name()), "package") ||
					strings.Contains(strings.ToLower(info.Name()), "component")) {

					packages = append(packages, PackageInfo{
						Name:      info.Name(),
						Version:   "已安装",
						Installed: true,
					})

					// 限制返回的包数量
					if len(packages) >= 10 {
						return filepath.SkipDir
					}
				}
				return nil
			})
		}
	}

	// 如果没有找到实际安装的包，添加一些常见的SAP ABAP包
	if len(packages) == 0 {
		standardPackages := []struct {
			name        string
			description string
		}{
			{"ABAP Development Tools", "用于Eclipse的ABAP开发工具"},
			{"SAP GUI", "SAP系统的图形用户界面"},
			{"SAP NetWeaver", "应用和集成平台"},
			{"SAP S/4HANA", "企业资源规划套件"},
			{"SAP BW/4HANA", "数据仓库解决方案"},
			{"SAP Cloud Platform ABAP Environment", "云平台ABAP环境"},
			{"SAP HANA Studio", "SAP HANA的管理工具"},
			{"SAPUI5", "用户界面开发工具包"},
			{"SAP Fiori", "用户体验设计"},
			{"SAP BusinessObjects", "商业智能工具"},
		}

		for _, pkg := range standardPackages {
			packages = append(packages, PackageInfo{
				Name:        pkg.name,
				Version:     "N/A",
				Description: pkg.description,
				Installed:   false,
			})
		}
	}

	return packages, nil
}

// 为ActionScript添加包检测功能
func (a *App) detectActionScriptWithPackages() LanguageInfo {
	info := a.detectActionScript()

	if info.Installed {
		// 获取已安装的ActionScript包
		packages, _ := a.listActionScriptPackages()
		info.Packages = packages
	}

	return info
}

// 列出ActionScript包
func (a *App) listActionScriptPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 检查Adobe Animate或Flash的安装目录
	animatePaths := []string{
		"C:\\Program Files\\Adobe\\Adobe Animate*",
		"C:\\Program Files (x86)\\Adobe\\Adobe Animate*",
		"/Applications/Adobe Animate*.app",
		filepath.Join(os.Getenv("HOME"), "Adobe", "Adobe Animate*"),
	}

	// 检查Apache Flex SDK的安装
	flexPaths := []string{
		"C:\\Program Files\\Adobe\\Flex SDK*",
		"C:\\Program Files (x86)\\Adobe\\Flex SDK*",
		"/opt/flex_sdk*",
		filepath.Join(os.Getenv("HOME"), "flex_sdk*"),
	}

	// 合并搜索路径
	searchPaths := append(animatePaths, flexPaths...)

	for _, pattern := range searchPaths {
		matches, _ := filepath.Glob(pattern)
		for _, dir := range matches {
			// 搜索libs或frameworks目录
			for _, subdir := range []string{"libs", "frameworks", "frameworks/libs"} {
				libsDir := filepath.Join(dir, subdir)
				if _, err := os.Stat(libsDir); err == nil {
					entries, err := os.ReadDir(libsDir)
					if err == nil {
						for _, entry := range entries {
							if !entry.IsDir() && (strings.HasSuffix(strings.ToLower(entry.Name()), ".swc") ||
								strings.HasSuffix(strings.ToLower(entry.Name()), ".swf")) {

								// 从文件名提取库名称和版本
								name := strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name()))
								version := "已安装"

								// 尝试从名称中提取版本
								versionRegex := regexp.MustCompile(`[-_](\d+(\.\d+)*)$`)
								matches := versionRegex.FindStringSubmatch(name)
								if len(matches) > 1 {
									version = matches[1]
									name = strings.TrimSuffix(name, matches[0])
								}

								packages = append(packages, PackageInfo{
									Name:      name,
									Version:   version,
									Installed: true,
								})

								// 限制返回的包数量
								if len(packages) >= 20 {
									return packages, nil
								}
							}
						}
					}
				}
			}
		}
	}

	// 如果没有找到实际安装的包，添加一些标准库
	if len(packages) == 0 {
		standardLibs := []struct {
			name    string
			version string
		}{
			{"fl.accessibility", "Standard"},
			{"fl.containers", "Standard"},
			{"fl.controls", "Standard"},
			{"fl.core", "Standard"},
			{"fl.data", "Standard"},
			{"fl.events", "Standard"},
			{"fl.livepreview", "Standard"},
			{"fl.managers", "Standard"},
			{"fl.motion", "Standard"},
			{"fl.transitions", "Standard"},
			{"fl.video", "Standard"},
			{"flash.accessibility", "Standard"},
			{"flash.display", "Standard"},
			{"flash.events", "Standard"},
			{"flash.external", "Standard"},
			{"flash.filters", "Standard"},
			{"flash.geom", "Standard"},
			{"flash.media", "Standard"},
			{"flash.net", "Standard"},
			{"flash.printing", "Standard"},
			{"flash.system", "Standard"},
			{"flash.text", "Standard"},
			{"flash.ui", "Standard"},
			{"flash.utils", "Standard"},
			{"flash.xml", "Standard"},
		}

		for _, lib := range standardLibs {
			packages = append(packages, PackageInfo{
				Name:      lib.name,
				Version:   lib.version,
				Installed: false,
			})
			if len(packages) >= 20 {
				break
			}
		}
	}

	return packages, nil
}

// 为APL添加包检测功能
func (a *App) detectAPLWithPackages() LanguageInfo {
	info := a.detectAPL()

	if info.Installed {
		// 获取已安装的APL包
		packages, _ := a.listAPLPackages()
		info.Packages = packages
	}

	return info
}

// 列出APL包
func (a *App) listAPLPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 检查Dyalog APL安装目录
	dyalogPaths := []string{
		"C:\\Program Files\\Dyalog\\Dyalog APL*",
		"C:\\Program Files (x86)\\Dyalog\\Dyalog APL*",
		"/opt/mdyalog/*",
		"/Applications/Dyalog*.app",
		filepath.Join(os.Getenv("HOME"), ".dyalog"),
	}

	for _, pattern := range dyalogPaths {
		matches, _ := filepath.Glob(pattern)
		for _, dir := range matches {
			// 搜索workspaces目录
			workspacesDir := filepath.Join(dir, "Workspaces")
			if _, err := os.Stat(workspacesDir); err == nil {
				entries, err := os.ReadDir(workspacesDir)
				if err == nil {
					for _, entry := range entries {
						if !entry.IsDir() && (strings.HasSuffix(strings.ToLower(entry.Name()), ".dws") ||
							strings.HasSuffix(strings.ToLower(entry.Name()), ".dyalog")) {

							name := strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name()))
							packages = append(packages, PackageInfo{
								Name:      name,
								Version:   "Dyalog工作区",
								Installed: true,
							})

							// 限制返回的包数量
							if len(packages) >= 20 {
								return packages, nil
							}
						}
					}
				}
			}

			// 搜索库目录
			libDir := filepath.Join(dir, "lib")
			if _, err := os.Stat(libDir); err == nil {
				entries, err := os.ReadDir(libDir)
				if err == nil {
					for _, entry := range entries {
						if !entry.IsDir() && (strings.HasSuffix(strings.ToLower(entry.Name()), ".dyalog") ||
							strings.HasSuffix(strings.ToLower(entry.Name()), ".dws") ||
							strings.HasSuffix(strings.ToLower(entry.Name()), ".apl")) {

							name := strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name()))
							packages = append(packages, PackageInfo{
								Name:      name,
								Version:   "系统库",
								Installed: true,
							})

							// 限制返回的包数量
							if len(packages) >= 20 {
								return packages, nil
							}
						}
					}
				}
			}
		}
	}

	// 如果没有找到实际安装的包，添加一些标准库
	if len(packages) == 0 {
		standardWorkspaces := []string{
			"dfns", "display", "eval", "groups", "plot", "regexp", "stats", "unicode", "utils", "ws2json",
		}

		for _, ws := range standardWorkspaces {
			packages = append(packages, PackageInfo{
				Name:      ws,
				Version:   "标准工作区",
				Installed: false,
			})
		}
	}

	return packages, nil
}

// 为Ballerina添加包检测功能
func (a *App) detectBallerinaWithPackages() LanguageInfo {
	info := a.detectBallerina()

	if info.Installed {
		// 获取已安装的Ballerina包
		packages, _ := a.listBallerinaPackages()
		info.Packages = packages
	}

	return info
}

// 列出Ballerina包
func (a *App) listBallerinaPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 使用Ballerina命令获取已安装的包
	if commandExists("bal") {
		output, err := executeCommandWithTimeout("bal", "list")
		if err == nil && len(output) > 0 {
			lines := strings.Split(output, "\n")
			for _, line := range lines {
				line = strings.TrimSpace(line)
				if line != "" && !strings.HasPrefix(line, "Compiling") && !strings.HasPrefix(line, "Running") {
					parts := strings.Fields(line)
					if len(parts) >= 2 {
						packages = append(packages, PackageInfo{
							Name:      parts[0],
							Version:   parts[1],
							Installed: true,
						})
					} else if len(parts) == 1 {
						packages = append(packages, PackageInfo{
							Name:      parts[0],
							Version:   "已安装",
							Installed: true,
						})
					}

					// 限制返回的包数量
					if len(packages) >= 20 {
						break
					}
				}
			}
		}

		// 检查中央存储库中的包
		output, err = executeCommandWithTimeout("bal", "search", "--limit", "20")
		if err == nil && len(output) > 0 && len(packages) < 20 {
			lines := strings.Split(output, "\n")
			for _, line := range lines {
				line = strings.TrimSpace(line)
				if line != "" && !strings.HasPrefix(line, "Compiling") && !strings.HasPrefix(line, "Running") && !strings.Contains(line, "NAME") {
					parts := strings.Fields(line)
					if len(parts) >= 2 {
						// 检查是否已经添加过
						exists := false
						for _, pkg := range packages {
							if pkg.Name == parts[0] {
								exists = true
								break
							}
						}

						if !exists {
							packages = append(packages, PackageInfo{
								Name:      parts[0],
								Version:   parts[1],
								Installed: false,
							})

							// 限制返回的包数量
							if len(packages) >= 20 {
								break
							}
						}
					}
				}
			}
		}
	}

	// 如果没有找到实际的包，添加一些标准库
	if len(packages) == 0 {
		standardLibs := []struct {
			name        string
			description string
		}{
			{"ballerina/http", "HTTP客户端和服务器实现"},
			{"ballerina/io", "I/O API"},
			{"ballerina/log", "日志API"},
			{"ballerina/math", "数学函数"},
			{"ballerina/mime", "MIME实用工具"},
			{"ballerina/oauth2", "OAuth 2.0支持"},
			{"ballerina/observe", "可观察性支持"},
			{"ballerina/sql", "SQL客户端"},
			{"ballerina/task", "任务调度"},
			{"ballerina/time", "时间相关函数"},
			{"ballerina/websocket", "WebSocket支持"},
			{"ballerina/grpc", "gRPC支持"},
			{"ballerina/crypto", "加密函数"},
			{"ballerina/email", "电子邮件客户端"},
			{"ballerina/file", "文件系统交互"},
			{"ballerina/jwt", "JWT处理"},
			{"ballerina/kafka", "Kafka客户端"},
			{"ballerina/kubernetes", "Kubernetes集成"},
			{"ballerina/mysql", "MySQL客户端"},
			{"ballerina/rabbitmq", "RabbitMQ客户端"},
		}

		for _, lib := range standardLibs {
			packages = append(packages, PackageInfo{
				Name:        lib.name,
				Version:     "标准库",
				Description: lib.description,
				Installed:   false,
			})
		}
	}

	return packages, nil
}

// 为BASIC添加包检测功能
func (a *App) detectBASICWithPackages() LanguageInfo {
	info := a.detectBASIC()

	if info.Installed {
		// 获取已安装的BASIC包
		packages, _ := a.listBASICPackages()
		info.Packages = packages
	}

	return info
}

// 列出BASIC包
func (a *App) listBASICPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 检查FreeBASIC安装
	if commandExists("fbc") {
		// 检查FreeBASIC包目录
		fbcPaths := []string{
			"C:\\Program Files\\FreeBASIC\\inc",
			"C:\\Program Files (x86)\\FreeBASIC\\inc",
			"/usr/local/include/freebasic",
			"/usr/include/freebasic",
		}

		for _, dir := range fbcPaths {
			if _, err := os.Stat(dir); err == nil {
				// 扫描包目录
				filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
					if err != nil {
						return nil
					}

					if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".bi") {
						name := strings.TrimSuffix(info.Name(), ".bi")

						// 检查是否已经添加过
						exists := false
						for _, pkg := range packages {
							if pkg.Name == name {
								exists = true
								break
							}
						}

						if !exists {
							packages = append(packages, PackageInfo{
								Name:      name,
								Version:   "FreeBASIC",
								Installed: true,
							})

							// 限制返回的包数量
							if len(packages) >= 20 {
								return filepath.SkipDir
							}
						}
					}
					return nil
				})
			}
		}
	}

	// 检查QB64安装
	qb64Paths := []string{
		"C:\\QB64\\internal\\c",
		filepath.Join(os.Getenv("HOME"), "qb64/internal/c"),
		"/usr/local/share/qb64/internal/c",
	}

	for _, dir := range qb64Paths {
		if _, err := os.Stat(dir); err == nil {
			entries, err := os.ReadDir(dir)
			if err == nil {
				for _, entry := range entries {
					if !entry.IsDir() && strings.HasSuffix(strings.ToLower(entry.Name()), ".h") {
						name := strings.TrimSuffix(entry.Name(), ".h")

						// 检查是否已经添加过
						exists := false
						for _, pkg := range packages {
							if pkg.Name == name {
								exists = true
								break
							}
						}

						if !exists {
							packages = append(packages, PackageInfo{
								Name:      name,
								Version:   "QB64",
								Installed: true,
							})

							// 限制返回的包数量
							if len(packages) >= 20 {
								return packages, nil
							}
						}
					}
				}
			}
		}
	}

	// 如果没有找到实际的包，添加一些标准库
	if len(packages) == 0 {
		// FreeBASIC标准库
		fbLibs := []string{
			"crt", "datetime", "file", "graphics", "math", "string", "system",
			"windows", "win32api", "opengl", "gfxlib2", "fbgfx", "dos", "devices",
		}

		for _, lib := range fbLibs {
			packages = append(packages, PackageInfo{
				Name:      lib,
				Version:   "FreeBASIC标准库",
				Installed: false,
			})
		}

		// QB64标准库
		qbLibs := []string{
			"graphics", "sound", "keyboard", "mouse", "file", "directory", "screen",
			"math", "string", "array", "date", "time", "inkey", "print", "input",
		}

		for _, lib := range qbLibs {
			// 检查是否已经添加过
			exists := false
			for _, pkg := range packages {
				if pkg.Name == lib {
					exists = true
					break
				}
			}

			if !exists {
				packages = append(packages, PackageInfo{
					Name:      lib,
					Version:   "QB64标准库",
					Installed: false,
				})

				// 限制返回的包数量
				if len(packages) >= 20 {
					break
				}
			}
		}
	}

	return packages, nil
}

// 为LOLCODE添加包检测功能
func (a *App) detectLOLCODEWithPackages() LanguageInfo {
	info := a.detectLOLCODE()

	if info.Installed {
		// 获取已安装的LOLCODE包
		packages, _ := a.listLOLCODEPackages()
		info.Packages = packages
	}

	return info
}

// 列出LOLCODE包
func (a *App) listLOLCODEPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// LOLCODE没有标准的包管理器，返回一些基本的标准库功能
	standardFunctions := []struct {
		name        string
		description string
	}{
		{"STDIO", "标准输入输出"},
		{"STDLIB", "标准库函数"},
		{"SMOOSH", "字符串连接"},
		{"STRCMP", "字符串比较"},
		{"MATHZ", "数学运算"},
		{"LOOPZ", "循环结构"},
		{"FILZ", "文件操作"},
		{"RANDOMZ", "随机数生成"},
		{"TIMEZ", "时间函数"},
		{"ARRAYIFICATION", "数组操作"},
	}

	for _, fn := range standardFunctions {
		packages = append(packages, PackageInfo{
			Name:        fn.name,
			Version:     "1.0",
			Description: fn.description,
			Installed:   false,
		})
	}

	return packages, nil
}
