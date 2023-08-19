#!/bin/bash

# 设置项目名称
PROJECT_NAME="NodeModulesSweeper"

# 创建输出目录
mkdir -p build/{windows,linux,macos}

# 为 Windows 编译
GOOS=windows GOARCH=amd64 go build -o build/windows/${PROJECT_NAME}.exe

# 为 Linux 编译
GOOS=linux GOARCH=amd64 go build -o build/linux/${PROJECT_NAME}

# 为 MacOS 编译
GOOS=darwin GOARCH=amd64 go build -o build/macos/${PROJECT_NAME}

echo "Compilation finished. Binaries are located under the 'build' directory."
