package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// 为Kotlin添加包检测功能
func (a *App) detectKotlinWithPackages() LanguageInfo {
	info := a.detectKotlin()

	if info.Installed {
		// 获取已安装的Kotlin包
		packages, _ := a.listKotlinPackages()
		info.Packages = packages
	}

	return info
}

// 列出已安装的Kotlin包
func (a *App) listKotlinPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 检查Gradle是否可用
	if !commandExists("gradle") && !commandExists("./gradlew") {
		return packages, nil
	}

	// 首先尝试从build.gradle文件中解析依赖
	gradleFiles, _ := findFilePatterns([]string{"build.gradle", "build.gradle.kts"})

	for _, gradleFile := range gradleFiles {
		content, err := os.ReadFile(gradleFile)
		if err != nil {
			continue
		}

		contentStr := string(content)

		// 检查是否为Kotlin项目
		if !strings.Contains(contentStr, "kotlin") && !strings.Contains(contentStr, "org.jetbrains.kotlin") {
			continue
		}

		// 提取依赖
		dependencyLines := extractDependenciesFromGradle(contentStr)
		for _, line := range dependencyLines {
			line = strings.TrimSpace(line)
			if line == "" || strings.HasPrefix(line, "//") {
				continue
			}

			// 解析依赖行
			pkg := parseGradleDependency(line)
			if pkg.Name != "" {
				packages = append(packages, pkg)
			}
		}
	}

	// 如果从项目文件中没有找到包，尝试使用Gradle缓存
	if len(packages) == 0 {
		cachePkgs := findKotlinPackagesFromGradleCache()
		if len(cachePkgs) > 0 {
			packages = append(packages, cachePkgs...)
		}
	}

	// 如果仍然没有找到包，添加一些常见的Kotlin包作为建议
	if len(packages) == 0 {
		packages = a.getRecommendedKotlinPackages()
	}

	return packages, nil
}

// 从Gradle缓存中查找Kotlin包
func findKotlinPackagesFromGradleCache() []PackageInfo {
	var packages []PackageInfo

	// 获取用户主目录
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return packages
	}

	// Gradle缓存路径
	gradleCachePath := filepath.Join(homeDir, ".gradle", "caches", "modules-2", "files-2.1")

	// 检查路径是否存在
	if _, err := os.Stat(gradleCachePath); err != nil {
		return packages
	}

	// Kotlin相关的包前缀
	kotlinPrefixes := []string{
		"org.jetbrains.kotlin",
		"org.jetbrains.kotlinx",
		"io.ktor",
		"org.jetbrains.exposed",
		"io.arrow-kt",
		"org.koin",
	}

	count := 0

	// 遍历目录查找Kotlin相关包
	filepath.Walk(gradleCachePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		// 限制数量
		if count >= 20 {
			return filepath.SkipDir
		}

		// 只处理目录
		if !info.IsDir() {
			return nil
		}

		// 路径深度检查，我们只关心group/name/version结构
		relPath, err := filepath.Rel(gradleCachePath, path)
		if err != nil {
			return nil
		}

		parts := strings.Split(relPath, string(filepath.Separator))
		if len(parts) == 3 {
			group := parts[0]
			name := parts[1]
			version := parts[2]

			// 检查是否为Kotlin相关包
			isKotlinRelated := false
			for _, prefix := range kotlinPrefixes {
				if strings.HasPrefix(group, prefix) {
					isKotlinRelated = true
					break
				}
			}

			if isKotlinRelated {
				description := getKotlinPackageDescription(group + ":" + name)
				packages = append(packages, PackageInfo{
					Name:        group + ":" + name,
					Version:     version,
					Description: description,
					Installed:   true,
				})

				count++
			}
		}

		return nil
	})

	return packages
}

// 从Gradle文件中提取依赖部分
func extractDependenciesFromGradle(content string) []string {
	var lines []string

	// 寻找dependencies块
	depsStart := strings.Index(content, "dependencies {")
	if depsStart == -1 {
		return lines
	}

	// 提取dependencies块内容
	depsContent := content[depsStart:]
	openBraces := 0
	endPos := 0

	for i, char := range depsContent {
		if char == '{' {
			openBraces++
		} else if char == '}' {
			openBraces--
			if openBraces == 0 {
				endPos = i
				break
			}
		}
	}

	if endPos > 0 {
		depsBlock := depsContent[:endPos+1]
		lines = strings.Split(depsBlock, "\n")
	}

	return lines
}

// 解析Gradle依赖行
func parseGradleDependency(line string) PackageInfo {
	var pkg PackageInfo

	// 移除注释
	if idx := strings.Index(line, "//"); idx != -1 {
		line = line[:idx]
	}

	line = strings.TrimSpace(line)
	if line == "" {
		return pkg
	}

	// 常见的依赖声明模式
	patterns := []string{
		`implementation\s+['"](.+):(.+):(.+)['"]`,     // implementation 'group:name:version'
		`api\s+['"](.+):(.+):(.+)['"]`,                // api 'group:name:version'
		`compile\s+['"](.+):(.+):(.+)['"]`,            // compile 'group:name:version'
		`testImplementation\s+['"](.+):(.+):(.+)['"]`, // testImplementation 'group:name:version'
	}

	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(line)
		if len(matches) >= 4 {
			group := matches[1]
			name := matches[2]
			version := matches[3]

			// 检查是否为Kotlin相关包
			if strings.Contains(group, "kotlin") || strings.Contains(name, "kotlin") {
				description := getKotlinPackageDescription(group + ":" + name)
				pkg = PackageInfo{
					Name:        group + ":" + name,
					Version:     version,
					Description: description,
					Installed:   true,
				}
				return pkg
			}
		}
	}

	return pkg
}

// 获取推荐的Kotlin包
func (a *App) getRecommendedKotlinPackages() []PackageInfo {
	return []PackageInfo{
		{
			Name:        "org.jetbrains.kotlin:kotlin-stdlib",
			Version:     "1.8.21",
			Description: "Kotlin标准库",
			Installed:   false,
		},
		{
			Name:        "org.jetbrains.kotlinx:kotlinx-coroutines-core",
			Version:     "1.7.1",
			Description: "Kotlin协程库",
			Installed:   false,
		},
		{
			Name:        "io.ktor:ktor-client-core",
			Version:     "2.3.1",
			Description: "Kotlin HTTP客户端",
			Installed:   false,
		},
		{
			Name:        "org.jetbrains.exposed:exposed-core",
			Version:     "0.41.1",
			Description: "Kotlin SQL框架",
			Installed:   false,
		},
		{
			Name:        "io.arrow-kt:arrow-core",
			Version:     "1.2.0",
			Description: "Kotlin函数式编程库",
			Installed:   false,
		},
		{
			Name:        "org.koin:koin-core",
			Version:     "3.4.0",
			Description: "Kotlin依赖注入框架",
			Installed:   false,
		},
		{
			Name:        "org.jetbrains.kotlin:kotlin-reflect",
			Version:     "1.8.21",
			Description: "Kotlin反射库",
			Installed:   false,
		},
		{
			Name:        "org.jetbrains.kotlinx:kotlinx-serialization-json",
			Version:     "1.5.1",
			Description: "Kotlin序列化库",
			Installed:   false,
		},
	}
}

// 获取Kotlin包的描述
func getKotlinPackageDescription(name string) string {
	descriptions := map[string]string{
		"org.jetbrains.kotlin:kotlin-stdlib":               "Kotlin标准库",
		"org.jetbrains.kotlinx:kotlinx-coroutines-core":    "Kotlin协程库",
		"io.ktor:ktor-client-core":                         "Kotlin HTTP客户端",
		"io.ktor:ktor-server-core":                         "Kotlin HTTP服务器",
		"org.jetbrains.exposed:exposed-core":               "Kotlin SQL框架",
		"io.arrow-kt:arrow-core":                           "Kotlin函数式编程库",
		"org.koin:koin-core":                               "Kotlin依赖注入框架",
		"org.jetbrains.kotlin:kotlin-reflect":              "Kotlin反射库",
		"org.jetbrains.kotlinx:kotlinx-serialization-json": "Kotlin序列化库",
		"org.jetbrains.kotlin:kotlin-test":                 "Kotlin测试库",
	}

	if desc, ok := descriptions[name]; ok {
		return desc
	}

	// 尝试分析包名
	if strings.Contains(name, "kotlin") {
		if strings.Contains(name, "test") {
			return "Kotlin测试相关库"
		} else if strings.Contains(name, "coroutine") {
			return "Kotlin协程相关库"
		} else if strings.Contains(name, "serialization") {
			return "Kotlin序列化相关库"
		}
		return "Kotlin相关库"
	}

	return "Kotlin包"
}

// 为Perl添加包检测功能
func (a *App) detectPerlWithPackages() LanguageInfo {
	info := a.detectPerl()

	if info.Installed {
		// 获取已安装的Perl包
		packages, _ := a.listPerlModules()
		info.Packages = packages
	}

	return info
}

// 列出已安装的Perl模块
func (a *App) listPerlModules() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 检查perl命令是否可用
	if !commandExists("perl") {
		return packages, nil
	}

	// 使用perl -M命令获取已安装的模块列表
	output, err := executeCommandWithTimeout("perl", "-e", "use ExtUtils::Installed; my $inst = ExtUtils::Installed->new(); foreach my $module ($inst->modules()) { print $module, \"\\n\"; }")
	if err != nil {
		// 如果上面的命令失败，尝试另一种方式
		output, err = executeCommandWithTimeout("perl", "-e", "foreach $path (@INC) { next if $path eq '.'; opendir(DIR, $path) or next; foreach $file (sort readdir(DIR)) { next unless $file =~ /\\.pm$/; $file =~ s/\\.pm$//; print \"$file\\n\"; } closedir DIR; }")
		if err != nil {
			return packages, err
		}
	}

	// 解析输出
	lines := strings.Split(output, "\n")
	moduleMap := make(map[string]bool) // 使用map去重

	for _, line := range lines {
		module := strings.TrimSpace(line)
		if module == "" || moduleMap[module] {
			continue
		}

		moduleMap[module] = true

		// 尝试获取模块版本
		version := "installed"
		versionOutput, _ := executeCommandWithTimeout("perl", "-e", "eval \"use "+module+"; print $"+module+"::VERSION;\" or print \"unknown\";")
		if versionOutput != "" && versionOutput != "unknown" {
			version = versionOutput
		}

		description := getPerlModuleDescription(module)
		packages = append(packages, PackageInfo{
			Name:        module,
			Version:     version,
			Description: description,
			Installed:   true,
		})

		// 限制数量
		if len(packages) >= 20 {
			break
		}
	}

	// 如果没有找到包，可能是因为权限问题，尝试使用cpan列出已安装的模块
	if len(packages) == 0 && commandExists("cpan") {
		output, err = executeCommandWithTimeout("cpan", "-l")
		if err == nil {
			lines := strings.Split(output, "\n")
			for _, line := range lines {
				module := strings.TrimSpace(line)
				if module == "" || moduleMap[module] {
					continue
				}

				moduleMap[module] = true
				description := getPerlModuleDescription(module)
				packages = append(packages, PackageInfo{
					Name:        module,
					Version:     "installed",
					Description: description,
					Installed:   true,
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

// 获取Perl模块的描述
func getPerlModuleDescription(name string) string {
	descriptions := map[string]string{
		"Moose":          "Perl 5的后现代对象系统",
		"DBI":            "Perl数据库接口模块",
		"Mojolicious":    "Perl的实时Web框架",
		"Dancer2":        "轻量级Perl Web应用框架",
		"Catalyst":       "Perl的MVC Web框架",
		"JSON":           "JSON (JavaScript Object Notation)编码器/解码器",
		"LWP::UserAgent": "Web用户代理类",
		"HTTP::Request":  "HTTP请求类",
		"XML::Simple":    "易于使用的XML解析器",
		"Test::More":     "Perl的单元测试框架",
		"DateTime":       "日期和时间对象",
		"DBIx::Class":    "可扩展和灵活的对象关系映射器",
		"Template":       "模板处理系统",
		"Plack":          "PSGI工具包和适配器",
		"File::Spec":     "可移植的文件名处理",
		"Module::Build":  "构建和安装Perl模块的系统",
		"Carp":           "错误报告模块",
		"Config::Tiny":   "小型简单的配置文件处理器",
		"Data::Dumper":   "将数据结构漂亮地打印出来",
		"CGI":            "通用网关接口模块",
	}

	if desc, ok := descriptions[name]; ok {
		return desc
	}

	// 尝试推断描述
	if strings.Contains(name, "::Test") || strings.Contains(name, "Test::") {
		return "测试相关模块"
	} else if strings.Contains(name, "::DB") || strings.Contains(name, "DB::") || strings.Contains(name, "::SQL") {
		return "数据库相关模块"
	} else if strings.Contains(name, "::Web") || strings.Contains(name, "HTTP::") || strings.Contains(name, "HTML::") {
		return "Web相关模块"
	} else if strings.Contains(name, "::JSON") || strings.Contains(name, "JSON::") {
		return "JSON处理模块"
	} else if strings.Contains(name, "::XML") || strings.Contains(name, "XML::") {
		return "XML处理模块"
	} else if strings.Contains(name, "::File") || strings.Contains(name, "File::") {
		return "文件处理模块"
	}

	return "Perl模块"
}

// 为Haxe添加包检测功能
func (a *App) detectHaxeWithPackages() LanguageInfo {
	info := a.detectHaxe()

	if info.Installed {
		// 获取已安装的Haxe包
		packages, _ := a.listHaxePackages()
		info.Packages = packages
	}

	return info
}

// 列出已安装的Haxe包
func (a *App) listHaxePackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 检查haxelib命令是否可用
	if !commandExists("haxelib") {
		return packages, nil
	}

	// 使用haxelib list命令获取已安装的包列表
	output, err := executeCommandWithTimeout("haxelib", "list")
	if err != nil {
		return packages, err
	}

	// 解析输出
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// 解析格式为 "packagename: version" 的行
		parts := strings.SplitN(line, ":", 2)
		if len(parts) >= 2 {
			name := strings.TrimSpace(parts[0])
			version := strings.TrimSpace(parts[1])

			description := getHaxePackageDescription(name)
			packages = append(packages, PackageInfo{
				Name:        name,
				Version:     version,
				Description: description,
				Installed:   true,
			})
		}
	}

	// 如果没有找到包，尝试从项目文件中查找
	if len(packages) == 0 {
		projectPkgs := findHaxePackagesFromProjects()
		if len(projectPkgs) > 0 {
			packages = append(packages, projectPkgs...)
		}
	}

	// 如果仍然没有找到包，添加一些常见的Haxe包作为建议
	if len(packages) == 0 {
		packages = getRecommendedHaxePackages()
	}

	return packages, nil
}

// 从项目文件中查找Haxe包
func findHaxePackagesFromProjects() []PackageInfo {
	var packages []PackageInfo

	// 查找project.xml或haxelib.json文件
	projectFiles, _ := findFilePatterns([]string{"project.xml", "haxelib.json", "*.hxml"})

	for _, path := range projectFiles {
		content, err := os.ReadFile(path)
		if err != nil {
			continue
		}

		contentStr := string(content)

		if strings.HasSuffix(path, ".json") {
			// 解析haxelib.json
			var haxelibJson map[string]interface{}
			if err := json.Unmarshal(content, &haxelibJson); err != nil {
				continue
			}

			// 检查dependencies字段
			if dependencies, ok := haxelibJson["dependencies"].(map[string]interface{}); ok {
				for name, version := range dependencies {
					versionStr := ""
					switch v := version.(type) {
					case string:
						versionStr = v
					case float64:
						versionStr = fmt.Sprintf("%v", v)
					}

					description := getHaxePackageDescription(name)
					packages = append(packages, PackageInfo{
						Name:        name,
						Version:     versionStr,
						Description: description,
						Installed:   true,
					})
				}
			}
		} else if strings.HasSuffix(path, ".xml") {
			// 简单解析project.xml中的haxelib标签
			if strings.Contains(contentStr, "<haxelib") {
				lines := strings.Split(contentStr, "\n")
				for _, line := range lines {
					if strings.Contains(line, "<haxelib") && strings.Contains(line, "name=") {
						nameStart := strings.Index(line, "name=\"")
						if nameStart != -1 {
							nameStart += 6
							nameEnd := strings.Index(line[nameStart:], "\"")
							if nameEnd != -1 {
								name := line[nameStart : nameStart+nameEnd]

								// 尝试提取版本
								version := "latest"
								versionStart := strings.Index(line, "version=\"")
								if versionStart != -1 {
									versionStart += 9
									versionEnd := strings.Index(line[versionStart:], "\"")
									if versionEnd != -1 {
										version = line[versionStart : versionStart+versionEnd]
									}
								}

								description := getHaxePackageDescription(name)
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
		} else if strings.HasSuffix(path, ".hxml") {
			// 解析.hxml文件中的-lib参数
			lines := strings.Split(contentStr, "\n")
			for _, line := range lines {
				line = strings.TrimSpace(line)
				if strings.HasPrefix(line, "-lib ") {
					name := strings.TrimSpace(line[5:])
					description := getHaxePackageDescription(name)
					packages = append(packages, PackageInfo{
						Name:        name,
						Version:     "specified in hxml",
						Description: description,
						Installed:   true,
					})
				}
			}
		}
	}

	return packages
}

// 获取推荐的Haxe包
func getRecommendedHaxePackages() []PackageInfo {
	return []PackageInfo{
		{
			Name:        "openfl",
			Version:     "9.2.1",
			Description: "跨平台应用程序框架",
			Installed:   false,
		},
		{
			Name:        "lime",
			Version:     "8.0.1",
			Description: "轻量级跨平台游戏框架",
			Installed:   false,
		},
		{
			Name:        "heaps",
			Version:     "1.10.0",
			Description: "高性能游戏框架",
			Installed:   false,
		},
		{
			Name:        "haxeui-core",
			Version:     "1.6.0",
			Description: "跨平台UI框架",
			Installed:   false,
		},
		{
			Name:        "flixel",
			Version:     "5.3.1",
			Description: "2D游戏引擎",
			Installed:   false,
		},
		{
			Name:        "hxcpp",
			Version:     "4.3.2",
			Description: "C++后端",
			Installed:   false,
		},
		{
			Name:        "tink_core",
			Version:     "2.0.2",
			Description: "函数式编程工具",
			Installed:   false,
		},
		{
			Name:        "hxnodejs",
			Version:     "12.1.0",
			Description: "Node.js支持",
			Installed:   false,
		},
	}
}

// 获取Haxe包的描述
func getHaxePackageDescription(name string) string {
	descriptions := map[string]string{
		"openfl":       "跨平台应用程序框架",
		"lime":         "轻量级跨平台游戏框架",
		"heaps":        "高性能游戏框架",
		"haxeui-core":  "跨平台UI框架",
		"flixel":       "2D游戏引擎",
		"hxcpp":        "C++后端",
		"tink_core":    "函数式编程工具",
		"hxnodejs":     "Node.js支持",
		"hxjava":       "Java后端",
		"hxcs":         "C#后端",
		"hxpy":         "Python后端",
		"hxphp":        "PHP后端",
		"actuate":      "动画库",
		"away3d":       "3D引擎",
		"hscript":      "脚本解释器",
		"hxmath":       "数学库",
		"polygonal-ds": "数据结构库",
		"format":       "各种文件格式的库",
		"hxssl":        "OpenSSL绑定",
		"haxe-strings": "字符串操作库",
	}

	if desc, ok := descriptions[name]; ok {
		return desc
	}

	// 尝试推断描述
	if strings.Contains(name, "ui") {
		return "UI相关库"
	} else if strings.Contains(name, "game") || strings.Contains(name, "engine") {
		return "游戏相关库"
	} else if strings.Contains(name, "net") || strings.Contains(name, "http") {
		return "网络相关库"
	} else if strings.Contains(name, "db") || strings.Contains(name, "sql") {
		return "数据库相关库"
	} else if strings.Contains(name, "test") {
		return "测试相关库"
	} else if strings.Contains(name, "format") || strings.Contains(name, "parser") {
		return "解析相关库"
	}

	return "Haxe库"
}
