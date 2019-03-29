package server

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/jharshman/simple-client-server/mock"
	"github.com/jharshman/simple-client-server/pkg/grpc"
	"github.com/matryer/is"
	"testing"
)

func TestHandler_Echo(t *testing.T) {
	is := is.New(t)
	mctl := gomock.NewController(t)
	defer mctl.Finish()

	mcli := mock.NewMockEchoServerClient(mctl)
	mcli.EXPECT().Echo(gomock.Any(), gomock.Any()).Return(&grpc.Message{
		Data: "test",
	}, nil)

	res, err := mcli.Echo(context.Background(), &grpc.Message{
		Data: "test",
	})

	is.NoErr(err)
	is.Equal(res.Data, "test")
}
