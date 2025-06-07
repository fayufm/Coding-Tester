package main

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// 列出Julia包
func (a *App) listJuliaPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 检查julia命令是否存在
	if !commandExists("julia") {
		return packages, nil
	}

	// 创建临时Julia脚本
	script := `
	using Pkg
	pkgs = Pkg.installed()
	for (name, ver) in pkgs
		println("$name|$ver")
	end
	`

	tempFile, err := os.CreateTemp("", "julia-pkg-list-*.jl")
	if err != nil {
		return packages, err
	}
	defer os.Remove(tempFile.Name())

	if _, err := tempFile.WriteString(script); err != nil {
		return packages, err
	}
	tempFile.Close()

	// 运行脚本
	output, err := executeCommandWithTimeout("julia", tempFile.Name())
	if err != nil {
		// 尝试使用更简单的命令
		output, err = executeCommandWithTimeout("julia", "-e", "using Pkg; pkgs = Pkg.installed(); for (name, ver) in pkgs println(\"$name|$ver\") end")
		if err != nil {
			return packages, err
		}
	}

	// 解析输出
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Split(line, "|")
		if len(parts) >= 2 {
			packages = append(packages, PackageInfo{
				Name:      parts[0],
				Version:   parts[1],
				Installed: true,
			})
		}
	}

	return packages, nil
}

// 列出Nim包
func (a *App) listNimPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 检查nimble命令是否存在
	if !commandExists("nimble") {
		return packages, nil
	}

	// 获取已安装的包列表
	output, err := executeCommandWithTimeout("nimble", "list", "--installed")
	if err != nil {
		return packages, err
	}

	// 解析输出
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		// 跳过标题行和空行
		if line == "" || strings.HasPrefix(line, "Installed packages") || strings.Contains(line, "====") {
			continue
		}

		// 解析格式为 "packagename -> version" 的行
		if strings.Contains(line, "->") {
			parts := strings.Split(line, "->")
			if len(parts) >= 2 {
				name := strings.TrimSpace(parts[0])
				version := strings.TrimSpace(parts[1])

				// 提取描述（如果有）
				description := ""
				if strings.Contains(version, "#") {
					descParts := strings.Split(version, "#")
					if len(descParts) >= 2 {
						version = strings.TrimSpace(descParts[0])
						description = strings.TrimSpace(descParts[1])
					}
				}

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

// 列出Crystal包
func (a *App) listCrystalPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 检查shards命令是否存在
	if !commandExists("shards") {
		return packages, nil
	}

	// 查找当前目录或项目目录中的shard.lock文件
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return packages, err
	}

	// 检查常见的Crystal项目目录
	crystalPaths := []string{
		".", // 当前目录
		filepath.Join(homeDir, ".crystal"),
		filepath.Join(homeDir, "crystal"),
		filepath.Join(homeDir, "projects", "crystal"),
	}

	for _, path := range crystalPaths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			continue
		}

		// 尝试读取shard.lock文件
		shardLockFile := filepath.Join(path, "shard.lock")
		if _, err := os.Stat(shardLockFile); err == nil {
			content, err := os.ReadFile(shardLockFile)
			if err != nil {
				continue
			}

			// 解析YAML格式的shard.lock
			// 由于YAML解析需要额外依赖，这里使用简单的正则表达式解析
			re := regexp.MustCompile(`(?m)^  (\S+):.*\s+version: ([\d\.]+)`)
			matches := re.FindAllStringSubmatch(string(content), -1)

			for _, match := range matches {
				if len(match) >= 3 {
					packages = append(packages, PackageInfo{
						Name:      match[1],
						Version:   match[2],
						Installed: true,
					})
				}
			}

			if len(packages) > 0 {
				return packages, nil
			}
		}

		// 尝试读取shard.yml文件
		shardYmlFile := filepath.Join(path, "shard.yml")
		if _, err := os.Stat(shardYmlFile); err == nil {
			content, err := os.ReadFile(shardYmlFile)
			if err != nil {
				continue
			}

			// 解析YAML格式的shard.yml
			re := regexp.MustCompile(`(?m)^dependencies:\s*\n([\s\S]*?)(\n\n|\z)`)
			match := re.FindStringSubmatch(string(content))

			if len(match) >= 2 {
				depsSection := match[1]
				depRe := regexp.MustCompile(`(?m)^\s+(\S+):\s*\n\s+github:\s+[\S/]+\s*\n\s+version:\s+([\d\.]+)`)
				depMatches := depRe.FindAllStringSubmatch(depsSection, -1)

				for _, depMatch := range depMatches {
					if len(depMatch) >= 3 {
						packages = append(packages, PackageInfo{
							Name:      depMatch[1],
							Version:   depMatch[2],
							Installed: true,
						})
					}
				}
			}

			if len(packages) > 0 {
				return packages, nil
			}
		}
	}

	// 如果找不到文件，尝试使用shards命令
	output, err := executeCommandWithTimeout("shards", "list")
	if err != nil {
		// 最后一次尝试：使用shards check命令
		output, err = executeCommandWithTimeout("shards", "check")
		if err != nil {
			return packages, err
		}

		// 从check输出中提取依赖
		re := regexp.MustCompile(`(?m)Resolving\s+(\S+)\s+.*\((\S+)\)`)
		matches := re.FindAllStringSubmatch(output, -1)

		for _, match := range matches {
			if len(match) >= 3 {
				packages = append(packages, PackageInfo{
					Name:      match[1],
					Version:   match[2],
					Installed: true,
				})
			}
		}
	} else {
		// 解析shards list输出
		lines := strings.Split(output, "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" || strings.HasPrefix(line, "Shards installed:") {
				continue
			}

			// 提取包名和版本
			re := regexp.MustCompile(`(\S+)\s+\((\S+)\)`)
			match := re.FindStringSubmatch(line)
			if len(match) >= 3 {
				packages = append(packages, PackageInfo{
					Name:      match[1],
					Version:   match[2],
					Installed: true,
				})
			}
		}
	}

	return packages, nil
}

// 列出Zig包
func (a *App) listZigPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 检查zig命令是否存在
	if !commandExists("zig") {
		return packages, nil
	}

	// 检查常见的Zig包管理器
	zigPackageManagers := []string{"zigmod", "gyro", "zpm"}

	for _, manager := range zigPackageManagers {
		if !commandExists(manager) {
			continue
		}

		var output string
		var err error

		// 根据不同的包管理器使用不同的命令
		switch manager {
		case "zigmod":
			output, err = executeCommandWithTimeout("zigmod", "sum")
		case "gyro":
			output, err = executeCommandWithTimeout("gyro", "list")
		case "zpm":
			output, err = executeCommandWithTimeout("zpm", "list")
		}

		if err != nil {
			continue
		}

		// 解析输出
		lines := strings.Split(output, "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}

			// 尝试提取包名和版本
			// 不同的包管理器有不同的输出格式，这里使用通用的正则表达式
			re := regexp.MustCompile(`(\S+)(?:@|\s+)(\S+)`)
			match := re.FindStringSubmatch(line)
			if len(match) >= 3 {
				packages = append(packages, PackageInfo{
					Name:      match[1],
					Version:   match[2],
					Installed: true,
				})
			}
		}

		if len(packages) > 0 {
			return packages, nil
		}
	}

	// 查找项目中的build.zig文件分析依赖
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return packages, err
	}

	// 检查常见的Zig项目目录
	zigPaths := []string{
		".", // 当前目录
		filepath.Join(homeDir, "zig"),
		filepath.Join(homeDir, "projects", "zig"),
	}

	for _, path := range zigPaths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			continue
		}

		// 查找build.zig文件
		err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// 如果找到build.zig文件
			if !info.IsDir() && info.Name() == "build.zig" {
				content, err := os.ReadFile(path)
				if err != nil {
					return nil
				}

				// 分析build.zig中的依赖
				// 这是一个简化版，实际上需要更复杂的解析
				re := regexp.MustCompile(`(?m)\.addPackage\(\s*\.{[^}]*name\s*=\s*"([^"]+)"[^}]*,`)
				matches := re.FindAllStringSubmatch(string(content), -1)

				for _, match := range matches {
					if len(match) >= 2 {
						packages = append(packages, PackageInfo{
							Name:      match[1],
							Version:   "unknown", // build.zig通常不包含版本信息
							Installed: true,
						})
					}
				}
			}

			return nil
		})

		if err == nil && len(packages) > 0 {
			return packages, nil
		}
	}

	// 检查是否安装了zls（Zig语言服务器）
	if commandExists("zls") {
		output, err := executeCommandWithTimeout("zls", "--version")
		if err == nil {
			packages = append(packages, PackageInfo{
				Name:        "zls",
				Description: "Zig语言服务器",
				Version:     strings.TrimSpace(output),
				Installed:   true,
			})
		}
	}

	return packages, nil
}

// 列出D语言包
func (a *App) listDPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 检查dub命令是否存在
	if !commandExists("dub") {
		return packages, nil
	}

	// 获取已安装的包列表
	output, err := executeCommandWithTimeout("dub", "list", "--all")
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

		// 解析格式为 "packagename 1.2.3: description" 的行
		parts := strings.SplitN(line, " ", 2)
		if len(parts) >= 2 {
			name := parts[0]
			rest := parts[1]

			version := "unknown"
			description := ""

			// 提取版本号和描述
			verMatch := regexp.MustCompile(`^([\d\.-]+):\s*(.*)$`).FindStringSubmatch(rest)
			if len(verMatch) >= 3 {
				version = verMatch[1]
				description = verMatch[2]
			}

			packages = append(packages, PackageInfo{
				Name:        name,
				Version:     version,
				Description: description,
				Installed:   true,
			})
		}
	}

	return packages, nil
}

// 修改现有检测函数，添加包检测支持
func (a *App) detectJuliaWithPackages() LanguageInfo {
	info := a.detectJulia()

	if info.Installed {
		// 获取已安装的Julia包
		packages, _ := a.listJuliaPackages()
		info.Packages = packages
	}

	return info
}

func (a *App) detectNimWithPackages() LanguageInfo {
	info := a.detectNim()

	if info.Installed {
		// 获取已安装的Nim包
		packages, _ := a.listNimPackages()
		info.Packages = packages
	}

	return info
}

func (a *App) detectCrystalWithPackages() LanguageInfo {
	info := a.detectCrystal()

	if info.Installed {
		// 获取已安装的Crystal包
		packages, _ := a.listCrystalPackages()
		info.Packages = packages
	}

	return info
}

func (a *App) detectZigWithPackages() LanguageInfo {
	info := a.detectZig()

	if info.Installed {
		// 获取已安装的Zig包
		packages, _ := a.listZigPackages()
		info.Packages = packages
	}

	return info
}

func (a *App) detectDWithPackages() LanguageInfo {
	info := a.detectD()

	if info.Installed {
		// 获取已安装的D包
		packages, _ := a.listDPackages()
		info.Packages = packages
	}

	return info
}
