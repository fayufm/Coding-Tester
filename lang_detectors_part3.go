package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

// 检测Objective-C
func (a *App) detectObjectiveC() LanguageInfo {
	info := LanguageInfo{
		Name:            "Objective-C",
		Installed:       false,
		DownloadURL:     "https://developer.apple.com/xcode/",
		InstallTutorial: "https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ProgrammingWithObjectiveC/Introduction/Introduction.html",
		PackageManager:  "CocoaPods",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "AFNetworking",
				Description: "Objective-C的网络库",
				Version:     "4.0.1",
			},
			{
				Name:        "SDWebImage",
				Description: "异步图像下载器，带有缓存支持",
				Version:     "5.16.0",
			},
			{
				Name:        "MBProgressHUD",
				Description: "为iOS应用程序添加闪亮、漂亮的HUD",
				Version:     "1.2.0",
			},
			{
				Name:        "Realm",
				Description: "移动数据库：SQLite和Core Data的替代品",
				Version:     "10.41.0",
			},
			{
				Name:        "Masonry",
				Description: "Objective-C的自动布局DSL",
				Version:     "1.1.0",
			},
		},
	}

	// 检测clang (macOS)
	if commandExists("clang") {
		output, err := executeCommandWithTimeout("clang", "--version")
		if err == nil && strings.Contains(output, "Apple") {
			info.Installed = true
			info.Version = "Xcode Clang: " + strings.Split(output, "\n")[0]

			// 检测已安装的CocoaPods包
			packages, _ := a.listObjectiveCPackages()
			info.Packages = packages
		}
	}

	return info
}

// 检测Groovy
func (a *App) detectGroovy() LanguageInfo {
	info := LanguageInfo{
		Name:            "Groovy",
		Installed:       false,
		DownloadURL:     "https://groovy.apache.org/download.html",
		InstallTutorial: "https://groovy.apache.org/documentation.html",
		PackageManager:  "Grape",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "spock-core",
				Description: "Spock测试框架，用于Java和Groovy应用程序",
				Version:     "2.3-groovy-4.0",
			},
			{
				Name:        "groovy-json",
				Description: "Groovy的JSON支持",
				Version:     "4.0.12",
			},
			{
				Name:        "groovy-xml",
				Description: "Groovy的XML支持",
				Version:     "4.0.12",
			},
			{
				Name:        "http-builder-ng-core",
				Description: "Groovy的HTTP客户端",
				Version:     "1.0.4",
			},
			{
				Name:        "ratpack-groovy",
				Description: "用于构建高性能Web应用程序的工具包",
				Version:     "1.9.0",
			},
		},
	}

	if !commandExists("groovy") {
		return info
	}

	output, err := executeCommandWithTimeout("groovy", "--version")
	if err == nil {
		info.Installed = true
		info.Version = output
	}

	return info
}

// 检测Clojure
func (a *App) detectClojure() LanguageInfo {
	info := LanguageInfo{
		Name:            "Clojure",
		Installed:       false,
		DownloadURL:     "https://clojure.org/guides/getting_started",
		InstallTutorial: "https://clojure.org/guides/getting_started",
		PackageManager:  "Leiningen/deps.edn",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "ring",
				Description: "Clojure的Web应用库",
				Version:     "1.10.0",
			},
			{
				Name:        "compojure",
				Description: "Clojure的路由库",
				Version:     "1.7.0",
			},
			{
				Name:        "hiccup",
				Description: "用于表示HTML的库",
				Version:     "2.0.0-alpha2",
			},
			{
				Name:        "clj-http",
				Description: "Clojure的HTTP客户端",
				Version:     "3.12.3",
			},
			{
				Name:        "core.async",
				Description: "Clojure的异步编程库",
				Version:     "1.6.673",
			},
		},
	}

	if !commandExists("clojure") {
		return info
	}

	output, err := executeCommandWithTimeout("clojure", "--version")
	if err == nil {
		info.Installed = true
		info.Version = output

		// 检查Leiningen是否安装
		if commandExists("lein") {
			// 列出Clojure包
			packages, _ := a.listClojurePackages()
			info.Packages = packages
		}
	}

	return info
}

// 列出Clojure包
func (a *App) listClojurePackages() ([]PackageInfo, error) {
	packages := []PackageInfo{}

	// 尝试从Leiningen项目中获取依赖信息
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	// 检查Leiningen的profiles.clj文件
	profilesPath := filepath.Join(homeDir, ".lein", "profiles.clj")
	if _, err := os.Stat(profilesPath); err == nil {
		// 读取profiles.clj文件内容
		content, err := os.ReadFile(profilesPath)
		if err == nil {
			contentStr := string(content)

			// 尝试提取依赖信息（简单解析，不是完整的Clojure解析器）
			// 查找类似 [org.clojure/clojure "1.11.1"] 的模式
			re := regexp.MustCompile(`\[([a-zA-Z0-9\.\-\_\/]+)\s+"([^"]+)"\]`)
			matches := re.FindAllStringSubmatch(contentStr, -1)

			for _, match := range matches {
				if len(match) >= 3 {
					packages = append(packages, PackageInfo{
						Name:      match[1],
						Version:   match[2],
						Installed: true,
					})
				}
			}
		}
	}

	// 检查deps.edn文件
	depsPath := filepath.Join(homeDir, ".clojure", "deps.edn")
	if _, err := os.Stat(depsPath); err == nil {
		// 读取deps.edn文件内容
		content, err := os.ReadFile(depsPath)
		if err == nil {
			contentStr := string(content)

			// 尝试提取依赖信息
			re := regexp.MustCompile(`([a-zA-Z0-9\.\-\_\/]+)\s+\{:mvn/version\s+"([^"]+)"\}`)
			matches := re.FindAllStringSubmatch(contentStr, -1)

			for _, match := range matches {
				if len(match) >= 3 {
					packages = append(packages, PackageInfo{
						Name:      match[1],
						Version:   match[2],
						Installed: true,
					})
				}
			}
		}
	}

	return packages, nil
}

// 检测Elixir
func (a *App) detectElixir() LanguageInfo {
	info := LanguageInfo{
		Name:            "Elixir",
		Installed:       false,
		DownloadURL:     "https://elixir-lang.org/install.html",
		InstallTutorial: "https://elixir-lang.org/getting-started/introduction.html",
		PackageManager:  "Hex/Mix",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "phoenix",
				Description: "Elixir的Web框架",
				Version:     "1.7.2",
			},
			{
				Name:        "ecto",
				Description: "数据库包装器和语言集成查询",
				Version:     "3.10.1",
			},
			{
				Name:        "absinthe",
				Description: "Elixir的GraphQL工具包",
				Version:     "1.7.1",
			},
			{
				Name:        "tesla",
				Description: "Elixir的HTTP客户端",
				Version:     "1.7.0",
			},
			{
				Name:        "ex_machina",
				Description: "Elixir的工厂库",
				Version:     "2.7.0",
			},
		},
	}

	if !commandExists("elixir") {
		return info
	}

	output, err := executeCommandWithTimeout("elixir", "--version")
	if err == nil {
		info.Installed = true
		info.Version = output

		// 检查mix是否安装
		if commandExists("mix") {
			// 列出Elixir包
			packages, _ := a.listElixirPackages()
			info.Packages = packages
		}
	}

	return info
}

// 列出Elixir包
func (a *App) listElixirPackages() ([]PackageInfo, error) {
	packages := []PackageInfo{}

	// 使用mix deps命令获取当前项目的依赖
	output, err := executeCommandWithTimeout("mix", "deps")
	if err == nil && output != "" {
		lines := strings.Split(output, "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}

			// 解析类似 "* phoenix 1.7.2" 的行
			re := regexp.MustCompile(`\* ([a-zA-Z0-9\_\-]+) \(([^\)]+)\)`)
			match := re.FindStringSubmatch(line)
			if len(match) >= 3 {
				packages = append(packages, PackageInfo{
					Name:      match[1],
					Version:   match[2],
					Installed: true,
				})
				continue
			}

			// 尝试另一种格式 "* phoenix 1.7.2"
			re = regexp.MustCompile(`\* ([a-zA-Z0-9\_\-]+) ([0-9\.]+)`)
			match = re.FindStringSubmatch(line)
			if len(match) >= 3 {
				packages = append(packages, PackageInfo{
					Name:      match[1],
					Version:   match[2],
					Installed: true,
				})
			}
		}
	}

	// 如果没有找到依赖，尝试使用hex list命令
	if len(packages) == 0 && commandExists("hex") {
		output, err := executeCommandWithTimeout("hex", "list")
		if err == nil && output != "" {
			lines := strings.Split(output, "\n")
			for _, line := range lines {
				line = strings.TrimSpace(line)
				if line == "" || strings.HasPrefix(line, "Name") {
					continue
				}

				parts := strings.Fields(line)
				if len(parts) >= 2 {
					packages = append(packages, PackageInfo{
						Name:      parts[0],
						Version:   parts[1],
						Installed: true,
					})
				}
			}
		}
	}

	return packages, nil
}

// 检测F#
func (a *App) detectFSharp() LanguageInfo {
	info := LanguageInfo{
		Name:            "F#",
		Installed:       false,
		DownloadURL:     "https://fsharp.org/use/",
		InstallTutorial: "https://docs.microsoft.com/en-us/dotnet/fsharp/get-started/",
		PackageManager:  "NuGet",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "FSharp.Core",
				Description: "F#核心库",
				Version:     "7.0.300",
			},
			{
				Name:        "Suave",
				Description: "F#的Web服务器库",
				Version:     "2.6.2",
			},
			{
				Name:        "FSharp.Data",
				Description: "F#的数据访问库",
				Version:     "6.2.0",
			},
			{
				Name:        "Fable",
				Description: "F#到JavaScript的编译器",
				Version:     "4.1.4",
			},
			{
				Name:        "Elmish",
				Description: "F#的UI库",
				Version:     "4.0.1",
			},
		},
	}

	if !commandExists("dotnet") {
		return info
	}

	output, err := executeCommandWithTimeout("dotnet", "fsi", "--version")
	if err == nil {
		info.Installed = true
		info.Version = output

		// 检测已安装的F#包
		packages, _ := a.listFSharpPackages()
		info.Packages = packages
	}

	return info
}

// 列出已安装的F#包
func (a *App) listFSharpPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 使用dotnet命令检查全局安装的包
	output, err := executeCommandWithTimeout("dotnet", "tool", "list", "--global")
	if err != nil {
		return packages, err
	}

	lines := strings.Split(output, "\n")
	for i, line := range lines {
		// 跳过标题行
		if i == 0 || strings.TrimSpace(line) == "" {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) >= 2 {
			// 检查是否是F#相关的包
			if strings.Contains(strings.ToLower(fields[0]), "fsharp") ||
				strings.Contains(strings.ToLower(fields[0]), "fs") ||
				strings.Contains(strings.ToLower(fields[0]), "fantomas") {
				packages = append(packages, PackageInfo{
					Name:      fields[0],
					Version:   fields[1],
					Installed: true,
				})
			}
		}
	}

	// 尝试查找用户项目中的F#包
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return packages, nil
	}

	// 检查常见的项目目录
	projectDirs := []string{
		filepath.Join(homeDir, "source", "repos"),
		filepath.Join(homeDir, "Projects"),
		filepath.Join(homeDir, "Documents", "Projects"),
	}

	for _, dir := range projectDirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			continue
		}

		// 查找F#项目文件
		err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}

			// 查找.fsproj文件
			if !info.IsDir() && strings.HasSuffix(path, ".fsproj") {
				// 读取项目文件
				content, err := os.ReadFile(path)
				if err != nil {
					return nil
				}

				// 提取PackageReference标签
				re := regexp.MustCompile(`<PackageReference\s+Include="([^"]+)"\s+Version="([^"]+)"`)
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

				// 限制项目数量
				if len(packages) >= 20 {
					return filepath.SkipDir
				}
			}
			return nil
		})
	}

	return packages, nil
}

// 检测Julia
func (a *App) detectJulia() LanguageInfo {
	info := LanguageInfo{
		Name:            "Julia",
		Installed:       false,
		DownloadURL:     "https://julialang.org/downloads/",
		InstallTutorial: "https://julialang.org/learning/",
		PackageManager:  "Pkg",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "DataFrames",
				Description: "用于处理表格数据的包",
				Version:     "1.5.0",
			},
			{
				Name:        "Plots",
				Description: "强大的可视化库",
				Version:     "1.38.12",
			},
			{
				Name:        "Flux",
				Description: "Julia的机器学习库",
				Version:     "0.14.4",
			},
			{
				Name:        "DifferentialEquations",
				Description: "微分方程求解器",
				Version:     "7.7.0",
			},
			{
				Name:        "JuMP",
				Description: "数学优化的领域特定语言",
				Version:     "1.11.1",
			},
		},
	}

	if !commandExists("julia") {
		return info
	}

	output, err := executeCommandWithTimeout("julia", "--version")
	if err == nil {
		info.Installed = true
		info.Version = output
	}

	return info
}

// 检测Prolog
func (a *App) detectProlog() LanguageInfo {
	info := LanguageInfo{
		Name:            "Prolog",
		Installed:       false,
		DownloadURL:     "https://www.swi-prolog.org/download/stable",
		InstallTutorial: "https://www.swi-prolog.org/pldoc/man?section=quickstart",
		PackageManager:  "SWI-Prolog Package Manager",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "http",
				Description: "HTTP客户端和服务器库",
				Version:     "latest",
			},
			{
				Name:        "clib",
				Description: "与C语言接口的库",
				Version:     "latest",
			},
			{
				Name:        "sgml",
				Description: "处理SGML、XML和HTML的库",
				Version:     "latest",
			},
			{
				Name:        "pce",
				Description: "ProLog的图形用户界面",
				Version:     "latest",
			},
			{
				Name:        "nlp",
				Description: "自然语言处理库",
				Version:     "latest",
			},
		},
	}

	if !commandExists("swipl") {
		return info
	}

	output, err := executeCommandWithTimeout("swipl", "--version")
	if err == nil {
		info.Installed = true
		info.Version = output

		// 获取已安装的Prolog包
		packages, _ := a.listPrologPackages()
		info.Packages = packages
	}

	return info
}

// 列出已安装的Prolog包
func (a *App) listPrologPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 创建一个临时Prolog脚本来列出已安装的包
	script := `
:- use_module(library(prolog_pack)).
:- initialization(main, main).

main(_) :-
    findall(pack(Name, Version), pack_info(Name, version(Version)), Packs),
    forall(member(pack(Name, Version), Packs),
           format('~w|~w~n', [Name, Version])),
    halt.
`
	tempFile, err := os.CreateTemp("", "prolog-pkg-list-*.pl")
	if err != nil {
		return packages, err
	}
	defer os.Remove(tempFile.Name())

	if _, err := tempFile.WriteString(script); err != nil {
		return packages, err
	}
	tempFile.Close()

	// 运行脚本
	output, err := executeCommandWithTimeout("swipl", "-q", "-f", tempFile.Name())
	if err != nil {
		// 尝试另一种方法：检查包目录
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return packages, err
		}

		// 检查SWI-Prolog的包目录
		packDirs := []string{
			filepath.Join(homeDir, ".local", "share", "swi-prolog", "pack"),
			filepath.Join(homeDir, ".swi-prolog", "pack"),
			"C:/Program Files/swipl/pack", // Windows
		}

		for _, dir := range packDirs {
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				continue
			}

			// 读取目录内容
			files, err := os.ReadDir(dir)
			if err != nil {
				continue
			}

			for _, file := range files {
				if file.IsDir() {
					// 尝试读取pack.pl文件获取版本信息
					packFile := filepath.Join(dir, file.Name(), "pack.pl")
					version := "latest"

					if content, err := os.ReadFile(packFile); err == nil {
						// 尝试提取版本信息
						re := regexp.MustCompile(`version\(['"](.*)['"]\)`)
						match := re.FindStringSubmatch(string(content))
						if len(match) >= 2 {
							version = match[1]
						}
					}

					packages = append(packages, PackageInfo{
						Name:      file.Name(),
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

		return packages, nil
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

// 检测Assembly
func (a *App) detectAssembly() LanguageInfo {
	info := LanguageInfo{
		Name:            "Assembly",
		Installed:       false,
		DownloadURL:     "https://www.nasm.us/",
		InstallTutorial: "https://www.nasm.us/doc/nasmdoc1.html",
		PackageManager:  "N/A",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "NASM",
				Description: "Netwide汇编器",
				Version:     "2.16.01",
			},
			{
				Name:        "MASM",
				Description: "Microsoft宏汇编器",
				Version:     "latest",
			},
			{
				Name:        "FASM",
				Description: "平面汇编器",
				Version:     "1.73.30",
			},
			{
				Name:        "YASM",
				Description: "NASM的重写版本",
				Version:     "1.3.0",
			},
			{
				Name:        "GAS",
				Description: "GNU汇编器",
				Version:     "latest",
			},
		},
	}

	// 检测NASM
	if commandExists("nasm") {
		output, err := executeCommandWithTimeout("nasm", "-v")
		if err == nil {
			info.Installed = true
			info.Version = "NASM: " + output

			// 获取已安装的Assembly相关工具
			packages, _ := a.listAssemblyTools()
			info.Packages = packages

			return info
		}
	}

	// 检测GAS
	if commandExists("as") {
		output, err := executeCommandWithTimeout("as", "--version")
		if err == nil {
			info.Installed = true
			info.Version = "GAS: " + output

			// 获取已安装的Assembly相关工具
			packages, _ := a.listAssemblyTools()
			info.Packages = packages

			return info
		}
	}

	// 检测FASM
	if commandExists("fasm") {
		_, err := executeCommandWithTimeout("fasm")
		if err == nil {
			info.Installed = true
			info.Version = "FASM: 已安装"

			// 获取已安装的Assembly相关工具
			packages, _ := a.listAssemblyTools()
			info.Packages = packages

			return info
		}
	}

	// 检测YASM
	if commandExists("yasm") {
		output, err := executeCommandWithTimeout("yasm", "--version")
		if err == nil {
			info.Installed = true
			info.Version = "YASM: " + output

			// 获取已安装的Assembly相关工具
			packages, _ := a.listAssemblyTools()
			info.Packages = packages

			return info
		}
	}

	return info
}

// 列出已安装的Assembly相关工具
func (a *App) listAssemblyTools() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 检测常见的汇编器和相关工具
	assemblyTools := []struct {
		command     string
		versionArgs []string
		name        string
	}{
		{"nasm", []string{"-v"}, "NASM (Netwide Assembler)"},
		{"as", []string{"--version"}, "GAS (GNU Assembler)"},
		{"fasm", []string{}, "FASM (Flat Assembler)"},
		{"yasm", []string{"--version"}, "YASM"},
		{"ml", []string{}, "MASM (Microsoft Macro Assembler)"},
		{"ml64", []string{}, "MASM64 (Microsoft Macro Assembler 64-bit)"},
		{"ld", []string{"--version"}, "LD (GNU Linker)"},
		{"objdump", []string{"--version"}, "objdump (显示目标文件信息)"},
		{"objcopy", []string{"--version"}, "objcopy (复制和转换目标文件)"},
		{"gdb", []string{"--version"}, "GDB (GNU调试器)"},
		{"ndisasm", []string{"-v"}, "ndisasm (NASM反汇编器)"},
	}

	for _, tool := range assemblyTools {
		if !commandExists(tool.command) {
			continue
		}

		var version string
		if len(tool.versionArgs) > 0 {
			output, err := executeCommandWithTimeout(tool.command, tool.versionArgs...)
			if err == nil {
				// 提取第一行作为版本信息
				lines := strings.Split(output, "\n")
				if len(lines) > 0 {
					version = lines[0]
				} else {
					version = "已安装"
				}
			} else {
				version = "已安装"
			}
		} else {
			version = "已安装"
		}

		packages = append(packages, PackageInfo{
			Name:      tool.name,
			Version:   version,
			Installed: true,
		})
	}

	// 检测常见的汇编库
	homeDir, err := os.UserHomeDir()
	if err == nil {
		// 检查常见的汇编库目录
		asmLibDirs := []string{
			filepath.Join(homeDir, "asm"),
			filepath.Join(homeDir, "assembly"),
			filepath.Join(homeDir, "nasm"),
			"C:/NASM", // Windows
		}

		for _, dir := range asmLibDirs {
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				continue
			}

			// 读取目录内容
			files, err := os.ReadDir(dir)
			if err != nil {
				continue
			}

			for _, file := range files {
				if !file.IsDir() && (strings.HasSuffix(file.Name(), ".inc") || strings.HasSuffix(file.Name(), ".mac")) {
					packages = append(packages, PackageInfo{
						Name:      file.Name(),
						Version:   "N/A",
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

// 检测COBOL
func (a *App) detectCOBOL() LanguageInfo {
	info := LanguageInfo{
		Name:            "COBOL",
		Installed:       false,
		DownloadURL:     "https://gnucobol.sourceforge.io/",
		InstallTutorial: "https://gnucobol.sourceforge.io/guides/GnuCOBOL%202.2%20NOV2017%20Programmers%20Guide.pdf",
		PackageManager:  "N/A",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "GnuCOBOL",
				Description: "开源COBOL编译器",
				Version:     "3.1.2",
			},
			{
				Name:        "OpenCOBOL",
				Description: "COBOL编译器",
				Version:     "latest",
			},
			{
				Name:        "COBOL-IT",
				Description: "企业级COBOL编译器",
				Version:     "latest",
			},
			{
				Name:        "Micro Focus COBOL",
				Description: "商业COBOL开发环境",
				Version:     "latest",
			},
			{
				Name:        "IBM COBOL",
				Description: "IBM大型机COBOL",
				Version:     "latest",
			},
		},
	}

	if !commandExists("cobc") {
		return info
	}

	output, err := executeCommandWithTimeout("cobc", "--version")
	if err == nil {
		info.Installed = true
		info.Version = output
	}

	return info
}

// 检测Fortran
func (a *App) detectFortran() LanguageInfo {
	info := LanguageInfo{
		Name:            "Fortran",
		Installed:       false,
		DownloadURL:     "https://gcc.gnu.org/fortran/",
		InstallTutorial: "https://gcc.gnu.org/wiki/GFortran",
		PackageManager:  "N/A",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "LAPACK",
				Description: "线性代数包",
				Version:     "3.11.0",
			},
			{
				Name:        "BLAS",
				Description: "基本线性代数子程序",
				Version:     "latest",
			},
			{
				Name:        "NetCDF",
				Description: "网络通用数据格式",
				Version:     "4.9.2",
			},
			{
				Name:        "HDF5",
				Description: "分层数据格式",
				Version:     "1.14.1",
			},
			{
				Name:        "MPI",
				Description: "消息传递接口",
				Version:     "latest",
			},
		},
	}

	if !commandExists("gfortran") {
		return info
	}

	output, err := executeCommandWithTimeout("gfortran", "--version")
	if err == nil {
		info.Installed = true
		info.Version = output
	}

	return info
}

// 检测Delphi/Pascal
func (a *App) detectDelphi() LanguageInfo {
	info := LanguageInfo{
		Name:            "Delphi/Pascal",
		Installed:       false,
		DownloadURL:     "https://www.embarcadero.com/products/delphi/starter",
		InstallTutorial: "https://docwiki.embarcadero.com/RADStudio/Alexandria/en/Installation_Notes",
		PackageManager:  "GetIt",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "VCL",
				Description: "Visual Component Library",
				Version:     "latest",
			},
			{
				Name:        "FMX",
				Description: "FireMonkey跨平台框架",
				Version:     "latest",
			},
			{
				Name:        "REST Debugger",
				Description: "RESTful服务调试工具",
				Version:     "latest",
			},
			{
				Name:        "DUnitX",
				Description: "Delphi的单元测试框架",
				Version:     "latest",
			},
			{
				Name:        "Spring4D",
				Description: "Delphi的对象容器和依赖注入框架",
				Version:     "latest",
			},
		},
	}

	// 检测Free Pascal
	if !commandExists("fpc") {
		return info
	}

	output, err := executeCommandWithTimeout("fpc", "-i")
	if err == nil {
		info.Installed = true
		info.Version = "Free Pascal: " + output
	}

	return info
}

// 检测Lisp
func (a *App) detectLisp() LanguageInfo {
	info := LanguageInfo{
		Name:            "Lisp",
		Installed:       false,
		DownloadURL:     "https://common-lisp.net/downloads/",
		InstallTutorial: "https://lisp-lang.org/learn/getting-started/",
		PackageManager:  "Quicklisp",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "alexandria",
				Description: "公共工具库",
				Version:     "latest",
			},
			{
				Name:        "cl-ppcre",
				Description: "正则表达式库",
				Version:     "latest",
			},
			{
				Name:        "bordeaux-threads",
				Description: "可移植的线程库",
				Version:     "latest",
			},
			{
				Name:        "hunchentoot",
				Description: "Web服务器",
				Version:     "latest",
			},
			{
				Name:        "quicklisp",
				Description: "包管理系统",
				Version:     "latest",
			},
		},
	}

	// 检查SBCL
	if commandExists("sbcl") {
		output, err := executeCommandWithTimeout("sbcl", "--version")
		if err == nil {
			info.Installed = true
			info.Version = "SBCL: " + output

			// 获取已安装的Lisp包
			packages, _ := a.listLispPackages()
			info.Packages = packages
		}
	}

	// 如果SBCL不存在，检查CLISP
	if !info.Installed && commandExists("clisp") {
		output, err := executeCommandWithTimeout("clisp", "--version")
		if err == nil {
			info.Installed = true
			info.Version = "CLISP: " + output

			// 获取已安装的Lisp包
			packages, _ := a.listLispPackages()
			info.Packages = packages
		}
	}

	return info
}

// 列出已安装的Lisp包
func (a *App) listLispPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 创建一个临时Lisp脚本来列出Quicklisp包
	script := `
(handler-case
    (progn
      (ignore-errors (load "~/quicklisp/setup.lisp"))
      (when (find-package :ql)
        (let ((systems (funcall (find-symbol "SYSTEM-LIST" :ql))))
          (dolist (system systems)
            (format t "~a~%" (getf system :name))))))
  (error () nil))
(exit)
`
	tempFile, err := os.CreateTemp("", "lisp-pkg-list-*.lisp")
	if err != nil {
		return packages, err
	}
	defer os.Remove(tempFile.Name())

	if _, err := tempFile.WriteString(script); err != nil {
		return packages, err
	}
	tempFile.Close()

	// 尝试使用SBCL运行脚本
	var output string
	if commandExists("sbcl") {
		output, err = executeCommandWithTimeout("sbcl", "--script", tempFile.Name())
	} else if commandExists("clisp") {
		output, err = executeCommandWithTimeout("clisp", tempFile.Name())
	} else {
		return packages, nil
	}

	if err != nil {
		// 如果脚本执行失败，尝试查找Quicklisp目录
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return packages, err
		}

		quicklispDir := filepath.Join(homeDir, "quicklisp", "dists", "quicklisp", "software")
		if _, err := os.Stat(quicklispDir); os.IsNotExist(err) {
			return packages, nil
		}

		// 遍历Quicklisp目录
		files, err := os.ReadDir(quicklispDir)
		if err != nil {
			return packages, err
		}

		for _, file := range files {
			if file.IsDir() {
				packages = append(packages, PackageInfo{
					Name:      file.Name(),
					Version:   "latest",
					Installed: true,
				})

				// 限制数量
				if len(packages) >= 20 {
					break
				}
			}
		}

		return packages, nil
	}

	// 解析输出
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		packages = append(packages, PackageInfo{
			Name:      line,
			Version:   "latest", // Quicklisp通常不显示版本号
			Installed: true,
		})

		// 限制数量
		if len(packages) >= 20 {
			break
		}
	}

	return packages, nil
}

// 列出已安装的Objective-C包
func (a *App) listObjectiveCPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 检查CocoaPods是否安装
	if !commandExists("pod") {
		return packages, nil
	}

	// 使用CocoaPods列出已安装的包
	output, err := executeCommandWithTimeout("pod", "list")
	if err == nil {
		lines := strings.Split(output, "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" || strings.HasPrefix(line, "->") || strings.HasPrefix(line, "Using") {
				continue
			}

			// 尝试提取包名和版本
			re := regexp.MustCompile(`^(\S+)\s+\(([^)]+)\)`)
			match := re.FindStringSubmatch(line)
			if len(match) >= 3 {
				packages = append(packages, PackageInfo{
					Name:      match[1],
					Version:   match[2],
					Installed: true,
				})

				// 限制数量
				if len(packages) >= 20 {
					break
				}
			}
		}
	}

	// 如果没有找到包，尝试查找Podfile文件
	if len(packages) == 0 {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return packages, err
		}

		// 检查常见的项目目录
		projectDirs := []string{
			filepath.Join(homeDir, "Documents"),
			filepath.Join(homeDir, "Projects"),
			filepath.Join(homeDir, "Developer"),
		}

		for _, dir := range projectDirs {
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				continue
			}

			// 递归查找Podfile文件
			err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return nil
				}

				if !info.IsDir() && (info.Name() == "Podfile" || info.Name() == "Podfile.lock") {
					// 读取Podfile文件
					content, err := os.ReadFile(path)
					if err != nil {
						return nil
					}

					// 解析Podfile内容
					if info.Name() == "Podfile.lock" {
						// 尝试提取PODS部分
						podSection := false
						lines := strings.Split(string(content), "\n")
						for _, line := range lines {
							if strings.TrimSpace(line) == "PODS:" {
								podSection = true
								continue
							}

							if podSection {
								if strings.HasPrefix(line, "  - ") {
									// 格式通常是 "  - AFNetworking (4.0.1)"
									parts := strings.SplitN(strings.TrimPrefix(line, "  - "), " ", 2)
									if len(parts) >= 2 {
										name := parts[0]
										version := strings.Trim(parts[1], "()")

										packages = append(packages, PackageInfo{
											Name:      name,
											Version:   version,
											Installed: true,
										})

										// 限制数量
										if len(packages) >= 20 {
											return filepath.SkipDir
										}
									}
								} else if !strings.HasPrefix(line, " ") && line != "" {
									// 离开PODS部分
									podSection = false
								}
							}
						}
					} else {
						// 解析普通Podfile
						re := regexp.MustCompile(`pod\s+['"]([^'"]+)['"](?:\s*,\s*['"]([^'"]+)['"])?`)
						matches := re.FindAllStringSubmatch(string(content), -1)
						for _, match := range matches {
							if len(match) >= 2 {
								version := "latest"
								if len(match) >= 3 && match[2] != "" {
									version = match[2]
								}

								packages = append(packages, PackageInfo{
									Name:      match[1],
									Version:   version,
									Installed: true,
								})

								// 限制数量
								if len(packages) >= 20 {
									return filepath.SkipDir
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

// 检测Bash
func (a *App) detectBash() LanguageInfo {
	info := LanguageInfo{
		Name:            "Bash",
		Installed:       false,
		DownloadURL:     "https://www.gnu.org/software/bash/",
		InstallTutorial: "https://www.gnu.org/software/bash/manual/bash.html",
		PackageManager:  "N/A",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "coreutils",
				Description: "GNU核心工具集",
				Version:     "latest",
			},
			{
				Name:        "findutils",
				Description: "GNU查找工具",
				Version:     "latest",
			},
			{
				Name:        "grep",
				Description: "GNU grep搜索工具",
				Version:     "latest",
			},
			{
				Name:        "sed",
				Description: "GNU流编辑器",
				Version:     "latest",
			},
			{
				Name:        "awk",
				Description: "文本处理语言",
				Version:     "latest",
			},
		},
	}

	// 检查bash是否安装
	if !commandExists("bash") {
		return info
	}

	output, err := executeCommandWithTimeout("bash", "--version")
	if err == nil {
		info.Installed = true
		// 提取第一行作为版本信息
		lines := strings.Split(output, "\n")
		if len(lines) > 0 {
			info.Version = lines[0]
		} else {
			info.Version = "Bash (已安装)"
		}

		// 获取已安装的Bash相关工具
		packages, _ := a.listBashPackages()
		info.Packages = packages
	}

	return info
}

// 列出已安装的Bash相关工具
func (a *App) listBashPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 常见的Bash工具
	bashTools := []struct {
		command string
		name    string
	}{
		{"grep", "GNU Grep"},
		{"sed", "GNU Sed"},
		{"awk", "GNU Awk"},
		{"find", "GNU Findutils"},
		{"ls", "GNU Coreutils"},
		{"cat", "GNU Coreutils"},
		{"cut", "GNU Coreutils"},
		{"sort", "GNU Coreutils"},
		{"uniq", "GNU Coreutils"},
		{"tr", "GNU Coreutils"},
		{"wc", "GNU Coreutils"},
		{"xargs", "GNU Findutils"},
		{"curl", "cURL"},
		{"wget", "GNU Wget"},
		{"jq", "jq JSON处理器"},
		{"yq", "YAML处理器"},
		{"parallel", "GNU Parallel"},
		{"rsync", "Rsync"},
		{"ssh", "OpenSSH"},
		{"git", "Git"},
	}

	for _, tool := range bashTools {
		if !commandExists(tool.command) {
			continue
		}

		var version string

		// 根据不同的命令获取版本信息
		switch tool.command {
		case "ls", "cat", "cut", "sort", "uniq", "tr", "wc":
			// 对于coreutils工具，使用--version
			output, err := executeCommandWithTimeout(tool.command, "--version")
			if err == nil {
				lines := strings.Split(output, "\n")
				if len(lines) > 0 {
					version = lines[0]
				} else {
					version = "已安装"
				}
			} else {
				version = "已安装"
			}
		default:
			// 对于其他工具，使用--version
			output, err := executeCommandWithTimeout(tool.command, "--version")
			if err == nil {
				lines := strings.Split(output, "\n")
				if len(lines) > 0 {
					version = lines[0]
				} else {
					version = "已安装"
				}
			} else {
				// 如果--version不起作用，尝试-v
				output, err = executeCommandWithTimeout(tool.command, "-v")
				if err == nil {
					lines := strings.Split(output, "\n")
					if len(lines) > 0 {
						version = lines[0]
					} else {
						version = "已安装"
					}
				} else {
					version = "已安装"
				}
			}
		}

		packages = append(packages, PackageInfo{
			Name:      tool.name,
			Version:   version,
			Installed: true,
		})

		// 限制数量
		if len(packages) >= 20 {
			break
		}
	}

	// 检查是否有bash配置文件中引用的库
	homeDir, err := os.UserHomeDir()
	if err == nil {
		// 检查常见的bash配置文件
		bashFiles := []string{
			filepath.Join(homeDir, ".bashrc"),
			filepath.Join(homeDir, ".bash_profile"),
			filepath.Join(homeDir, ".profile"),
		}

		for _, file := range bashFiles {
			if _, err := os.Stat(file); os.IsNotExist(err) {
				continue
			}

			// 读取文件内容
			content, err := os.ReadFile(file)
			if err != nil {
				continue
			}

			// 查找source或.命令引用的库
			re := regexp.MustCompile(`(?:source|\.)\s+([^\s;]+)`)
			matches := re.FindAllStringSubmatch(string(content), -1)

			for _, match := range matches {
				if len(match) >= 2 {
					libPath := match[1]
					// 展开~为home目录
					if strings.HasPrefix(libPath, "~/") {
						libPath = filepath.Join(homeDir, libPath[2:])
					}

					// 获取库名称
					libName := filepath.Base(libPath)

					// 检查文件是否存在
					if _, err := os.Stat(libPath); err == nil {
						packages = append(packages, PackageInfo{
							Name:      libName,
							Version:   "已加载",
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

	return packages, nil
}

// 检测VBA
func (a *App) detectVBA() LanguageInfo {
	info := LanguageInfo{
		Name:            "VBA",
		Installed:       false,
		DownloadURL:     "https://www.microsoft.com/en-us/microsoft-365/excel",
		InstallTutorial: "https://support.microsoft.com/en-us/office/getting-started-with-vba-in-office-4caac695-2398-4fbc-b020-e9982a9f8b4d",
		PackageManager:  "References",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "Microsoft Excel Object Library",
				Description: "Excel对象模型",
				Version:     "latest",
			},
			{
				Name:        "Microsoft Office Object Library",
				Description: "Office共享功能",
				Version:     "latest",
			},
			{
				Name:        "Microsoft Forms 2.0 Object Library",
				Description: "用户界面控件",
				Version:     "latest",
			},
			{
				Name:        "Microsoft Scripting Runtime",
				Description: "文件系统访问",
				Version:     "latest",
			},
			{
				Name:        "Microsoft Visual Basic for Applications Extensibility",
				Description: "VBA IDE访问",
				Version:     "latest",
			},
		},
	}

	// 检查Office应用程序是否安装
	officePaths := []string{
		"C:/Program Files/Microsoft Office",
		"C:/Program Files (x86)/Microsoft Office",
		"C:/Program Files/Microsoft Office 15", // Office 2013
		"C:/Program Files/Microsoft Office 16", // Office 2016/2019/365
	}

	for _, path := range officePaths {
		if _, err := os.Stat(path); err == nil {
			info.Installed = true
			info.Version = "Microsoft Office VBA"

			// 获取已安装的VBA引用
			packages, _ := a.listVBAPackages()
			info.Packages = packages

			break
		}
	}

	// 如果没有找到Office目录，尝试检查注册表
	if !info.Installed && runtime.GOOS == "windows" {
		// 使用PowerShell检查Office安装
		output, err := executeCommandWithTimeout("powershell", "-Command",
			"Get-ItemProperty HKLM:\\Software\\Microsoft\\Office\\*\\Excel\\InstallRoot -ErrorAction SilentlyContinue | Select-Object -ExpandProperty Path")

		if err == nil && output != "" {
			info.Installed = true
			info.Version = "Microsoft Office VBA"

			// 获取已安装的VBA引用
			packages, _ := a.listVBAPackages()
			info.Packages = packages
		}
	}

	return info
}

// 列出已安装的VBA引用
func (a *App) listVBAPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 常见的VBA引用
	commonReferences := []struct {
		name    string
		regPath string
	}{
		{"Microsoft Excel Object Library", "Excel.Application"},
		{"Microsoft Word Object Library", "Word.Application"},
		{"Microsoft Office Object Library", "Office.Application"},
		{"Microsoft Forms 2.0 Object Library", "Forms.Form"},
		{"Microsoft Scripting Runtime", "Scripting.FileSystemObject"},
		{"Microsoft Visual Basic for Applications Extensibility", "VBIDE.VBE"},
		{"Microsoft ActiveX Data Objects", "ADODB.Connection"},
		{"Microsoft XML", "MSXML2.DOMDocument"},
		{"Microsoft Outlook Object Library", "Outlook.Application"},
		{"Microsoft PowerPoint Object Library", "PowerPoint.Application"},
		{"Microsoft Access Object Library", "Access.Application"},
		{"Microsoft Publisher Object Library", "Publisher.Application"},
		{"Microsoft DAO Object Library", "DAO.DBEngine"},
		{"OLE Automation", "stdole.StdFont"},
		{"Microsoft Windows Common Controls", "MSComCtl.ListViewCtrl"},
	}

	// 在Windows系统上，尝试从注册表获取版本信息
	if runtime.GOOS == "windows" {
		for _, ref := range commonReferences {
			// 使用PowerShell检查注册表
			output, err := executeCommandWithTimeout("powershell", "-Command",
				fmt.Sprintf("Get-ItemProperty 'HKLM:\\SOFTWARE\\Classes\\%s\\CLSID' -ErrorAction SilentlyContinue", ref.regPath))

			if err == nil && output != "" {
				// 尝试提取版本信息
				versionOutput, err := executeCommandWithTimeout("powershell", "-Command",
					fmt.Sprintf("(Get-Item (Get-ItemProperty 'HKLM:\\SOFTWARE\\Classes\\%s\\CLSID' -ErrorAction SilentlyContinue).'@' -ErrorAction SilentlyContinue).VersionInfo.FileVersion", ref.regPath))

				version := "已安装"
				if err == nil && versionOutput != "" {
					version = versionOutput
				}

				packages = append(packages, PackageInfo{
					Name:      ref.name,
					Version:   version,
					Installed: true,
				})
			}

			// 限制数量
			if len(packages) >= 15 {
				break
			}
		}
	}

	// 如果没有找到任何引用，添加一些默认的
	if len(packages) == 0 {
		// 检查Office是否安装
		officePaths := []string{
			"C:/Program Files/Microsoft Office",
			"C:/Program Files (x86)/Microsoft Office",
		}

		officeInstalled := false
		for _, path := range officePaths {
			if _, err := os.Stat(path); err == nil {
				officeInstalled = true
				break
			}
		}

		if officeInstalled {
			packages = append(packages, PackageInfo{
				Name:      "VBA Standard Library",
				Version:   "内置",
				Installed: true,
			})
		}
	}

	return packages, nil
}

// 检测PowerShell
func (a *App) detectPowerShell() LanguageInfo {
	info := LanguageInfo{
		Name:            "PowerShell",
		Installed:       false,
		DownloadURL:     "https://github.com/PowerShell/PowerShell/releases",
		InstallTutorial: "https://docs.microsoft.com/en-us/powershell/scripting/install/installing-powershell",
		PackageManager:  "PowerShellGet",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "PSReadLine",
				Description: "命令行编辑体验增强",
				Version:     "2.2.6",
			},
			{
				Name:        "Az",
				Description: "Azure PowerShell模块",
				Version:     "9.3.0",
			},
			{
				Name:        "AWS.Tools.Common",
				Description: "AWS PowerShell模块",
				Version:     "4.1.118",
			},
			{
				Name:        "ImportExcel",
				Description: "Excel导入导出模块",
				Version:     "7.8.5",
			},
			{
				Name:        "Pester",
				Description: "PowerShell测试框架",
				Version:     "5.4.1",
			},
		},
	}

	// 检查PowerShell是否安装
	if !commandExists("powershell") && !commandExists("pwsh") {
		return info
	}

	// 首先尝试PowerShell Core (pwsh)
	psCommand := "pwsh"
	if !commandExists(psCommand) {
		psCommand = "powershell"
	}

	output, err := executeCommandWithTimeout(psCommand, "-Command", "$PSVersionTable.PSVersion.ToString()")
	if err == nil {
		info.Installed = true
		info.Version = "PowerShell " + strings.TrimSpace(output)

		// 获取已安装的PowerShell模块
		packages, _ := a.listPowerShellPackages(psCommand)
		info.Packages = packages
	}

	return info
}

// 列出已安装的PowerShell模块
func (a *App) listPowerShellPackages(psCommand string) ([]PackageInfo, error) {
	var packages []PackageInfo

	// 使用PowerShell命令获取已安装的模块
	output, err := executeCommandWithTimeout(psCommand, "-Command",
		"Get-Module -ListAvailable | Sort-Object -Property Name | Select-Object -First 20 | ForEach-Object { $_.Name + '|' + $_.Version }")

	if err != nil {
		return packages, err
	}

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

	// 如果没有找到模块，尝试另一种方法
	if len(packages) == 0 {
		// 尝试获取常见模块
		commonModules := []string{
			"PSReadLine",
			"Microsoft.PowerShell.Management",
			"Microsoft.PowerShell.Utility",
			"Microsoft.PowerShell.Security",
			"Microsoft.PowerShell.Host",
			"PackageManagement",
			"PowerShellGet",
			"Az",
			"AWS.Tools.Common",
			"ImportExcel",
			"Pester",
		}

		for _, module := range commonModules {
			output, err := executeCommandWithTimeout(psCommand, "-Command",
				fmt.Sprintf("(Get-Module -Name %s -ListAvailable | Select-Object -First 1).Version.ToString()", module))

			if err == nil && output != "" {
				packages = append(packages, PackageInfo{
					Name:      module,
					Version:   strings.TrimSpace(output),
					Installed: true,
				})
			}
		}
	}

	return packages, nil
}

// 检测SQL
func (a *App) detectSQL() LanguageInfo {
	info := LanguageInfo{
		Name:            "SQL",
		Installed:       false,
		DownloadURL:     "https://www.mysql.com/downloads/",
		InstallTutorial: "https://dev.mysql.com/doc/mysql-installation-excerpt/5.7/en/",
		PackageManager:  "数据库扩展",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "MySQL",
				Description: "开源关系型数据库",
				Version:     "8.0.33",
			},
			{
				Name:        "PostgreSQL",
				Description: "高级开源关系型数据库",
				Version:     "15.3",
			},
			{
				Name:        "SQLite",
				Description: "轻量级文件数据库",
				Version:     "3.42.0",
			},
			{
				Name:        "SQL Server",
				Description: "Microsoft企业级数据库",
				Version:     "2022",
			},
			{
				Name:        "Oracle Database",
				Description: "企业级关系型数据库",
				Version:     "21c",
			},
		},
	}

	// 检查各种SQL数据库是否安装
	dbClients := []struct {
		name    string
		command string
		args    []string
	}{
		{"MySQL", "mysql", []string{"--version"}},
		{"PostgreSQL", "psql", []string{"--version"}},
		{"SQLite", "sqlite3", []string{"--version"}},
		{"SQL Server", "sqlcmd", []string{"-?"}},
	}

	for _, client := range dbClients {
		if commandExists(client.command) {
			info.Installed = true
			output, err := executeCommandWithTimeout(client.command, client.args...)
			if err == nil {
				// 如果至少有一个客户端安装，则标记为已安装
				if info.Version == "" {
					info.Version = client.name + ": " + output
				} else {
					info.Version += ", " + client.name
				}
			} else {
				if info.Version == "" {
					info.Version = client.name + " (已安装)"
				} else {
					info.Version += ", " + client.name
				}
			}
		}
	}

	// 如果找到了数据库客户端，获取扩展信息
	if info.Installed {
		packages, _ := a.listSQLPackages()
		info.Packages = packages
	}

	return info
}

// 列出已安装的SQL数据库和扩展
func (a *App) listSQLPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 检查MySQL
	if commandExists("mysql") {
		// 尝试获取MySQL版本
		output, err := executeCommandWithTimeout("mysql", "--version")
		if err == nil {
			packages = append(packages, PackageInfo{
				Name:      "MySQL",
				Version:   output,
				Installed: true,
			})

			// 尝试获取已安装的MySQL插件
			pluginsOutput, err := executeCommandWithTimeout("mysql", "-e", "SHOW PLUGINS")
			if err == nil {
				lines := strings.Split(pluginsOutput, "\n")
				for i, line := range lines {
					// 跳过标题行
					if i == 0 || strings.TrimSpace(line) == "" {
						continue
					}

					fields := strings.Fields(line)
					if len(fields) >= 2 {
						packages = append(packages, PackageInfo{
							Name:      "MySQL Plugin: " + fields[0],
							Version:   "已启用",
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

	// 检查PostgreSQL
	if commandExists("psql") {
		// 尝试获取PostgreSQL版本
		output, err := executeCommandWithTimeout("psql", "--version")
		if err == nil {
			packages = append(packages, PackageInfo{
				Name:      "PostgreSQL",
				Version:   output,
				Installed: true,
			})

			// 尝试获取已安装的PostgreSQL扩展
			extOutput, err := executeCommandWithTimeout("psql", "-c", "SELECT name, default_version FROM pg_available_extensions LIMIT 10")
			if err == nil {
				lines := strings.Split(extOutput, "\n")
				for i, line := range lines {
					// 跳过标题行和分隔行
					if i <= 1 || strings.TrimSpace(line) == "" || strings.Contains(line, "----") {
						continue
					}

					fields := strings.Fields(line)
					if len(fields) >= 2 {
						packages = append(packages, PackageInfo{
							Name:      "PostgreSQL Extension: " + fields[0],
							Version:   fields[1],
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

	// 检查SQLite
	if commandExists("sqlite3") {
		output, err := executeCommandWithTimeout("sqlite3", "--version")
		if err == nil {
			packages = append(packages, PackageInfo{
				Name:      "SQLite",
				Version:   output,
				Installed: true,
			})
		}
	}

	// 检查SQL Server
	if commandExists("sqlcmd") {
		packages = append(packages, PackageInfo{
			Name:      "SQL Server Client Tools",
			Version:   "已安装",
			Installed: true,
		})
	}

	// 检查Oracle
	if commandExists("sqlplus") {
		output, err := executeCommandWithTimeout("sqlplus", "-v")
		if err == nil {
			packages = append(packages, PackageInfo{
				Name:      "Oracle Database Client",
				Version:   output,
				Installed: true,
			})
		}
	}

	// 检查常见的SQL工具
	sqlTools := []struct {
		name    string
		command string
		args    []string
	}{
		{"DBeaver", "dbeaver", []string{"--version"}},
		{"MySQL Workbench", "mysql-workbench", []string{"--version"}},
		{"pgAdmin", "pgadmin4", []string{"--version"}},
		{"SQL Server Management Studio", "ssms", []string{}},
		{"Oracle SQL Developer", "sqldeveloper", []string{}},
	}

	for _, tool := range sqlTools {
		if commandExists(tool.command) {
			version := "已安装"
			if len(tool.args) > 0 {
				output, err := executeCommandWithTimeout(tool.command, tool.args...)
				if err == nil {
					version = output
				}
			}

			packages = append(packages, PackageInfo{
				Name:      tool.name,
				Version:   version,
				Installed: true,
			})

			// 限制数量
			if len(packages) >= 20 {
				return packages, nil
			}
		}
	}

	return packages, nil
}

// 检测HTML/CSS
func (a *App) detectHTML() LanguageInfo {
	info := LanguageInfo{
		Name:            "HTML/CSS",
		Installed:       true, // HTML/CSS总是可用的，不需要安装
		DownloadURL:     "https://developer.mozilla.org/en-US/docs/Web/HTML",
		InstallTutorial: "https://developer.mozilla.org/en-US/docs/Learn/HTML/Introduction_to_HTML",
		PackageManager:  "npm/yarn",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "Bootstrap",
				Description: "流行的CSS框架",
				Version:     "5.3.0",
			},
			{
				Name:        "Tailwind CSS",
				Description: "功能类优先的CSS框架",
				Version:     "3.3.2",
			},
			{
				Name:        "Sass",
				Description: "CSS预处理器",
				Version:     "1.62.1",
			},
			{
				Name:        "Less",
				Description: "CSS预处理器",
				Version:     "4.1.3",
			},
			{
				Name:        "PostCSS",
				Description: "CSS转换工具",
				Version:     "8.4.24",
			},
		},
	}

	// 设置版本为当前支持的HTML和CSS版本
	info.Version = "HTML5, CSS3"

	// 获取已安装的HTML/CSS相关工具和框架
	packages, _ := a.listHTMLPackages()
	info.Packages = packages

	return info
}

// 列出已安装的HTML/CSS相关工具和框架
func (a *App) listHTMLPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 检查常见的前端工具
	frontendTools := []struct {
		name    string
		command string
		args    []string
	}{
		{"Sass", "sass", []string{"--version"}},
		{"Less", "lessc", []string{"--version"}},
		{"PostCSS", "postcss", []string{"--version"}},
		{"Stylelint", "stylelint", []string{"--version"}},
		{"PurgeCSS", "purgecss", []string{"--version"}},
		{"Autoprefixer", "autoprefixer", []string{"--version"}},
		{"Browsersync", "browser-sync", []string{"--version"}},
		{"Gulp", "gulp", []string{"--version"}},
		{"Webpack", "webpack", []string{"--version"}},
		{"Parcel", "parcel", []string{"--version"}},
	}

	for _, tool := range frontendTools {
		if commandExists(tool.command) {
			version := "已安装"
			if len(tool.args) > 0 {
				output, err := executeCommandWithTimeout(tool.command, tool.args...)
				if err == nil {
					version = output
				}
			}

			packages = append(packages, PackageInfo{
				Name:      tool.name,
				Version:   version,
				Installed: true,
			})

			// 限制数量
			if len(packages) >= 20 {
				return packages, nil
			}
		}
	}

	// 检查Node.js是否安装，如果安装了，检查全局安装的前端包
	if commandExists("npm") {
		// 检查常见的前端框架和库
		frontendPackages := []string{
			"bootstrap",
			"tailwindcss",
			"@angular/cli",
			"@vue/cli",
			"create-react-app",
			"next",
			"nuxt",
			"svelte",
			"ember-cli",
		}

		for _, pkg := range frontendPackages {
			// 检查全局安装的包
			output, err := executeCommandWithTimeout("npm", "list", "-g", pkg, "--depth=0")
			if err == nil && !strings.Contains(output, "empty") {
				// 提取版本信息
				re := regexp.MustCompile(pkg + "@([\\d\\.]+)")
				match := re.FindStringSubmatch(output)
				version := "已安装"
				if len(match) >= 2 {
					version = match[1]
				}

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

	// 检查用户主目录下是否有前端项目
	homeDir, err := os.UserHomeDir()
	if err == nil {
		// 检查常见的项目目录
		projectDirs := []string{
			filepath.Join(homeDir, "Projects"),
			filepath.Join(homeDir, "Documents", "Projects"),
			filepath.Join(homeDir, "source", "repos"),
		}

		for _, dir := range projectDirs {
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				continue
			}

			// 递归查找package.json文件
			err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return nil
				}

				if !info.IsDir() && info.Name() == "package.json" {
					// 读取package.json文件
					content, err := os.ReadFile(path)
					if err != nil {
						return nil
					}

					// 解析JSON
					var packageJSON map[string]interface{}
					if err := json.Unmarshal(content, &packageJSON); err != nil {
						return nil
					}

					// 检查依赖
					dependencies := make(map[string]interface{})
					if deps, ok := packageJSON["dependencies"].(map[string]interface{}); ok {
						for k, v := range deps {
							dependencies[k] = v
						}
					}
					if devDeps, ok := packageJSON["devDependencies"].(map[string]interface{}); ok {
						for k, v := range devDeps {
							dependencies[k] = v
						}
					}

					// 查找前端相关的包
					frontendKeywords := []string{
						"bootstrap", "tailwind", "css", "sass", "less", "postcss", "style", "html",
						"angular", "react", "vue", "svelte", "ember", "next", "nuxt", "gatsby",
					}

					for pkg, version := range dependencies {
						for _, keyword := range frontendKeywords {
							if strings.Contains(strings.ToLower(pkg), keyword) {
								packages = append(packages, PackageInfo{
									Name:      pkg,
									Version:   fmt.Sprintf("%v", version),
									Installed: true,
								})

								// 限制数量
								if len(packages) >= 20 {
									return filepath.SkipDir
								}
								break
							}
						}
					}
				}
				return nil
			})
		}
	}

	// 如果没有找到任何包，添加一些默认的
	if len(packages) == 0 {
		defaultPackages := []struct {
			name    string
			version string
		}{
			{"HTML5", "标准"},
			{"CSS3", "标准"},
			{"Web浏览器", "已安装"},
		}

		for _, pkg := range defaultPackages {
			packages = append(packages, PackageInfo{
				Name:      pkg.name,
				Version:   pkg.version,
				Installed: true,
			})
		}
	}

	return packages, nil
}
