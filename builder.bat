@echo off
echo Build scriptgen_win_x86-64.exe
set GOOS=windows
set GOARCH=amd64
go build -ldflags="-s -w" -o bin\scriptgen_win_x86-64.exe
echo Build scriptgen_win_arm64.exe
set GOOS=windows
set GOARCH=arm64
go build -ldflags="-s -w" -o bin\scriptgen_win_arm64.exe
echo Build scriptgen_mac_x86-64
set GOOS=darwin
set GOARCH=amd64
go build -ldflags="-s -w" -o bin\scriptgen_mac_x86-64
echo Build scriptgen_mac_arm64
set GOOS=darwin
set GOARCH=arm64
go build -ldflags="-s -w" -o bin\scriptgen_mac_arm64
echo Build scriptgen_lnx_x86-64
set GOOS=linux
set GOARCH=amd64
go build -ldflags="-s -w" -o bin\scriptgen_lnx_x86-64
echo Build scriptgen_lnx_arm64
set GOOS=linux
set GOARCH=arm64
go build -ldflags="-s -w" -o bin\scriptgen_lnx_arm64
echo Done!