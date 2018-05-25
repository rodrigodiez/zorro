.PHONY: protobuf clean

protobuf:
	mkdir -p pkg/protobuf
	protoc -I pb/ pb/*.proto --go_out=plugins=grpc:pkg/protobuf