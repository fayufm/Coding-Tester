package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// 为Swift添加包检测功能
func (a *App) detectSwiftWithPackages() LanguageInfo {
	info := a.detectSwift()

	if info.Installed {
		// 获取已安装的Swift包
		packages, _ := a.listSwiftPackages()
		info.Packages = packages
	}

	return info
}

// 列出已安装的Swift包（增强版）
func (a *App) listSwiftPackagesEnhanced() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 检查swift package命令是否可用
	if !commandExists("swift") {
		return packages, nil
	}

	// 方法1: 检查全局Swift包
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
		return a.findExistingSwiftPackages()
	}

	// 添加Package.swift文件内容，包含一些常用依赖
	packageSwift := `
// swift-tools-version:5.5
import PackageDescription

let package = Package(
    name: "SwiftPkgChecker",
    dependencies: [
        .package(url: "https://github.com/Alamofire/Alamofire.git", from: "5.0.0"),
        .package(url: "https://github.com/SwiftyJSON/SwiftyJSON.git", from: "5.0.0"),
        .package(url: "https://github.com/onevcat/Kingfisher.git", from: "7.0.0"),
        .package(url: "https://github.com/ReactiveX/RxSwift.git", from: "6.0.0"),
        .package(url: "https://github.com/SnapKit/SnapKit.git", from: "5.0.0")
    ],
    targets: [
        .target(
            name: "SwiftPkgChecker",
            dependencies: ["Alamofire", "SwiftyJSON", "Kingfisher", "RxSwift", "SnapKit"]),
    ]
)
`

	err = os.WriteFile(filepath.Join(tempDir, "Package.swift"), []byte(packageSwift), 0644)
	if err != nil {
		return packages, err
	}

	// 尝试解析依赖
	resolveCmd := exec.Command("swift", "package", "resolve")
	resolveCmd.Dir = tempDir
	_ = resolveCmd.Run() // 忽略错误，继续尝试获取信息

	// 获取依赖信息
	output, err := executeCommandWithTimeout("swift", "package", "--package-path", tempDir, "show-dependencies", "--format", "json")
	if err != nil {
		// 如果JSON格式失败，尝试文本格式
		output, err = executeCommandWithTimeout("swift", "package", "--package-path", tempDir, "show-dependencies")
		if err != nil {
			return a.findExistingSwiftPackages() // 备选方案
		}

		// 解析文本输出
		return a.parseSwiftDependenciesText(output)
	}

	// 解析JSON输出
	return a.parseSwiftDependenciesJSON(output)
}

// 查找现有的Swift项目并提取依赖
func (a *App) findExistingSwiftPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return packages, err
	}

	// 检查常见的Swift项目目录
	swiftDirs := []string{
		filepath.Join(homeDir, "Documents", "SwiftProjects"),
		filepath.Join(homeDir, "Projects", "Swift"),
		filepath.Join(homeDir, "Developer", "Swift"),
		filepath.Join(homeDir, "Library", "Developer", "Xcode", "DerivedData"),
	}

	for _, dir := range swiftDirs {
		if _, err := os.Stat(dir); err != nil {
			continue
		}

		// 查找Package.swift文件
		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}

			if !info.IsDir() && filepath.Base(path) == "Package.swift" {
				// 找到Swift包文件，读取内容
				content, err := os.ReadFile(path)
				if err != nil {
					return nil
				}

				// 解析Package.swift文件中的依赖
				pkgs := parseSwiftPackageFile(string(content))
				packages = append(packages, pkgs...)

				// 限制处理的文件数量
				if len(packages) > 20 {
					return filepath.SkipDir
				}
			}
			return nil
		})

		if err == nil && len(packages) > 0 {
			break
		}
	}

	// 如果没有找到包，添加一些常见的Swift包作为建议
	if len(packages) == 0 {
		packages = append(packages, PackageInfo{
			Name:        "Alamofire",
			Version:     "5.7.1",
			Description: "Swift的优雅网络库",
			Installed:   false,
		})
		packages = append(packages, PackageInfo{
			Name:        "SwiftyJSON",
			Version:     "5.0.1",
			Description: "处理JSON数据的更好方式",
			Installed:   false,
		})
		packages = append(packages, PackageInfo{
			Name:        "Kingfisher",
			Version:     "7.8.1",
			Description: "强大的纯Swift库，用于从网络下载和缓存图像",
			Installed:   false,
		})
	}

	return packages, nil
}

// 解析Swift依赖的JSON输出
func (a *App) parseSwiftDependenciesJSON(output string) ([]PackageInfo, error) {
	var packages []PackageInfo
	var result map[string]interface{}

	err := json.Unmarshal([]byte(output), &result)
	if err != nil {
		return packages, err
	}

	// 解析JSON结构
	if dependencies, ok := result["dependencies"].([]interface{}); ok {
		for _, dep := range dependencies {
			if depMap, ok := dep.(map[string]interface{}); ok {
				name, _ := depMap["name"].(string)
				version, _ := depMap["version"].(string)
				if name != "" {
					description := getSwiftPackageDescription(name)
					packages = append(packages, PackageInfo{
						Name:        name,
						Version:     version,
						Description: description,
						Installed:   true,
					})
				}
			}
		}
	}

	return packages, nil
}

// 解析Swift依赖的文本输出
func (a *App) parseSwiftDependenciesText(output string) ([]PackageInfo, error) {
	var packages []PackageInfo
	lines := strings.Split(output, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// 尝试匹配格式: PackageName X.Y.Z
		parts := strings.Fields(line)
		if len(parts) >= 2 {
			name := parts[0]
			version := parts[len(parts)-1]

			// 过滤版本号（通常是数字和点）
			if strings.ContainsAny(version, "0123456789.") {
				description := getSwiftPackageDescription(name)
				packages = append(packages, PackageInfo{
					Name:        name,
					Version:     version,
					Description: description,
					Installed:   true,
				})
			}
		}
	}

	return packages, nil
}

// 解析Package.swift文件内容
func parseSwiftPackageFile(content string) []PackageInfo {
	var packages []PackageInfo

	// 查找.package行
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.Contains(line, ".package(") {
			// 提取包名
			urlStart := strings.Index(line, "\"")
			if urlStart == -1 {
				continue
			}
			urlEnd := strings.Index(line[urlStart+1:], "\"")
			if urlEnd == -1 {
				continue
			}
			url := line[urlStart+1 : urlStart+1+urlEnd]

			// 从URL提取名称
			parts := strings.Split(url, "/")
			if len(parts) < 2 {
				continue
			}
			name := parts[len(parts)-1]
			name = strings.TrimSuffix(name, ".git")

			// 尝试提取版本
			version := "latest"
			if strings.Contains(line, "from:") {
				versionStart := strings.Index(line, "from:")
				if versionStart != -1 {
					versionQuoteStart := strings.Index(line[versionStart:], "\"")
					if versionQuoteStart != -1 {
						versionStart += versionQuoteStart + 1
						versionQuoteEnd := strings.Index(line[versionStart:], "\"")
						if versionQuoteEnd != -1 {
							version = line[versionStart : versionStart+versionQuoteEnd]
						}
					}
				}
			}

			description := getSwiftPackageDescription(name)
			packages = append(packages, PackageInfo{
				Name:        name,
				Version:     version,
				Description: description,
				Installed:   true,
			})
		}
	}

	return packages
}

// 获取Swift包的描述
func getSwiftPackageDescription(name string) string {
	descriptions := map[string]string{
		"Alamofire":  "Swift的优雅网络库",
		"SwiftyJSON": "处理JSON数据的更好方式",
		"Kingfisher": "强大的纯Swift库，用于从网络下载和缓存图像",
		"SnapKit":    "Swift的自动布局DSL",
		"RxSwift":    "Swift的响应式编程",
		"Moya":       "基于Alamofire的网络抽象层",
		"SwiftUI":    "Apple的声明式UI框架",
		"Combine":    "Apple的响应式框架",
		"Charts":     "强大的图表库",
		"Realm":      "移动数据库",
		"Firebase":   "Google的移动平台",
		"SwiftLint":  "Swift代码规范检查工具",
		"Swinject":   "Swift的依赖注入框架",
		"Quick":      "Swift的行为驱动开发框架",
		"Nimble":     "Swift的匹配器框架",
	}

	if desc, ok := descriptions[name]; ok {
		return desc
	}

	// 针对不在预定义列表中的包，从名称推断描述
	if strings.Contains(strings.ToLower(name), "network") || strings.Contains(strings.ToLower(name), "http") {
		return "网络相关库"
	} else if strings.Contains(strings.ToLower(name), "ui") || strings.Contains(strings.ToLower(name), "view") {
		return "UI相关库"
	} else if strings.Contains(strings.ToLower(name), "data") || strings.Contains(strings.ToLower(name), "db") {
		return "数据库相关库"
	}

	return "Swift包"
}
