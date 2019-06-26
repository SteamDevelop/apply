
@echo on

set batPath=%~dp0
cd %batPath%
%~d0

protoc --go_out=plugins=grpc:./ apply.proto
go build

cd %batPath%