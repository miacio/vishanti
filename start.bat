@echo off

setlocal

if exist start.bat goto ok
echo start.bat must be run from it's folder
goto end

:ok

start /b bin\vishanti.exe >> log\runner.log 2>&1 &

echo start successfully

:end