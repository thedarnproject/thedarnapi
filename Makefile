PROTO_DIR := api
protoc:
	protoc -I ${PROTO_DIR} ${PROTO_DIR}/darn.proto --go_out=plugins=grpc:${PROTO_DIR}
