package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"syscall"
	"time"
)

// PackageInfo 存储包信息
type PackageInfo struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Installed   bool   `json:"installed"`
	InstallLink string `json:"installLink"`
	DownloadURL string `json:"downloadUrl"`
}

// LanguageInfo 存储编程语言的信息
type LanguageInfo struct {
	Name            string        `json:"name"`
	Installed       bool          `json:"installed"`
	Version         string        `json:"version"`
	MissingDeps     []string      `json:"missingDeps"`
	DownloadURL     string        `json:"downloadUrl"`
	InstallTutorial string        `json:"installTutorial"`
	PackageManager  string        `json:"packageManager"`
	Packages        []PackageInfo `json:"packages"`
	Extensions      []PackageInfo `json:"extensions"`
	RecommendedPkgs []PackageInfo `json:"recommendedPkgs"`
}

// PackageTutorial 存储包管理器教程
type PackageTutorial struct {
	Name        string `json:"name"`
	InstallCmd  string `json:"installCmd"`
	SearchCmd   string `json:"searchCmd"`
	UpdateCmd   string `json:"updateCmd"`
	TutorialURL string `json:"tutorialUrl"`
}

// App 应用程序结构体
type App struct {
}

// NewApp 创建一个新的App实例
func NewApp() *App {
	return &App{}
}

// 应用程序启动时调用
func (a *App) startup(ctx context.Context) {
	fmt.Println("应用程序已启动")
}

// TestFunction 测试函数，用于验证前后端通信
func (a *App) TestFunction() string {
	fmt.Println("TestFunction被调用")
	return "测试成功！后端正常工作。"
}

// SystemInfo 存储系统信息
type SystemInfo struct {
	OS   string `json:"os"`
	Arch string `json:"arch"`
	CPUs string `json:"cpus"`
}

// GetSystemInfo 获取系统信息
func (a *App) GetSystemInfo() SystemInfo {
	fmt.Println("GetSystemInfo被调用")

	info := SystemInfo{
		OS:   runtime.GOOS,
		Arch: runtime.GOARCH,
		CPUs: fmt.Sprintf("%d", runtime.NumCPU()),
	}

	return info
}

// 添加超时上下文和错误处理辅助函数
func executeCommandWithTimeout(name string, args ...string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, name, args...)

	// 在Windows系统上隐藏命令窗口
	if runtime.GOOS == "windows" {
		cmd.SysProcAttr = &syscall.SysProcAttr{
			HideWindow:    true,
			CreationFlags: 0x08000000, // CREATE_NO_WINDOW
		}
	}

	output, err := cmd.CombinedOutput()

	if ctx.Err() == context.DeadlineExceeded {
		return "", fmt.Errorf("命令执行超时: %s %v", name, args)
	}

	return strings.TrimSpace(string(output)), err
}

// 安全地检查命令是否存在
func commandExists(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

// DetectLanguages 并行检测系统中安装的编程语言
func (a *App) DetectLanguages() []LanguageInfo {
	fmt.Println("开始检测编程语言...")

	// 定义所有要检测的语言函数
	detectors := []func() LanguageInfo{
		a.detectGo,
		a.detectPython,
		a.detectNode,
		a.detectJava,
		a.detectCSharp,
		a.detectRubyWithPackages,
		a.detectPHP,
		a.detectRust,
		a.detectCppWithPackages,
		a.detectSwiftWithPackages,
		a.detectKotlinWithPackages,
		a.detectDartWithPackages,
		a.detectTypeScriptWithPackages,
		a.detectPerlWithPackages,
		a.detectLuaWithPackages,
		a.detectRWithPackages,
		a.detectMatlabWithPackages,
		a.detectScalaWithPackages,
		a.detectHaskellWithPackages,
		a.detectObjectiveCWithPackages,
		a.detectGroovyWithPackages,
		a.detectClojureWithPackages,
		a.detectElixirWithPackages,
		a.detectFSharpWithPackages,
		a.detectJuliaWithPackages,
		a.detectPrologWithPackages,
		a.detectAssemblyWithPackages,
		a.detectCOBOLWithPackages,
		a.detectFortranWithPackages,
		a.detectDelphiWithPackages,
		a.detectLispWithPackages,
		a.detectSchemeWithPackages,
		a.detectCrystalWithPackages,
		a.detectNimWithPackages,
		a.detectDWithPackages,
		a.detectVHDLWithPackages,
		a.detectErlangWithPackages,
		a.detectSmalltalkWithPackages,
		a.detectOCamlWithPackages,
		a.detectTclWithPackages,
		a.detectBashWithPackages,
		a.detectPowerShellWithPackages,
		a.detectVBAWithPackages,
		a.detectSQLWithPackages,
		a.detectHTMLWithPackages,
		a.detectApexWithPackages,
		a.detectSolidityWithPackages,
		a.detectWebAssemblyWithPackages,
		a.detectZigWithPackages,
		a.detectHaxeWithPackages,
		a.detectABAPWithPackages,
		a.detectActionScriptWithPackages,
		a.detectAPLWithPackages,
		a.detectBallerinaWithPackages,
		a.detectBASICWithPackages,
		a.detectBooWithPackages,
		a.detectCeylonWithPackages,
		a.detectCoffeeScriptWithPackages,
		a.detectElmWithPackages,
		a.detectHackWithPackages,
		a.detectJWithPackages,
		a.detectJythonWithPackages,
		a.detectLOLCODEWithPackages,
		a.detectPureScriptWithPackages,
		a.detectQSharpWithPackages,
		a.detectRedWithPackages,
		a.detectReScriptWithPackages,
		a.detectScratchWithPackages,
		a.detectValaWithPackages,
		a.detectXSLTWithPackages,
	}

	// 创建一个通道来接收检测结果
	resultChan := make(chan LanguageInfo, len(detectors))

	// 创建一个WaitGroup来等待所有goroutine完成
	var wg sync.WaitGroup

	// 设置最大并发数，避免系统负载过高
	maxConcurrency := 10
	semaphore := make(chan struct{}, maxConcurrency)

	// 启动goroutine来检测每种语言
	for _, detector := range detectors {
		wg.Add(1)
		go func(detect func() LanguageInfo) {
			defer wg.Done()

			// 获取信号量，限制并发数
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			// 执行检测并发送结果到通道
			result := detect()
			resultChan <- result

			// 打印检测进度信息
			fmt.Printf("已检测: %s\n", result.Name)
		}(detector)
	}

	// 启动一个goroutine等待所有检测完成并关闭结果通道
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// 从通道中收集结果
	languages := []LanguageInfo{}
	for lang := range resultChan {
		languages = append(languages, lang)
	}

	fmt.Printf("已完成所有语言检测，共检测 %d 种语言\n", len(languages))
	return languages
}

// GetPackageTutorials 获取包管理器教程
func (a *App) GetPackageTutorials() []PackageTutorial {
	tutorials := []PackageTutorial{
		{
			Name:        "npm (Node.js)",
			InstallCmd:  "npm install [包名]",
			SearchCmd:   "npm search [关键词]",
			UpdateCmd:   "npm update [包名]",
			TutorialURL: "https://docs.npmjs.com/",
		},
		{
			Name:        "pip (Python)",
			InstallCmd:  "pip install [包名]",
			SearchCmd:   "pip search [关键词]",
			UpdateCmd:   "pip install --upgrade [包名]",
			TutorialURL: "https://pip.pypa.io/en/stable/",
		},
		{
			Name:        "go (Go)",
			InstallCmd:  "go get [包名]",
			SearchCmd:   "go list -m -versions [包名]",
			UpdateCmd:   "go get -u [包名]",
			TutorialURL: "https://go.dev/doc/modules/managing-dependencies",
		},
		{
			Name:        "maven (Java)",
			InstallCmd:  "mvn install",
			SearchCmd:   "mvn dependency:get -Dartifact=[groupId]:[artifactId]:[version]",
			UpdateCmd:   "mvn versions:use-latest-versions",
			TutorialURL: "https://maven.apache.org/guides/getting-started/",
		},
		{
			Name:        "nuget (.NET)",
			InstallCmd:  "dotnet add package [包名]",
			SearchCmd:   "dotnet nuget search [关键词]",
			UpdateCmd:   "dotnet add package [包名] --version [版本号]",
			TutorialURL: "https://learn.microsoft.com/en-us/nuget/",
		},
		{
			Name:        "gem (Ruby)",
			InstallCmd:  "gem install [包名]",
			SearchCmd:   "gem search [关键词]",
			UpdateCmd:   "gem update [包名]",
			TutorialURL: "https://guides.rubygems.org/",
		},
		{
			Name:        "composer (PHP)",
			InstallCmd:  "composer require [包名]",
			SearchCmd:   "composer search [关键词]",
			UpdateCmd:   "composer update [包名]",
			TutorialURL: "https://getcomposer.org/doc/",
		},
		{
			Name:        "cargo (Rust)",
			InstallCmd:  "cargo add [包名]",
			SearchCmd:   "cargo search [关键词]",
			UpdateCmd:   "cargo update [包名]",
			TutorialURL: "https://doc.rust-lang.org/cargo/",
		},
		{
			Name:        "elm (Elm)",
			InstallCmd:  "elm install [包名]",
			SearchCmd:   "elm search [关键词]",
			UpdateCmd:   "elm install [包名]@[版本]",
			TutorialURL: "https://guide.elm-lang.org/install/elm.html",
		},
		{
			Name:        "spago (PureScript)",
			InstallCmd:  "spago install [包名]",
			SearchCmd:   "spago search [关键词]",
			UpdateCmd:   "spago upgrade-set",
			TutorialURL: "https://github.com/purescript/spago",
		},
		{
			Name:        "pacman (J)",
			InstallCmd:  "load 'pacman'",
			SearchCmd:   "load 'pacman' \n require 'find'",
			UpdateCmd:   "load 'pacman' \n 'update' jpkg '*'",
			TutorialURL: "https://code.jsoftware.com/wiki/Pacman",
		},
		{
			Name:        "Pkg (Julia)",
			InstallCmd:  "using Pkg; Pkg.add(\"[包名]\")",
			SearchCmd:   "using Pkg; Pkg.add(PackageSpec(name=\"[包名]\"))",
			UpdateCmd:   "using Pkg; Pkg.update(\"[包名]\")",
			TutorialURL: "https://pkgdocs.julialang.org/v1/",
		},
		{
			Name:        "nimble (Nim)",
			InstallCmd:  "nimble install [包名]",
			SearchCmd:   "nimble search [关键词]",
			UpdateCmd:   "nimble install [包名]@#head",
			TutorialURL: "https://github.com/nim-lang/nimble#readme",
		},
		{
			Name:        "shards (Crystal)",
			InstallCmd:  "shards install",
			SearchCmd:   "查找Crystal包: https://crystalshards.org/",
			UpdateCmd:   "shards update",
			TutorialURL: "https://crystal-lang.org/reference/man/shards/",
		},
		{
			Name:        "zigmod (Zig)",
			InstallCmd:  "zigmod fetch",
			SearchCmd:   "查找Zig包: https://astrolabe.pm/",
			UpdateCmd:   "zigmod update",
			TutorialURL: "https://github.com/nektro/zigmod",
		},
		{
			Name:        "dub (D语言)",
			InstallCmd:  "dub add [包名]",
			SearchCmd:   "dub search [关键词]",
			UpdateCmd:   "dub upgrade",
			TutorialURL: "https://dub.pm/getting_started",
		},
	}
	return tutorials
}

// AIProvider 存储AI提供商信息
type AIProvider struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	EndpointURL  string `json:"endpointUrl"`
	DocumentURL  string `json:"documentUrl"`
	ApiKeyHeader string `json:"apiKeyHeader"`
	ApiKeyPrefix string `json:"apiKeyPrefix"`
}

// AIConfig 存储AI配置信息
type AIConfig struct {
	SelectedProviderID string `json:"selectedProviderId"`
	APIKey             string `json:"apiKey"`
	CustomEndpoint     string `json:"customEndpoint"`
}

// AIResponse 存储AI响应信息
type AIResponse struct {
	Content  string `json:"content"`
	Error    string `json:"error"`
	Success  bool   `json:"success"`
	Provider string `json:"provider"`
}

// ThemeConfig 存储主题配置
type ThemeConfig struct {
	Theme string `json:"theme"`
}

// LanguageConfig 存储语言配置
type LanguageConfig struct {
	Language string `json:"language"`
}

// GetLanguageConfig 获取语言配置
func (a *App) GetLanguageConfig() LanguageConfig {
	configPath := a.getLanguageConfigPath()

	// 如果配置文件不存在，返回默认配置
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return LanguageConfig{
			Language: "zh",
		}
	}

	// 读取配置文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		return LanguageConfig{
			Language: "zh",
		}
	}

	// 解析配置
	var config LanguageConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return LanguageConfig{
			Language: "zh",
		}
	}

	// 验证语言是否有效
	if config.Language != "zh" && config.Language != "en" && config.Language != "ru" {
		config.Language = "zh" // 如果无效，设置为默认中文
	}

	return config
}

// SaveLanguageConfig 保存语言配置
func (a *App) SaveLanguageConfig(config LanguageConfig) error {
	configPath := a.getLanguageConfigPath()
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}

// getLanguageConfigPath 获取语言配置文件路径
func (a *App) getLanguageConfigPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "language_config.json"
	}
	configDir := filepath.Join(homeDir, ".networ_tester")

	// 确保目录存在
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		os.MkdirAll(configDir, 0755)
	}

	return filepath.Join(configDir, "language_config.json")
}

// GetThemeConfig 获取主题配置
func (a *App) GetThemeConfig() ThemeConfig {
	configPath := a.getThemeConfigPath()

	// 如果配置文件不存在，返回默认配置
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return ThemeConfig{
			Theme: "default",
		}
	}

	// 读取配置文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		return ThemeConfig{
			Theme: "default",
		}
	}

	// 解析配置
	var config ThemeConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return ThemeConfig{
			Theme: "default",
		}
	}

	return config
}

// SaveThemeConfig 保存主题配置
func (a *App) SaveThemeConfig(config ThemeConfig) error {
	configPath := a.getThemeConfigPath()
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}

// getThemeConfigPath 获取主题配置文件路径
func (a *App) getThemeConfigPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "theme_config.json"
	}
	configDir := filepath.Join(homeDir, ".networ_tester")

	// 确保目录存在
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		os.MkdirAll(configDir, 0755)
	}

	return filepath.Join(configDir, "theme_config.json")
}

// GetAIProviders 获取支持的AI提供商列表
func (a *App) GetAIProviders() []AIProvider {
	providers := []AIProvider{
		{
			ID:           "openai",
			Name:         "OpenAI (ChatGPT)",
			EndpointURL:  "https://api.openai.com/v1/chat/completions",
			DocumentURL:  "https://platform.openai.com/docs/api-reference",
			ApiKeyHeader: "Authorization",
			ApiKeyPrefix: "Bearer ",
		},
		{
			ID:           "azure_openai",
			Name:         "Azure OpenAI",
			EndpointURL:  "https://YOUR_RESOURCE_NAME.openai.azure.com/openai/deployments/YOUR_DEPLOYMENT_NAME/chat/completions?api-version=2023-05-15",
			DocumentURL:  "https://learn.microsoft.com/zh-cn/azure/ai-services/openai/reference",
			ApiKeyHeader: "api-key",
			ApiKeyPrefix: "",
		},
		{
			ID:           "anthropic",
			Name:         "Anthropic (Claude)",
			EndpointURL:  "https://api.anthropic.com/v1/messages",
			DocumentURL:  "https://docs.anthropic.com/claude/reference/getting-started-with-the-api",
			ApiKeyHeader: "x-api-key",
			ApiKeyPrefix: "",
		},
		{
			ID:           "baidu",
			Name:         "百度文心一言",
			EndpointURL:  "https://aip.baidubce.com/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/completions",
			DocumentURL:  "https://cloud.baidu.com/doc/WENXINWORKSHOP/s/jlil56u11",
			ApiKeyHeader: "Authorization",
			ApiKeyPrefix: "",
		},
		{
			ID:           "aliyun",
			Name:         "阿里通义千问",
			EndpointURL:  "https://dashscope.aliyuncs.com/api/v1/services/aigc/text-generation/generation",
			DocumentURL:  "https://help.aliyun.com/document_detail/2400395.html",
			ApiKeyHeader: "Authorization",
			ApiKeyPrefix: "Bearer ",
		},
		{
			ID:           "custom",
			Name:         "自定义API",
			EndpointURL:  "",
			DocumentURL:  "",
			ApiKeyHeader: "Authorization",
			ApiKeyPrefix: "Bearer ",
		},
	}
	return providers
}

// 获取配置文件路径
func (a *App) getAIConfigPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "ai_config.json"
	}
	configDir := filepath.Join(homeDir, ".networ_tester")

	// 确保目录存在
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		os.MkdirAll(configDir, 0755)
	}

	return filepath.Join(configDir, "ai_config.json")
}

// SaveAIConfig 保存AI配置
func (a *App) SaveAIConfig(config AIConfig) error {
	configPath := a.getAIConfigPath()
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}

// GetAIConfig 获取AI配置
func (a *App) GetAIConfig() AIConfig {
	configPath := a.getAIConfigPath()

	// 如果配置文件不存在，返回默认配置
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return AIConfig{
			SelectedProviderID: "openai",
			APIKey:             "",
			CustomEndpoint:     "",
		}
	}

	// 读取配置文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		return AIConfig{
			SelectedProviderID: "openai",
			APIKey:             "",
			CustomEndpoint:     "",
		}
	}

	// 解析配置
	var config AIConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return AIConfig{
			SelectedProviderID: "openai",
			APIKey:             "",
			CustomEndpoint:     "",
		}
	}

	return config
}

// QueryAI 向AI发送查询
func (a *App) QueryAI(providerID string, apiKey string, customEndpoint string, query string) AIResponse {
	// 获取提供商信息
	var selectedProvider AIProvider
	found := false

	for _, provider := range a.GetAIProviders() {
		if provider.ID == providerID {
			selectedProvider = provider
			found = true
			break
		}
	}

	if !found {
		return AIResponse{
			Success: false,
			Error:   "未找到指定的AI提供商",
		}
	}

	// 使用自定义端点（如果有）
	endpoint := selectedProvider.EndpointURL
	if providerID == "custom" && customEndpoint != "" {
		endpoint = customEndpoint
	} else if providerID == "azure_openai" && customEndpoint != "" {
		endpoint = customEndpoint
	}

	// 准备请求
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	var reqBody []byte
	var err error

	// 根据不同提供商准备请求体
	switch providerID {
	case "openai":
		reqBody, err = json.Marshal(map[string]interface{}{
			"model": "gpt-3.5-turbo",
			"messages": []map[string]string{
				{
					"role":    "system",
					"content": "你是一个专注于编程开发问题的AI助手，请提供准确、具体的编程帮助。",
				},
				{
					"role":    "user",
					"content": query,
				},
			},
		})
	case "azure_openai":
		reqBody, err = json.Marshal(map[string]interface{}{
			"messages": []map[string]string{
				{
					"role":    "system",
					"content": "你是一个专注于编程开发问题的AI助手，请提供准确、具体的编程帮助。",
				},
				{
					"role":    "user",
					"content": query,
				},
			},
		})
	case "anthropic":
		reqBody, err = json.Marshal(map[string]interface{}{
			"model":      "claude-3-haiku-20240307",
			"max_tokens": 1000,
			"messages": []map[string]string{
				{
					"role":    "user",
					"content": query,
				},
			},
			"system": "你是一个专注于编程开发问题的AI助手，请提供准确、具体的编程帮助。",
		})
	case "baidu":
		reqBody, err = json.Marshal(map[string]interface{}{
			"messages": []map[string]string{
				{
					"role":    "user",
					"content": query,
				},
			},
		})
	case "aliyun":
		reqBody, err = json.Marshal(map[string]interface{}{
			"model": "qwen-max",
			"input": map[string]interface{}{
				"messages": []map[string]string{
					{
						"role":    "system",
						"content": "你是一个专注于编程开发问题的AI助手，请提供准确、具体的编程帮助。",
					},
					{
						"role":    "user",
						"content": query,
					},
				},
			},
		})
	case "custom":
		// 对于自定义API，使用OpenAI格式作为默认
		reqBody, err = json.Marshal(map[string]interface{}{
			"model": "gpt-3.5-turbo",
			"messages": []map[string]string{
				{
					"role":    "system",
					"content": "你是一个专注于编程开发问题的AI助手，请提供准确、具体的编程帮助。",
				},
				{
					"role":    "user",
					"content": query,
				},
			},
		})
	}

	if err != nil {
		return AIResponse{
			Success: false,
			Error:   "准备请求失败: " + err.Error(),
		}
	}

	// 创建请求
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(string(reqBody)))
	if err != nil {
		return AIResponse{
			Success: false,
			Error:   "创建请求失败: " + err.Error(),
		}
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 添加API密钥
	apiKeyValue := selectedProvider.ApiKeyPrefix + apiKey
	req.Header.Set(selectedProvider.ApiKeyHeader, apiKeyValue)

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return AIResponse{
			Success: false,
			Error:   "发送请求失败: " + err.Error(),
		}
	}
	defer resp.Body.Close()

	// 读取响应
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return AIResponse{
			Success: false,
			Error:   "读取响应失败: " + err.Error(),
		}
	}

	// 检查状态码
	if resp.StatusCode != http.StatusOK {
		return AIResponse{
			Success: false,
			Error:   fmt.Sprintf("API返回错误 (状态码: %d): %s", resp.StatusCode, string(respBody)),
		}
	}

	// 解析响应
	var content string

	// 根据不同提供商解析响应
	switch providerID {
	case "openai", "azure_openai":
		var openaiResp struct {
			Choices []struct {
				Message struct {
					Content string `json:"content"`
				} `json:"message"`
			} `json:"choices"`
		}
		if err := json.Unmarshal(respBody, &openaiResp); err != nil || len(openaiResp.Choices) == 0 {
			return AIResponse{
				Success: false,
				Error:   "解析响应失败: " + err.Error(),
			}
		}
		content = openaiResp.Choices[0].Message.Content
	case "anthropic":
		var anthropicResp struct {
			Content []struct {
				Text string `json:"text"`
			} `json:"content"`
		}
		if err := json.Unmarshal(respBody, &anthropicResp); err != nil || len(anthropicResp.Content) == 0 {
			return AIResponse{
				Success: false,
				Error:   "解析响应失败: " + err.Error(),
			}
		}
		content = anthropicResp.Content[0].Text
	case "baidu":
		var baiduResp struct {
			Result string `json:"result"`
		}
		if err := json.Unmarshal(respBody, &baiduResp); err != nil {
			return AIResponse{
				Success: false,
				Error:   "解析响应失败: " + err.Error(),
			}
		}
		content = baiduResp.Result
	case "aliyun":
		var aliyunResp struct {
			Output struct {
				Choices []struct {
					Message struct {
						Content string `json:"content"`
					} `json:"message"`
				} `json:"choices"`
			} `json:"output"`
		}
		if err := json.Unmarshal(respBody, &aliyunResp); err != nil || len(aliyunResp.Output.Choices) == 0 {
			return AIResponse{
				Success: false,
				Error:   "解析响应失败: " + err.Error(),
			}
		}
		content = aliyunResp.Output.Choices[0].Message.Content
	case "custom":
		// 尝试使用OpenAI格式解析
		var openaiResp struct {
			Choices []struct {
				Message struct {
					Content string `json:"content"`
				} `json:"message"`
			} `json:"choices"`
		}
		if err := json.Unmarshal(respBody, &openaiResp); err != nil || len(openaiResp.Choices) == 0 {
			// 如果失败，返回原始响应
			content = string(respBody)
		} else {
			content = openaiResp.Choices[0].Message.Content
		}
	}

	return AIResponse{
		Content:  content,
		Success:  true,
		Provider: selectedProvider.Name,
	}
}

// SearchPackage 搜索指定包管理器中的包
func (a *App) SearchPackage(packageManager string, packageName string) []PackageInfo {
	fmt.Printf("搜索包: %s (使用 %s)\n", packageName, packageManager)

	if packageName == "" {
		return []PackageInfo{}
	}

	var packages []PackageInfo
	var err error

	switch packageManager {
	case "npm":
		packages, err = a.searchNpmPackage(packageName)
	case "pip":
		packages, err = a.searchPipPackage(packageName)
	case "gem":
		packages, err = a.searchGemPackage(packageName)
	case "cargo":
		packages, err = a.searchCargoPackage(packageName)
	case "composer":
		packages, err = a.searchComposerPackage(packageName)
	case "nuget":
		packages, err = a.searchNuGetPackage(packageName)
	case "maven":
		packages, err = a.searchMavenPackage(packageName)
	case "go":
		packages, err = a.searchGoPackage(packageName)
	case "dub":
		packages, err = a.searchDubPackage(packageName)
	case "hex":
		packages, err = a.searchHexPackage(packageName)
	case "nimble":
		packages, err = a.searchNimblePackage(packageName)
	case "brew":
		packages, err = a.searchBrewPackage(packageName)
	default:
		return []PackageInfo{}
	}

	if err != nil {
		fmt.Printf("搜索包时出错: %v\n", err)
		return []PackageInfo{}
	}

	// 为包添加安装链接
	for i := range packages {
		a.addPackageLinks(&packages[i], packageManager)
	}

	// 尝试使用AI获取更多信息
	a.enrichPackageInfoWithAI(packages, packageManager, packageName)

	return packages
}

// addPackageLinks 为包添加安装链接
func (a *App) addPackageLinks(pkg *PackageInfo, packageManager string) {
	switch packageManager {
	case "npm":
		pkg.InstallLink = fmt.Sprintf("npm install %s", pkg.Name)
		pkg.DownloadURL = fmt.Sprintf("https://www.npmjs.com/package/%s", pkg.Name)
	case "pip":
		pkg.InstallLink = fmt.Sprintf("pip install %s", pkg.Name)
		pkg.DownloadURL = fmt.Sprintf("https://pypi.org/project/%s", pkg.Name)
	case "gem":
		pkg.InstallLink = fmt.Sprintf("gem install %s", pkg.Name)
		pkg.DownloadURL = fmt.Sprintf("https://rubygems.org/gems/%s", pkg.Name)
	case "cargo":
		pkg.InstallLink = fmt.Sprintf("cargo add %s", pkg.Name)
		pkg.DownloadURL = fmt.Sprintf("https://crates.io/crates/%s", pkg.Name)
	case "composer":
		pkg.InstallLink = fmt.Sprintf("composer require %s", pkg.Name)
		pkg.DownloadURL = fmt.Sprintf("https://packagist.org/packages/%s", pkg.Name)
	case "nuget":
		pkg.InstallLink = fmt.Sprintf("dotnet add package %s", pkg.Name)
		pkg.DownloadURL = fmt.Sprintf("https://www.nuget.org/packages/%s", pkg.Name)
	case "maven":
		parts := strings.Split(pkg.Name, ":")
		groupId := parts[0]
		artifactId := ""
		if len(parts) > 1 {
			artifactId = parts[1]
		}
		pkg.InstallLink = fmt.Sprintf("<dependency>\n  <groupId>%s</groupId>\n  <artifactId>%s</artifactId>\n  <version>%s</version>\n</dependency>", groupId, artifactId, pkg.Version)
		pkg.DownloadURL = fmt.Sprintf("https://mvnrepository.com/artifact/%s/%s", groupId, artifactId)
	case "go":
		pkg.InstallLink = fmt.Sprintf("go get %s", pkg.Name)
		pkg.DownloadURL = fmt.Sprintf("https://pkg.go.dev/%s", pkg.Name)
	case "dub":
		pkg.InstallLink = fmt.Sprintf("dub add %s", pkg.Name)
		pkg.DownloadURL = fmt.Sprintf("https://code.dlang.org/packages/%s", pkg.Name)
	case "hex":
		pkg.InstallLink = fmt.Sprintf("{:deps, [{:%s, \"~> %s\"}]}", pkg.Name, pkg.Version)
		pkg.DownloadURL = fmt.Sprintf("https://hex.pm/packages/%s", pkg.Name)
	case "nimble":
		pkg.InstallLink = fmt.Sprintf("nimble install %s", pkg.Name)
		pkg.DownloadURL = fmt.Sprintf("https://nimble.directory/pkg/%s", pkg.Name)
	case "brew":
		pkg.InstallLink = fmt.Sprintf("brew install %s", pkg.Name)
		pkg.DownloadURL = fmt.Sprintf("https://formulae.brew.sh/formula/%s", pkg.Name)
	}
}

// enrichPackageInfoWithAI 使用AI API丰富包信息
func (a *App) enrichPackageInfoWithAI(packages []PackageInfo, packageManager, packageName string) {
	// 只处理前5个包，避免过多的API调用
	if len(packages) > 5 {
		packages = packages[:5]
	}

	// 随机选择一个API
	apis := []struct {
		name   string
		apiKey string
		url    string
	}{
		{
			name:   "通义千问",
			apiKey: "sk-07ef4701031d41668beebb521e80eaf0",
			url:    "https://dashscope.aliyuncs.com/api/v1/services/aigc/text-generation/generation",
		},
		{
			name:   "DeepSeek",
			apiKey: "sk-0b2be14756fe4195a7bc2bcb78d19f8f",
			url:    "https://api.deepseek.com/v1/chat/completions",
		},
	}

	// 随机选择一个API
	apiIndex := time.Now().Unix() % int64(len(apis))
	selectedAPI := apis[apiIndex]

	for i, pkg := range packages {
		// 构建查询
		query := fmt.Sprintf("请提供%s包管理器中%s包的最新下载链接和安装命令，只返回JSON格式，包含downloadUrl和installCommand字段", packageManager, pkg.Name)

		// 调用AI API获取信息
		result := a.callPackageInfoAPI(selectedAPI.name, selectedAPI.apiKey, selectedAPI.url, query)
		if result != nil {
			// 尝试解析结果
			downloadURL, _ := result["downloadUrl"].(string)
			installCommand, _ := result["installCommand"].(string)

			// 如果有有效信息，则更新包
			if downloadURL != "" && downloadURL != "unknown" {
				packages[i].DownloadURL = downloadURL
			}
			if installCommand != "" && installCommand != "unknown" {
				packages[i].InstallLink = installCommand
			}
		}
	}
}

// callPackageInfoAPI 调用AI API获取包信息
func (a *App) callPackageInfoAPI(apiName, apiKey, apiURL, query string) map[string]interface{} {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	var reqBody []byte
	var err error

	// 根据不同API构建请求体
	switch apiName {
	case "通义千问":
		reqBody, err = json.Marshal(map[string]interface{}{
			"model": "qwen-max",
			"input": map[string]interface{}{
				"messages": []map[string]string{
					{
						"role":    "system",
						"content": "你是一个编程助手，请提供准确的包信息。只返回JSON格式，包含downloadUrl和installCommand字段。",
					},
					{
						"role":    "user",
						"content": query,
					},
				},
			},
		})
	case "DeepSeek":
		reqBody, err = json.Marshal(map[string]interface{}{
			"model": "deepseek-chat",
			"messages": []map[string]string{
				{
					"role":    "system",
					"content": "你是一个编程助手，请提供准确的包信息。只返回JSON格式，包含downloadUrl和installCommand字段。",
				},
				{
					"role":    "user",
					"content": query,
				},
			},
		})
	default:
		return nil
	}

	if err != nil {
		fmt.Printf("准备AI请求失败: %v\n", err)
		return nil
	}

	// 创建请求
	req, err := http.NewRequest("POST", apiURL, strings.NewReader(string(reqBody)))
	if err != nil {
		fmt.Printf("创建AI请求失败: %v\n", err)
		return nil
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 添加API密钥
	switch apiName {
	case "通义千问":
		req.Header.Set("Authorization", "Bearer "+apiKey)
	case "DeepSeek":
		req.Header.Set("Authorization", "Bearer "+apiKey)
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("发送AI请求失败: %v\n", err)
		return nil
	}
	defer resp.Body.Close()

	// 读取响应
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取AI响应失败: %v\n", err)
		return nil
	}

	// 检查状态码
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("AI API返回错误 (状态码: %d): %s\n", resp.StatusCode, string(respBody))
		return nil
	}

	// 解析响应
	var content string

	// 根据不同API解析响应
	switch apiName {
	case "通义千问":
		var aliyunResp struct {
			Output struct {
				Choices []struct {
					Message struct {
						Content string `json:"content"`
					} `json:"message"`
				} `json:"choices"`
			} `json:"output"`
		}
		if err := json.Unmarshal(respBody, &aliyunResp); err != nil || len(aliyunResp.Output.Choices) == 0 {
			fmt.Printf("解析通义千问响应失败: %v\n", err)
			return nil
		}
		content = aliyunResp.Output.Choices[0].Message.Content
	case "DeepSeek":
		var deepseekResp struct {
			Choices []struct {
				Message struct {
					Content string `json:"content"`
				} `json:"message"`
			} `json:"choices"`
		}
		if err := json.Unmarshal(respBody, &deepseekResp); err != nil || len(deepseekResp.Choices) == 0 {
			fmt.Printf("解析DeepSeek响应失败: %v\n", err)
			return nil
		}
		content = deepseekResp.Choices[0].Message.Content
	}

	// 尝试从内容中提取JSON
	jsonStart := strings.Index(content, "{")
	jsonEnd := strings.LastIndex(content, "}")
	if jsonStart >= 0 && jsonEnd > jsonStart {
		jsonStr := content[jsonStart : jsonEnd+1]
		var result map[string]interface{}
		if err := json.Unmarshal([]byte(jsonStr), &result); err == nil {
			return result
		}
	}

	return nil
}

// searchNpmPackage 搜索npm包
func (a *App) searchNpmPackage(packageName string) ([]PackageInfo, error) {
	output, err := executeCommandWithTimeout("npm", "search", "--json", packageName)
	if err != nil {
		return nil, err
	}

	var results []map[string]interface{}
	if err := json.Unmarshal([]byte(output), &results); err != nil {
		// 尝试解析非JSON输出
		lines := strings.Split(output, "\n")
		packages := []PackageInfo{}

		for _, line := range lines {
			if strings.Contains(line, packageName) {
				fields := strings.Fields(line)
				if len(fields) >= 2 {
					packages = append(packages, PackageInfo{
						Name:        fields[0],
						Version:     fields[1],
						Description: strings.Join(fields[2:], " "),
					})
				}
			}
		}

		return packages, nil
	}

	packages := []PackageInfo{}
	for _, result := range results {
		name, _ := result["name"].(string)
		version, _ := result["version"].(string)
		description, _ := result["description"].(string)

		packages = append(packages, PackageInfo{
			Name:        name,
			Version:     version,
			Description: description,
		})

		// 限制结果数量
		if len(packages) >= 20 {
			break
		}
	}

	return packages, nil
}

// searchPipPackage 搜索pip包
func (a *App) searchPipPackage(packageName string) ([]PackageInfo, error) {
	// pip search已被弃用，使用PyPI API
	url := fmt.Sprintf("https://pypi.org/pypi/%s/json", packageName)
	resp, err := http.Get(url)
	if err != nil {
		// 尝试使用pip命令
		output, err := executeCommandWithTimeout("pip", "search", packageName)
		if err != nil {
			return nil, err
		}

		lines := strings.Split(output, "\n")
		packages := []PackageInfo{}

		for _, line := range lines {
			if strings.Contains(strings.ToLower(line), strings.ToLower(packageName)) {
				parts := strings.SplitN(line, " - ", 2)
				if len(parts) == 2 {
					nameVer := strings.TrimSpace(parts[0])
					desc := strings.TrimSpace(parts[1])

					nameParts := strings.Split(nameVer, " ")
					name := nameParts[0]
					version := ""
					if len(nameParts) > 1 {
						version = strings.Trim(nameParts[1], "()")
					}

					packages = append(packages, PackageInfo{
						Name:        name,
						Version:     version,
						Description: desc,
					})

					if len(packages) >= 20 {
						break
					}
				}
			}
		}

		return packages, nil
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	info, ok := result["info"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("无效的PyPI响应格式")
	}

	name, _ := info["name"].(string)
	version, _ := info["version"].(string)
	summary, _ := info["summary"].(string)

	return []PackageInfo{
		{
			Name:        name,
			Version:     version,
			Description: summary,
		},
	}, nil
}

// searchGemPackage 搜索Ruby gem包
func (a *App) searchGemPackage(packageName string) ([]PackageInfo, error) {
	output, err := executeCommandWithTimeout("gem", "search", "-r", packageName)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(output, "\n")
	packages := []PackageInfo{}

	for _, line := range lines {
		if strings.Contains(line, packageName) {
			parts := strings.SplitN(line, " ", 2)
			if len(parts) == 2 {
				name := strings.TrimSpace(parts[0])
				versionDesc := strings.TrimSpace(parts[1])

				version := ""
				description := ""

				// 提取版本号
				versionMatch := regexp.MustCompile(`\((.*?)\)`).FindStringSubmatch(versionDesc)
				if len(versionMatch) > 1 {
					version = versionMatch[1]
				}

				// 提取描述
				descMatch := regexp.MustCompile(`\) (.*)`).FindStringSubmatch(versionDesc)
				if len(descMatch) > 1 {
					description = descMatch[1]
				}

				packages = append(packages, PackageInfo{
					Name:        name,
					Version:     version,
					Description: description,
				})

				if len(packages) >= 20 {
					break
				}
			}
		}
	}

	return packages, nil
}

// searchCargoPackage 搜索Rust cargo包
func (a *App) searchCargoPackage(packageName string) ([]PackageInfo, error) {
	output, err := executeCommandWithTimeout("cargo", "search", packageName, "--limit", "20")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(output, "\n")
	packages := []PackageInfo{}

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		// 匹配 "name = "0.1.0" # 描述" 格式
		parts := strings.SplitN(line, " = ", 2)
		if len(parts) == 2 {
			name := strings.TrimSpace(parts[0])

			// 提取版本和描述
			versionDesc := strings.TrimSpace(parts[1])
			version := ""
			description := ""

			// 提取版本
			versionMatch := regexp.MustCompile(`"([^"]+)"`).FindStringSubmatch(versionDesc)
			if len(versionMatch) > 1 {
				version = versionMatch[1]
			}

			// 提取描述
			descMatch := regexp.MustCompile(`#\s*(.*)`).FindStringSubmatch(versionDesc)
			if len(descMatch) > 1 {
				description = descMatch[1]
			}

			packages = append(packages, PackageInfo{
				Name:        name,
				Version:     version,
				Description: description,
			})
		}
	}

	return packages, nil
}

// searchComposerPackage 搜索PHP Composer包
func (a *App) searchComposerPackage(packageName string) ([]PackageInfo, error) {
	output, err := executeCommandWithTimeout("composer", "search", packageName)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(output, "\n")
	packages := []PackageInfo{}

	for _, line := range lines {
		if strings.Contains(line, packageName) {
			parts := strings.SplitN(line, " ", 2)
			if len(parts) == 2 {
				name := strings.TrimSpace(parts[0])
				description := strings.TrimSpace(parts[1])

				packages = append(packages, PackageInfo{
					Name:        name,
					Description: description,
				})

				if len(packages) >= 20 {
					break
				}
			}
		}
	}

	return packages, nil
}

// searchNuGetPackage 搜索.NET NuGet包
func (a *App) searchNuGetPackage(packageName string) ([]PackageInfo, error) {
	output, err := executeCommandWithTimeout("dotnet", "nuget", "search", packageName)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(output, "\n")
	packages := []PackageInfo{}
	inResultSection := false

	for _, line := range lines {
		// 跳过标题行
		if strings.Contains(line, "------") {
			inResultSection = true
			continue
		}

		if inResultSection && strings.TrimSpace(line) != "" {
			fields := strings.Fields(line)
			if len(fields) >= 3 {
				name := fields[0]
				version := fields[1]
				description := strings.Join(fields[2:], " ")

				packages = append(packages, PackageInfo{
					Name:        name,
					Version:     version,
					Description: description,
				})

				if len(packages) >= 20 {
					break
				}
			}
		}
	}

	return packages, nil
}

// searchMavenPackage 搜索Java Maven包
func (a *App) searchMavenPackage(packageName string) ([]PackageInfo, error) {
	// 使用Maven Central API
	url := fmt.Sprintf("https://search.maven.org/solrsearch/select?q=%s&rows=20&wt=json", packageName)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	response, ok := result["response"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("无效的Maven响应格式")
	}

	docs, ok := response["docs"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("无效的Maven响应文档格式")
	}

	packages := []PackageInfo{}
	for _, doc := range docs {
		docMap, ok := doc.(map[string]interface{})
		if !ok {
			continue
		}

		groupId, _ := docMap["g"].(string)
		artifactId, _ := docMap["a"].(string)
		version, _ := docMap["latestVersion"].(string)

		name := groupId + ":" + artifactId

		packages = append(packages, PackageInfo{
			Name:        name,
			Version:     version,
			Description: "",
		})
	}

	return packages, nil
}

// searchGoPackage 搜索Go包
func (a *App) searchGoPackage(packageName string) ([]PackageInfo, error) {
	// 使用pkg.go.dev API
	url := fmt.Sprintf("https://pkg.go.dev/search?q=%s", packageName)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 简单解析HTML响应
	content := string(body)

	// 使用正则表达式提取包信息
	packagePattern := regexp.MustCompile(`<a class="[^"]*" href="/([^"]+)"[^>]*>([^<]+)</a>`)
	matches := packagePattern.FindAllStringSubmatch(content, -1)

	packages := []PackageInfo{}
	for _, match := range matches {
		if len(match) >= 3 {
			path := match[1]
			name := match[2]

			packages = append(packages, PackageInfo{
				Name:        name,
				Version:     "",
				Description: path,
			})

			if len(packages) >= 20 {
				break
			}
		}
	}

	return packages, nil
}

// searchDubPackage 搜索D语言的Dub包
func (a *App) searchDubPackage(packageName string) ([]PackageInfo, error) {
	output, err := executeCommandWithTimeout("dub", "search", packageName)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(output, "\n")
	packages := []PackageInfo{}

	for _, line := range lines {
		if strings.Contains(line, packageName) {
			parts := strings.SplitN(line, " - ", 2)
			if len(parts) == 2 {
				nameVer := strings.TrimSpace(parts[0])
				desc := strings.TrimSpace(parts[1])

				nameParts := strings.Split(nameVer, " ")
				name := nameParts[0]
				version := ""
				if len(nameParts) > 1 {
					version = strings.Trim(nameParts[1], "()")
				}

				packages = append(packages, PackageInfo{
					Name:        name,
					Version:     version,
					Description: desc,
				})

				if len(packages) >= 20 {
					break
				}
			}
		}
	}

	return packages, nil
}

// searchHexPackage 搜索Elixir的Hex包
func (a *App) searchHexPackage(packageName string) ([]PackageInfo, error) {
	output, err := executeCommandWithTimeout("mix", "hex.search", packageName)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(output, "\n")
	packages := []PackageInfo{}

	for _, line := range lines {
		if strings.Contains(line, packageName) {
			parts := strings.SplitN(line, " - ", 2)
			if len(parts) == 2 {
				name := strings.TrimSpace(parts[0])
				desc := strings.TrimSpace(parts[1])

				packages = append(packages, PackageInfo{
					Name:        name,
					Description: desc,
				})

				if len(packages) >= 20 {
					break
				}
			}
		}
	}

	return packages, nil
}

// searchNimblePackage 搜索Nim的Nimble包
func (a *App) searchNimblePackage(packageName string) ([]PackageInfo, error) {
	output, err := executeCommandWithTimeout("nimble", "search", packageName)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(output, "\n")
	packages := []PackageInfo{}

	for _, line := range lines {
		if strings.Contains(line, packageName) {
			parts := strings.SplitN(line, " - ", 2)
			if len(parts) == 2 {
				nameVer := strings.TrimSpace(parts[0])
				desc := strings.TrimSpace(parts[1])

				nameParts := strings.Split(nameVer, " ")
				name := nameParts[0]
				version := ""
				if len(nameParts) > 1 {
					version = strings.Join(nameParts[1:], " ")
					version = strings.Trim(version, "()")
				}

				packages = append(packages, PackageInfo{
					Name:        name,
					Version:     version,
					Description: desc,
				})

				if len(packages) >= 20 {
					break
				}
			}
		}
	}

	return packages, nil
}

// searchBrewPackage 搜索Homebrew包
func (a *App) searchBrewPackage(packageName string) ([]PackageInfo, error) {
	output, err := executeCommandWithTimeout("brew", "search", packageName)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(output, "\n")
	packages := []PackageInfo{}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && !strings.HasPrefix(line, "==>") {
			packages = append(packages, PackageInfo{
				Name: line,
			})

			if len(packages) >= 20 {
				break
			}
		}
	}

	// 获取更多信息
	for i, pkg := range packages {
		infoOutput, err := executeCommandWithTimeout("brew", "info", "--json=v1", pkg.Name)
		if err != nil {
			continue
		}

		var infoResults []map[string]interface{}
		if err := json.Unmarshal([]byte(infoOutput), &infoResults); err != nil {
			continue
		}

		if len(infoResults) > 0 {
			info := infoResults[0]
			version, _ := info["versions"].(map[string]interface{})["stable"].(string)
			desc, _ := info["desc"].(string)

			packages[i].Version = version
			packages[i].Description = desc
		}
	}

	return packages, nil
}

// GetMissingPackages 获取语言的缺失推荐包
func (a *App) GetMissingPackages(languageName string) []PackageInfo {
	// 这里可以实现获取推荐包的逻辑
	// 目前返回一些示例数据
	recommendedPackages := map[string][]PackageInfo{
		"Python": {
			{Name: "requests", Version: "2.31.0", Description: "简单易用的HTTP库"},
			{Name: "numpy", Version: "1.26.0", Description: "科学计算基础库"},
			{Name: "pandas", Version: "2.1.1", Description: "数据分析工具"},
		},
		"Node.js": {
			{Name: "express", Version: "4.18.2", Description: "Web应用框架"},
			{Name: "lodash", Version: "4.17.21", Description: "实用工具库"},
			{Name: "axios", Version: "1.5.1", Description: "HTTP客户端"},
		},
		"Go": {
			{Name: "gin", Version: "v1.9.1", Description: "Web框架"},
			{Name: "gorm", Version: "v1.25.4", Description: "ORM库"},
			{Name: "cobra", Version: "v1.7.0", Description: "CLI应用框架"},
		},
		"Ruby": {
			{Name: "rails", Version: "7.1.1", Description: "Web应用框架"},
			{Name: "rspec", Version: "3.12.0", Description: "测试框架"},
			{Name: "pry", Version: "0.14.2", Description: "增强的交互式shell"},
		},
		"Rust": {
			{Name: "tokio", Version: "1.32.0", Description: "异步运行时"},
			{Name: "serde", Version: "1.0.189", Description: "序列化框架"},
			{Name: "clap", Version: "4.4.6", Description: "命令行参数解析"},
		},
		"PHP": {
			{Name: "laravel/framework", Version: "10.x", Description: "Web应用框架"},
			{Name: "symfony/symfony", Version: "6.3", Description: "组件库和框架"},
			{Name: "guzzlehttp/guzzle", Version: "7.8", Description: "HTTP客户端"},
		},
		"Java": {
			{Name: "org.springframework.boot:spring-boot-starter", Version: "3.1.4", Description: "Spring Boot框架"},
			{Name: "com.google.guava:guava", Version: "32.1.2", Description: "Google核心库"},
			{Name: "org.projectlombok:lombok", Version: "1.18.30", Description: "减少样板代码"},
		},
		"C#": {
			{Name: "Newtonsoft.Json", Version: "13.0.3", Description: "JSON处理库"},
			{Name: "Microsoft.EntityFrameworkCore", Version: "7.0.11", Description: "ORM框架"},
			{Name: "Serilog", Version: "3.0.1", Description: "结构化日志"},
		},
	}

	if packages, ok := recommendedPackages[languageName]; ok {
		return packages
	}

	return []PackageInfo{}
}
