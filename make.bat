@echo off
setlocal EnableDelayedExpansion

set "MODULE_PATH=github.com/mattia37773"
set "APP_NAME=mt"
set "BUILD_METHOD=source"

for /f "delims=" %%i in ('git describe --tags --always 2^>nul') do set "GIT_TAG=%%i"
if not defined GIT_TAG set "GIT_TAG=dev"

for /f "delims=" %%i in ('go env GOOS') do set "GOOS=%%i"


if "%~1"=="" goto help
if "%~1"=="help" goto help
if "%~1"=="watch" goto watch
if "%~1"=="build" goto build
if "%~1"=="test" goto test
if "%~1"=="test-basic" goto test-basic
if "%~1"=="test-cover" goto test-cover
if "%~1"=="clear-docker" goto clear-docker
if "%~1"=="docker-show" goto docker-show
if "%~1"=="check-docker-compose" goto check-docker-compose

echo Unknown target: %~1
echo.
goto help


:help
echo Usage: make.bat [target]
echo.
echo   help                 Shows help for all command
echo   watch                Recompile on filechange
echo   build                Build the binary
echo   test                 Run the testsuite
echo   test-basic           Run the testsuite wihout the verbose flag
echo   test-cover           Show the test coverage
echo   clear-docker         This removes everything in docker! from all namespaces
echo   docker-show          Show all containers, images ^& volumes
echo   check-docker-compose Check Docker Compose
goto end


:watch
echo watch the code
call .\bin\watch.bat
goto end


:build
echo Building...
set "GIT_TAG=%GIT_TAG%"
go build -ldflags="-X '%MODULE_PATH%/%APP_NAME%/config.OverrideCurrentVersion=%GIT_TAG%' -X '%MODULE_PATH%/%APP_NAME%/config.OverrideBuildMethod=%BUILD_METHOD%'"
goto end


:test
call :check-docker-compose
if errorlevel 1 goto end

echo Testing with unittests...
echo.

echo Removing the binaries
if exist "tests\bin\%GOOS%" rmdir /s /q "tests\bin\%GOOS%"
if not exist "tests\bin" mkdir "tests\bin"
if not exist "tests\bin\%GOOS%" mkdir "tests\bin\%GOOS%"

echo Building binarie for testing update
echo.

REM source install
set "GIT_TAG=%GIT_TAG%"
go build ^
    -ldflags="-X '%MODULE_PATH%/%APP_NAME%/config.OverrideCurrentVersion=v1.0.0' -X '%MODULE_PATH%/%APP_NAME%/config.OverrideBuildMethod=source'" ^
    -o "tests\bin\%GOOS%\source-install.exe"

if errorlevel 1 goto end

REM go install
set "GIT_TAG=%GIT_TAG%"
go build ^
    -ldflags="-X '%MODULE_PATH%/%APP_NAME%/config.OverrideCurrentVersion=v1.0.0'" ^
    -o "tests\bin\%GOOS%\go-install.exe"

if errorlevel 1 goto end

REM homebrew install
set "GIT_TAG=%GIT_TAG%"
go build ^
    -ldflags="-X '%MODULE_PATH%/%APP_NAME%/config.OverrideCurrentVersion=v1.0.0' -X '%MODULE_PATH%/%APP_NAME%/config.OverrideBuildMethod=homebrew'" ^
    -o "tests\bin\%GOOS%\homebrew-install.exe"

if errorlevel 1 goto end

go clean -cache

go test -v -p 1 ./tests/cmd/... ^
    -ldflags="-X '%MODULE_PATH%/%APP_NAME%/config.OverrideCurrentVersion=v1.0.0'" ^
    ./...

if errorlevel 1 goto end

echo.
echo Removing the binaries
if exist "tests\bin\%GOOS%" rmdir /s /q "tests\bin\%GOOS%"

goto end


:test-basic
echo Testing with unittests without verbose mode...
echo.

echo Removing the binaries
if exist "tests\bin\homebrew-install.exe" del /f /q "tests\bin\homebrew-install.exe"
if exist "tests\bin\go-install.exe" del /f /q "tests\bin\go-install.exe"
if exist "tests\bin\source-install.exe" del /f /q "tests\bin\source-install.exe"

if not exist "tests\bin" mkdir "tests\bin"

echo Building binarie for testing update
echo.

REM source install
set "GIT_TAG=%GIT_TAG%"
go build ^
    -ldflags="-X '%MODULE_PATH%/%APP_NAME%/config.OverrideCurrentVersion=v1.0.0' -X '%MODULE_PATH%/%APP_NAME%/config.OverrideBuildMethod=source'" ^
    -o "tests\bin\source-install.exe"

if errorlevel 1 goto end

REM go install
set "GIT_TAG=%GIT_TAG%"
go build ^
    -ldflags="-X '%MODULE_PATH%/%APP_NAME%/config.OverrideCurrentVersion=v1.0.0'" ^
    -o "tests\bin\go-install.exe"

if errorlevel 1 goto end

REM homebrew install
set "GIT_TAG=%GIT_TAG%"
go build ^
    -ldflags="-X '%MODULE_PATH%/%APP_NAME%/config.OverrideCurrentVersion=v1.0.0' -X '%MODULE_PATH%/%APP_NAME%/config.OverrideBuildMethod=homebrew'" ^
    -o "tests\bin\homebrew-install.exe"

if errorlevel 1 goto end

go clean -cache

go test -p 1 ./tests/cmd/... ^
    -ldflags="-X '%MODULE_PATH%/%APP_NAME%/config.OverrideCurrentVersion=v1.0.0'" ^
    ./...

if errorlevel 1 goto end

echo.
echo Removing the binaries
if exist "tests\bin\homebrew-install.exe" del /f /q "tests\bin\homebrew-install.exe"
if exist "tests\bin\go-install.exe" del /f /q "tests\bin\go-install.exe"
if exist "tests\bin\source-install.exe" del /f /q "tests\bin\source-install.exe"

goto end


:test-cover
echo Shows unittest coverage...
echo Testing with unittests...
echo.

echo Removing the binaries
if exist "tests\bin\homebrew-install.exe" del /f /q "tests\bin\homebrew-install.exe"
if exist "tests\bin\go-install.exe" del /f /q "tests\bin\go-install.exe"
if exist "tests\bin\source-install.exe" del /f /q "tests\bin\source-install.exe"

if not exist "tests\bin" mkdir "tests\bin"

echo Building binarie for testing update
echo.

REM source install
set "GIT_TAG=%GIT_TAG%"
go build ^
    -ldflags="-X '%MODULE_PATH%/%APP_NAME%/config.OverrideCurrentVersion=v1.0.0' -X '%MODULE_PATH%/%APP_NAME%/config.OverrideBuildMethod=source'" ^
    -o "tests\bin\source-install.exe"

if errorlevel 1 goto end

REM go install
set "GIT_TAG=%GIT_TAG%"
go build ^
    -ldflags="-X '%MODULE_PATH%/%APP_NAME%/config.OverrideCurrentVersion=v1.0.0'" ^
    -o "tests\bin\go-install.exe"

if errorlevel 1 goto end

REM homebrew install
set "GIT_TAG=%GIT_TAG%"
go build ^
    -ldflags="-X '%MODULE_PATH%/%APP_NAME%/config.OverrideCurrentVersion=v1.0.0' -X '%MODULE_PATH%/%APP_NAME%/config.OverrideBuildMethod=homebrew'" ^
    -o "tests\bin\homebrew-install.exe"

if errorlevel 1 goto end

go clean -cache

go test -v -p 1 ./tests/cmd/... ^
    -ldflags="-X '%MODULE_PATH%/%APP_NAME%/config.OverrideCurrentVersion=v1.0.0'" ^
    ./...

if errorlevel 1 goto end

echo.
echo Removing the binaries
if exist "tests\bin\homebrew-install.exe" del /f /q "tests\bin\homebrew-install.exe"
if exist "tests\bin\go-install.exe" del /f /q "tests\bin\go-install.exe"
if exist "tests\bin\source-install.exe" del /f /q "tests\bin\source-install.exe"

goto end


:clear-docker
echo Removing everything in docker from all namespaces...

docker rm -f $(docker ps -aq) 2>nul
docker system prune -a --volumes -f

goto end


:docker-show
echo All running docker containers
docker ps
echo.
echo All docker images
docker image ls
echo.
echo All docker volumes
docker volume ls

goto end


:check-docker-compose
docker compose version >nul 2>&1
if errorlevel 1 (
    echo Docker Compose isn't installed
    echo please install it
    exit /b 1
)

docker info >nul 2>&1
if errorlevel 1 (
    echo The docker engine isn't running
    exit /b 1
)

exit /b 0


:end
endlocal
