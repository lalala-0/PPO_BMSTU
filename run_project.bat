@echo off
echo Starting project...

REM Запуск трех инстансов Go приложения и сохранение их PID
start "Instance 1" cmd /c "go run main.go config1.json"
start "Instance 2" cmd /c "go run main.go config2.json"
start "Instance 3" cmd /c "go run main.go config3.json"

REM Пауза, чтобы дать время на запуск инстансов
timeout /t 5 /nobreak > nul

REM Запуск Nginx в новом терминале из нужной папки
echo Starting Nginx...
start "Nginx" cmd /k "cd /d C:\nginx-1.27.2 && nginx.exe"
echo Nginx started successfully!

REM Запуск встроенного PHP сервера на порте 8000
echo Starting PHP built-in server...
start "PHP Server" cmd /c "cd /d C:\nginx-1.27.2 && C:\php\php.exe -S localhost:8000"

REM Открыть страницу проекта в браузере
timeout /t 3 /nobreak > nul
echo Opening project in browser...
start "" "http://localhost/"

echo Project started successfully!
echo Press any key to terminate all processes...
pause > nul

REM Завершение всех процессов
echo Terminating all processes...

REM Завершение Nginx
taskkill /F /IM nginx.exe > nul 2>&1
echo Terminated nginx process.

REM Завершение PHP сервера
taskkill /F /IM php.exe > nul 2>&1
echo Terminated php process.

REM Завершение Go инстансов
taskkill /F /IM go.exe > nul 2>&1
echo Terminated Go processes.

echo All processes terminated.
