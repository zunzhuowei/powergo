::ֻ��ʾecho
@echo off
::����DOS������������ɫ
@color 06


::����GOPATH·����golandIDE��ͬ
SET projectName=powergo
SET projectPath=D:\GOPATH\src\powergo
SET GOPATH=%projectPath%

;SET CGO_ENABLED=0
;SET GOOS=linux
;SET GOARCH=amd64
::;go build  -v -o build/%projectName% %CD%\src\server\main.go
;go build  -v -o %projectName% %projectPath%\src\main.go
echo "--------- %DATE%%TIME% Build For %projectName% Success!"

SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=386
::go build  -v -o build/%projectName%.exe %CD%\src\server\main.go
go build  -v -o %projectName%.exe %projectPath%\src\main.go
echo "--------- %DATE%%TIME% Build For %projectName%.exe Success!"
pause


::xcopy G:\GoProject\github.com\*.* D:\W\ /s /e
::rmdir /q /s G:\GoProject\github.com\
::xcopy G:\GoProject\github.com\*.txt D:\W\ 