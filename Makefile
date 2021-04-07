.PHONY: generate
build: vendor-proto .generate

PHONY: .generate
.generate:
	mkdir -p pkg/ocp-task-api
	protoc -I vendor.protogen \
		--go_out=pkg/ocp-task-api --go_opt=paths=import \
		--go-grpc_out=pkg/ocp-task-api --go-grpc_opt=paths=import \
		--grpc-gateway_out=pkg/ocp-task-api \
		--grpc-gateway_opt=logtostderr=true \
    --grpc-gateway_opt=paths=import \
		--validate_out lang=go:pkg/ocp-task-api \
		api/ocp-task-api/ocp-task-api.proto
	mv pkg/ocp-task-api/gihtub.com/ozoncp/ocp-task-api/pkg/ocp-task-api/* pkg/ocp-task-api/
	rm -rf pkg/ocp-task-api/gihtub.com
	mkdir -p cmd/ocp-task-api

PHONY: vendor-proto
vendor-proto: .vendor-proto

PHONY: .vendor-proto
.vendor-proto:
		mkdir -p vendor.protogen
		mkdir -p vendor.protogen/api/ocp-task-api
		cp api/ocp-task-api/ocp-task-api.proto vendor.protogen/api/ocp-task-api
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi
		@if [ ! -d vendor.protogen/github.com/envoyproxy ]; then \
			mkdir -p vendor.protogen/github.com/envoyproxy &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/github.com/envoyproxy/protoc-gen-validate ;\
		fi


.PHONY: deps
deps: install-go-deps

.PHONY: install-go-deps
install-go-deps: .install-go-deps

.PHONY: .install-go-deps
.install-go-deps:
		ls go.mod || go mod init
		go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
		go get -u github.com/golang/protobuf/proto
		go get -u github.com/golang/protobuf/protoc-gen-go
		go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
		tmpdir=$$(mktemp -d); cd $$tmpdir && export GO111MODULE=off && go get -d github.com/envoyproxy/protoc-gen-validate \
			&& go build -o $$GOPATH/bin/protoc-gen-validate $$GOPATH/src/github.com/envoyproxy/protoc-gen-validate/main.go && cd -
