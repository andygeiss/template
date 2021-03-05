@echo off
cls
setlocal EnableDelayedExpansion

echo getting info from Git ...
git pull
git rev-parse --short HEAD > build.txt
git describe --tags > version.txt
for %%* in (.) do @echo %%~n* > name.txt
echo.

SET /p BUILD=<build.txt
SET /p NAME=<name.txt

echo start building ...
go build --ldflags "-s -w -X=main.build=%BUILD% -X=main.name=%NAME% -X=main.version=%VERSION%" -o main.exe main.go

echo cleaning up ...
del build.txt name.txt version.txt

echo done.
