package main

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

// 为PureScript添加包检测功能
func (a *App) detectPureScriptWithPackages() LanguageInfo {
	info := a.detectPureScript()

	if info.Installed {
		// 获取已安装的PureScript包
		packages, _ := a.listPureScriptPackages()
		info.Packages = packages
	}

	return info
}

// 列出已安装的PureScript包
func (a *App) listPureScriptPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 检查PureScript命令是否可用
	if !commandExists("purs") && !commandExists("spago") {
		return packages, nil
	}

	// 首先尝试从spago.dhall文件解析依赖（Spago是PureScript的包管理器）
	packages = a.findPureScriptPackagesFromProjects()

	// 尝试通过spago命令获取包列表
	if len(packages) == 0 && commandExists("spago") {
		output, err := executeCommandWithTimeout("spago", "ls", "packages")
		if err == nil {
			pkgs := a.parseSpagoPackagesList(output)
			if len(pkgs) > 0 {
				packages = append(packages, pkgs...)
			}
		}
	}

	// 尝试通过bower命令获取PureScript包（旧版PureScript使用bower）
	if len(packages) == 0 && commandExists("bower") {
		output, err := executeCommandWithTimeout("bower", "list", "--json")
		if err == nil {
			pkgs := a.parseBowerPackagesList(output)
			if len(pkgs) > 0 {
				packages = append(packages, pkgs...)
			}
		}
	}

	// 尝试通过npm查找purs相关包（因为许多PureScript工具通过npm安装）
	if len(packages) == 0 && commandExists("npm") {
		output, err := executeCommandWithTimeout("npm", "list", "-g", "--json")
		if err == nil {
			var result map[string]interface{}
			if json.Unmarshal([]byte(output), &result) == nil {
				if deps, ok := result["dependencies"].(map[string]interface{}); ok {
					for name, info := range deps {
						if strings.Contains(strings.ToLower(name), "purescript") || strings.Contains(strings.ToLower(name), "purs") {
							version := ""
							if infoMap, ok := info.(map[string]interface{}); ok {
								if v, ok := infoMap["version"].(string); ok {
									version = v
								}
							}

							description := getPureScriptPackageDescription(name)
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
		}
	}

	// 如果没有找到包，添加一些常见的PureScript包作为建议
	if len(packages) == 0 {
		packages = a.getRecommendedPureScriptPackages()
	}

	return packages, nil
}

// 从项目中查找PureScript包
func (a *App) findPureScriptPackagesFromProjects() []PackageInfo {
	var packages []PackageInfo

	// 查找spago.dhall配置文件
	dhallPaths, _ := findFilePatterns([]string{"spago.dhall", "*.dhall"})

	// 如果找到dhall文件，尝试解析内容
	if len(dhallPaths) > 0 {
		// Dhall文件是文本格式，我们需要查找dependencies部分
		for _, path := range dhallPaths {
			content, err := os.ReadFile(path)
			if err != nil {
				continue
			}

			contentStr := string(content)
			// 非常简单的解析，查找依赖部分
			if idx := strings.Index(contentStr, "dependencies"); idx != -1 {
				// 提取依赖部分的内容
				depStart := idx + strings.Index(contentStr[idx:], "[")
				if depStart != -1 {
					depEnd := depStart + strings.Index(contentStr[depStart:], "]")
					if depEnd != -1 {
						depContent := contentStr[depStart+1 : depEnd]
						// 分割每个依赖项
						deps := strings.Split(depContent, ",")
						for _, dep := range deps {
							dep = strings.TrimSpace(dep)
							// 移除引号
							dep = strings.Trim(dep, "\"'")
							if dep != "" {
								description := getPureScriptPackageDescription(dep)
								packages = append(packages, PackageInfo{
									Name:        dep,
									Version:     "latest",
									Description: description,
									Installed:   true,
								})
							}
						}
					}
				}
			}
		}
	}

	// 查找bower.json文件（旧版PureScript包管理）
	bowerPaths, _ := findFilePatterns([]string{"bower.json"})

	for _, path := range bowerPaths {
		content, err := os.ReadFile(path)
		if err != nil {
			continue
		}

		var bowerJson map[string]interface{}
		if err := json.Unmarshal(content, &bowerJson); err != nil {
			continue
		}

		// 检查dependencies和devDependencies
		deps := make(map[string]interface{})

		if dependencies, ok := bowerJson["dependencies"].(map[string]interface{}); ok {
			for k, v := range dependencies {
				deps[k] = v
			}
		}

		if devDependencies, ok := bowerJson["devDependencies"].(map[string]interface{}); ok {
			for k, v := range devDependencies {
				deps[k] = v
			}
		}

		// 添加所有找到的依赖
		for name, version := range deps {
			versionStr := ""
			switch v := version.(type) {
			case string:
				versionStr = v
			case float64:
				versionStr = formatFloatVersion(v)
			}

			description := getPureScriptPackageDescription(name)
			packages = append(packages, PackageInfo{
				Name:        name,
				Version:     versionStr,
				Description: description,
				Installed:   true,
			})
		}
	}

	// 查找package.json文件中的PureScript依赖
	npmPaths, _ := findFilePatterns([]string{"package.json"})

	for _, path := range npmPaths {
		content, err := os.ReadFile(path)
		if err != nil {
			continue
		}

		var packageJson map[string]interface{}
		if err := json.Unmarshal(content, &packageJson); err != nil {
			continue
		}

		// 检查dependencies和devDependencies
		deps := make(map[string]interface{})

		if dependencies, ok := packageJson["dependencies"].(map[string]interface{}); ok {
			for k, v := range dependencies {
				if strings.Contains(strings.ToLower(k), "purescript") || strings.Contains(strings.ToLower(k), "purs") {
					deps[k] = v
				}
			}
		}

		if devDependencies, ok := packageJson["devDependencies"].(map[string]interface{}); ok {
			for k, v := range devDependencies {
				if strings.Contains(strings.ToLower(k), "purescript") || strings.Contains(strings.ToLower(k), "purs") {
					deps[k] = v
				}
			}
		}

		// 添加所有找到的PureScript相关依赖
		for name, version := range deps {
			versionStr := ""
			switch v := version.(type) {
			case string:
				versionStr = v
			case float64:
				versionStr = formatFloatVersion(v)
			}

			description := getPureScriptPackageDescription(name)
			packages = append(packages, PackageInfo{
				Name:        name,
				Version:     versionStr,
				Description: description,
				Installed:   true,
			})
		}
	}

	return packages
}

// 解析spago包列表输出
func (a *App) parseSpagoPackagesList(output string) []PackageInfo {
	var packages []PackageInfo

	lines := strings.Split(output, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// spago通常输出形式为 "package-name"
		// 移除可能的前缀符号和引号
		name := strings.Trim(strings.TrimLeft(line, "+-* "), "\"'")

		// 有时会包含版本信息
		var version string
		if idx := strings.LastIndex(name, "@"); idx != -1 {
			version = name[idx+1:]
			name = name[:idx]
		} else {
			version = "latest"
		}

		description := getPureScriptPackageDescription(name)
		packages = append(packages, PackageInfo{
			Name:        name,
			Version:     version,
			Description: description,
			Installed:   true,
		})
	}

	return packages
}

// 解析bower包列表输出（JSON格式）
func (a *App) parseBowerPackagesList(output string) []PackageInfo {
	var packages []PackageInfo

	var result map[string]interface{}
	if err := json.Unmarshal([]byte(output), &result); err != nil {
		return packages
	}

	if dependencies, ok := result["dependencies"].(map[string]interface{}); ok {
		for name, info := range dependencies {
			// 只关注PureScript相关包
			if !strings.Contains(strings.ToLower(name), "purescript") && !strings.Contains(strings.ToLower(name), "purs") {
				continue
			}

			version := "latest"
			if infoMap, ok := info.(map[string]interface{}); ok {
				if pkgMeta, ok := infoMap["pkgMeta"].(map[string]interface{}); ok {
					if v, ok := pkgMeta["version"].(string); ok {
						version = v
					}
				}
			}

			description := getPureScriptPackageDescription(name)
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

// 获取推荐的PureScript包
func (a *App) getRecommendedPureScriptPackages() []PackageInfo {
	return []PackageInfo{
		{
			Name:        "purescript-prelude",
			Version:     "6.0.1",
			Description: "PureScript的基本函数和类型",
			Installed:   false,
		},
		{
			Name:        "purescript-effect",
			Version:     "4.0.0",
			Description: "副作用处理",
			Installed:   false,
		},
		{
			Name:        "purescript-console",
			Version:     "6.0.0",
			Description: "控制台输出",
			Installed:   false,
		},
		{
			Name:        "purescript-aff",
			Version:     "7.1.0",
			Description: "异步效果",
			Installed:   false,
		},
		{
			Name:        "purescript-halogen",
			Version:     "7.0.0",
			Description: "UI库",
			Installed:   false,
		},
		{
			Name:        "purescript-node-fs",
			Version:     "8.1.0",
			Description: "Node.js文件系统API",
			Installed:   false,
		},
		{
			Name:        "purescript-arrays",
			Version:     "7.1.0",
			Description: "数组操作",
			Installed:   false,
		},
		{
			Name:        "purescript-strings",
			Version:     "6.0.1",
			Description: "字符串操作",
			Installed:   false,
		},
		{
			Name:        "purescript-foreign",
			Version:     "7.0.0",
			Description: "外部数据处理",
			Installed:   false,
		},
		{
			Name:        "purescript-argonaut",
			Version:     "9.0.0",
			Description: "JSON处理",
			Installed:   false,
		},
	}
}

// 获取PureScript包的描述
func getPureScriptPackageDescription(name string) string {
	descriptions := map[string]string{
		"purescript-prelude":           "PureScript的基本函数和类型",
		"purescript-effect":            "副作用处理",
		"purescript-console":           "控制台输出",
		"purescript-aff":               "异步效果",
		"purescript-halogen":           "UI库",
		"purescript-node-fs":           "Node.js文件系统API",
		"purescript-arrays":            "数组操作",
		"purescript-strings":           "字符串操作",
		"purescript-foreign":           "外部数据处理",
		"purescript-argonaut":          "JSON处理",
		"purescript-web-dom":           "DOM API绑定",
		"purescript-web-html":          "HTML API绑定",
		"purescript-web-events":        "事件API绑定",
		"purescript-web-uievents":      "UI事件API绑定",
		"purescript-routing":           "客户端路由",
		"purescript-react":             "React绑定",
		"purescript-transformers":      "单子变换器",
		"purescript-parsing":           "解析器组合器",
		"purescript-profunctor-lenses": "函数式透镜",
		"purescript-integers":          "整数操作",
		"purescript-datetime":          "日期和时间处理",
		"purescript-node-buffer":       "Node.js缓冲区API",
		"purescript-node-http":         "Node.js HTTP API",
		"purescript-node-process":      "Node.js进程API",
		"purescript-react-basic":       "React基础绑定",
		"spago":                        "PureScript的包管理器和构建工具",
		"purs":                         "PureScript编译器",
		"purescript":                   "PureScript语言",
	}

	if desc, ok := descriptions[name]; ok {
		return desc
	}

	// 尝试从名称推断描述
	if strings.HasPrefix(name, "purescript-") {
		baseName := strings.TrimPrefix(name, "purescript-")

		if strings.Contains(baseName, "dom") || strings.Contains(baseName, "html") || strings.Contains(baseName, "css") {
			return "DOM/HTML相关库"
		} else if strings.Contains(baseName, "json") || strings.Contains(baseName, "argonaut") {
			return "JSON处理库"
		} else if strings.Contains(baseName, "http") || strings.Contains(baseName, "fetch") || strings.Contains(baseName, "ajax") {
			return "HTTP客户端库"
		} else if strings.Contains(baseName, "test") || strings.Contains(baseName, "spec") {
			return "测试框架"
		} else if strings.Contains(baseName, "react") || strings.Contains(baseName, "vue") {
			return "前端框架绑定"
		} else if strings.Contains(baseName, "node") {
			return "Node.js API绑定"
		} else if strings.Contains(baseName, "web") {
			return "Web API绑定"
		}

		return baseName + "相关功能"
	}

	return "PureScript包"
}

// 查找多个文件模式
func findFilePatterns(patterns []string) ([]string, error) {
	var results []string

	for _, pattern := range patterns {
		files, err := findFiles(pattern)
		if err == nil && len(files) > 0 {
			results = append(results, files...)
		}
	}

	return results, nil
}

// 格式化浮点数版本
func formatFloatVersion(v float64) string {
	return strconv.FormatFloat(v, 'f', -1, 64)
}
