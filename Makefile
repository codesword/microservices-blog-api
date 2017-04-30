

gen:
	cd api-spec && \
	protoc  -I. -I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	 --go_out=plugins=grpc:../pb authorization/*.proto && \
	 protoc  -I. -I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	 --go_out=plugins=grpc:../pb blog/*.proto
