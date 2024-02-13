protoc:
	protoc --go_out=pkg/core --go_opt=paths=source_relative --go-grpc_out=pkg/core --go-grpc_opt=paths=source_relative proto/audit.proto

build:
	docker-compose build

run:
	docker-compose up