@echo off

go get -u github.com/gorilla/mux

REM go build ./src/hello
go build ./src/simple-api
