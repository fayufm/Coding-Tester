package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

// 检测C/C++
func (a *App) detectCpp() LanguageInfo {
	info := LanguageInfo{
		Name:            "C/C++",
		Installed:       false,
		DownloadURL:     "https://visualstudio.microsoft.com/downloads/",
		InstallTutorial: "https://docs.microsoft.com/en-us/cpp/build/vscpp-step-0-installation",
		PackageManager:  "vcpkg/conan",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "boost",
				Description: "C++的可移植库，提供从线程到数学、算法等各种功能",
				Version:     "1.82.0",
			},
			{
				Name:        "opencv",
				Description: "开源计算机视觉库",
				Version:     "4.7.0",
			},
			{
				Name:        "eigen",
				Description: "C++的线性代数库",
				Version:     "3.4.0",
			},
			{
				Name:        "qt",
				Description: "跨平台应用程序和UI框架",
				Version:     "6.5.1",
			},
			{
				Name:        "sfml",
				Description: "简单快速的多媒体库",
				Version:     "2.6.0",
			},
		},
	}

	// 检测GCC
	if commandExists("gcc") {
		output, err := executeCommandWithTimeout("gcc", "--version")
		if err == nil {
			info.Installed = true
			info.Version = "GCC: " + strings.Split(output, "\n")[0]

			// 获取已安装的C/C++包
			packages, _ := a.listCppPackages()
			info.Packages = packages

			return info
		}
	}

	// 检测Clang
	if commandExists("clang") {
		output, err := executeCommandWithTimeout("clang", "--version")
		if err == nil {
			info.Installed = true
			info.Version = "Clang: " + strings.Split(output, "\n")[0]

			// 获取已安装的C/C++包
			packages, _ := a.listCppPackages()
			info.Packages = packages

			return info
		}
	}

	// 检测MSVC (Windows)
	if commandExists("cl") {
		output, err := executeCommandWithTimeout("cl")
		if err == nil {
			info.Installed = true
			info.Version = "MSVC: " + output

			// 获取已安装的C/C++包
			packages, _ := a.listCppPackages()
			info.Packages = packages

			return info
		}
	}

	return info
}

// 列出已安装的C/C++包
func (a *App) listCppPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 检查vcpkg是否安装
	if commandExists("vcpkg") {
		output, err := executeCommandWithTimeout("vcpkg", "list")
		if err == nil {
			lines := strings.Split(output, "\n")
			for _, line := range lines {
				line = strings.TrimSpace(line)
				if line == "" || strings.HasPrefix(line, "vcpkg") || strings.Contains(line, "packages are compatible") {
					continue
				}

				parts := strings.Fields(line)
				if len(parts) >= 2 {
					name := parts[0]
					version := parts[1]
					// 移除可能的特殊字符
					name = strings.TrimSuffix(name, ":")

					packages = append(packages, PackageInfo{
						Name:      name,
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

	// 检查conan是否安装
	if len(packages) == 0 && commandExists("conan") {
		output, err := executeCommandWithTimeout("conan", "list", "*:*", "--format=json")
		if err == nil {
			var result map[string]interface{}
			if err := json.Unmarshal([]byte(output), &result); err == nil {
				if items, ok := result["results"].([]interface{}); ok {
					for _, item := range items {
						if pkg, ok := item.(map[string]interface{}); ok {
							name, _ := pkg["name"].(string)
							version, _ := pkg["version"].(string)

							if name != "" && version != "" {
								packages = append(packages, PackageInfo{
									Name:      name,
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

		// 如果JSON格式失败，尝试普通列表
		if len(packages) == 0 {
			output, err := executeCommandWithTimeout("conan", "search", "--raw")
			if err == nil {
				lines := strings.Split(output, "\n")
				for _, line := range lines {
					line = strings.TrimSpace(line)
					if line == "" {
						continue
					}

					// 尝试提取名称和版本
					if strings.Contains(line, "/") {
						parts := strings.Split(line, "/")
						if len(parts) >= 2 {
							name := parts[0]
							version := parts[1]

							packages = append(packages, PackageInfo{
								Name:      name,
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

	// 检查系统库目录
	if len(packages) == 0 {
		// 检查常见的系统库目录
		libDirs := []string{
			"/usr/lib",
			"/usr/local/lib",
			"C:/Program Files/Common Files",
			"C:/Program Files (x86)/Common Files",
		}

		for _, dir := range libDirs {
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				continue
			}

			// 检查常见的C/C++库
			commonLibs := []struct {
				name    string
				pattern string
			}{
				{"boost", "boost*"},
				{"opencv", "opencv*"},
				{"qt", "qt*"},
				{"eigen", "eigen*"},
				{"sfml", "sfml*"},
				{"glfw", "glfw*"},
				{"glew", "glew*"},
				{"sdl", "sdl*"},
			}

			for _, lib := range commonLibs {
				matches, _ := filepath.Glob(filepath.Join(dir, lib.pattern))
				if len(matches) > 0 {
					// 尝试提取版本信息
					version := "installed"
					for _, match := range matches {
						baseName := filepath.Base(match)
						// 尝试从名称中提取版本号
						re := regexp.MustCompile(`\d+(\.\d+)+`)
						versionMatch := re.FindString(baseName)
						if versionMatch != "" {
							version = versionMatch
							break
						}
					}

					packages = append(packages, PackageInfo{
						Name:      lib.name,
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

	return packages, nil
}

// 检测Swift
func (a *App) detectSwift() LanguageInfo {
	info := LanguageInfo{
		Name:            "Swift",
		Installed:       false,
		DownloadURL:     "https://swift.org/download/",
		InstallTutorial: "https://swift.org/getting-started/",
		PackageManager:  "Swift Package Manager",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "Alamofire",
				Description: "Swift的优雅网络库",
				Version:     "5.7.1",
			},
			{
				Name:        "SwiftyJSON",
				Description: "处理JSON数据的更好方式",
				Version:     "5.0.1",
			},
			{
				Name:        "Kingfisher",
				Description: "强大的纯Swift库，用于从网络下载和缓存图像",
				Version:     "7.8.1",
			},
			{
				Name:        "SnapKit",
				Description: "Swift的自动布局DSL",
				Version:     "5.6.0",
			},
			{
				Name:        "RxSwift",
				Description: "Swift的响应式编程",
				Version:     "6.5.0",
			},
		},
	}

	if !commandExists("swift") {
		return info
	}

	output, err := executeCommandWithTimeout("swift", "--version")
	if err == nil {
		info.Installed = true
		info.Version = output

		// 检测已安装的Swift包
		packages, _ := a.listSwiftPackages()
		info.Packages = packages
	}

	return info
}

// 列出已安装的Swift包
func (a *App) listSwiftPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 检查是否有Swift项目目录
	// 创建临时目录
	tempDir, err := os.MkdirTemp("", "swift-pkg-check")
	if err != nil {
		return packages, err
	}
	defer os.RemoveAll(tempDir)

	// 初始化Swift包
	initCmd := exec.Command("swift", "package", "init")
	initCmd.Dir = tempDir
	if err := initCmd.Run(); err != nil {
		// 如果无法初始化包，尝试查找现有的Swift项目
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return packages, err
		}

		// 检查常见的Swift项目目录
		swiftDirs := []string{
			filepath.Join(homeDir, "Documents", "SwiftProjects"),
			filepath.Join(homeDir, "Projects", "Swift"),
			filepath.Join(homeDir, "Developer", "Swift"),
		}

		for _, dir := range swiftDirs {
			if _, err := os.Stat(dir); err == nil {
				// 找到可能的Swift项目目录，尝试列出包
				files, err := os.ReadDir(dir)
				if err != nil {
					continue
				}

				for _, file := range files {
					if file.IsDir() {
						projectDir := filepath.Join(dir, file.Name())
						if _, err := os.Stat(filepath.Join(projectDir, "Package.swift")); err == nil {
							// 找到Swift包，尝试获取依赖信息
							cmd := exec.Command("swift", "package", "show-dependencies", "--format", "json")
							cmd.Dir = projectDir
							output, err := cmd.Output()
							if err == nil {
								// 解析JSON输出
								var result map[string]interface{}
								if err := json.Unmarshal(output, &result); err == nil {
									if dependencies, ok := result["dependencies"].([]interface{}); ok {
										for _, dep := range dependencies {
											if depMap, ok := dep.(map[string]interface{}); ok {
												name, _ := depMap["name"].(string)
												version, _ := depMap["version"].(string)
												if name != "" {
													packages = append(packages, PackageInfo{
														Name:      name,
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
						}
					}
				}
			}
		}

		return packages, nil
	}

	// 添加一些常用依赖
	packageSwift := `
// swift-tools-version:5.5
import PackageDescription

let package = Package(
    name: "SwiftPkgCheck",
    dependencies: [
        .package(url: "https://github.com/Alamofire/Alamofire.git", from: "5.0.0"),
        .package(url: "https://github.com/SwiftyJSON/SwiftyJSON.git", from: "5.0.0"),
    ],
    targets: [
        .target(
            name: "SwiftPkgCheck",
            dependencies: ["Alamofire", "SwiftyJSON"]),
    ]
)
`

	err = os.WriteFile(filepath.Join(tempDir, "Package.swift"), []byte(packageSwift), 0644)
	if err != nil {
		return packages, err
	}

	// 尝试解析依赖
	cmd := exec.Command("swift", "package", "resolve")
	cmd.Dir = tempDir
	if err := cmd.Run(); err != nil {
		return packages, err
	}

	// 获取依赖信息
	cmd = exec.Command("swift", "package", "show-dependencies", "--format", "json")
	cmd.Dir = tempDir
	output, err := cmd.Output()
	if err != nil {
		return packages, err
	}

	// 解析JSON输出
	var result map[string]interface{}
	if err := json.Unmarshal(output, &result); err != nil {
		return packages, err
	}

	if dependencies, ok := result["dependencies"].([]interface{}); ok {
		for _, dep := range dependencies {
			if depMap, ok := dep.(map[string]interface{}); ok {
				name, _ := depMap["name"].(string)
				version, _ := depMap["version"].(string)
				if name != "" {
					packages = append(packages, PackageInfo{
						Name:      name,
						Version:   version,
						Installed: true,
					})
				}
			}
		}
	}

	return packages, nil
}

// 检测Kotlin
func (a *App) detectKotlin() LanguageInfo {
	info := LanguageInfo{
		Name:            "Kotlin",
		Installed:       false,
		DownloadURL:     "https://kotlinlang.org/docs/command-line.html",
		InstallTutorial: "https://kotlinlang.org/docs/tutorials/command-line.html",
		PackageManager:  "Gradle",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "kotlinx.coroutines",
				Description: "Kotlin协程库",
				Version:     "1.7.1",
			},
			{
				Name:        "ktor",
				Description: "Kotlin的异步Web框架",
				Version:     "2.3.1",
			},
			{
				Name:        "exposed",
				Description: "Kotlin SQL框架",
				Version:     "0.41.1",
			},
			{
				Name:        "arrow-kt",
				Description: "Kotlin的函数式编程库",
				Version:     "1.2.0",
			},
			{
				Name:        "koin",
				Description: "Kotlin的依赖注入框架",
				Version:     "3.4.0",
			},
		},
	}

	if !commandExists("kotlin") {
		return info
	}

	output, err := executeCommandWithTimeout("kotlin", "-version")
	if err == nil {
		info.Installed = true
		info.Version = output
	}

	return info
}

// 检测Dart
func (a *App) detectDart() LanguageInfo {
	info := LanguageInfo{
		Name:            "Dart",
		Installed:       false,
		DownloadURL:     "https://dart.dev/get-dart",
		InstallTutorial: "https://dart.dev/tutorials/server/get-started",
		PackageManager:  "pub",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "flutter",
				Description: "Google的UI工具包，用于从单一代码库构建应用",
				Version:     "3.10.5",
			},
			{
				Name:        "http",
				Description: "用于HTTP请求的可组合库",
				Version:     "1.1.0",
			},
			{
				Name:        "provider",
				Description: "Flutter的状态管理库",
				Version:     "6.0.5",
			},
			{
				Name:        "dio",
				Description: "强大的Dart HTTP客户端",
				Version:     "5.2.1+1",
			},
			{
				Name:        "hive",
				Description: "轻量级且快速的键值数据库",
				Version:     "2.2.3",
			},
		},
	}

	if !commandExists("dart") {
		return info
	}

	output, err := executeCommandWithTimeout("dart", "--version")
	if err == nil {
		info.Installed = true
		info.Version = output

		// 检测已安装的Dart包
		packages, _ := a.listDartPackages()
		info.Packages = packages
	}

	return info
}

// 列出已安装的Dart包
func (a *App) listDartPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 检查pub命令是否可用
	if !commandExists("pub") && !commandExists("dart") {
		return packages, nil
	}

	// 尝试使用dart pub命令
	var output string
	var err error

	if commandExists("dart") {
		output, err = executeCommandWithTimeout("dart", "pub", "global", "list", "--json")
	} else {
		output, err = executeCommandWithTimeout("pub", "global", "list", "--json")
	}

	if err != nil {
		// 如果无法获取全局包，尝试查找Flutter/Dart项目
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return packages, err
		}

		// 检查Flutter/Dart项目目录
		dartDirs := []string{
			filepath.Join(homeDir, "Documents", "Flutter"),
			filepath.Join(homeDir, "Documents", "Dart"),
			filepath.Join(homeDir, "Projects", "Flutter"),
			filepath.Join(homeDir, "Projects", "Dart"),
		}

		for _, dir := range dartDirs {
			if _, err := os.Stat(dir); err == nil {
				// 递归查找pubspec.yaml文件
				err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
					if err != nil {
						return nil
					}

					if !info.IsDir() && info.Name() == "pubspec.yaml" {
						// 找到Dart项目，读取pubspec.yaml
						content, err := os.ReadFile(path)
						if err != nil {
							return nil
						}

						// 简单解析YAML内容，查找dependencies部分
						lines := strings.Split(string(content), "\n")
						inDependencies := false

						for _, line := range lines {
							line = strings.TrimSpace(line)

							if strings.HasPrefix(line, "dependencies:") {
								inDependencies = true
								continue
							}

							if inDependencies {
								if strings.HasPrefix(line, "  ") && !strings.HasPrefix(line, "    ") {
									parts := strings.SplitN(line, ":", 2)
									if len(parts) >= 2 {
										name := strings.TrimSpace(parts[0])
										version := strings.TrimSpace(parts[1])

										// 清理版本字符串
										version = strings.Trim(version, "^~>=< ")

										if name != "" {
											packages = append(packages, PackageInfo{
												Name:      name,
												Version:   version,
												Installed: true,
											})

											// 限制数量
											if len(packages) >= 20 {
												return filepath.SkipDir
											}
										}
									}
								} else if !strings.HasPrefix(line, " ") && line != "" {
									// 离开dependencies部分
									inDependencies = false
								}
							}
						}
					}

					return nil
				})
			}
		}

		return packages, nil
	}

	// 解析JSON输出
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(output), &result); err != nil {
		return packages, err
	}

	// 提取包信息
	if pkgs, ok := result["packages"].([]interface{}); ok {
		for _, pkg := range pkgs {
			if pkgMap, ok := pkg.(map[string]interface{}); ok {
				name, _ := pkgMap["name"].(string)
				version, _ := pkgMap["version"].(string)

				if name != "" {
					packages = append(packages, PackageInfo{
						Name:      name,
						Version:   version,
						Installed: true,
					})

					// 限制数量
					if len(packages) >= 20 {
						break
					}
				}
			}
		}
	}

	return packages, nil
}

// 检测TypeScript
func (a *App) detectTypeScript() LanguageInfo {
	info := LanguageInfo{
		Name:            "TypeScript",
		Installed:       false,
		DownloadURL:     "https://www.typescriptlang.org/download",
		InstallTutorial: "https://www.typescriptlang.org/docs/handbook/typescript-tooling-in-5-minutes.html",
		PackageManager:  "npm",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "typescript",
				Description: "TypeScript是JavaScript的超集，添加了类型系统",
				Version:     "5.1.6",
			},
			{
				Name:        "ts-node",
				Description: "在Node.js中执行TypeScript",
				Version:     "10.9.1",
			},
			{
				Name:        "eslint",
				Description: "可插拔的JavaScript和TypeScript代码检查工具",
				Version:     "8.43.0",
			},
			{
				Name:        "prettier",
				Description: "代码格式化工具",
				Version:     "3.0.0",
			},
			{
				Name:        "tslint",
				Description: "可扩展的TypeScript静态分析工具",
				Version:     "6.1.3",
			},
		},
	}

	if !commandExists("tsc") {
		return info
	}

	output, err := executeCommandWithTimeout("tsc", "--version")
	if err == nil {
		info.Installed = true
		info.Version = output

		// 检查npm是否安装
		if commandExists("npm") {
			// 列出TypeScript相关的全局包
			packages, _ := a.listTypeScriptPackages()
			info.Packages = packages
		}
	}

	return info
}

// 列出TypeScript相关包
func (a *App) listTypeScriptPackages() ([]PackageInfo, error) {
	// 使用npm list命令获取全局安装的包
	output, _ := executeCommandWithTimeout("npm", "list", "--global", "--json", "--depth=0")

	var result struct {
		Dependencies map[string]struct {
			Version string `json:"version"`
		} `json:"dependencies"`
	}

	if err := json.Unmarshal([]byte(output), &result); err != nil {
		return nil, err
	}

	// TypeScript相关的包名前缀
	tsRelatedPrefixes := []string{
		"typescript",
		"ts-",
		"@types/",
		"tslint",
		"eslint-plugin-typescript",
		"eslint-config-typescript",
	}

	packages := make([]PackageInfo, 0)
	for name, info := range result.Dependencies {
		// 检查是否为TypeScript相关包
		isTypeScriptRelated := false
		for _, prefix := range tsRelatedPrefixes {
			if strings.HasPrefix(name, prefix) {
				isTypeScriptRelated = true
				break
			}
		}

		if isTypeScriptRelated {
			packages = append(packages, PackageInfo{
				Name:      name,
				Version:   info.Version,
				Installed: true,
			})
		}
	}

	return packages, nil
}

// 检测Perl
func (a *App) detectPerl() LanguageInfo {
	info := LanguageInfo{
		Name:            "Perl",
		Installed:       false,
		DownloadURL:     "https://www.perl.org/get.html",
		InstallTutorial: "https://www.perl.org/get.html#unix_like",
		PackageManager:  "cpan",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "Moose",
				Description: "Perl 5的后现代对象系统",
				Version:     "2.2015",
			},
			{
				Name:        "DBI",
				Description: "Perl数据库接口模块",
				Version:     "1.643",
			},
			{
				Name:        "Mojolicious",
				Description: "Perl的实时Web框架",
				Version:     "9.31",
			},
			{
				Name:        "Dancer2",
				Description: "轻量级Perl Web应用框架",
				Version:     "0.400000",
			},
			{
				Name:        "Catalyst",
				Description: "Perl的MVC Web框架",
				Version:     "5.90131",
			},
		},
	}

	if !commandExists("perl") {
		return info
	}

	output, err := executeCommandWithTimeout("perl", "--version")
	if err == nil {
		info.Installed = true
		info.Version = output
	}

	return info
}

// 检测Lua
func (a *App) detectLua() LanguageInfo {
	info := LanguageInfo{
		Name:            "Lua",
		Installed:       false,
		DownloadURL:     "https://www.lua.org/download.html",
		InstallTutorial: "https://www.lua.org/manual/5.4/readme.html",
		PackageManager:  "LuaRocks",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "luasocket",
				Description: "Lua的网络支持",
				Version:     "3.1.0",
			},
			{
				Name:        "luafilesystem",
				Description: "Lua的文件系统操作",
				Version:     "1.8.0",
			},
			{
				Name:        "penlight",
				Description: "Lua的实用程序库",
				Version:     "1.13.1",
			},
			{
				Name:        "luasql",
				Description: "Lua的SQL数据库接口",
				Version:     "2.6.0",
			},
			{
				Name:        "lua-cjson",
				Description: "Lua的快速JSON解析器/编码器",
				Version:     "2.1.0",
			},
		},
	}

	if !commandExists("lua") {
		return info
	}

	output, err := executeCommandWithTimeout("lua", "-v")
	if err == nil {
		info.Installed = true
		info.Version = output

		// 检测已安装的Lua包
		packages, _ := a.listLuaPackages()
		info.Packages = packages
	}

	return info
}

// 列出已安装的Lua包
func (a *App) listLuaPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 检查LuaRocks是否安装
	if !commandExists("luarocks") {
		return packages, nil
	}

	// 使用LuaRocks列出已安装的包
	output, err := executeCommandWithTimeout("luarocks", "list", "--porcelain")
	if err != nil {
		return packages, err
	}

	lines := strings.Split(output, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// LuaRocks porcelain输出格式是：name version path arch
		parts := strings.Fields(line)
		if len(parts) >= 2 {
			packages = append(packages, PackageInfo{
				Name:      parts[0],
				Version:   parts[1],
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

// 检测R
func (a *App) detectR() LanguageInfo {
	info := LanguageInfo{
		Name:            "R",
		Installed:       false,
		DownloadURL:     "https://cran.r-project.org/",
		InstallTutorial: "https://cran.r-project.org/doc/manuals/r-release/R-admin.html",
		PackageManager:  "install.packages",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "ggplot2",
				Description: "数据可视化包",
				Version:     "3.4.2",
			},
			{
				Name:        "dplyr",
				Description: "数据操作包",
				Version:     "1.1.2",
			},
			{
				Name:        "tidyr",
				Description: "用于整理数据的工具",
				Version:     "1.3.0",
			},
			{
				Name:        "shiny",
				Description: "交互式Web应用框架",
				Version:     "1.7.4",
			},
			{
				Name:        "caret",
				Description: "分类和回归训练",
				Version:     "6.0-94",
			},
		},
	}

	if !commandExists("R") {
		return info
	}

	output, err := executeCommandWithTimeout("R", "--version")
	if err == nil {
		info.Installed = true
		info.Version = output

		// 检测已安装的R包
		packages, _ := a.listRPackages()
		info.Packages = packages
	}

	return info
}

// 列出已安装的R包
func (a *App) listRPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 使用R命令获取已安装的包列表
	// 创建一个R脚本来输出已安装的包
	rScript := `
	pkgs <- installed.packages()[, c("Package", "Version")]
	cat(paste(pkgs[,"Package"], pkgs[,"Version"], sep="|"), sep="\n")
	`

	// 创建临时文件
	tempFile, err := os.CreateTemp("", "r-pkg-list-*.R")
	if err != nil {
		return packages, err
	}
	defer os.Remove(tempFile.Name())

	// 写入R脚本
	if _, err := tempFile.WriteString(rScript); err != nil {
		return packages, err
	}
	tempFile.Close()

	// 执行R脚本
	output, err := executeCommandWithTimeout("R", "--vanilla", "-f", tempFile.Name(), "--quiet")
	if err != nil {
		// 尝试另一种方法
		output, err = executeCommandWithTimeout("R", "--vanilla", "-e", "installed.packages()[, c(\"Package\", \"Version\")]")
		if err != nil {
			return packages, err
		}
	}

	// 解析输出
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "[") || strings.HasPrefix(line, "Package") {
			continue
		}

		var name, version string

		// 尝试分隔符"|"
		if strings.Contains(line, "|") {
			parts := strings.Split(line, "|")
			if len(parts) >= 2 {
				name = strings.TrimSpace(parts[0])
				version = strings.TrimSpace(parts[1])
			}
		} else {
			// 尝试提取引号中的内容
			re := regexp.MustCompile(`"([^"]+)"`)
			matches := re.FindAllStringSubmatch(line, -1)
			if len(matches) >= 2 {
				name = matches[0][1]
				version = matches[1][1]
			}
		}

		if name != "" {
			packages = append(packages, PackageInfo{
				Name:      name,
				Version:   version,
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

// 检测MATLAB
func (a *App) detectMatlab() LanguageInfo {
	info := LanguageInfo{
		Name:            "MATLAB",
		Installed:       false,
		DownloadURL:     "https://www.mathworks.com/products/matlab.html",
		InstallTutorial: "https://www.mathworks.com/help/install/",
		PackageManager:  "Add-On Explorer",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "Signal Processing Toolbox",
				Description: "提供用于分析、预处理和提取信号特征的行业标准算法",
				Version:     "R2023a",
			},
			{
				Name:        "Statistics and Machine Learning Toolbox",
				Description: "提供用于分析和建模数据的函数和应用程序",
				Version:     "R2023a",
			},
			{
				Name:        "Image Processing Toolbox",
				Description: "提供用于图像处理、分析和算法开发的一组参考标准算法和工作流程应用程序",
				Version:     "R2023a",
			},
			{
				Name:        "Optimization Toolbox",
				Description: "提供用于寻找参数以最小化或最大化目标的函数",
				Version:     "R2023a",
			},
			{
				Name:        "Deep Learning Toolbox",
				Description: "提供用于设计和实现深度神经网络的框架",
				Version:     "R2023a",
			},
		},
	}

	// 检测MATLAB (Windows)
	if commandExists("matlab") {
		info.Installed = true
		info.Version = "MATLAB (版本无法自动检测)"
		return info
	}

	// 在Windows上尝试使用where命令
	output, err := executeCommandWithTimeout("where", "matlab")
	if err == nil && output != "" {
		info.Installed = true
		info.Version = "MATLAB (版本无法自动检测)"
	}

	return info
}

// 检测Scala
func (a *App) detectScala() LanguageInfo {
	info := LanguageInfo{
		Name:            "Scala",
		Installed:       false,
		DownloadURL:     "https://www.scala-lang.org/download/",
		InstallTutorial: "https://docs.scala-lang.org/getting-started/index.html",
		PackageManager:  "sbt",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "akka",
				Description: "构建高度并发、分布式和弹性消息驱动应用程序的工具包",
				Version:     "2.8.0",
			},
			{
				Name:        "play",
				Description: "高生产力的Java和Scala Web应用程序框架",
				Version:     "2.8.19",
			},
			{
				Name:        "cats",
				Description: "Scala的函数式编程库",
				Version:     "2.9.0",
			},
			{
				Name:        "spark",
				Description: "用于大规模数据处理的统一分析引擎",
				Version:     "3.4.1",
			},
			{
				Name:        "zio",
				Description: "用于异步和并发编程的类型安全、可组合库",
				Version:     "2.0.15",
			},
		},
	}

	if !commandExists("scala") {
		return info
	}

	output, err := executeCommandWithTimeout("scala", "-version")
	if err == nil {
		info.Installed = true
		info.Version = output

		// 检测已安装的Scala包
		packages, _ := a.listScalaPackages()
		info.Packages = packages
	}

	return info
}

// 列出已安装的Scala包
func (a *App) listScalaPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 检查sbt是否安装
	if !commandExists("sbt") {
		return packages, nil
	}

	// 创建临时目录用于sbt项目
	tempDir, err := os.MkdirTemp("", "scala-pkg-check")
	if err != nil {
		return packages, err
	}
	defer os.RemoveAll(tempDir)

	// 创建简单的build.sbt文件
	buildSbt := `
scalaVersion := "2.13.10"
libraryDependencies ++= Seq()
`
	err = os.WriteFile(filepath.Join(tempDir, "build.sbt"), []byte(buildSbt), 0644)
	if err != nil {
		return packages, err
	}

	// 创建项目目录结构
	os.MkdirAll(filepath.Join(tempDir, "project"), 0755)

	// 运行sbt命令获取依赖信息
	cmd := exec.Command("sbt", "dependencyList")
	cmd.Dir = tempDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		// 尝试另一种方法：检查Ivy缓存
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return packages, err
		}

		// Ivy缓存路径
		ivyCachePath := filepath.Join(homeDir, ".ivy2", "cache")
		if _, err := os.Stat(ivyCachePath); err != nil {
			return packages, nil
		}

		// 遍历Ivy缓存目录
		count := 0
		err = filepath.Walk(ivyCachePath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}

			if info.IsDir() {
				// 检查是否是一个Scala库目录（通常有org.scala-lang等前缀）
				relPath, err := filepath.Rel(ivyCachePath, path)
				if err != nil {
					return nil
				}

				parts := strings.Split(relPath, string(filepath.Separator))
				if len(parts) == 2 {
					// 检查版本子目录
					versionDirs, err := os.ReadDir(path)
					if err != nil {
						return nil
					}

					if len(versionDirs) > 0 {
						// 使用最后一个目录作为最新版本
						latestVersion := versionDirs[len(versionDirs)-1].Name()

						packages = append(packages, PackageInfo{
							Name:      parts[0] + ":" + parts[1],
							Version:   latestVersion,
							Installed: true,
						})

						count++
						if count >= 20 {
							return filepath.SkipDir
						}
					}
				}
			}

			return nil
		})

		return packages, nil
	}

	// 解析sbt输出
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.Contains(line, ":") && !strings.Contains(line, "Success") {
			// 尝试提取依赖信息，格式通常是 "org:name:version"
			parts := strings.Split(line, ":")
			if len(parts) >= 3 {
				name := parts[0] + ":" + parts[1]
				version := parts[2]

				// 移除可能的额外信息
				if idx := strings.Index(version, " "); idx > 0 {
					version = version[:idx]
				}

				packages = append(packages, PackageInfo{
					Name:      name,
					Version:   version,
					Installed: true,
				})

				// 限制数量
				if len(packages) >= 20 {
					break
				}
			}
		}
	}

	return packages, nil
}

// 检测Haskell
func (a *App) detectHaskell() LanguageInfo {
	info := LanguageInfo{
		Name:            "Haskell",
		Installed:       false,
		DownloadURL:     "https://www.haskell.org/downloads/",
		InstallTutorial: "https://www.haskell.org/ghcup/",
		PackageManager:  "cabal/stack",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "aeson",
				Description: "快速的JSON解析和生成",
				Version:     "2.1.2.1",
			},
			{
				Name:        "servant",
				Description: "声明式Web API的类型级框架",
				Version:     "0.19.1",
			},
			{
				Name:        "lens",
				Description: "Haskell的函数式引用",
				Version:     "5.2.2",
			},
			{
				Name:        "pandoc",
				Description: "通用文档转换器",
				Version:     "3.1.3",
			},
			{
				Name:        "yesod",
				Description: "Haskell的Web框架",
				Version:     "1.6.2.1",
			},
		},
	}

	if !commandExists("ghc") {
		return info
	}

	output, err := executeCommandWithTimeout("ghc", "--version")
	if err == nil {
		info.Installed = true
		info.Version = output

		// 检测已安装的Haskell包
		packages, _ := a.listHaskellPackages()
		info.Packages = packages
	}

	return info
}

// 列出已安装的Haskell包
func (a *App) listHaskellPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 尝试使用cabal命令
	if commandExists("cabal") {
		output, err := executeCommandWithTimeout("cabal", "list", "--installed", "--simple-output")
		if err == nil {
			lines := strings.Split(output, "\n")
			for _, line := range lines {
				line = strings.TrimSpace(line)
				if line == "" {
					continue
				}

				// cabal输出格式通常是 "package-name version"
				parts := strings.Split(line, " ")
				if len(parts) >= 2 {
					packages = append(packages, PackageInfo{
						Name:      parts[0],
						Version:   parts[1],
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

	// 如果cabal没有结果或不存在，尝试使用stack
	if len(packages) == 0 && commandExists("stack") {
		output, err := executeCommandWithTimeout("stack", "list-dependencies")
		if err == nil {
			lines := strings.Split(output, "\n")
			for _, line := range lines {
				line = strings.TrimSpace(line)
				if line == "" || strings.HasPrefix(line, "base ") {
					continue
				}

				// stack输出格式通常是 "package-name version"
				parts := strings.Split(line, " ")
				if len(parts) >= 2 {
					packages = append(packages, PackageInfo{
						Name:      parts[0],
						Version:   parts[1],
						Installed: true,
					})

					// 限制数量
					if len(packages) >= 20 {
						break
					}
				}
			}
		}
	}

	return packages, nil
}
