package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	Port     string `mapstructure:"port"`
	LogLvl   string `mapstructure:"loglvl"`
	CertFile string `mapstructure:"cert"`
	KeyFile  string `mapstructure:"key"`
}

type Client struct {
	Addr     string `mapstructure:"server"`
	Port     string `mapstructure:"server-port"`
	Message  string `mapstructure:"msg"`
	CertFile string `mapstructure:"client-cert"`
}

func (s *Server) BindPort() string {
	return fmt.Sprintf(":%s", s.Port)
}

func (s *Server) LogLevel() log.Level {
	var l log.Level
	switch s.LogLvl {
	case "info":
		l = log.InfoLevel
	case "warn":
		l = log.WarnLevel
	case "error":
		l = log.ErrorLevel
	case "debug":
		l = log.DebugLevel
	default:
		l = log.InfoLevel
	}
	return l
}
