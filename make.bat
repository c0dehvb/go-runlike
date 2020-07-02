@echo off

if "%~2"=="image" (
    set MAKE_IMAGE=1
)

set ALL_PARAMS=
:param
set str=%~1
if "%str%"=="" (
    goto end
)
if "%ALL_PARAMS%"=="" (
    set ALL_PARAMS=%str%
) else (
    set ALL_PARAMS=%ALL_PARAMS% %str%
)
shift /0
goto param
:end


set MODULE_NAME=gitlab.iwhalecloud.com/rancher/gpushare-device-plugin

::@echo MODULE_NAME=%MODULE_NAME%
::@echo GOPATH=%GOPATH%
::@echo ALL_PARAMS=%ALL_PARAMS%

if "%MAKE_IMAGE%" neq "1" (
docker run --rm^
 -v %cd%:/go/src/%MODULE_NAME%^
 -w /go/src/%MODULE_NAME%^
 -e GOPROXY=https://goproxy.cn^
 golang:1.13^
 /bin/bash -c "make %ALL_PARAMS%"
)