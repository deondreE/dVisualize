@echo off
setlocal enabledelayedexpansion

set "files="
for %%f in (src\*.go) do (
    set "files=!files! %%f"
)

go build -o main.exe !files!
echo Build completed for all Go files in src directory.