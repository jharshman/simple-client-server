package server

import (
	"context"
	"github.com/jharshman/simple-client-server/pkg/grpc"
)

type handler struct{}

func (h *handler) Echo(ctx context.Context, in *grpc.Message) (*grpc.Message, error) {
	return in, nil
}
