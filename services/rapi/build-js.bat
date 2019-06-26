
@echo on

set batPath=%~dp0
cd %batPath%
%~d0

protoc --js_out=import_style=commonjs:./ --grpc-web_out=import_style=commonjs,mode=grpcweb:./ apply.proto

cd %batPath%