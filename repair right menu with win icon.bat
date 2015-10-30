reg add "HKEY_CLASSES_ROOT\lnkfile" /v IsShortcut /f
reg add "HKEY_CLASSES_ROOT\piffile" /v IsShortcut /f
taskkill /f /im explorer.exe & explorer.exe
pause