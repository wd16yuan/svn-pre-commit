@echo off

setlocal
set REPOS=%1
set TXN=%2
::验证程序目录
set CheckProHome=""
::SVN Server 命令目录
set SvnServerBin=""

cd /d %CheckProHome%

%SvnServerBin%\svnlook.exe log %REPOS% -t %TXN% | findstr ".........." > nul
if %errorlevel% gtr 0 goto nul_err

%SvnServerBin%\svnlook.exe changed %REPOS% -t %TXN% | .\svn-pre-commit.exe
if %errorlevel% gtr 0 exit 1
exit 0


:nul_err
echo 请输入10个字符以上信息 >&2
exit 1