package server

import (
	"github.com/jharshman/simple-client-server/config"
	"github.com/jharshman/simple-client-server/pkg/grpc"
	log "github.com/sirupsen/logrus"
	rpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
	"os"
	"os/signal"
	"syscall"
)

// Start initializes the gRPC server process.
func Start(cfg *config.Server) error {
	initLogger(cfg.LogLevel())

	//debug logging
	log.Debugf("port %s\n", cfg.Port)
	log.Debugf("log level %s\n", cfg.LogLvl)
	log.Debugf("cert file %s\n", cfg.CertFile)
	log.Debugf("key file %s\n", cfg.KeyFile)

	lis, err := net.Listen("tcp", cfg.BindPort())
	if err != nil {
		return err
	}
	defer lis.Close()

	tls, _ := credentials.NewServerTLSFromFile(cfg.CertFile, cfg.KeyFile)
	opts := []rpc.ServerOption{
		rpc.Creds(tls),
	}
	grpcServer := rpc.NewServer(opts...)

	srv := &handler{}
	grpc.RegisterEchoServerServer(grpcServer, srv)

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	log.Infof("starting grpc server: %s\n", cfg.BindPort())
	go grpcServer.Serve(lis)

	<-done
	grpcServer.Stop()
	log.Infof("stopped server\n")

	return nil
}

func initLogger(level log.Level) {
	log.SetLevel(level)
}
