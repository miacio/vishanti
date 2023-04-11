@echo off

setlocal

if exist build.bat goto ok
echo build.bat must be run from it's folder
goto end

:ok

set GOPROXY=https://goproxy.cn,direct
set GO111MODULE=on

if not exist log mkdir log

gofmt -w -s .

go build -o bin/vishanti.exe github.com/miacio/vishanti/runner

:end
echo finished