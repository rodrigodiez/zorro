package main

import (
	"context"

	"google.golang.org/grpc/codes"

	"github.com/rodrigodiez/zorro/pkg/generator"
	"github.com/rodrigodiez/zorro/pkg/protobuf"
	"github.com/rodrigodiez/zorro/pkg/service"
	"github.com/rodrigodiez/zorro/pkg/storage"
	"google.golang.org/grpc/status"
)

type server struct {
	zorro     service.Zorro
	storage   storage.Storage
	generator generator.Generator
}

func (s *server) Mask(ctx context.Context, req *protobuf.MaskRequest) (*protobuf.MaskResponse, error) {
	return &protobuf.MaskResponse{
		Value: s.zorro.Mask(req.GetKey()),
	}, nil
}

func (s *server) Unmask(ctx context.Context, req *protobuf.UnmaskRequest) (*protobuf.UnmaskResponse, error) {
	key, ok := s.zorro.Unmask(req.GetValue())

	if !ok {
		return nil, status.Errorf(codes.NotFound, "Not Found")
	}

	return &protobuf.UnmaskResponse{
		Key: key,
	}, nil
}
