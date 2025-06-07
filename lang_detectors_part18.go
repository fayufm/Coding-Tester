package main

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// 为Groovy添加包检测功能
func (a *App) detectGroovyWithPackages() LanguageInfo {
	info := a.detectGroovy()

	if info.Installed {
		// 获取已安装的Groovy包
		packages, _ := a.listGroovyPackages()
		info.Packages = packages
	}

	return info
}

// 列出Groovy包
func (a *App) listGroovyPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 获取用户主目录
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return packages, err
	}

	// 检查Grape缓存目录
	grapeCacheDir := filepath.Join(homeDir, ".groovy", "grapes")
	if _, err := os.Stat(grapeCacheDir); os.IsNotExist(err) {
		// 尝试另一个可能的位置
		grapeCacheDir = filepath.Join(homeDir, ".gradle", "caches", "modules-2", "files-2.1")
	}

	if _, err := os.Stat(grapeCacheDir); err == nil {
		// 遍历目录查找jar文件
		count := 0
		_ = filepath.Walk(grapeCacheDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}

			// 查找jar文件，提取包信息
			if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".jar") {
				fileName := info.Name()

				// 尝试从文件名中提取包名和版本
				re := regexp.MustCompile(`^(.+)-(\d+\.\d+.*?)\.jar$`)
				matches := re.FindStringSubmatch(fileName)

				if len(matches) >= 3 {
					packageName := matches[1]
					version := matches[2]

					packages = append(packages, PackageInfo{
						Name:      packageName,
						Version:   version,
						Installed: true,
					})

					count++
					if count >= 20 {
						return filepath.SkipDir
					}
				}
			}
			return nil
		})
	}

	return packages, nil
}

// 为MATLAB添加包检测功能
func (a *App) detectMatlabWithPackages() LanguageInfo {
	info := a.detectMatlab()

	if info.Installed {
		// 获取已安装的MATLAB包
		packages, _ := a.listMatlabPackages()
		info.Packages = packages
	}

	return info
}

// 列出MATLAB包
func (a *App) listMatlabPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 获取用户主目录
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return packages, err
	}

	// MATLAB附加组件目录
	toolboxDirs := []string{
		filepath.Join(homeDir, "Documents", "MATLAB", "Add-Ons"),
		"C:\\Program Files\\MATLAB\\R2023a\\toolbox",
		"C:\\Program Files\\MATLAB\\R2022b\\toolbox",
		"C:\\Program Files\\MATLAB\\R2022a\\toolbox",
		"C:\\Program Files\\MATLAB\\R2021b\\toolbox",
		"/Applications/MATLAB_R2023a.app/toolbox",
		"/usr/local/MATLAB/R2023a/toolbox",
	}

	for _, dir := range toolboxDirs {
		if _, err := os.Stat(dir); err == nil {
			// 目录存在，查找工具箱
			dirs, err := os.ReadDir(dir)
			if err != nil {
				continue
			}

			for _, entry := range dirs {
				if entry.IsDir() {
					dirName := entry.Name()
					// 跳过一些常见的系统目录
					if dirName == "local" || dirName == "matlab" || dirName == "shared" {
						continue
					}

					// 尝试从info.xml或类似文件中获取版本信息
					versionFiles := []string{
						filepath.Join(dir, dirName, "info.xml"),
						filepath.Join(dir, dirName, "version.txt"),
					}

					version := "安装"
					for _, vFile := range versionFiles {
						if _, err := os.Stat(vFile); err == nil {
							content, err := os.ReadFile(vFile)
							if err == nil {
								// 尝试从文件中提取版本信息
								re := regexp.MustCompile(`[Vv]ersion[>\s":=]+([0-9\.]+)`)
								matches := re.FindSubmatch(content)
								if len(matches) >= 2 {
									version = string(matches[1])
									break
								}
							}
						}
					}

					packages = append(packages, PackageInfo{
						Name:      dirName + " Toolbox",
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

	return packages, nil
}

// 为Scheme添加包检测功能
func (a *App) detectSchemeWithPackages() LanguageInfo {
	info := a.detectScheme()

	if info.Installed {
		// 获取已安装的Scheme包
		packages, _ := a.listSchemePackages()
		info.Packages = packages
	}

	return info
}

// 列出Scheme包
func (a *App) listSchemePackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 检查是否安装了Guile包管理器
	if commandExists("guile") {
		output, err := executeCommandWithTimeout("guile", "-c", "(display (map car (map cdr (append-map cdr (hash-map->list cons (module-submodules (resolve-interface '(ice-9 session))))))))")
		if err == nil && output != "" {
			modules := strings.Split(strings.Trim(output, "()"), " ")
			for _, module := range modules {
				module = strings.TrimSpace(module)
				if module != "" && !strings.HasPrefix(module, ".") {
					packages = append(packages, PackageInfo{
						Name:      module,
						Version:   "内置",
						Installed: true,
					})

					if len(packages) >= 20 {
						break
					}
				}
			}
		}
	}

	// 检查是否安装了Racket
	if len(packages) == 0 && commandExists("raco") {
		output, err := executeCommandWithTimeout("raco", "pkg", "show")
		if err == nil && output != "" {
			lines := strings.Split(output, "\n")
			for _, line := range lines {
				line = strings.TrimSpace(line)
				if line != "" {
					parts := strings.Fields(line)
					if len(parts) > 0 {
						pkgName := parts[0]
						version := "已安装"

						if len(parts) > 1 {
							version = strings.Join(parts[1:], " ")
						}

						packages = append(packages, PackageInfo{
							Name:      pkgName,
							Version:   version,
							Installed: true,
						})

						if len(packages) >= 20 {
							break
						}
					}
				}
			}
		}
	}

	// 检查是否安装了Akku包管理器
	if len(packages) == 0 && commandExists("akku") {
		output, err := executeCommandWithTimeout("akku", "list")
		if err == nil && output != "" {
			lines := strings.Split(output, "\n")
			for _, line := range lines {
				line = strings.TrimSpace(line)
				if line != "" && !strings.HasPrefix(line, "#") {
					parts := strings.Fields(line)
					if len(parts) >= 2 {
						packages = append(packages, PackageInfo{
							Name:      parts[0],
							Version:   parts[1],
							Installed: true,
						})

						if len(packages) >= 20 {
							break
						}
					}
				}
			}
		}
	}

	return packages, nil
}
