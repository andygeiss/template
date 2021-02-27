@echo off
cls
setlocal EnableDelayedExpansion

echo getting info from Git ...
git pull
git rev-parse --short HEAD > build.txt
git describe --tags > version.txtfor %%* in (.) do @echo %%~n*
echo.

SET /p BUILD=<build.txt
SET /p NAME=<name.txt
for %%F in ($PWD) do SET NAME=%%~dpF

echo NAME = %NAME%
exit

echo start building ...
go build --ldflags "-s -w -X=main.build=%BUILD% -X=main.name=%NAME% -X=main.version=%VERSION%" -o %GOPATH%\bin\create-go-app.exe main.go

echo cleaning up ...
del build.txt name.txt version.txt

echo done.
