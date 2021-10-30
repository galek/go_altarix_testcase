# Backend server(queue manager). Unmantained
Written as test cases for Altarix in 2017, without experience with Go and Microservices architecture.
It Will be rewritten for demonstration my level at 2021

## Features
* RabbitMQ.
* Working as service
* Get/Write to Postgres

License: GPL 2.0

## Status on 2020
Unmantained, is need check and refactor if needed.

## TODO
- Docker

## Running
- go to testcase\src\altarix
### Under Windows:
- Run build
- Turn off the UAC(!!!!!!!!!!!!!!!!!!!!)
- replace testacase/bin/nssm.exe to 32bit version if your system is used 32bit addresses
- Unpack DenwerKit and run denwer\Run.exe
- run RabbitMQ and postgres
- run install bat
- run start bat

### Under Linux:
- Run build
- run RabbitMQ and postgres
- run sudo altarix install
- run sudo altarix start

### Debug/Release build
- see debughelper.go 13 and 14 lines:
- ISDebug - set to true for heavy debug. All errors will written to console, else to error.log and info.log 
- ISShowSendGetReq - only rabbitmq requested will writted to error.log and info.log
