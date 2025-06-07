package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

// 为ReScript添加包检测功能
func (a *App) detectReScriptWithPackages() LanguageInfo {
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
		packages, _ := a.listReScriptPackagesFixed()
		info.Packages = packages
	}

	return info
}

// 列出已安装的ReScript相关包（修复版本）
func (a *App) listReScriptPackagesFixed() ([]PackageInfo, error) {
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
