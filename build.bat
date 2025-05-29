@echo off
REM Build script for code2txt on Windows
REM Creates binaries for all supported platforms

setlocal enabledelayedexpansion

set VERSION=%1
if "%VERSION%"=="" set VERSION=dev

set LDFLAGS=-s -w -X main.version=%VERSION%

echo Building code2txt %VERSION%...

REM Create builds directory
if not exist builds mkdir builds

REM Windows AMD64
echo Building Windows AMD64...
set GOOS=windows
set GOARCH=amd64
go build -ldflags="%LDFLAGS%" -o builds/code2txt-windows-amd64.exe

REM Linux AMD64
echo Building Linux AMD64...
set GOOS=linux
set GOARCH=amd64
go build -ldflags="%LDFLAGS%" -o builds/code2txt-linux-amd64

REM macOS Intel
echo Building macOS Intel...
set GOOS=darwin
set GOARCH=amd64
go build -ldflags="%LDFLAGS%" -o builds/code2txt-macos-intel

REM macOS Apple Silicon
echo Building macOS Apple Silicon...
set GOOS=darwin
set GOARCH=arm64
go build -ldflags="%LDFLAGS%" -o builds/code2txt-macos-arm64

echo Build complete! Binaries available in builds/
dir builds\