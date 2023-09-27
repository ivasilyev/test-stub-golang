@echo off

set "VER=1.0"

mkdir "target"

set "GOOS=linux"
set "GOARCH=amd64"
echo Compiling for %GOOS% %GOARCH%
go build -o "target/test-stub-golang_%GOOS%_%GOARCH%_v%VER%" "src/main/go/main.go"

set "GOOS=windows"
set "GOARCH=amd64"
echo Compiling for %GOOS% %GOARCH%
go build -o "target/test-stub-golang_%GOOS%_%GOARCH%_v%VER%.exe" "src/main/go/main.go"
