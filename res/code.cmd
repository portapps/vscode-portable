@echo off
setlocal

set VSCODE_APPDATA=%~dp0\data\appdata
set VSCODE_LOGS=%~dp0\data\logs
set VSCODE_EXTENSIONS=%~dp0\data\extensions

set VSCODE_DEV=
set ELECTRON_RUN_AS_NODE=1

call "%~dp0app\Code.exe" "%~dp0app\resources\app\out\cli.js" %*
endlocal
