@echo off
SETLOCAL

:: 设置项目名称
SET PROJECT_NAME=NodeModulesSweeper

:: 创建输出目录
IF NOT EXIST build\windows mkdir build\windows
IF NOT EXIST build\linux mkdir build\linux
IF NOT EXIST build\macos mkdir build\macos

:: 为 Windows 编译
SET GOOS=windows
SET GOARCH=amd64
go build -o build\windows\%PROJECT_NAME%.exe

:: 为 Linux 编译
SET GOOS=linux
SET GOARCH=amd64
go build -o build\linux\%PROJECT_NAME%

:: 为 MacOS 编译
SET GOOS=darwin
SET GOARCH=amd64
go build -o build\macos\%PROJECT_NAME%

echo Compilation finished. Binaries are located under the 'build' directory.
ENDLOCAL
