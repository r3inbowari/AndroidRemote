for /f "delims=" %%i in ('go version') do (set goVersion=%%i)
@REM echo %goVersion%

for /f "delims=" %%i in ('git show -s --format^=%%H') do (set gitHash=%%i)
@REM echo %gitHash%

for /f "delims=" %%i in ('git show -s --format^=%%cd') do (set buildTime=%%i)
@REM echo %buildTime%

set GOOS=windows
go build -ldflags "-X 'main.goVersion=%goVersion%' -X 'main.gitHash=%gitHash%' -X 'main.buildTime=%buildTime%'" -o RemoteServer.exe main.go

set GOOS=linux
go build -ldflags "-X 'main.goVersion=%goVersion%' -X 'main.gitHash=%gitHash%' -X 'main.buildTime=%buildTime%'" -o RemoteServer main.go