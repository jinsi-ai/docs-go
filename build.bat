@echo off
chcp 65001 >nul
echo ==========================================
echo DocsGo Cross-Platform Build Script
echo ==========================================
echo.

:: Create bin directory
if not exist bin mkdir bin

:: Clean old builds
echo Cleaning old builds...
if exist bin\*.exe del /Q bin\*.exe
if exist bin\docs-go-* del /Q bin\docs-go-*
echo.

:: Get version from git tag or use default
for /f "tokens=*" %%a in ('git describe --tags --always 2^>nul') do set VERSION=%%a
if "%VERSION%"=="" set VERSION=v0.1
echo Building version: %VERSION%
echo.

:: Build flags
set LDFLAGS=-ldflags "-s -w -X main.version=%VERSION%"

echo ==========================================
echo Building for Windows (amd64)...
echo ==========================================
set GOOS=windows
set GOARCH=amd64
go build %LDFLAGS% -o bin\docs-go-windows-amd64.exe
if %errorlevel% neq 0 (
    echo Failed to build for Windows amd64
    exit /b 1
)
echo ✓ bin\docs-go-windows-amd64.exe
echo.

echo ==========================================
echo Building for Windows (arm64)...
echo ==========================================
set GOOS=windows
set GOARCH=arm64
go build %LDFLAGS% -o bin\docs-go-windows-arm64.exe
if %errorlevel% neq 0 (
    echo Failed to build for Windows arm64
    exit /b 1
)
echo ✓ bin\docs-go-windows-arm64.exe
echo.

echo ==========================================
echo Building for Linux (amd64)...
echo ==========================================
set GOOS=linux
set GOARCH=amd64
go build %LDFLAGS% -o bin\docs-go-linux-amd64
if %errorlevel% neq 0 (
    echo Failed to build for Linux amd64
    exit /b 1
)
echo ✓ bin\docs-go-linux-amd64
echo.

echo ==========================================
echo Building for Linux (arm64)...
echo ==========================================
set GOOS=linux
set GOARCH=arm64
go build %LDFLAGS% -o bin\docs-go-linux-arm64
if %errorlevel% neq 0 (
    echo Failed to build for Linux arm64
    exit /b 1
)
echo ✓ bin\docs-go-linux-arm64
echo.

echo ==========================================
echo Building for macOS (amd64)...
echo ==========================================
set GOOS=darwin
set GOARCH=amd64
go build %LDFLAGS% -o bin\docs-go-darwin-amd64
if %errorlevel% neq 0 (
    echo Failed to build for macOS amd64
    exit /b 1
)
echo ✓ bin\docs-go-darwin-amd64
echo.

echo ==========================================
echo Building for macOS (arm64)...
echo ==========================================
set GOOS=darwin
set GOARCH=arm64
go build %LDFLAGS% -o bin\docs-go-darwin-arm64
if %errorlevel% neq 0 (
    echo Failed to build for macOS arm64
    exit /b 1
)
echo ✓ bin\docs-go-darwin-arm64
echo.

echo ==========================================
echo Build Summary
echo ==========================================
dir /b bin\docs-go-* 2>nul
echo.
echo Total builds completed successfully!
pause
