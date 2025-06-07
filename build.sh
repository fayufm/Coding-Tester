#!/bin/bash

echo "正在构建 Networ Tester..."

# 确保已安装Go
if ! command -v go &> /dev/null; then
    echo "错误: 未找到Go。请安装Go 1.21或更高版本。"
    exit 1
fi

# 确保已安装Wails
if ! command -v wails &> /dev/null; then
    echo "正在安装Wails..."
    go install github.com/wailsapp/wails/v2/cmd/wails@latest
fi

# 下载依赖
echo "正在下载依赖..."
go mod tidy

# 构建应用程序
echo "正在构建应用程序..."
wails build

echo "构建完成！可执行文件位于 build/bin 目录。" 