package main

import (
	"os"
	"path/filepath"
	"strings"
)

// 为Q#添加包检测功能
func (a *App) detectQSharpWithPackages() LanguageInfo {
	info := a.detectQSharp()

	if info.Installed {
		// 获取已安装的Q#包
		packages, _ := a.listQSharpPackages()
		info.Packages = packages
	}

	return info
}

// 列出Q#包
func (a *App) listQSharpPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 如果dotnet命令存在
	if commandExists("dotnet") {
		// 尝试检查已安装的Quantum开发工具包
		output, err := executeCommandWithTimeout("dotnet", "tool", "list", "--global")
		if err == nil {
			lines := strings.Split(output, "\n")
			for _, line := range lines {
				if strings.Contains(line, "quantum") || strings.Contains(line, "qsharp") {
					parts := strings.Fields(line)
					if len(parts) >= 2 {
						packages = append(packages, PackageInfo{
							Name:      parts[0],
							Version:   parts[1],
							Installed: true,
						})
					}
				}
			}
		}

		// 尝试检查当前目录中的Q#项目
		if _, err := os.Stat("*.csproj"); err == nil {
			// 使用dotnet list查看项目引用
			output, err := executeCommandWithTimeout("dotnet", "list", "package")
			if err == nil {
				lines := strings.Split(output, "\n")
				inPackageSection := false
				for _, line := range lines {
					line = strings.TrimSpace(line)

					if strings.HasPrefix(line, "Top-level Package") {
						inPackageSection = true
						continue
					}

					if inPackageSection && line != "" && !strings.HasPrefix(line, ">") && !strings.HasPrefix(line, "Project") {
						parts := strings.Fields(line)
						if len(parts) >= 2 && (strings.Contains(parts[0], "Quantum") || strings.Contains(parts[0], "Microsoft.Quantum")) {
							packages = append(packages, PackageInfo{
								Name:      parts[0],
								Version:   parts[1],
								Installed: true,
							})
						}
					}
				}
			}
		}
	}

	// 如果没有找到任何包，添加标准的Quantum包
	if len(packages) == 0 {
		standardPackages := []struct {
			name        string
			version     string
			description string
		}{
			{"Microsoft.Quantum.Standard", "0.28.302812", "Q#标准库"},
			{"Microsoft.Quantum.Development.Kit", "0.28.302812", "Q#开发工具包"},
			{"Microsoft.Quantum.Numerics", "0.28.302812", "数值计算库"},
			{"Microsoft.Quantum.Simulation.Runtime", "0.28.302812", "量子模拟运行时"},
			{"Microsoft.Quantum.EntryPoint", "0.28.302812", "Q#入口点实用程序"},
			{"Microsoft.Quantum.QSharp.Core", "0.28.302812", "Q# Core库"},
			{"Microsoft.Quantum.Runtime", "0.28.302812", "Q#运行时库"},
			{"Microsoft.Quantum.Compiler", "0.28.302812", "Q#编译器"},
			{"Microsoft.Quantum.Type.Operators", "0.28.302812", "Q#类型操作符"},
			{"Microsoft.Quantum.Diagnostics", "0.28.302812", "诊断工具"},
			{"Microsoft.Quantum.Chemistry", "0.28.302812", "量子化学库"},
			{"Microsoft.Quantum.Measurement", "0.28.302812", "量子测量库"},
			{"Microsoft.Quantum.Convert", "0.28.302812", "类型转换库"},
		}

		for _, pkg := range standardPackages {
			packages = append(packages, PackageInfo{
				Name:        pkg.name,
				Version:     pkg.version,
				Description: pkg.description,
				Installed:   false,
			})
		}
	}

	return packages, nil
}

// 为Red添加包检测功能
func (a *App) detectRedWithPackages() LanguageInfo {
	info := a.detectRed()

	if info.Installed {
		// 获取已安装的Red包
		packages, _ := a.listRedPackages()
		info.Packages = packages
	}

	return info
}

// 列出Red包
func (a *App) listRedPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 如果Red命令存在
	if commandExists("red") {
		// 尝试获取已安装的模块
		output, err := executeCommandWithTimeout("red", "-h")
		if err == nil {
			// Red没有标准的包管理器，但我们可以检查是否有某些标准模块可用
			if strings.Contains(output, "modules") || strings.Contains(output, "import") {
				// 尝试查找Red用户目录中的模块
				userDir, err := os.UserHomeDir()
				if err == nil {
					redPaths := []string{
						filepath.Join(userDir, ".red"),
						filepath.Join(userDir, "red"),
						filepath.Join(userDir, ".config", "red"),
					}

					for _, dir := range redPaths {
						if _, err := os.Stat(dir); err == nil {
							// 遍历查找.red文件
							err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
								if err != nil {
									return nil
								}

								if !info.IsDir() && (strings.HasSuffix(path, ".red") || strings.HasSuffix(path, ".reds")) {
									name := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))

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
											Version:   "用户模块",
											Installed: true,
										})
									}
								}
								return nil
							})

							if err == nil && len(packages) > 0 {
								break
							}
						}
					}
				}
			}
		}
	}

	// 如果没有找到任何模块，添加标准Red模块
	if len(packages) == 0 {
		standardModules := []struct {
			name        string
			description string
		}{
			{"redbin", "Red二进制格式"},
			{"view", "GUI系统"},
			{"parse", "解析方言"},
			{"object", "对象系统"},
			{"reactivity", "响应式编程支持"},
			{"redbin", "二进制序列化格式"},
			{"mezz", "中间辅助函数"},
			{"help", "帮助系统"},
			{"debug", "调试工具"},
			{"face", "图形界面元素"},
			{"vid", "可视化界面方言"},
			{"draw", "绘图DSL"},
			{"vector", "向量图形"},
			{"shape", "形状DSL"},
			{"pair", "坐标处理"},
			{"event", "事件系统"},
			{"font", "字体处理"},
			{"text", "文本处理"},
			{"image", "图像处理"},
			{"color", "颜色处理"},
		}

		for _, module := range standardModules {
			packages = append(packages, PackageInfo{
				Name:        module.name,
				Version:     "内置",
				Description: module.description,
				Installed:   false,
			})
		}
	}

	return packages, nil
}

// 为Scratch添加包检测功能
func (a *App) detectScratchWithPackages() LanguageInfo {
	info := a.detectScratch()

	if info.Installed {
		// 获取已安装的Scratch扩展
		packages, _ := a.listScratchExtensions()
		info.Packages = packages
	}

	return info
}

// 列出Scratch扩展
func (a *App) listScratchExtensions() ([]PackageInfo, error) {
	var packages []PackageInfo

	// Scratch没有传统意义上的包，但有扩展和库
	// 检查是否有Scratch Desktop或相关应用
	scratchPaths := []string{
		"C:\\Program Files\\Scratch Desktop",
		"C:\\Program Files (x86)\\Scratch Desktop",
		"/Applications/Scratch Desktop.app",
		"/Applications/Scratch 3.app",
		filepath.Join(os.Getenv("HOME"), "Applications", "Scratch Desktop.app"),
		filepath.Join(os.Getenv("HOME"), "Applications", "Scratch 3.app"),
	}

	for _, path := range scratchPaths {
		if _, err := os.Stat(path); err == nil {
			// 找到Scratch安装，添加默认扩展
			packages = append(packages, PackageInfo{
				Name:      "Scratch Desktop",
				Version:   "已安装",
				Installed: true,
			})
			break
		}
	}

	// 添加标准Scratch扩展
	standardExtensions := []struct {
		name        string
		description string
	}{
		{"pen", "绘图功能"},
		{"music", "音乐功能"},
		{"text-to-speech", "文字转语音"},
		{"translate", "翻译功能"},
		{"video-sensing", "视频感应"},
		{"microbit", "micro:bit支持"},
		{"ev3", "LEGO MINDSTORMS EV3支持"},
		{"boost", "LEGO BOOST支持"},
		{"wedo2", "LEGO WeDo 2.0支持"},
		{"makeymakey", "Makey Makey支持"},
		{"gdxfor", "LEGO Force & Acceleration传感器支持"},
		{"tello", "Tello无人机支持"},
		{"ml", "机器学习支持"},
		{"speech", "语音识别"},
	}

	for _, ext := range standardExtensions {
		packages = append(packages, PackageInfo{
			Name:        ext.name,
			Version:     "标准扩展",
			Description: ext.description,
			Installed:   false,
		})
	}

	// 添加第三方扩展库
	thirdPartyExtensions := []struct {
		name        string
		description string
	}{
		{"TurboWarp", "Scratch的高性能分支"},
		{"ClipCC", "Scratch修改版"},
		{"PenguinMod", "扩展的Scratch修改版"},
		{"SheepTester's Extensions", "社区扩展集"},
		{"Snap!", "类Scratch可视化编程语言"},
		{"griffpatch's Extensions", "社区成员griffpatch的扩展"},
		{"E羊icques Extensions", "社区扩展集"},
		{"Turbowarp Extensions", "Turbowarp扩展集"},
	}

	for _, ext := range thirdPartyExtensions {
		packages = append(packages, PackageInfo{
			Name:        ext.name,
			Version:     "第三方",
			Description: ext.description,
			Installed:   false,
		})
	}

	return packages, nil
}
