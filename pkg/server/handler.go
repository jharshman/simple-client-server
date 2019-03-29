package server

import (
	"context"
	"github.com/jharshman/simple-client-server/pkg/grpc"
	log "github.com/sirupsen/logrus"
)

type handler struct{}

func (h *handler) Echo(ctx context.Context, in *grpc.Message) (*grpc.Message, error) {
	log.Infof("received %s\n", in.Data)
	return in, nil
}
