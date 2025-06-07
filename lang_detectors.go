package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

// 检测Go语言
func (a *App) detectGo() LanguageInfo {
	info := LanguageInfo{
		Name:            "Go",
		Installed:       false,
		DownloadURL:     "https://golang.org/dl/",
		InstallTutorial: "https://golang.org/doc/install",
		PackageManager:  "go get",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "gin",
				Description: "Gin是一个用Go编写的HTTP Web框架，具有类似Martini的API，但性能更好",
				Version:     "v1.9.1",
			},
			{
				Name:        "gorm",
				Description: "GORM是Go语言的优秀ORM库，开发人员友好，支持关联查询、事务等特性",
				Version:     "v1.25.2",
			},
			{
				Name:        "cobra",
				Description: "用于创建强大的现代CLI应用程序的库",
				Version:     "v1.7.0",
			},
			{
				Name:        "viper",
				Description: "Go应用程序的完整配置解决方案",
				Version:     "v1.16.0",
			},
			{
				Name:        "fiber",
				Description: "Express风格的、基于Fasthttp的Web框架",
				Version:     "v2.47.0",
			},
		},
	}

	if !commandExists("go") {
		return info
	}

	output, err := executeCommandWithTimeout("go", "version")
	if err == nil {
		info.Installed = true
		info.Version = output

		// 检测已安装的包
		goPath := os.Getenv("GOPATH")
		if goPath != "" {
			pkgDir := filepath.Join(goPath, "pkg", "mod")
			if _, err := os.Stat(pkgDir); err == nil {
				// 列出已安装的包
				packages, _ := a.listGoPackages(pkgDir)
				info.Packages = packages
			}
		}
	}

	return info
}

// 列出Go包
func (a *App) listGoPackages(pkgDir string) ([]PackageInfo, error) {
	var packages []PackageInfo

	// 获取最多10个包作为示例
	count := 0
	err := filepath.Walk(pkgDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && path != pkgDir {
			relPath, _ := filepath.Rel(pkgDir, path)
			parts := strings.Split(relPath, "@")
			if len(parts) > 1 {
				pkg := PackageInfo{
					Name:      parts[0],
					Version:   parts[1],
					Installed: true,
				}
				packages = append(packages, pkg)
				count++

				// 只获取前10个包
				if count >= 10 {
					return filepath.SkipDir
				}
			}
		}
		return nil
	})

	return packages, err
}

// 检测Python
func (a *App) detectPython() LanguageInfo {
	info := LanguageInfo{
		Name:            "Python",
		Installed:       false,
		DownloadURL:     "https://www.python.org/downloads/",
		InstallTutorial: "https://www.python.org/about/gettingstarted/",
		PackageManager:  "pip",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "numpy",
				Description: "Python科学计算的基础包，支持大型多维数组和矩阵运算",
				Version:     "1.24.3",
			},
			{
				Name:        "pandas",
				Description: "数据分析和操作工具，提供数据结构和操作数据的工具",
				Version:     "2.0.2",
			},
			{
				Name:        "requests",
				Description: "简单而优雅的HTTP库，用于发送HTTP请求",
				Version:     "2.31.0",
			},
			{
				Name:        "flask",
				Description: "轻量级Web应用框架",
				Version:     "2.3.2",
			},
			{
				Name:        "django",
				Description: "高级Python Web框架，鼓励快速开发和清洁实用的设计",
				Version:     "4.2.2",
			},
		},
	}

	// 尝试Python 3
	pythonCommands := []string{"python", "python3"}
	for _, cmd := range pythonCommands {
		if !commandExists(cmd) {
			continue
		}

		output, err := executeCommandWithTimeout(cmd, "--version")
		if err == nil {
			info.Installed = true
			info.Version = output

			// 检查pip是否安装
			pipCommands := []string{"pip", "pip3"}
			for _, pipCmd := range pipCommands {
				if !commandExists(pipCmd) {
					continue
				}

				_, pipErr := executeCommandWithTimeout(pipCmd, "--version")
				if pipErr == nil {
					// 列出已安装的Python包
					packages, _ := a.listPipPackages(pipCmd)
					info.Packages = packages
					break
				}
			}

			// 如果没有找到pip
			if len(info.Packages) == 0 {
				info.MissingDeps = append(info.MissingDeps, "pip")
			}

			break
		}
	}

	return info
}

// 列出Python包
func (a *App) listPipPackages(pipCmd string) ([]PackageInfo, error) {
	output, err := executeCommandWithTimeout(pipCmd, "list", "--format=json")
	if err != nil {
		return nil, err
	}

	var pipPackages []struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	}

	if err := json.Unmarshal([]byte(output), &pipPackages); err != nil {
		return nil, err
	}

	packages := make([]PackageInfo, 0, len(pipPackages))
	for _, pkg := range pipPackages {
		packages = append(packages, PackageInfo{
			Name:      pkg.Name,
			Version:   pkg.Version,
			Installed: true,
		})
	}

	return packages, nil
}

// 检测Node.js
func (a *App) detectNode() LanguageInfo {
	info := LanguageInfo{
		Name:            "Node.js",
		Installed:       false,
		DownloadURL:     "https://nodejs.org/en/download/",
		InstallTutorial: "https://nodejs.org/en/learn/getting-started/how-to-install-nodejs",
		PackageManager:  "npm",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "express",
				Description: "快速、unopinionated的Node.js Web框架",
				Version:     "4.18.2",
			},
			{
				Name:        "react",
				Description: "用于构建用户界面的JavaScript库",
				Version:     "18.2.0",
			},
			{
				Name:        "vue",
				Description: "渐进式JavaScript框架，用于构建用户界面",
				Version:     "3.3.4",
			},
			{
				Name:        "axios",
				Description: "基于Promise的HTTP客户端，用于浏览器和Node.js",
				Version:     "1.4.0",
			},
			{
				Name:        "typescript",
				Description: "JavaScript的超集，添加了类型系统",
				Version:     "5.1.6",
			},
		},
	}

	if !commandExists("node") {
		return info
	}

	output, err := executeCommandWithTimeout("node", "--version")
	if err == nil {
		info.Installed = true
		info.Version = output

		// 检查npm是否安装
		if !commandExists("npm") {
			info.MissingDeps = append(info.MissingDeps, "npm")
		} else {
			// 列出全局安装的npm包
			packages, _ := a.listNpmPackages()
			info.Packages = packages
		}
	}

	return info
}

// 列出Node.js包
func (a *App) listNpmPackages() ([]PackageInfo, error) {
	output, _ := executeCommandWithTimeout("npm", "list", "--global", "--json", "--depth=0")
	// npm list 命令即使成功也可能返回非零退出码，所以我们继续处理输出

	var result struct {
		Dependencies map[string]struct {
			Version string `json:"version"`
		} `json:"dependencies"`
	}

	if err := json.Unmarshal([]byte(output), &result); err != nil {
		return nil, err
	}

	packages := make([]PackageInfo, 0, len(result.Dependencies))
	for name, info := range result.Dependencies {
		packages = append(packages, PackageInfo{
			Name:      name,
			Version:   info.Version,
			Installed: true,
		})
	}

	return packages, nil
}

// 检测Java
func (a *App) detectJava() LanguageInfo {
	info := LanguageInfo{
		Name:            "Java",
		Installed:       false,
		DownloadURL:     "https://www.oracle.com/java/technologies/downloads/",
		InstallTutorial: "https://docs.oracle.com/en/java/javase/17/install/overview-jdk-installation.html",
		PackageManager:  "Maven/Gradle",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "spring-boot-starter",
				Description: "Spring Boot核心启动器，包括自动配置支持、日志和YAML",
				Version:     "3.1.1",
			},
			{
				Name:        "junit",
				Description: "Java编程语言的单元测试框架",
				Version:     "5.9.3",
			},
			{
				Name:        "lombok",
				Description: "自动插入编辑器并构建工具，简化Java开发",
				Version:     "1.18.28",
			},
			{
				Name:        "gson",
				Description: "Google的JSON解析库",
				Version:     "2.10.1",
			},
			{
				Name:        "log4j",
				Description: "Java的可靠、快速和灵活的日志框架",
				Version:     "2.20.0",
			},
		},
	}

	if !commandExists("java") {
		return info
	}

	// 尝试不同的Java版本命令格式
	javaCommands := [][]string{
		{"java", "--version"},
		{"java", "-version"},
	}

	for _, cmdArgs := range javaCommands {
		output, err := executeCommandWithTimeout(cmdArgs[0], cmdArgs[1])
		if err == nil {
			info.Installed = true
			info.Version = output
			break
		}
	}

	if info.Installed {
		// 检查Maven是否安装
		if commandExists("mvn") {
			// 列出Maven已安装的包
			mavenPackages, _ := a.listMavenPackages()
			info.Packages = append(info.Packages, mavenPackages...)
		} else {
			info.MissingDeps = append(info.MissingDeps, "Maven")
		}

		// 检查Gradle是否安装
		if commandExists("gradle") {
			// 列出Gradle已安装的包
			gradlePackages, _ := a.listGradlePackages()
			info.Packages = append(info.Packages, gradlePackages...)
		} else {
			info.MissingDeps = append(info.MissingDeps, "Gradle")
		}
	}

	return info
}

// 列出Maven包
func (a *App) listMavenPackages() ([]PackageInfo, error) {
	// 尝试获取Maven本地仓库中的包信息
	// 这里我们使用Maven帮助命令来获取一些基本信息
	output, err := executeCommandWithTimeout("mvn", "help:evaluate", "-Dexpression=settings.localRepository", "-q", "-DforceStdout")
	if err != nil {
		return nil, err
	}

	// 获取Maven本地仓库路径
	localRepo := strings.TrimSpace(output)
	if localRepo == "" {
		// 如果无法获取本地仓库路径，使用默认路径
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		localRepo = filepath.Join(homeDir, ".m2", "repository")
	}

	// 查找最常用的Java库
	commonJavaLibs := []struct {
		groupId    string
		artifactId string
	}{
		{"org.springframework.boot", "spring-boot-starter"},
		{"org.springframework", "spring-core"},
		{"junit", "junit"},
		{"org.junit.jupiter", "junit-jupiter-api"},
		{"com.google.code.gson", "gson"},
		{"org.apache.logging.log4j", "log4j-core"},
		{"org.projectlombok", "lombok"},
		{"com.fasterxml.jackson.core", "jackson-databind"},
		{"org.hibernate", "hibernate-core"},
		{"org.apache.commons", "commons-lang3"},
	}

	packages := []PackageInfo{}

	// 检查每个常用库是否已安装
	for _, lib := range commonJavaLibs {
		libPath := filepath.Join(localRepo, strings.Replace(lib.groupId, ".", string(filepath.Separator), -1), lib.artifactId)
		if _, err := os.Stat(libPath); err == nil {
			// 尝试查找最新版本
			versions, err := os.ReadDir(libPath)
			if err != nil {
				continue
			}

			// 简单地取最后一个目录作为最新版本
			if len(versions) > 0 {
				latestVersion := versions[len(versions)-1].Name()
				packages = append(packages, PackageInfo{
					Name:      lib.groupId + ":" + lib.artifactId,
					Version:   latestVersion,
					Installed: true,
				})
			}
		}
	}

	return packages, nil
}

// 列出Gradle包
func (a *App) listGradlePackages() ([]PackageInfo, error) {
	// Gradle没有简单的方式列出全局安装的包，但我们可以检查Gradle缓存
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	// Gradle缓存路径
	gradleCachePath := filepath.Join(homeDir, ".gradle", "caches", "modules-2", "files-2.1")

	// 检查路径是否存在
	if _, err := os.Stat(gradleCachePath); err != nil {
		return nil, err
	}

	packages := []PackageInfo{}
	count := 0

	// 遍历目录查找常用的Gradle包
	err = filepath.Walk(gradleCachePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 限制数量，避免遍历过多
		if count >= 10 {
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
			// 这可能是一个group/name/version结构
			group := parts[0]
			name := parts[1]
			version := parts[2]

			// 添加到结果
			packages = append(packages, PackageInfo{
				Name:      group + ":" + name,
				Version:   version,
				Installed: true,
			})

			count++
		}

		return nil
	})

	return packages, nil
}

// 检测C#
func (a *App) detectCSharp() LanguageInfo {
	info := LanguageInfo{
		Name:            "C# (.NET)",
		Installed:       false,
		DownloadURL:     "https://dotnet.microsoft.com/download",
		InstallTutorial: "https://dotnet.microsoft.com/learn/dotnet/hello-world-tutorial/intro",
		PackageManager:  "NuGet",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "Newtonsoft.Json",
				Description: ".NET的流行高性能JSON框架",
				Version:     "13.0.3",
			},
			{
				Name:        "Microsoft.EntityFrameworkCore",
				Description: "Entity Framework Core是现代对象数据库映射器，用于.NET",
				Version:     "7.0.8",
			},
			{
				Name:        "AutoMapper",
				Description: "对象到对象的映射器，简化对象之间的映射",
				Version:     "12.0.1",
			},
			{
				Name:        "Serilog",
				Description: ".NET的诊断日志记录库",
				Version:     "3.0.1",
			},
			{
				Name:        "xunit",
				Description: ".NET的免费、开源、社区关注的单元测试工具",
				Version:     "2.5.0",
			},
		},
	}

	if !commandExists("dotnet") {
		return info
	}

	output, err := executeCommandWithTimeout("dotnet", "--version")
	if err == nil {
		info.Installed = true
		info.Version = output

		// 列出全局安装的.NET工具
		packages, _ := a.listDotNetTools()
		info.Packages = packages
	}

	return info
}

// 列出.NET工具
func (a *App) listDotNetTools() ([]PackageInfo, error) {
	output, err := executeCommandWithTimeout("dotnet", "tool", "list", "--global")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(output, "\n")
	packages := []PackageInfo{}

	// 跳过前两行（标题行）
	for i := 2; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) >= 2 {
			packages = append(packages, PackageInfo{
				Name:      fields[0],
				Version:   fields[1],
				Installed: true,
			})
		}
	}

	return packages, nil
}

// 检测Ruby
func (a *App) detectRuby() LanguageInfo {
	info := LanguageInfo{
		Name:            "Ruby",
		Installed:       false,
		DownloadURL:     "https://www.ruby-lang.org/en/downloads/",
		InstallTutorial: "https://www.ruby-lang.org/en/documentation/installation/",
		PackageManager:  "gem",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "rails",
				Description: "Ruby on Rails是一个全栈Web应用程序框架",
				Version:     "7.0.5",
			},
			{
				Name:        "sinatra",
				Description: "创建Web应用程序的DSL",
				Version:     "3.0.6",
			},
			{
				Name:        "rspec",
				Description: "Ruby的行为驱动开发框架",
				Version:     "3.12.0",
			},
			{
				Name:        "puma",
				Description: "Ruby的现代并发Web服务器",
				Version:     "6.3.0",
			},
			{
				Name:        "devise",
				Description: "Rails的灵活认证解决方案",
				Version:     "4.9.2",
			},
		},
	}

	if !commandExists("ruby") {
		return info
	}

	output, err := executeCommandWithTimeout("ruby", "--version")
	if err == nil {
		info.Installed = true
		info.Version = output

		// 检查gem是否安装
		if !commandExists("gem") {
			info.MissingDeps = append(info.MissingDeps, "RubyGems")
		} else {
			// 列出已安装的gem
			packages, _ := a.listGems()
			info.Packages = packages
		}
	}

	return info
}

// 列出Ruby gems
func (a *App) listGems() ([]PackageInfo, error) {
	output, err := executeCommandWithTimeout("gem", "list", "--local")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(output, "\n")
	packages := []PackageInfo{}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")
		if len(parts) >= 2 {
			name := parts[0]
			// 版本通常是 (x.y.z) 格式
			versionPart := strings.Join(parts[1:], " ")
			version := strings.Trim(versionPart, "()")

			packages = append(packages, PackageInfo{
				Name:      name,
				Version:   version,
				Installed: true,
			})
		}
	}

	return packages, nil
}

// 检测PHP
func (a *App) detectPHP() LanguageInfo {
	info := LanguageInfo{
		Name:            "PHP",
		Installed:       false,
		DownloadURL:     "https://www.php.net/downloads.php",
		InstallTutorial: "https://www.php.net/manual/en/install.php",
		PackageManager:  "Composer",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "laravel/framework",
				Description: "Laravel框架",
				Version:     "v10.13.5",
			},
			{
				Name:        "symfony/symfony",
				Description: "Symfony框架",
				Version:     "v6.3.0",
			},
			{
				Name:        "guzzlehttp/guzzle",
				Description: "PHP的HTTP客户端",
				Version:     "7.7.0",
			},
			{
				Name:        "phpunit/phpunit",
				Description: "PHP的单元测试框架",
				Version:     "10.2.2",
			},
			{
				Name:        "doctrine/orm",
				Description: "PHP的对象关系映射器",
				Version:     "2.15.1",
			},
		},
	}

	if !commandExists("php") {
		return info
	}

	output, err := executeCommandWithTimeout("php", "--version")
	if err == nil {
		info.Installed = true
		info.Version = output

		// 检查Composer是否安装
		if !commandExists("composer") {
			info.MissingDeps = append(info.MissingDeps, "Composer")
		} else {
			// Composer已安装，但列出全局包需要特定命令
			packages, _ := a.listComposerPackages()
			info.Packages = packages
		}
	}

	return info
}

// 列出Composer包
func (a *App) listComposerPackages() ([]PackageInfo, error) {
	// 这个命令会列出全局安装的包，但需要在有composer.json的目录中执行
	// 这里简化处理，实际应用可能需要更复杂的逻辑
	output, err := executeCommandWithTimeout("composer", "global", "show", "--format=json")
	if err != nil {
		return nil, err
	}

	var result struct {
		Installed []struct {
			Name    string `json:"name"`
			Version string `json:"version"`
		} `json:"installed"`
	}

	if err := json.Unmarshal([]byte(output), &result); err != nil {
		return nil, err
	}

	packages := make([]PackageInfo, 0, len(result.Installed))
	for _, pkg := range result.Installed {
		packages = append(packages, PackageInfo{
			Name:      pkg.Name,
			Version:   pkg.Version,
			Installed: true,
		})
	}

	return packages, nil
}

// 检测Rust
func (a *App) detectRust() LanguageInfo {
	info := LanguageInfo{
		Name:            "Rust",
		Installed:       false,
		DownloadURL:     "https://www.rust-lang.org/tools/install",
		InstallTutorial: "https://doc.rust-lang.org/book/ch01-01-installation.html",
		PackageManager:  "cargo",
		RecommendedPkgs: []PackageInfo{
			{
				Name:        "tokio",
				Description: "Rust的异步运行时",
				Version:     "1.29.1",
			},
			{
				Name:        "serde",
				Description: "Rust的序列化框架",
				Version:     "1.0.164",
			},
			{
				Name:        "actix-web",
				Description: "Rust的强大、实用、快速的Web框架",
				Version:     "4.3.1",
			},
			{
				Name:        "clap",
				Description: "Rust的命令行参数解析器",
				Version:     "4.3.10",
			},
			{
				Name:        "diesel",
				Description: "Rust的安全、可扩展的ORM和查询构建器",
				Version:     "2.1.0",
			},
		},
	}

	if !commandExists("rustc") {
		return info
	}

	output, err := executeCommandWithTimeout("rustc", "--version")
	if err == nil {
		info.Installed = true
		info.Version = output

		// 检查Cargo是否安装
		if !commandExists("cargo") {
			info.MissingDeps = append(info.MissingDeps, "Cargo")
		} else {
			// 列出已安装的crate
			packages, _ := a.listCrates()
			info.Packages = packages
		}
	}

	return info
}

// 列出Rust crates
func (a *App) listCrates() ([]PackageInfo, error) {
	// 这个命令会列出已安装的crate，但需要在Rust项目中执行
	// 这里我们列出已安装的工具包
	output, err := executeCommandWithTimeout("cargo", "install", "--list")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(output, "\n")
	packages := []PackageInfo{}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "cargo-") || strings.Contains(line, "v") {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				name := parts[0]
				version := strings.TrimPrefix(parts[1], "v")

				packages = append(packages, PackageInfo{
					Name:      name,
					Version:   version,
					Installed: true,
				})
			}
		}
	}

	return packages, nil
}
