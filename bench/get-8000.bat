@echo off

set y=%date:~0,4%
set m=%date:~5,2%
set d=%date:~8,2%
set h=%time:~0,2%
set h=%h: =0%
set i=%time:~3,2%
set s=%time:~6,2%
set time=%y%-%m%-%d%T%h%-%i%-%s%
set name=%~n0

set  url=http://127.0.0.1:8000/
echo get %url%
echo.
ab -n 1000 -c 100 %url%  | tee %name%_%time%.txt
echo.

pause