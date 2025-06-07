package main

import (
	"encoding/json"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

// 为Elm添加包检测功能
func (a *App) detectElmWithPackages() LanguageInfo {
	info := a.detectElm()

	if info.Installed {
		// 获取已安装的Elm包
		packages, _ := a.listElmPackages()
		info.Packages = packages
	}

	return info
}

// 列出已安装的Elm包
func (a *App) listElmPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 检查elm命令是否可用
	if !commandExists("elm") {
		return packages, nil
	}

	// 首先尝试从elm.json文件解析依赖
	packages = a.findElmPackagesFromProjects()

	// 如果没有找到项目文件，尝试通过elm命令获取
	if len(packages) == 0 {
		// 尝试使用elm命令获取包列表
		output, err := executeCommandWithTimeout("elm", "install", "--help")
		if err == nil {
			// 解析输出寻找可能提到的包
			packages = a.parseElmPackagesFromHelp(output)
		}
	}

	// 尝试使用npm查找elm相关包（因为很多Elm工具是通过npm安装的）
	if len(packages) == 0 && commandExists("npm") {
		output, err := executeCommandWithTimeout("npm", "list", "-g", "--json")
		if err == nil {
			npmPackages := a.parseNpmGlobalPackages(output)
			for _, pkg := range npmPackages {
				if strings.Contains(pkg.Name, "elm") {
					pkg.Description = getElmPackageDescription(pkg.Name)
					packages = append(packages, pkg)
				}
			}
		}
	}

	// 如果仍然没有找到包，添加常见Elm包作为建议
	if len(packages) == 0 {
		packages = a.getRecommendedElmPackages()
	}

	return packages, nil
}

// 从项目中查找Elm包
func (a *App) findElmPackagesFromProjects() []PackageInfo {
	var packages []PackageInfo

	// 常见的elm.json位置
	elmJsonPaths, _ := findFiles("elm.json")

	for _, elmJsonPath := range elmJsonPaths {
		content, err := readFile(elmJsonPath)
		if err != nil {
			continue
		}

		// 解析JSON
		var elmJson map[string]interface{}
		if err := json.Unmarshal([]byte(content), &elmJson); err != nil {
			continue
		}

		// 检查依赖
		dependencies, ok := elmJson["dependencies"].(map[string]interface{})
		if !ok {
			// 尝试应用类型结构
			if deps, ok := elmJson["dependencies"].(map[string]interface{}); ok {
				if direct, ok := deps["direct"].(map[string]interface{}); ok {
					dependencies = direct
				}
			}
		}

		if dependencies != nil {
			for name, version := range dependencies {
				versionStr := ""
				switch v := version.(type) {
				case string:
					versionStr = v
				case float64:
					versionStr = strconv.FormatFloat(v, 'f', -1, 64)
				}

				description := getElmPackageDescription(name)
				packages = append(packages, PackageInfo{
					Name:        name,
					Version:     versionStr,
					Description: description,
					Installed:   true,
				})
			}
		}
	}

	return packages
}

// 从帮助信息中解析可能的Elm包
func (a *App) parseElmPackagesFromHelp(output string) []PackageInfo {
	var packages []PackageInfo

	// 查找形如 author/package 的模式
	re := regexp.MustCompile(`([a-zA-Z0-9-]+)/([a-zA-Z0-9-]+)`)
	matches := re.FindAllStringSubmatch(output, -1)

	seen := make(map[string]bool)
	for _, match := range matches {
		if len(match) >= 3 {
			fullName := match[1] + "/" + match[2]
			if seen[fullName] {
				continue
			}
			seen[fullName] = true

			description := getElmPackageDescription(fullName)
			packages = append(packages, PackageInfo{
				Name:        fullName,
				Version:     "latest",
				Description: description,
				Installed:   false,
			})
		}
	}

	return packages
}

// 解析全局npm包列表
func (a *App) parseNpmGlobalPackages(output string) []PackageInfo {
	var packages []PackageInfo

	var result map[string]interface{}
	if err := json.Unmarshal([]byte(output), &result); err != nil {
		return packages
	}

	if deps, ok := result["dependencies"].(map[string]interface{}); ok {
		for name, info := range deps {
			version := ""
			if infoMap, ok := info.(map[string]interface{}); ok {
				if v, ok := infoMap["version"].(string); ok {
					version = v
				}
			}

			packages = append(packages, PackageInfo{
				Name:      name,
				Version:   version,
				Installed: true,
			})
		}
	}

	return packages
}

// 获取推荐的Elm包
func (a *App) getRecommendedElmPackages() []PackageInfo {
	return []PackageInfo{
		{
			Name:        "elm/core",
			Version:     "1.0.5",
			Description: "Elm语言的核心库",
			Installed:   false,
		},
		{
			Name:        "elm/browser",
			Version:     "1.0.2",
			Description: "与浏览器交互",
			Installed:   false,
		},
		{
			Name:        "elm/html",
			Version:     "1.0.0",
			Description: "HTML元素的Elm表示",
			Installed:   false,
		},
		{
			Name:        "elm/json",
			Version:     "1.1.3",
			Description: "处理JSON数据",
			Installed:   false,
		},
		{
			Name:        "elm/http",
			Version:     "2.0.0",
			Description: "HTTP请求处理",
			Installed:   false,
		},
		{
			Name:        "elm/url",
			Version:     "1.0.0",
			Description: "处理URL",
			Installed:   false,
		},
		{
			Name:        "elm/time",
			Version:     "1.0.0",
			Description: "时间处理库",
			Installed:   false,
		},
		{
			Name:        "elm/svg",
			Version:     "1.0.1",
			Description: "SVG元素的Elm表示",
			Installed:   false,
		},
		{
			Name:        "elm-explorations/test",
			Version:     "1.2.2",
			Description: "单元测试框架",
			Installed:   false,
		},
		{
			Name:        "mdgriffith/elm-ui",
			Version:     "1.1.8",
			Description: "Elm的布局库",
			Installed:   false,
		},
	}
}

// 获取Elm包的描述
func getElmPackageDescription(name string) string {
	descriptions := map[string]string{
		"elm/core":                   "Elm语言的核心库",
		"elm/browser":                "与浏览器交互",
		"elm/html":                   "HTML元素的Elm表示",
		"elm/json":                   "处理JSON数据",
		"elm/http":                   "HTTP请求处理",
		"elm/url":                    "处理URL",
		"elm/time":                   "时间处理库",
		"elm/svg":                    "SVG元素的Elm表示",
		"elm-explorations/test":      "单元测试框架",
		"mdgriffith/elm-ui":          "Elm的布局库",
		"elm/virtual-dom":            "虚拟DOM实现",
		"elm/random":                 "随机数生成",
		"elm/parser":                 "解析器组合器库",
		"elm/file":                   "文件操作",
		"elm/bytes":                  "二进制数据处理",
		"elm/regex":                  "正则表达式",
		"elm-community/list-extra":   "List操作的额外函数",
		"elm-community/maybe-extra":  "Maybe类型的额外函数",
		"elm-community/string-extra": "String操作的额外函数",
		"elm-community/json-extra":   "JSON处理的额外帮助函数",
		"rtfeldman/elm-css":          "CSS的类型安全表示",
		"krisajenkins/remotedata":    "处理远程数据加载状态",
		"elm-community/result-extra": "Result类型的额外函数",
		"elm-community/dict-extra":   "Dict操作的额外函数",
		"elm-community/array-extra":  "Array操作的额外函数",
	}

	if desc, ok := descriptions[name]; ok {
		return desc
	}

	// 尝试分析包名
	parts := strings.Split(name, "/")
	if len(parts) == 2 {
		packageName := parts[1]

		if strings.Contains(packageName, "json") {
			return "JSON处理相关库"
		} else if strings.Contains(packageName, "http") || strings.Contains(packageName, "request") {
			return "HTTP/网络请求相关库"
		} else if strings.Contains(packageName, "ui") || strings.Contains(packageName, "html") || strings.Contains(packageName, "css") {
			return "UI/界面相关库"
		} else if strings.Contains(packageName, "test") {
			return "测试相关库"
		} else if strings.Contains(packageName, "parser") {
			return "解析相关库"
		} else if strings.Contains(packageName, "router") || strings.Contains(packageName, "navigation") {
			return "路由/导航相关库"
		}

		return parts[0] + "的" + packageName + "库"
	}

	return "Elm包"
}

// 辅助函数：查找文件
func findFiles(pattern string) ([]string, error) {
	// 执行系统查找命令
	var output string
	var err error

	switch runtime.GOOS {
	case "windows":
		output, err = executeCommandWithTimeout("powershell", "-Command", "Get-ChildItem -Path . -Filter "+pattern+" -Recurse | Select-Object -ExpandProperty FullName")
	default:
		output, err = executeCommandWithTimeout("find", ".", "-name", pattern, "-type", "f")
	}

	if err != nil {
		return nil, err
	}

	// 解析结果
	files := strings.Split(output, "\n")
	var result []string
	for _, file := range files {
		if file = strings.TrimSpace(file); file != "" {
			result = append(result, file)
		}
	}

	return result, nil
}

// 辅助函数：读取文件内容
func readFile(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
