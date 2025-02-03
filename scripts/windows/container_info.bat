@echo off
docker ps -l > container_info.txt
if %errorlevel% neq 0 (
    echo Error executing docker command. Please check if Docker is installed and running.
    exit /b %errorlevel%
)
echo Image info saved to container_info.txt