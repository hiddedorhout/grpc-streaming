PROJECT_ROOT = ${PWD}
CLIENT_DIR = $(PROJECT_ROOT)/client
SERVER_DIR = ${PROJECT_ROOT}/server
JS_DIR = js/client/src/proto
PROXY_DIR = $(PROJECT_ROOT)/proxy
LOCAL_HOST_NAME = localhost
PROTO_DIR = $(PROJECT_ROOT)/proto
PROTOCOL_NAME = service
DOCKER_ROOT_DIR = /src

srcGen:
	protoc -I $(PROTO_DIR) $(PROTOCOL_NAME).proto --go_out=plugins=grpc:$(CLIENT_DIR)
	protoc -I $(PROTO_DIR) $(PROTOCOL_NAME).proto --go_out=plugins=grpc:$(SERVER_DIR)
	
	rm -rf $(JS_DIR)
	mkdir -p $(JS_DIR)
	protoc -I $(PROTO_DIR) $(PROTOCOL_NAME).proto \
	--js_out=import_style=commonjs:$(PROJECT_ROOT)/${JS_DIR} --grpc-web_out=import_style=commonjs,mode=grpcwebtext:$(PROJECT_ROOT)/${JS_DIR}

	docker run --rm -it \
	--mount type=bind,source="$(PWD)",target=$(DOCKER_ROOT_DIR) \
	debian:stretch sed -i \
		-e '1s;^;/* eslint-disable */\n;' \
		-e "s/\(var jspb =\)/\/\/ @ts-ignore\n\1/" \
		$(DOCKER_ROOT_DIR)/$(JS_DIR)/$(PROTOCOL_NAME)_pb.js


run-proxy:
	sed 's/LOCAL_HOST_NAME/$(LOCAL_HOST_NAME)/' $(PROXY_DIR)/envoy.yaml > $(PROXY_DIR)/envoy.local.yaml
	docker run -ti --rm --net=host \
		--mount type=bind,source="$(PROXY_DIR)",target=/src \
		-p 8080:8080 -p 9901:9901 \
		envoyproxy/envoy:v1.10.0 envoy \
			--config-path /src/envoy.local.yaml \
			--log-level trace

run-server:
	go build ./server/cmd
	./cmd