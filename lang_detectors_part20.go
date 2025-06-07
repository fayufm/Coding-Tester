package main

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// 为Erlang添加包检测功能
func (a *App) detectErlangWithPackages() LanguageInfo {
	info := a.detectErlang()

	if info.Installed {
		// 获取已安装的Erlang包
		packages, _ := a.listErlangPackages()
		info.Packages = packages
	}

	return info
}

// 列出Erlang包
func (a *App) listErlangPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 检查rebar3是否可用（Erlang包管理器）
	if commandExists("rebar3") {
		output, err := executeCommandWithTimeout("rebar3", "plugins", "list")
		if err == nil {
			lines := strings.Split(output, "\n")
			for _, line := range lines {
				line = strings.TrimSpace(line)
				if line != "" && !strings.HasPrefix(line, "==") && !strings.HasPrefix(line, "Available") {
					parts := strings.Fields(line)
					if len(parts) > 0 {
						name := parts[0]
						version := "已安装"
						if len(parts) > 1 {
							version = parts[1]
						}

						packages = append(packages, PackageInfo{
							Name:      name,
							Version:   version,
							Installed: true,
						})
					}
				}
			}
		}
	}

	// 检查Erlang应用程序目录
	erlangAppDirs := []string{
		"/usr/lib/erlang/lib",
		"C:\\Program Files\\erl-*\\lib",
		"C:\\Program Files\\Erlang*\\lib",
	}

	for _, pattern := range erlangAppDirs {
		matches, _ := filepath.Glob(pattern)
		for _, dir := range matches {
			if _, err := os.Stat(dir); err == nil {
				entries, err := os.ReadDir(dir)
				if err != nil {
					continue
				}

				for _, entry := range entries {
					if entry.IsDir() {
						// 提取应用名称和版本
						nameParts := strings.Split(entry.Name(), "-")
						if len(nameParts) >= 2 {
							name := nameParts[0]
							version := strings.Join(nameParts[1:], "-")

							packages = append(packages, PackageInfo{
								Name:      name,
								Version:   version,
								Installed: true,
							})
						} else {
							packages = append(packages, PackageInfo{
								Name:      entry.Name(),
								Version:   "已安装",
								Installed: true,
							})
						}

						// 限制返回的包数量
						if len(packages) >= 20 {
							return packages, nil
						}
					}
				}
			}
		}
	}

	// 检查Mix (Elixir/Erlang包管理器)
	if len(packages) < 20 && commandExists("mix") {
		_, err := executeCommandWithTimeout("mix", "local.hex", "--if-missing")
		if err == nil {
			output, err := executeCommandWithTimeout("mix", "deps")
			if err == nil {
				lines := strings.Split(output, "\n")
				for _, line := range lines {
					line = strings.TrimSpace(line)
					if line != "" && strings.HasPrefix(line, "* ") {
						parts := strings.Fields(line[2:]) // 跳过"* "前缀
						if len(parts) > 0 {
							name := parts[0]
							version := "已安装"

							// 尝试提取版本信息
							for i, part := range parts {
								if i > 0 && strings.HasPrefix(part, "v") && regexp.MustCompile(`v\d+\.\d+`).MatchString(part) {
									version = part
									break
								}
							}

							packages = append(packages, PackageInfo{
								Name:      name,
								Version:   version,
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
		}
	}

	return packages, nil
}

// 为Smalltalk添加包检测功能
func (a *App) detectSmalltalkWithPackages() LanguageInfo {
	info := a.detectSmalltalk()

	if info.Installed {
		// 获取已安装的Smalltalk包
		packages, _ := a.listSmalltalkPackages()
		info.Packages = packages
	}

	return info
}

// 列出Smalltalk包
func (a *App) listSmalltalkPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 检查Pharo Smalltalk
	if commandExists("pharo") || commandExists("pharo-vm") || commandExists("pharolauncher") {
		// 查找Pharo安装目录
		pharoPaths := []string{
			"/Applications/Pharo*.app",
			"C:\\Program Files\\Pharo*",
			filepath.Join(os.Getenv("HOME"), "Pharo"),
		}

		for _, pattern := range pharoPaths {
			matches, _ := filepath.Glob(pattern)
			for _, dir := range matches {
				// 尝试查找包信息
				packagesDir := filepath.Join(dir, "package-cache")
				if _, err := os.Stat(packagesDir); err == nil {
					// 读取包目录
					entries, err := os.ReadDir(packagesDir)
					if err == nil {
						for _, entry := range entries {
							name := entry.Name()
							// 过滤一些系统文件
							if !strings.HasPrefix(name, ".") && !strings.HasSuffix(name, ".zip") {
								version := "已安装"
								// 尝试从名称中提取版本
								versionRegex := regexp.MustCompile(`-v?(\d+\.\d+(\.\d+)?)`)
								matches := versionRegex.FindStringSubmatch(name)
								if len(matches) > 1 {
									version = matches[1]
									name = strings.TrimSuffix(name, matches[0])
								}

								packages = append(packages, PackageInfo{
									Name:      name,
									Version:   version,
									Installed: true,
								})
							}

							// 限制返回的包数量
							if len(packages) >= 20 {
								return packages, nil
							}
						}
					}
				}
			}
		}
	}

	// 检查Squeak Smalltalk
	squeakPaths := []string{
		"/Applications/Squeak*.app",
		"C:\\Program Files\\Squeak*",
		filepath.Join(os.Getenv("HOME"), ".squeak"),
	}

	for _, pattern := range squeakPaths {
		matches, _ := filepath.Glob(pattern)
		for _, dir := range matches {
			// 尝试查找包目录
			packagesFile := filepath.Join(dir, "package-cache")
			if _, err := os.Stat(packagesFile); err == nil {
				content, err := os.ReadFile(packagesFile)
				if err == nil {
					lines := strings.Split(string(content), "\n")
					for _, line := range lines {
						line = strings.TrimSpace(line)
						if line != "" && !strings.HasPrefix(line, "#") {
							parts := strings.Fields(line)
							if len(parts) > 0 {
								name := parts[0]
								version := "已安装"
								if len(parts) > 1 {
									version = parts[1]
								}

								packages = append(packages, PackageInfo{
									Name:      name,
									Version:   version,
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
			}
		}
	}

	return packages, nil
}

// 为Tcl添加包检测功能
func (a *App) detectTclWithPackages() LanguageInfo {
	info := a.detectTcl()

	if info.Installed {
		// 获取已安装的Tcl包
		packages, _ := a.listTclPackages()
		info.Packages = packages
	}

	return info
}

// 列出Tcl包
func (a *App) listTclPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 使用Tcl的package命令获取已安装的包
	if commandExists("tclsh") {
		// 创建一个临时的Tcl脚本来列出包
		tempDir, err := os.MkdirTemp("", "tcl-pkg")
		if err == nil {
			defer os.RemoveAll(tempDir)

			scriptPath := filepath.Join(tempDir, "list_packages.tcl")
			script := `
puts "TCL_PACKAGES_BEGIN"
foreach pkg [lsort [package names]] {
    if {[catch {package require $pkg} version]} {
        puts "$pkg <未知>"
    } else {
        puts "$pkg $version"
    }
}
puts "TCL_PACKAGES_END"
`
			err = os.WriteFile(scriptPath, []byte(script), 0644)
			if err == nil {
				output, err := executeCommandWithTimeout("tclsh", scriptPath)
				if err == nil {
					// 解析输出
					lines := strings.Split(output, "\n")
					inPackageSection := false
					for _, line := range lines {
						line = strings.TrimSpace(line)
						if line == "TCL_PACKAGES_BEGIN" {
							inPackageSection = true
							continue
						}
						if line == "TCL_PACKAGES_END" {
							inPackageSection = false
							continue
						}

						if inPackageSection && line != "" {
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
		}
	}

	// 检查Tcl库目录
	tclLibDirs := []string{
		"/usr/lib/tcl*",
		"/usr/local/lib/tcl*",
		"/usr/share/tcltk",
		"C:\\Program Files\\Tcl\\lib",
		"C:\\Tcl\\lib",
	}

	for _, pattern := range tclLibDirs {
		matches, _ := filepath.Glob(pattern)
		for _, dir := range matches {
			entries, err := os.ReadDir(dir)
			if err != nil {
				continue
			}

			for _, entry := range entries {
				if entry.IsDir() {
					// 检查pkgIndex.tcl文件
					pkgIndexPath := filepath.Join(dir, entry.Name(), "pkgIndex.tcl")
					if _, err := os.Stat(pkgIndexPath); err == nil {
						// 找到包，提取名称
						name := entry.Name()
						// 尝试提取版本
						version := "已安装"
						versionRegex := regexp.MustCompile(`(\d+\.\d+(\.\d+)?)`)
						matches := versionRegex.FindStringSubmatch(name)
						if len(matches) > 0 {
							version = matches[0]
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
							packages = append(packages, PackageInfo{
								Name:      name,
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
		}
	}

	return packages, nil
}

// 为Apex添加包检测功能
func (a *App) detectApexWithPackages() LanguageInfo {
	info := a.detectApex()

	if info.Installed {
		// 获取已安装的Apex包
		packages, _ := a.listApexPackages()
		info.Packages = packages
	}

	return info
}

// 列出Apex包
func (a *App) listApexPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 检查Salesforce CLI工具
	if commandExists("sfdx") {
		// 获取已安装的Salesforce CLI插件
		output, err := executeCommandWithTimeout("sfdx", "plugins", "--core")
		if err == nil {
			lines := strings.Split(output, "\n")
			for _, line := range lines {
				line = strings.TrimSpace(line)
				if line != "" && !strings.HasPrefix(line, "=") && !strings.Contains(line, "Usage:") && !strings.Contains(line, "CLI.") {
					parts := strings.Fields(line)
					if len(parts) >= 2 {
						packages = append(packages, PackageInfo{
							Name:      parts[0],
							Version:   parts[1],
							Installed: true,
						})
					} else if len(parts) == 1 && parts[0] != "Name" {
						packages = append(packages, PackageInfo{
							Name:      parts[0],
							Version:   "已安装",
							Installed: true,
						})
					}
				}
			}
		}

		// 获取已安装的Apex库
		output, err = executeCommandWithTimeout("sfdx", "force:apex:class:list", "--json")
		if err == nil && strings.Contains(output, "\"records\"") {
			// 解析JSON输出
			re := regexp.MustCompile(`"Name"\s*:\s*"([^"]+)"`)
			matches := re.FindAllStringSubmatch(output, -1)

			for _, match := range matches {
				if len(match) >= 2 {
					// 检查是否已经添加过
					exists := false
					for _, pkg := range packages {
						if pkg.Name == match[1] {
							exists = true
							break
						}
					}

					if !exists {
						packages = append(packages, PackageInfo{
							Name:      match[1],
							Version:   "Apex类",
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
	}

	return packages, nil
}

// 为Solidity添加包检测功能
func (a *App) detectSolidityWithPackages() LanguageInfo {
	info := a.detectSolidity()

	if info.Installed {
		// 获取已安装的Solidity包
		packages, _ := a.listSolidityPackages()
		info.Packages = packages
	}

	return info
}

// 列出Solidity包
func (a *App) listSolidityPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 检查Truffle
	if commandExists("truffle") {
		// 检查当前目录是否包含Truffle项目
		if _, err := os.Stat("truffle-config.js"); err == nil {
			// 获取已安装的包
			output, err := executeCommandWithTimeout("truffle", "version")
			if err == nil {
				// 提取Truffle版本信息
				lines := strings.Split(output, "\n")
				for _, line := range lines {
					line = strings.TrimSpace(line)
					if line != "" {
						parts := strings.Split(line, ":")
						if len(parts) >= 2 {
							name := strings.TrimSpace(parts[0])
							version := strings.TrimSpace(parts[1])

							packages = append(packages, PackageInfo{
								Name:      name,
								Version:   version,
								Installed: true,
							})
						}
					}
				}
			}
		}
	}

	// 检查npm中与Solidity相关的包
	if commandExists("npm") {
		output, err := executeCommandWithTimeout("npm", "list", "--global", "--depth=0")
		if err == nil {
			lines := strings.Split(output, "\n")
			for _, line := range lines {
				line = strings.TrimSpace(line)
				// 过滤Solidity相关的包
				if strings.Contains(line, "solidity") ||
					strings.Contains(line, "ethereum") ||
					strings.Contains(line, "web3") ||
					strings.Contains(line, "truffle") ||
					strings.Contains(line, "hardhat") ||
					strings.Contains(line, "ganache") {

					// 提取包名和版本
					re := regexp.MustCompile(`([a-zA-Z0-9\-_@\/]+)@(\d+\.\d+\.\d+)`)
					match := re.FindStringSubmatch(line)
					if len(match) >= 3 {
						packages = append(packages, PackageInfo{
							Name:      match[1],
							Version:   match[2],
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
	}

	// 检查OpenZeppelin合约
	if len(packages) < 20 {
		nodeModulesDir := "node_modules"
		ozDir := filepath.Join(nodeModulesDir, "@openzeppelin")
		if _, err := os.Stat(ozDir); err == nil {
			entries, err := os.ReadDir(ozDir)
			if err == nil {
				for _, entry := range entries {
					if entry.IsDir() {
						// 尝试从package.json获取版本
						packageJson := filepath.Join(ozDir, entry.Name(), "package.json")
						if _, err := os.Stat(packageJson); err == nil {
							content, err := os.ReadFile(packageJson)
							if err == nil {
								versionRegex := regexp.MustCompile(`"version"\s*:\s*"([^"]+)"`)
								match := versionRegex.FindSubmatch(content)
								version := "已安装"
								if len(match) >= 2 {
									version = string(match[1])
								}

								packages = append(packages, PackageInfo{
									Name:      "@openzeppelin/" + entry.Name(),
									Version:   version,
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
			}
		}
	}

	return packages, nil
}

// 为WebAssembly添加包检测功能
func (a *App) detectWebAssemblyWithPackages() LanguageInfo {
	info := a.detectWebAssembly()

	if info.Installed {
		// 获取已安装的WebAssembly工具
		packages, _ := a.listWebAssemblyTools()
		info.Packages = packages
	}

	return info
}
