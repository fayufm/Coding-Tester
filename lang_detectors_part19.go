package main

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// 为Assembly添加包检测功能
func (a *App) detectAssemblyWithPackages() LanguageInfo {
	info := a.detectAssembly()

	if info.Installed {
		// 获取已安装的Assembly相关工具
		packages, _ := a.listAssemblyTools()
		info.Packages = packages
	}

	return info
}

// 为COBOL添加包检测功能
func (a *App) detectCOBOLWithPackages() LanguageInfo {
	info := a.detectCOBOL()

	if info.Installed {
		// 获取已安装的COBOL包
		packages, _ := a.listCOBOLPackages()
		info.Packages = packages
	}

	return info
}

// 列出COBOL包
func (a *App) listCOBOLPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 检查GnuCOBOL
	if commandExists("cobc") {
		output, err := executeCommandWithTimeout("cobc", "--info")
		if err == nil {
			// 尝试提取COBOL库信息
			lines := strings.Split(output, "\n")
			for _, line := range lines {
				if strings.Contains(line, "COB_LIBRARY_PATH") ||
					strings.Contains(line, "COB_MODULE_PATH") ||
					strings.Contains(line, "COB_COPY_DIR") {
					parts := strings.SplitN(line, ":", 2)
					if len(parts) == 2 {
						paths := strings.Split(strings.TrimSpace(parts[1]), ":")
						for _, path := range paths {
							path = strings.TrimSpace(path)
							if path != "" && path != "." {
								// 提取目录名作为库名
								name := filepath.Base(path)
								packages = append(packages, PackageInfo{
									Name:      name + " (库路径)",
									Version:   path,
									Installed: true,
								})
							}
						}
					}
				}
			}
		}
	}

	// 检查MicroFocus COBOL
	mfCobolPaths := []string{
		"C:\\Program Files\\Micro Focus\\Visual COBOL",
		"C:\\Program Files (x86)\\Micro Focus\\Visual COBOL",
		"/opt/microfocus/VisualCOBOL",
	}

	for _, path := range mfCobolPaths {
		if _, err := os.Stat(path); err == nil {
			// 扫描子目录获取库信息
			entries, err := os.ReadDir(path)
			if err == nil {
				for _, entry := range entries {
					if entry.IsDir() {
						// 排除一些常见的非库目录
						name := entry.Name()
						if !strings.EqualFold(name, "bin") &&
							!strings.EqualFold(name, "etc") &&
							!strings.EqualFold(name, "include") {
							packages = append(packages, PackageInfo{
								Name:      name,
								Version:   "已安装",
								Installed: true,
							})
						}
					}
				}
			}
		}
	}

	return packages, nil
}

// 为Fortran添加包检测功能
func (a *App) detectFortranWithPackages() LanguageInfo {
	info := a.detectFortran()

	if info.Installed {
		// 获取已安装的Fortran包
		packages, _ := a.listFortranPackages()
		info.Packages = packages
	}

	return info
}

// 列出Fortran包
func (a *App) listFortranPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 如果安装了fpm (Fortran Package Manager)
	if commandExists("fpm") {
		output, err := executeCommandWithTimeout("fpm", "list", "--registry")
		if err == nil {
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
					} else if len(parts) == 1 {
						packages = append(packages, PackageInfo{
							Name:      parts[0],
							Version:   "已安装",
							Installed: true,
						})
					}
				}
			}
		}
	}

	// 检查常见的Fortran库目录
	fortranLibDirs := []string{
		"/usr/lib/gcc/x86_64-linux-gnu",
		"/usr/local/lib/gcc",
		"C:\\Program Files\\mingw-w64\\x86_64-8.1.0-posix-seh-rt_v6-rev0\\mingw64\\lib\\gcc",
	}

	for _, baseDir := range fortranLibDirs {
		if _, err := os.Stat(baseDir); err == nil {
			// 遍历查找Fortran库
			filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return nil
				}

				// 查找Fortran库文件
				if !info.IsDir() && (strings.HasSuffix(strings.ToLower(info.Name()), ".mod") ||
					strings.HasSuffix(strings.ToLower(info.Name()), ".a") &&
						strings.Contains(strings.ToLower(info.Name()), "fortran")) {
					name := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))

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
							Version:   "系统库",
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

	return packages, nil
}

// 为Delphi添加包检测功能
func (a *App) detectDelphiWithPackages() LanguageInfo {
	info := a.detectDelphi()

	if info.Installed {
		// 获取已安装的Delphi包
		packages, _ := a.listDelphiPackages()
		info.Packages = packages
	}

	return info
}

// 列出Delphi包
func (a *App) listDelphiPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 获取用户主目录
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return packages, err
	}

	// 检查Delphi库路径
	delphiPaths := []string{
		filepath.Join(homeDir, "Documents", "Embarcadero", "Studio"),
		"C:\\Program Files (x86)\\Embarcadero\\Studio",
		"C:\\Program Files\\Embarcadero\\Studio",
	}

	for _, baseDir := range delphiPaths {
		if _, err := os.Stat(baseDir); err == nil {
			// 遍历子目录查找组件包
			filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return nil
				}

				// 查找Delphi包文件 (.bpl, .dcp, .dpk)
				if !info.IsDir() && (strings.HasSuffix(strings.ToLower(info.Name()), ".bpl") ||
					strings.HasSuffix(strings.ToLower(info.Name()), ".dcp") ||
					strings.HasSuffix(strings.ToLower(info.Name()), ".dpk")) {

					name := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))

					// 排除一些系统包
					if strings.HasPrefix(strings.ToLower(name), "rtl") ||
						strings.HasPrefix(strings.ToLower(name), "vcl") ||
						strings.HasPrefix(strings.ToLower(name), "fmx") {
						return nil
					}

					// 检查是否已经添加过
					exists := false
					for _, pkg := range packages {
						if pkg.Name == name {
							exists = true
							break
						}
					}

					if !exists {
						// 提取版本信息（如果可能）
						version := "已安装"
						versionRegex := regexp.MustCompile(`\d+\.\d+`)
						versionMatch := versionRegex.FindString(path)
						if versionMatch != "" {
							version = versionMatch
						}

						packages = append(packages, PackageInfo{
							Name:      name,
							Version:   version,
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

	return packages, nil
}

// 为Ada添加包检测功能
func (a *App) detectAdaWithPackages() LanguageInfo {
	info := a.detectAda()

	if info.Installed {
		// 获取已安装的Ada包
		packages, _ := a.listAdaPackages()
		info.Packages = packages
	}

	return info
}

// 列出Ada包
func (a *App) listAdaPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 检查GNAT (GNU Ada)
	if commandExists("gnat") {
		// 检查Ada库目录
		adaLibDirs := []string{
			"/usr/lib/gcc/x86_64-linux-gnu/*/adalib",
			"/usr/local/lib/gcc/*/adalib",
			"C:\\GNAT\\*/lib\\gcc\\*\\adalib",
		}

		for _, pattern := range adaLibDirs {
			matches, _ := filepath.Glob(pattern)
			for _, dir := range matches {
				if _, err := os.Stat(dir); err == nil {
					// 查找.ali文件（Ada库信息文件）
					entries, err := os.ReadDir(dir)
					if err != nil {
						continue
					}

					for _, entry := range entries {
						if !entry.IsDir() && strings.HasSuffix(strings.ToLower(entry.Name()), ".ali") {
							name := strings.TrimSuffix(entry.Name(), ".ali")

							// 排除一些系统库
							if strings.HasPrefix(name, "s-") ||
								strings.HasPrefix(name, "a-") ||
								strings.HasPrefix(name, "g-") {
								continue
							}

							packages = append(packages, PackageInfo{
								Name:      name,
								Version:   "已安装",
								Installed: true,
							})

							// 限制返回的包数量
							if len(packages) >= 20 {
								break
							}
						}
					}
				}
			}

			if len(packages) >= 20 {
				break
			}
		}
	}

	// 检查Alire包管理器
	if len(packages) < 20 && commandExists("alr") {
		output, err := executeCommandWithTimeout("alr", "list", "--installed")
		if err == nil {
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
	}

	return packages, nil
}

// 为VHDL添加包检测功能
func (a *App) detectVHDLWithPackages() LanguageInfo {
	info := a.detectVHDL()

	if info.Installed {
		// 获取已安装的VHDL包
		packages, _ := a.listVHDLPackages()
		info.Packages = packages
	}

	return info
}

// 列出VHDL包
func (a *App) listVHDLPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 检查GHDL
	if commandExists("ghdl") {
		// 尝试获取GHDL库路径
		output, err := executeCommandWithTimeout("ghdl", "--version")
		if err == nil {
			// 检查VHDL库目录
			vhdlLibDirs := []string{
				"/usr/local/lib/ghdl",
				"/usr/lib/ghdl",
				"C:\\Program Files\\GHDL\\lib",
			}

			// 从输出尝试提取GHDL库路径
			lines := strings.Split(output, "\n")
			for _, line := range lines {
				if strings.Contains(line, "lib") && strings.Contains(line, "ghdl") {
					parts := strings.Fields(line)
					for _, part := range parts {
						if strings.Contains(part, "lib") && strings.Contains(part, "ghdl") {
							vhdlLibDirs = append(vhdlLibDirs, part)
							break
						}
					}
				}
			}

			for _, dir := range vhdlLibDirs {
				if _, err := os.Stat(dir); err == nil {
					// 查找库目录
					entries, err := os.ReadDir(dir)
					if err != nil {
						continue
					}

					for _, entry := range entries {
						if entry.IsDir() {
							// 排除常见的非库目录
							name := entry.Name()
							if name != "work" && name != "temp" && name != "src" {
								packages = append(packages, PackageInfo{
									Name:      name,
									Version:   "GHDL库",
									Installed: true,
								})
							}
						}
					}
				}
			}
		}
	}

	// 检查ModelSim/Questa
	modelsimPaths := []string{
		"C:\\ModelSim",
		"C:\\Questa",
		"C:\\intelFPGA",
		"C:\\altera",
		"/opt/modelsim",
		"/opt/questa",
	}

	for _, baseDir := range modelsimPaths {
		if _, err := os.Stat(baseDir); err == nil {
			// 递归查找libraries目录或.lib文件
			_ = filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return nil
				}

				if info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), "libraries") {
					libEntries, err := os.ReadDir(path)
					if err != nil {
						return nil
					}

					for _, libEntry := range libEntries {
						if libEntry.IsDir() {
							name := libEntry.Name()
							// 排除常见的工作目录
							if name != "work" && name != "temp" {
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
										Version:   "ModelSim/Questa库",
										Installed: true,
									})

									if len(packages) >= 20 {
										return filepath.SkipDir
									}
								}
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
