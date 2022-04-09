Install docker-compose first
docker-compose up -d

update
docker-compose build /or/
docker-compose up --build

iniciar GO
go mod init [pasta local]

Acessar docker
docker exec -it teste_app_1 bash

GERAR PROTO p/ GRPC
protoc --go_out=application/grpc/pb --go_opt=paths=source_relative --go-grpc_out=application/grpc/pb --go-grpc_opt=paths=source_relative --proto_path=application/grpc/protofiles application/grpc/protofiles/*.proto

Evans
evans -r repl
