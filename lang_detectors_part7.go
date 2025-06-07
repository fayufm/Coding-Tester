package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// 为Hack添加包检测功能
func (a *App) detectHackWithPackages() LanguageInfo {
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

		// 获取已安装的Hack包
		packages, _ := a.listHackPackages()
		info.Packages = packages
	}

	return info
}

// 列出已安装的Hack包
func (a *App) listHackPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 检查composer是否安装
	if !commandExists("composer") {
		return packages, nil
	}

	// 尝试在常见目录中查找composer.json
	projectDirs := []string{
		".", "./src", "./app",
	}

	for _, dir := range projectDirs {
		composerJsonPath := filepath.Join(dir, "composer.json")
		composerLockPath := filepath.Join(dir, "composer.lock")

		// 首先尝试从composer.lock获取更准确的已安装版本信息
		if _, err := os.Stat(composerLockPath); err == nil {
			content, err := os.ReadFile(composerLockPath)
			if err != nil {
				continue
			}

			// 解析JSON内容
			var composerLock struct {
				Packages []struct {
					Name    string `json:"name"`
					Version string `json:"version"`
				} `json:"packages"`
			}

			if err := json.Unmarshal(content, &composerLock); err != nil {
				continue
			}

			// 添加Hack相关包
			for _, pkg := range composerLock.Packages {
				if strings.Contains(pkg.Name, "hhvm") ||
					strings.Contains(pkg.Name, "hack") ||
					strings.Contains(pkg.Name, "facebook") {
					packages = append(packages, PackageInfo{
						Name:      pkg.Name,
						Version:   pkg.Version,
						Installed: true,
					})

					// 限制数量
					if len(packages) >= 20 {
						return packages, nil
					}
				}
			}
		}

		// 如果没有找到lock文件或没有找到足够的包，尝试从composer.json获取
		if len(packages) < 5 {
			if _, err := os.Stat(composerJsonPath); err == nil {
				content, err := os.ReadFile(composerJsonPath)
				if err != nil {
					continue
				}

				// 解析JSON内容
				var composerJson struct {
					Require    map[string]string `json:"require"`
					RequireDev map[string]string `json:"require-dev"`
				}

				if err := json.Unmarshal(content, &composerJson); err != nil {
					continue
				}

				// 添加Hack相关包
				for _, deps := range []map[string]string{composerJson.Require, composerJson.RequireDev} {
					for pkg, version := range deps {
						if strings.Contains(pkg, "hhvm") ||
							strings.Contains(pkg, "hack") ||
							strings.Contains(pkg, "facebook") {
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
	}

	return packages, nil
}

// 为Jython添加包检测功能
func (a *App) detectJythonWithPackages() LanguageInfo {
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

		// 获取已安装的Jython包
		packages, _ := a.listJythonPackages()
		info.Packages = packages
	}

	return info
}

// 列出已安装的Jython包
func (a *App) listJythonPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 尝试使用jython的pip列出已安装的包
	output, err := executeCommandWithTimeout("jython", "-m", "pip", "list")
	if err == nil {
		lines := strings.Split(output, "\n")
		// 跳过标题行
		for i := 2; i < len(lines); i++ {
			line := strings.TrimSpace(lines[i])
			if line == "" {
				continue
			}

			parts := strings.Fields(line)
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
					break
				}
			}
		}
	}

	// 如果没有找到包，尝试使用easy_install
	if len(packages) == 0 {
		output, err := executeCommandWithTimeout("jython", "-m", "easy_install", "--list")
		if err == nil {
			lines := strings.Split(output, "\n")
			for _, line := range lines {
				line = strings.TrimSpace(line)
				if strings.Contains(line, " ") {
					parts := strings.Fields(line)
					if len(parts) >= 1 {
						name := parts[0]
						version := "installed"

						// 尝试提取版本
						if len(parts) >= 2 && strings.HasPrefix(parts[1], "(") {
							version = strings.Trim(parts[1], "()")
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
		}
	}

	// 检查requirements.txt文件
	requirementsFiles := []string{
		"requirements.txt", "jython-requirements.txt", "./src/requirements.txt",
	}

	for _, reqFile := range requirementsFiles {
		if _, err := os.Stat(reqFile); err == nil {
			content, err := os.ReadFile(reqFile)
			if err != nil {
				continue
			}

			lines := strings.Split(string(content), "\n")
			for _, line := range lines {
				line = strings.TrimSpace(line)
				if line == "" || strings.HasPrefix(line, "#") {
					continue
				}

				// 解析包名和版本
				var name, version string
				if strings.Contains(line, "==") {
					parts := strings.Split(line, "==")
					name = parts[0]
					if len(parts) > 1 {
						version = parts[1]
					} else {
						version = "specified in requirements"
					}
				} else if strings.Contains(line, ">=") {
					parts := strings.Split(line, ">=")
					name = parts[0]
					if len(parts) > 1 {
						version = ">=" + parts[1]
					} else {
						version = "specified in requirements"
					}
				} else {
					name = line
					version = "specified in requirements"
				}

				// 检查是否已添加
				alreadyAdded := false
				for _, existingPkg := range packages {
					if existingPkg.Name == name {
						alreadyAdded = true
						break
					}
				}

				if !alreadyAdded {
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

	return packages, nil
}

// 为Boo添加包检测功能
func (a *App) detectBooWithPackages() LanguageInfo {
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
	if !commandExists("booc") {
		return info
	}

	output, err := executeCommandWithTimeout("booc", "-version")
	if err == nil {
		info.Installed = true
		info.Version = output

		// 获取已安装的Boo包
		packages, _ := a.listBooPackages()
		info.Packages = packages
	}

	return info
}

// 列出已安装的Boo包
func (a *App) listBooPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 检查是否安装了NuGet
	if !commandExists("nuget") && !commandExists("dotnet") {
		return packages, nil
	}

	// 尝试在常见目录中查找packages.config或*.csproj文件
	projectFiles := []string{}

	// 查找packages.config文件
	packagesConfigs, _ := filepath.Glob("**/packages.config")
	projectFiles = append(projectFiles, packagesConfigs...)

	// 查找*.csproj文件
	csprojFiles, _ := filepath.Glob("**/*.csproj")
	projectFiles = append(projectFiles, csprojFiles...)

	for _, projectFile := range projectFiles {
		content, err := os.ReadFile(projectFile)
		if err != nil {
			continue
		}

		// 检查是否包含Boo相关包
		contentStr := string(content)
		if strings.Contains(contentStr, "Boo.Lang") ||
			strings.Contains(contentStr, "boo-lang") ||
			strings.Contains(contentStr, "booc") {

			// 尝试提取包信息
			if strings.HasSuffix(projectFile, ".config") {
				// 从packages.config提取
				// 简单解析XML
				lines := strings.Split(contentStr, "\n")
				for _, line := range lines {
					if strings.Contains(line, "package") && strings.Contains(line, "id=") {
						var name, version string

						// 提取id
						idStart := strings.Index(line, "id=\"")
						if idStart != -1 {
							idStart += 4
							idEnd := strings.Index(line[idStart:], "\"")
							if idEnd != -1 {
								name = line[idStart : idStart+idEnd]
							}
						}

						// 提取version
						versionStart := strings.Index(line, "version=\"")
						if versionStart != -1 {
							versionStart += 9
							versionEnd := strings.Index(line[versionStart:], "\"")
							if versionEnd != -1 {
								version = line[versionStart : versionStart+versionEnd]
							}
						}

						if name != "" && (strings.Contains(name, "Boo") || strings.Contains(name, "boo")) {
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
			} else if strings.HasSuffix(projectFile, ".csproj") {
				// 从.csproj提取
				lines := strings.Split(contentStr, "\n")
				for _, line := range lines {
					if strings.Contains(line, "PackageReference") && strings.Contains(line, "Include=") {
						var name, version string

						// 提取Include
						includeStart := strings.Index(line, "Include=\"")
						if includeStart != -1 {
							includeStart += 9
							includeEnd := strings.Index(line[includeStart:], "\"")
							if includeEnd != -1 {
								name = line[includeStart : includeStart+includeEnd]
							}
						}

						// 提取Version
						versionStart := strings.Index(line, "Version=\"")
						if versionStart != -1 {
							versionStart += 9
							versionEnd := strings.Index(line[versionStart:], "\"")
							if versionEnd != -1 {
								version = line[versionStart : versionStart+versionEnd]
							}
						}

						if name != "" && (strings.Contains(name, "Boo") || strings.Contains(name, "boo")) {
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

	// 如果没有找到包，添加默认的Boo包
	if len(packages) == 0 {
		defaultBooPackages := []struct {
			Name        string
			Description string
			Version     string
		}{
			{"Boo.Lang", "Boo语言运行时库", "installed"},
			{"Boo.Lang.Compiler", "Boo编译器", "installed"},
			{"Boo.Lang.Parser", "Boo解析器", "installed"},
		}

		for _, pkg := range defaultBooPackages {
			packages = append(packages, PackageInfo{
				Name:        pkg.Name,
				Description: pkg.Description,
				Version:     pkg.Version,
				Installed:   true,
			})
		}
	}

	return packages, nil
}

// 为Ceylon添加包检测功能
func (a *App) detectCeylonWithPackages() LanguageInfo {
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

		// 获取已安装的Ceylon包
		packages, _ := a.listCeylonPackages()
		info.Packages = packages
	}

	return info
}

// 列出已安装的Ceylon包
func (a *App) listCeylonPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 尝试列出已安装的Ceylon模块
	output, err := executeCommandWithTimeout("ceylon", "info", "--list-modules")
	if err == nil {
		lines := strings.Split(output, "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" || !strings.Contains(line, "/") {
				continue
			}

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
					break
				}
			}
		}
	}

	// 检查module.ceylon文件
	moduleFiles, _ := filepath.Glob("**/module.ceylon")
	for _, moduleFile := range moduleFiles {
		content, err := os.ReadFile(moduleFile)
		if err != nil {
			continue
		}

		contentStr := string(content)
		lines := strings.Split(contentStr, "\n")
		for _, line := range lines {
			if strings.Contains(line, "import") && strings.Contains(line, "\"") {
				// 提取模块名
				moduleStart := strings.Index(line, "\"")
				if moduleStart != -1 {
					moduleStart++
					moduleEnd := strings.Index(line[moduleStart:], "\"")
					if moduleEnd != -1 {
						moduleName := line[moduleStart : moduleStart+moduleEnd]

						// 提取版本
						version := "specified in module.ceylon"
						versionStart := strings.Index(line[moduleStart+moduleEnd:], "\"")
						if versionStart != -1 {
							versionStart = moduleStart + moduleEnd + versionStart + 1
							versionEnd := strings.Index(line[versionStart:], "\"")
							if versionEnd != -1 {
								version = line[versionStart : versionStart+versionEnd]
							}
						}

						// 检查是否已添加
						alreadyAdded := false
						for _, existingPkg := range packages {
							if existingPkg.Name == moduleName {
								alreadyAdded = true
								break
							}
						}

						if !alreadyAdded {
							packages = append(packages, PackageInfo{
								Name:      moduleName,
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

	// 如果没有找到包，添加默认的Ceylon包
	if len(packages) == 0 {
		defaultCeylonPackages := []struct {
			Name    string
			Version string
		}{
			{"ceylon.language", "1.3.3"},
			{"ceylon.collection", "1.3.3"},
			{"ceylon.file", "1.3.3"},
		}

		for _, pkg := range defaultCeylonPackages {
			packages = append(packages, PackageInfo{
				Name:      pkg.Name,
				Version:   pkg.Version,
				Installed: true,
			})
		}
	}

	return packages, nil
}

// 为XSLT添加包检测功能
func (a *App) detectXSLTWithPackages() LanguageInfo {
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

			// 添加默认的XSLT处理器
			packages := []PackageInfo{
				{
					Name:        "libxslt",
					Description: "XSLT库",
					Version:     extractVersionFromString(output),
					Installed:   true,
				},
			}

			info.Packages = packages
			return info
		}
	}

	// 检查Saxon
	if commandExists("saxon") {
		output, err := executeCommandWithTimeout("saxon", "--version")
		if err == nil {
			info.Installed = true
			info.Version = "Saxon: " + output

			// 添加Saxon包
			packages := []PackageInfo{
				{
					Name:        "Saxon",
					Description: "XSLT和XQuery处理器",
					Version:     extractVersionFromString(output),
					Installed:   true,
				},
			}

			info.Packages = packages
			return info
		}
	}

	return info
}

// 从字符串中提取版本号
func extractVersionFromString(input string) string {
	// 简单的版本提取
	for _, word := range strings.Fields(input) {
		if strings.ContainsAny(word, "0123456789") && strings.Contains(word, ".") {
			// 清理非版本字符
			version := strings.Trim(word, "()[]{},:;\"'")
			if strings.ContainsAny(version, "0123456789") && strings.Contains(version, ".") {
				return version
			}
		}
	}

	return "installed"
}

// 为Ruby添加包检测功能
func (a *App) detectRubyWithPackages() LanguageInfo {
	info := a.detectRuby()

	if info.Installed {
		// 获取已安装的Ruby包
		packages, _ := a.listRubyGems()
		info.Packages = packages
	}

	return info
}

// 列出已安装的Ruby包（gem）
func (a *App) listRubyGems() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 检查gem命令是否可用
	if !commandExists("gem") {
		return packages, nil
	}

	// 获取已安装的gem列表
	output, err := executeCommandWithTimeout("gem", "list", "--local")
	if err != nil {
		return packages, err
	}

	// 解析gem列表
	return a.parseGemList(output)
}

// 解析gem列表输出
func (a *App) parseGemList(output string) ([]PackageInfo, error) {
	var packages []PackageInfo

	// 使用正则表达式解析gem list的输出
	// 格式通常为: gem_name (version1, version2, ...)
	re := regexp.MustCompile(`([^\s]+)\s+\(([^)]+)\)`)
	matches := re.FindAllStringSubmatch(output, -1)

	for _, match := range matches {
		if len(match) >= 3 {
			name := match[1]
			versions := strings.Split(match[2], ", ")
			latestVersion := versions[0] // 通常第一个是最新的版本

			description := getRubyGemDescription(name)
			packages = append(packages, PackageInfo{
				Name:        name,
				Version:     latestVersion,
				Description: description,
				Installed:   true,
			})
		}
	}

	// 如果没有找到包，可能格式不同，尝试另一种解析方式
	if len(packages) == 0 {
		lines := strings.Split(output, "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}

			parts := strings.Fields(line)
			if len(parts) >= 1 {
				name := parts[0]
				version := ""
				if len(parts) >= 2 {
					// 移除版本号周围的括号
					version = strings.Trim(parts[1], "()")
				}

				description := getRubyGemDescription(name)
				packages = append(packages, PackageInfo{
					Name:        name,
					Version:     version,
					Description: description,
					Installed:   true,
				})
			}
		}
	}

	// 获取一些包的详细信息
	for i := 0; i < len(packages) && i < 10; i++ { // 限制为前10个包以提高性能
		pkgInfo, err := a.getGemInfo(packages[i].Name)
		if err == nil && pkgInfo.Description != "" {
			packages[i].Description = pkgInfo.Description
		}
	}

	return packages, nil
}

// 获取Ruby包（gem）的详细信息
func (a *App) getGemInfo(name string) (PackageInfo, error) {
	var pkgInfo PackageInfo
	pkgInfo.Name = name
	pkgInfo.Installed = true

	// 使用gem info命令获取详细信息
	output, err := executeCommandWithTimeout("gem", "info", name)
	if err != nil {
		return pkgInfo, err
	}

	// 解析输出
	lines := strings.Split(output, "\n")
	for i, line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "Version:") {
			pkgInfo.Version = strings.TrimSpace(strings.TrimPrefix(line, "Version:"))
		} else if strings.HasPrefix(line, "Summary:") || strings.HasPrefix(line, "Description:") {
			pkgInfo.Description = strings.TrimSpace(strings.TrimPrefix(strings.TrimPrefix(line, "Summary:"), "Description:"))
		}

		// 如果找不到明确的描述行，尝试获取第二行作为描述
		if pkgInfo.Description == "" && i == 1 {
			pkgInfo.Description = line
		}
	}

	return pkgInfo, nil
}

// 获取Ruby包的描述
func getRubyGemDescription(name string) string {
	descriptions := map[string]string{
		"rails":        "Ruby on Rails是一个全栈Web应用框架",
		"rake":         "Ruby构建工具和任务运行器",
		"bundler":      "Ruby依赖管理器",
		"rspec":        "Ruby的行为驱动开发测试框架",
		"sinatra":      "轻量级Web应用框架",
		"devise":       "灵活的Rails身份验证解决方案",
		"activerecord": "Rails的ORM框架",
		"puma":         "Ruby/Rack服务器",
		"nokogiri":     "HTML、XML、SAX和Reader解析器",
		"sidekiq":      "Ruby的后台作业处理框架",
		"capybara":     "Web应用的集成测试工具",
		"rubocop":      "Ruby代码风格检查和静态分析工具",
		"webpacker":    "Rails中使用Webpack管理JavaScript的工具",
		"faraday":      "HTTP客户端库",
		"minitest":     "Ruby的完整套件测试工具",
	}

	if desc, ok := descriptions[name]; ok {
		return desc
	}

	// 根据包名推断描述
	if strings.Contains(name, "rails") {
		return "Rails相关库"
	} else if strings.Contains(name, "json") {
		return "JSON解析或生成库"
	} else if strings.Contains(name, "http") || strings.Contains(name, "net") {
		return "网络或HTTP相关库"
	} else if strings.Contains(name, "test") || strings.Contains(name, "spec") {
		return "测试相关库"
	} else if strings.Contains(name, "db") || strings.Contains(name, "sql") {
		return "数据库相关库"
	}

	return "Ruby gem"
}
