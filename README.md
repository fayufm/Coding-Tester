# Networ Tester

一个用于检测系统中安装的编程语言及其依赖的工具。

## 功能特点

- 检测系统中已安装的编程语言
- 显示编程语言的版本信息
- 检查编程语言的依赖是否完整
- 为未安装的编程语言提供下载链接和安装教程
- 现代、美观的用户界面

## 支持的编程语言

- Go
- Python
- Node.js
- Java
- C# (.NET)
- Ruby
- PHP
- Rust

## 安装说明

### 从源代码构建

1. 确保已安装 Go 1.21 或更高版本
2. 安装 Wails CLI:
   ```
   go install github.com/wailsapp/wails/v2/cmd/wails@latest
   ```
3. 克隆此仓库:
   ```
   git clone https://github.com/your-username/networ-tester.git
   cd networ-tester
   ```
4. 构建应用程序:
   ```
   wails build
   ```
5. 可执行文件将在 `build/bin` 目录中生成

### 直接下载

访问 [Releases](https://github.com/your-username/networ-tester/releases) 页面下载最新版本的可执行文件。

## 使用方法

1. 启动应用程序
2. 点击"扫描编程语言"按钮
3. 查看已安装和未安装的编程语言列表
4. 点击语言卡片查看详细信息
5. 对于未安装的语言，可以通过提供的链接下载并按照教程安装

## 开发

### 前提条件

- Go 1.21+
- Wails v2.7.1+

### 开发命令

- 启动开发模式:
  ```
  wails dev
  ```

- 构建应用程序:
  ```
  wails build
  ```

## 许可证

MIT 