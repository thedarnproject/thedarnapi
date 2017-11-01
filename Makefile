PROTO_DIR := api
BIN_DIR := bin
SERVER_DIR := server
PROTOBUF_DOWNLOAD_URL := https://github.com/google/protobuf/releases/download/v3.4.0/protoc-3.4.0-linux-x86_64.zip
PROTOC_FILE := /tmp/protoc.zip
PROTOC_TEMP_LOCATION := /tmp/protoc
PROTOC_BINARY_LOCATION := /usr/local/bin/

default: binary

setup-protobuf:
	curl -L -o ${PROTOC_FILE} https://github.com/google/protobuf/releases/download/v3.4.0/protoc-3.4.0-linux-x86_64.zip
	mkdir -p ${PROTOC_TEMP_LOCATION}
	unzip ${PROTOC_FILE} -d ${PROTOC_TEMP_LOCATION}
	mv ${PROTOC_TEMP_LOCATION}/bin/protoc ${PROTOC_BINARY_LOCATION}
	rm -rf ${PROTOC_FILE} ${PROTOC_TEMP_LOCATION}

protoc:
	protoc \
		--proto_path=/usr/local/include \
		--proto_path=${GOPATH}/src \
		--proto_path=${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--proto_path=${PROTO_DIR} \
		--go_out=plugins=grpc:${PROTO_DIR} \
		${PROTO_DIR}/darn.proto

reverse-protoc:
	protoc \
		--proto_path=/usr/local/include \
		--proto_path=${GOPATH}/src \
		--proto_path=${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--proto_path=${PROTO_DIR} \
		--grpc-gateway_out=logtostderr=true:${PROTO_DIR} \
		${PROTO_DIR}/darn.proto

binary:
	mkdir -p bin
	go build -o ${BIN_DIR}/thedarnapi ${SERVER_DIR}/main.go

install:
	cd ${SERVER_DIR}/ && go install

clean:
	rm -rf bin/

glide:
	glide update --strip-vendor

glide-vc:
	glide-vc --use-lock-file --only-code --no-tests

glide-hard:
	glide cache-clear
	glide update --strip-vendor

build-docker-image:
	docker build -t containscafeine/thedarnapi:master .
