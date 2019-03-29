package config

import (
	"github.com/matryer/is"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestServer_BindPort(t *testing.T) {

	tests := []struct {
		before string
		after  string
	}{
		{before: "8080", after: ":8080"},
		{before: "443", after: ":443"},
		{before: "80", after: ":80"},
	}

	is := is.New(t)

	for _, v := range tests {
		t.Run(v.before, func(t *testing.T) {
			c := Server{
				Port: v.before,
			}
			p := c.BindPort()
			is.Equal(p, v.after)
		})
	}
}

func TestServer_LogLevel(t *testing.T) {

	tests := []struct {
		before string
		after  log.Level
	}{
		{before: "info", after: log.InfoLevel},
		{before: "warn", after: log.WarnLevel},
		{before: "error", after: log.ErrorLevel},
		{before: "debug", after: log.DebugLevel},
		{before: "bogus", after: log.InfoLevel},
	}

	is := is.New(t)

	for _, v := range tests {
		t.Run(v.before, func(t *testing.T) {
			c := Server{
				LogLvl: v.before,
			}
			l := c.LogLevel()
			is.Equal(l, v.after)
		})
	}
}
