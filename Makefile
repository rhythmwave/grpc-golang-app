.PHONY:
	build run

generate_model:
	go build -o bin/gorm-gen cmd/gorm-gen/main.go
	bin/gorm-gen generate

doc:
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	for api in proto/*; do \
		protoc --swagger_out=logtostderr=true:docs $$api ; \
	done

generate:
	protoc --go_out=. --go-grpc_out=. proto/*.proto

build: 
	go build -o bin/main cmd/server/main.go

run: build
	./bin/main

dev: generate
	go run cmd/main.go