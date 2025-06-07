package main

import (
	"regexp"
	"strings"
)

// 为OCaml添加包检测功能
func (a *App) detectOCamlWithPackages() LanguageInfo {
	info := a.detectOCaml()

	if info.Installed {
		// 获取已安装的OCaml包
		packages, _ := a.listOCamlPackages()
		info.Packages = packages
	}

	return info
}

// 列出已安装的OCaml包
func (a *App) listOCamlPackages() ([]PackageInfo, error) {
	var packages []PackageInfo

	// 检查opam命令是否可用（OCaml包管理器）
	if !commandExists("opam") {
		return packages, nil
	}

	// 使用opam list获取已安装的包
	output, err := executeCommandWithTimeout("opam", "list", "--installed", "--columns=name,version,synopsis", "--normalise", "--color=never")
	if err != nil {
		return packages, err
	}

	// 解析opam输出
	return a.parseOPAMList(output)
}

// 解析OPAM列表输出
func (a *App) parseOPAMList(output string) ([]PackageInfo, error) {
	var packages []PackageInfo

	// 跳过标题行
	lines := strings.Split(output, "\n")
	if len(lines) < 2 {
		return packages, nil
	}

	// 从第二行开始处理
	for i := 1; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}

		// 解析行内容
		parts := strings.SplitN(line, " ", 3)
		if len(parts) >= 2 {
			name := strings.TrimSpace(parts[0])
			version := strings.TrimSpace(parts[1])

			description := ""
			if len(parts) >= 3 {
				description = strings.TrimSpace(parts[2])
			}

			// 如果没有描述，尝试获取预定义描述
			if description == "" {
				description = getOCamlPackageDescription(name)
			}

			packages = append(packages, PackageInfo{
				Name:        name,
				Version:     version,
				Description: description,
				Installed:   true,
			})
		}
	}

	// 如果上面的方法没有获取到包，尝试另一种解析方式
	if len(packages) == 0 {
		// 尝试使用正则表达式解析
		re := regexp.MustCompile(`([^\s]+)\s+([^\s]+)\s+(.*)`)
		for _, line := range lines {
			if strings.TrimSpace(line) == "" {
				continue
			}

			matches := re.FindStringSubmatch(line)
			if len(matches) >= 3 {
				name := matches[1]
				version := matches[2]
				description := ""
				if len(matches) >= 4 {
					description = matches[3]
				}

				if description == "" {
					description = getOCamlPackageDescription(name)
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

	// 尝试直接获取包列表
	if len(packages) == 0 {
		output, err := executeCommandWithTimeout("opam", "list", "--installed", "--short")
		if err == nil {
			lines := strings.Split(output, "\n")
			for _, line := range lines {
				line = strings.TrimSpace(line)
				if line == "" {
					continue
				}

				parts := strings.Split(line, ".")
				if len(parts) >= 2 {
					name := parts[0]
					version := parts[1]
					description := getOCamlPackageDescription(name)

					packages = append(packages, PackageInfo{
						Name:        name,
						Version:     version,
						Description: description,
						Installed:   true,
					})
				} else {
					// 如果没有版本信息
					name := line
					description := getOCamlPackageDescription(name)

					packages = append(packages, PackageInfo{
						Name:        name,
						Version:     "installed",
						Description: description,
						Installed:   true,
					})
				}
			}
		}
	}

	// 如果仍然没有找到包，添加一些常见的OCaml包作为建议
	if len(packages) == 0 {
		packages = append(packages, PackageInfo{
			Name:        "core",
			Version:     "latest",
			Description: "OCaml的替代标准库",
			Installed:   false,
		})
		packages = append(packages, PackageInfo{
			Name:        "dune",
			Version:     "latest",
			Description: "OCaml的构建系统",
			Installed:   false,
		})
		packages = append(packages, PackageInfo{
			Name:        "ocaml-lsp-server",
			Version:     "latest",
			Description: "OCaml语言服务器",
			Installed:   false,
		})
		packages = append(packages, PackageInfo{
			Name:        "merlin",
			Version:     "latest",
			Description: "OCaml的编辑器支持",
			Installed:   false,
		})
		packages = append(packages, PackageInfo{
			Name:        "utop",
			Version:     "latest",
			Description: "改进的OCaml REPL",
			Installed:   false,
		})
	}

	return packages, nil
}

// 获取OCaml包的描述
func getOCamlPackageDescription(name string) string {
	descriptions := map[string]string{
		"core":                    "OCaml的替代标准库",
		"dune":                    "OCaml的构建系统",
		"ocaml-lsp-server":        "OCaml语言服务器",
		"merlin":                  "OCaml的编辑器支持",
		"utop":                    "改进的OCaml REPL",
		"base":                    "OCaml全功能标准库替代品的精简版",
		"async":                   "OCaml的异步编程库",
		"lwt":                     "OCaml的轻量级线程库",
		"cmdliner":                "用于命令行接口的声明式定义",
		"ppx_deriving":            "类型导出的编译时代码生成",
		"yojson":                  "OCaml的JSON库",
		"cohttp":                  "OCaml的HTTP客户端和服务器",
		"reason":                  "OCaml的替代语法",
		"ocamlformat":             "OCaml的代码格式化工具",
		"ocamlgraph":              "OCaml的图形库",
		"opam":                    "OCaml包管理器",
		"ounit":                   "OCaml的单元测试框架",
		"ocaml-migrate-parsetree": "AST转换工具",
		"batteries":               "OCaml的替代标准库",
		"angstrom":                "OCaml的解析器组合器库",
	}

	if desc, ok := descriptions[name]; ok {
		return desc
	}

	// 根据包名推断描述
	if strings.Contains(name, "parser") {
		return "解析相关库"
	} else if strings.Contains(name, "json") {
		return "JSON处理库"
	} else if strings.Contains(name, "net") || strings.Contains(name, "http") {
		return "网络相关库"
	} else if strings.Contains(name, "test") {
		return "测试相关库"
	} else if strings.Contains(name, "db") || strings.Contains(name, "sql") {
		return "数据库相关库"
	} else if strings.Contains(name, "graph") {
		return "图形处理库"
	}

	return "OCaml包"
}
