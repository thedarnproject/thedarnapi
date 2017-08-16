PROTO_DIR := api
BIN_DIR := bin
SERVER_DIR := server

default: binary

protoc:
	protoc -I ${PROTO_DIR} ${PROTO_DIR}/darn.proto --go_out=plugins=grpc:${PROTO_DIR}

binary:
	mkdir -p bin
	go build -o ${BIN_DIR}/thedarnapi ${SERVER_DIR}/main.go

clean:
	rm -rf bin/

glide:
	glide update --strip-vendor

glide-vc:
	glide-vc --only-code --no-tests

glide-hard:
	rm -rf ~/.glide/
	glide update --strip-vendor
