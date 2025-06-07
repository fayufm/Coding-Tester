package main

import (
	"strings"
)

// 检测Scheme
func (a *App) detectScheme() LanguageInfo {
	info := LanguageInfo{
		Name:            "Scheme",
		Installed:       false,
		DownloadURL:     "https://www.scheme.com/download/",
		InstallTutorial: "https://www.scheme.com/tspl4/",
		PackageManager:  "Akku",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "chez-scheme",
				Description: "Chez Scheme实现",
				Version:     "9.5.8",
			},
			{
				Name:        "racket",
				Description: "Racket编程语言",
				Version:     "8.9",
			},
			{
				Name:        "guile",
				Description: "GNU的Scheme实现",
				Version:     "3.0.9",
			},
		},
	}

	// 检测Chez Scheme
	if commandExists("scheme") {
		output, err := executeCommandWithTimeout("scheme", "--version")
		if err == nil {
			info.Installed = true
			info.Version = "Chez Scheme: " + output
			return info
		}
	}

	// 检测Guile
	if commandExists("guile") {
		output, err := executeCommandWithTimeout("guile", "--version")
		if err == nil {
			info.Installed = true
			info.Version = "Guile: " + output
			return info
		}
	}

	// 检测Racket
	if commandExists("racket") {
		output, err := executeCommandWithTimeout("racket", "--version")
		if err == nil {
			info.Installed = true
			info.Version = "Racket: " + output
			return info
		}
	}

	return info
}

// 检测Crystal
func (a *App) detectCrystal() LanguageInfo {
	info := LanguageInfo{
		Name:            "Crystal",
		Installed:       false,
		DownloadURL:     "https://crystal-lang.org/install/",
		InstallTutorial: "https://crystal-lang.org/reference/",
		PackageManager:  "shards",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "kemal",
				Description: "Crystal的快速、简单的Web框架",
				Version:     "1.4.0",
			},
			{
				Name:        "lucky",
				Description: "全栈Web框架",
				Version:     "1.0.0",
			},
			{
				Name:        "amber",
				Description: "Crystal的Web应用框架",
				Version:     "0.36.0",
			},
		},
	}

	if !commandExists("crystal") {
		return info
	}

	output, err := executeCommandWithTimeout("crystal", "--version")
	if err == nil {
		info.Installed = true
		info.Version = output
	}

	return info
}

// 检测Nim
func (a *App) detectNim() LanguageInfo {
	info := LanguageInfo{
		Name:            "Nim",
		Installed:       false,
		DownloadURL:     "https://nim-lang.org/install.html",
		InstallTutorial: "https://nim-lang.org/docs/tut1.html",
		PackageManager:  "nimble",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "jester",
				Description: "Nim的Web框架",
				Version:     "0.5.0",
			},
			{
				Name:        "karax",
				Description: "单页应用框架",
				Version:     "1.2.2",
			},
			{
				Name:        "nimx",
				Description: "跨平台GUI框架",
				Version:     "0.3.0",
			},
		},
	}

	if !commandExists("nim") {
		return info
	}

	output, err := executeCommandWithTimeout("nim", "--version")
	if err == nil {
		info.Installed = true
		info.Version = output
	}

	return info
}

// 检测D语言
func (a *App) detectD() LanguageInfo {
	info := LanguageInfo{
		Name:            "D",
		Installed:       false,
		DownloadURL:     "https://dlang.org/download.html",
		InstallTutorial: "https://dlang.org/getting_started.html",
		PackageManager:  "dub",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "vibe-d",
				Description: "D语言的异步I/O和Web框架",
				Version:     "0.9.5",
			},
			{
				Name:        "mir-algorithm",
				Description: "数值算法和数据结构",
				Version:     "3.20.0",
			},
			{
				Name:        "dxml",
				Description: "XML解析库",
				Version:     "0.4.3",
			},
		},
	}

	// 检测DMD编译器
	if commandExists("dmd") {
		output, err := executeCommandWithTimeout("dmd", "--version")
		if err == nil {
			info.Installed = true
			info.Version = "DMD: " + output
			return info
		}
	}

	// 检测LDC编译器
	if commandExists("ldc2") {
		output, err := executeCommandWithTimeout("ldc2", "--version")
		if err == nil {
			info.Installed = true
			info.Version = "LDC: " + output
			return info
		}
	}

	// 检测GDC编译器
	if commandExists("gdc") {
		output, err := executeCommandWithTimeout("gdc", "--version")
		if err == nil {
			info.Installed = true
			info.Version = "GDC: " + output
			return info
		}
	}

	return info
}

// 检测Ada
func (a *App) detectAda() LanguageInfo {
	info := LanguageInfo{
		Name:            "Ada",
		Installed:       false,
		DownloadURL:     "https://www.adacore.com/download",
		InstallTutorial: "https://learn.adacore.com/courses/intro-to-ada/index.html",
		PackageManager:  "Alire",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "gnatcoll",
				Description: "GNAT组件集合",
				Version:     "23.0.0",
			},
			{
				Name:        "aws",
				Description: "Ada Web Server",
				Version:     "23.0.0",
			},
			{
				Name:        "aunit",
				Description: "Ada单元测试框架",
				Version:     "23.0.0",
			},
		},
	}

	if !commandExists("gnat") {
		return info
	}

	output, err := executeCommandWithTimeout("gnat", "--version")
	if err == nil {
		info.Installed = true
		info.Version = output
	}

	return info
}

// 检测VHDL
func (a *App) detectVHDL() LanguageInfo {
	info := LanguageInfo{
		Name:            "VHDL",
		Installed:       false,
		DownloadURL:     "https://ghdl.github.io/ghdl/",
		InstallTutorial: "https://ghdl.github.io/ghdl/getting/",
		PackageManager:  "N/A",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "GHDL",
				Description: "开源VHDL模拟器",
				Version:     "3.0.0",
			},
			{
				Name:        "ModelSim",
				Description: "VHDL/Verilog模拟器",
				Version:     "latest",
			},
			{
				Name:        "Vivado",
				Description: "Xilinx综合和分析工具",
				Version:     "latest",
			},
		},
	}

	// 检测GHDL
	if commandExists("ghdl") {
		output, err := executeCommandWithTimeout("ghdl", "--version")
		if err == nil {
			info.Installed = true
			info.Version = "GHDL: " + output
			return info
		}
	}

	return info
}

// 检测Erlang
func (a *App) detectErlang() LanguageInfo {
	info := LanguageInfo{
		Name:            "Erlang",
		Installed:       false,
		DownloadURL:     "https://www.erlang.org/downloads",
		InstallTutorial: "https://www.erlang.org/doc/getting_started/intro.html",
		PackageManager:  "rebar3",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "cowboy",
				Description: "小型、快速、模块化的HTTP服务器",
				Version:     "2.10.0",
			},
			{
				Name:        "lager",
				Description: "Erlang/OTP的日志框架",
				Version:     "3.9.2",
			},
			{
				Name:        "jiffy",
				Description: "JSON解码器/编码器",
				Version:     "1.1.1",
			},
		},
	}

	if !commandExists("erl") {
		return info
	}

	output, err := executeCommandWithTimeout("erl", "-eval", "io:format(\"~s\", [erlang:system_info(otp_release)]), halt().", "-noshell")
	if err == nil {
		info.Installed = true
		info.Version = "OTP " + output
	}

	return info
}

// 检测Smalltalk
func (a *App) detectSmalltalk() LanguageInfo {
	info := LanguageInfo{
		Name:            "Smalltalk",
		Installed:       false,
		DownloadURL:     "https://squeak.org/downloads/",
		InstallTutorial: "https://squeak.org/documentation/",
		PackageManager:  "Monticello",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "Squeak",
				Description: "开源Smalltalk实现",
				Version:     "6.0",
			},
			{
				Name:        "Pharo",
				Description: "现代、开源的Smalltalk实现",
				Version:     "10.0",
			},
			{
				Name:        "GNU Smalltalk",
				Description: "GNU项目的Smalltalk实现",
				Version:     "3.2.5",
			},
		},
	}

	// 检测GNU Smalltalk
	if commandExists("gst") {
		output, err := executeCommandWithTimeout("gst", "--version")
		if err == nil {
			info.Installed = true
			info.Version = "GNU Smalltalk: " + output
			return info
		}
	}

	// Squeak和Pharo通常不通过命令行调用
	return info
}

// 检测OCaml
func (a *App) detectOCaml() LanguageInfo {
	info := LanguageInfo{
		Name:            "OCaml",
		Installed:       false,
		DownloadURL:     "https://ocaml.org/docs/install.html",
		InstallTutorial: "https://ocaml.org/learn/tutorials/",
		PackageManager:  "OPAM",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "core",
				Description: "替代OCaml标准库的工业级库",
				Version:     "v0.15.1",
			},
			{
				Name:        "dune",
				Description: "OCaml的构建系统",
				Version:     "3.7.0",
			},
			{
				Name:        "lwt",
				Description: "协作线程库",
				Version:     "5.6.1",
			},
		},
	}

	if !commandExists("ocaml") {
		return info
	}

	output, err := executeCommandWithTimeout("ocaml", "-version")
	if err == nil {
		info.Installed = true
		info.Version = output
	}

	return info
}

// 检测Tcl
func (a *App) detectTcl() LanguageInfo {
	info := LanguageInfo{
		Name:            "Tcl",
		Installed:       false,
		DownloadURL:     "https://www.tcl.tk/software/tcltk/",
		InstallTutorial: "https://www.tcl.tk/doc/",
		PackageManager:  "Teapot",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "Tk",
				Description: "Tcl的GUI工具包",
				Version:     "8.6.12",
			},
			{
				Name:        "Expect",
				Description: "自动化交互应用程序",
				Version:     "5.45.4",
			},
			{
				Name:        "TclOO",
				Description: "Tcl的面向对象系统",
				Version:     "1.1.0",
			},
		},
	}

	if !commandExists("tclsh") {
		return info
	}

	output, err := executeCommandWithTimeout("tclsh", "-c", "puts $tcl_version")
	if err == nil {
		info.Installed = true
		info.Version = "Tcl " + output
	}

	return info
}

// 检测Apex
func (a *App) detectApex() LanguageInfo {
	info := LanguageInfo{
		Name:            "Apex",
		Installed:       false,
		DownloadURL:     "https://developer.salesforce.com/docs/atlas.en-us.apexcode.meta/apexcode/apex_intro_get_started.htm",
		InstallTutorial: "https://trailhead.salesforce.com/en/content/learn/trails/force_com_dev_beginner",
		PackageManager:  "Salesforce CLI",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "Salesforce CLI",
				Description: "命令行界面，用于与Salesforce组织交互",
				Version:     "latest",
			},
			{
				Name:        "Salesforce Extensions for VS Code",
				Description: "用于Apex开发的VS Code扩展",
				Version:     "latest",
			},
		},
	}

	// 检测Salesforce CLI
	if commandExists("sfdx") {
		output, err := executeCommandWithTimeout("sfdx", "--version")
		if err == nil {
			info.Installed = true
			info.Version = "Salesforce CLI: " + output
		}
	}

	return info
}

// 检测Solidity
func (a *App) detectSolidity() LanguageInfo {
	info := LanguageInfo{
		Name:            "Solidity",
		Installed:       false,
		DownloadURL:     "https://soliditylang.org/",
		InstallTutorial: "https://docs.soliditylang.org/en/latest/installing-solidity.html",
		PackageManager:  "npm",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "solc",
				Description: "Solidity编译器",
				Version:     "0.8.19",
			},
			{
				Name:        "truffle",
				Description: "以太坊开发环境",
				Version:     "5.9.2",
			},
			{
				Name:        "hardhat",
				Description: "以太坊开发环境",
				Version:     "2.16.1",
			},
		},
	}

	// 检测solc编译器
	if commandExists("solc") {
		output, err := executeCommandWithTimeout("solc", "--version")
		if err == nil {
			info.Installed = true
			info.Version = output
			return info
		}
	}

	// 检测通过npm安装的solc
	if commandExists("npm") {
		output, err := executeCommandWithTimeout("npm", "list", "-g", "solc")
		if err == nil && !strings.Contains(output, "empty") {
			info.Installed = true
			info.Version = "solc (npm)"
			return info
		}
	}

	return info
}

// 检测WebAssembly
func (a *App) detectWebAssembly() LanguageInfo {
	info := LanguageInfo{
		Name:            "WebAssembly",
		Installed:       false,
		DownloadURL:     "https://webassembly.org/getting-started/developers-guide/",
		InstallTutorial: "https://developer.mozilla.org/en-US/docs/WebAssembly",
		PackageManager:  "N/A",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "emscripten",
				Description: "将C/C++编译为WebAssembly的工具链",
				Version:     "3.1.42",
			},
			{
				Name:        "wasm-pack",
				Description: "将Rust编译为WebAssembly的打包工具",
				Version:     "0.12.0",
			},
			{
				Name:        "wabt",
				Description: "WebAssembly二进制工具包",
				Version:     "1.0.33",
			},
		},
	}

	// 检测emscripten
	if commandExists("emcc") {
		output, err := executeCommandWithTimeout("emcc", "--version")
		if err == nil {
			info.Installed = true
			info.Version = "Emscripten: " + output

			// 获取已安装的WebAssembly相关工具
			packages, _ := a.listWebAssemblyTools()
			info.Packages = packages

			return info
		}
	}

	// 检测wasm-pack
	if commandExists("wasm-pack") {
		output, err := executeCommandWithTimeout("wasm-pack", "--version")
		if err == nil {
			info.Installed = true
			info.Version = "wasm-pack: " + output

			// 获取已安装的WebAssembly相关工具
			packages, _ := a.listWebAssemblyTools()
			info.Packages = packages

			return info
		}
	}

	return info
}

// 列出已安装的WebAssembly相关工具
func (a *App) listWebAssemblyTools() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 检测emscripten
	if commandExists("emcc") {
		output, err := executeCommandWithTimeout("emcc", "--version")
		if err == nil {
			packages = append(packages, PackageInfo{
				Name:        "emscripten",
				Description: "将C/C++编译为WebAssembly的工具链",
				Version:     strings.TrimSpace(output),
				Installed:   true,
			})
		}
	}

	// 检测wasm-pack
	if commandExists("wasm-pack") {
		output, err := executeCommandWithTimeout("wasm-pack", "--version")
		if err == nil {
			packages = append(packages, PackageInfo{
				Name:        "wasm-pack",
				Description: "将Rust编译为WebAssembly的打包工具",
				Version:     strings.TrimSpace(output),
				Installed:   true,
			})
		}
	}

	// 检测wabt工具包
	wabtTools := []string{"wasm2wat", "wat2wasm", "wasm-objdump", "wasm-interp", "wasm-validate"}
	for _, tool := range wabtTools {
		if commandExists(tool) {
			output, err := executeCommandWithTimeout(tool, "--version")
			if err == nil {
				packages = append(packages, PackageInfo{
					Name:      tool,
					Version:   strings.TrimSpace(output),
					Installed: true,
				})
			}
		}
	}

	// 检测wasmer
	if commandExists("wasmer") {
		output, err := executeCommandWithTimeout("wasmer", "--version")
		if err == nil {
			packages = append(packages, PackageInfo{
				Name:        "wasmer",
				Description: "WebAssembly运行时",
				Version:     strings.TrimSpace(output),
				Installed:   true,
			})
		}
	}

	// 检测wasmtime
	if commandExists("wasmtime") {
		output, err := executeCommandWithTimeout("wasmtime", "--version")
		if err == nil {
			packages = append(packages, PackageInfo{
				Name:        "wasmtime",
				Description: "WebAssembly运行时",
				Version:     strings.TrimSpace(output),
				Installed:   true,
			})
		}
	}

	// 检测AssemblyScript
	if commandExists("asc") {
		output, err := executeCommandWithTimeout("asc", "--version")
		if err == nil {
			packages = append(packages, PackageInfo{
				Name:        "AssemblyScript",
				Description: "TypeScript到WebAssembly编译器",
				Version:     strings.TrimSpace(output),
				Installed:   true,
			})
		}
	}

	// 通过npm检查AssemblyScript
	if commandExists("npm") {
		output, err := executeCommandWithTimeout("npm", "list", "-g", "assemblyscript")
		if err == nil && strings.Contains(output, "assemblyscript") {
			version := "unknown"
			lines := strings.Split(output, "\n")
			for _, line := range lines {
				if strings.Contains(line, "assemblyscript@") {
					parts := strings.Split(line, "@")
					if len(parts) > 1 {
						version = parts[len(parts)-1]
					}
					break
				}
			}

			// 检查是否已经添加过
			found := false
			for _, pkg := range packages {
				if pkg.Name == "AssemblyScript" {
					found = true
					break
				}
			}

			if !found {
				packages = append(packages, PackageInfo{
					Name:        "AssemblyScript",
					Description: "TypeScript到WebAssembly编译器",
					Version:     version,
					Installed:   true,
				})
			}
		}
	}

	return packages, nil
}

// 检测Zig
func (a *App) detectZig() LanguageInfo {
	info := LanguageInfo{
		Name:            "Zig",
		Installed:       false,
		DownloadURL:     "https://ziglang.org/download/",
		InstallTutorial: "https://ziglang.org/learn/",
		PackageManager:  "zig build",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "zls",
				Description: "Zig语言服务器",
				Version:     "latest",
			},
			{
				Name:        "zigmod",
				Description: "Zig的包管理器",
				Version:     "latest",
			},
			{
				Name:        "gyro",
				Description: "Zig的包管理器",
				Version:     "latest",
			},
		},
	}

	if !commandExists("zig") {
		return info
	}

	output, err := executeCommandWithTimeout("zig", "version")
	if err == nil {
		info.Installed = true
		info.Version = output
	}

	return info
}

// 检测Haxe
func (a *App) detectHaxe() LanguageInfo {
	info := LanguageInfo{
		Name:            "Haxe",
		Installed:       false,
		DownloadURL:     "https://haxe.org/download/",
		InstallTutorial: "https://haxe.org/documentation/introduction/",
		PackageManager:  "haxelib",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "openfl",
				Description: "跨平台应用程序框架",
				Version:     "9.2.1",
			},
			{
				Name:        "lime",
				Description: "轻量级跨平台游戏框架",
				Version:     "8.0.1",
			},
			{
				Name:        "heaps",
				Description: "高性能游戏框架",
				Version:     "1.10.0",
			},
		},
	}

	if !commandExists("haxe") {
		return info
	}

	output, err := executeCommandWithTimeout("haxe", "--version")
	if err == nil {
		info.Installed = true
		info.Version = output
	}

	return info
}
