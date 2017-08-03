PROTO_DIR := api
BIN_DIR := bin
SERVER_DIR := server

protoc:
	protoc -I ${PROTO_DIR} ${PROTO_DIR}/darn.proto --go_out=plugins=grpc:${PROTO_DIR}

binary:
	go build -o ${BIN_DIR}/thedarnapi ${SERVER_DIR}/main.go
