package server

import (
	"github.com/jharshman/simple-client-server/config"
	"github.com/jharshman/simple-client-server/pkg/grpc"
	log "github.com/sirupsen/logrus"
	rpc "google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func Start(cfg *config.Server) error {
	initLogger(cfg.LogLevel())

	//debug logging
	log.Debugf("port %s\n", cfg.Port)
	log.Debugf("log level %s\n", cfg.LogLvl)

	lis, err := net.Listen("tcp", cfg.BindPort())
	if err != nil {
		return err
	}
	defer lis.Close()

	srv := &handler{}
	grpcServer := rpc.NewServer()
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
