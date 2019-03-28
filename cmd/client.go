/*
Copyright © 2019 Joshua Harshman

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"context"
	"github.com/jharshman/simple-client-server/config"
	"github.com/jharshman/simple-client-server/pkg/grpc"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	rpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"

	"github.com/spf13/cobra"
)

var (
	clientCfg  config.Client
	serverAddr string
	serverPort string
	msg        string
	clientCert string
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Echo Client",
	Long: `
Facilitates communication to Echo Server.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := viper.Unmarshal(&clientCfg); err != nil {
			log.Fatalf("unmarshal client config: %v\n", err)
		}
		// run client
		if clientCfg.Message == "" {
			log.Fatalln("no message to send")
		}

		tls, err := credentials.NewClientTLSFromFile(clientCfg.CertFile, "")
		if err != nil {
			log.Fatalf("%v\n", err)
		}
		opts := []rpc.DialOption{
			rpc.WithTransportCredentials(tls),
		}
		conn, err := rpc.Dial(net.JoinHostPort(clientCfg.Addr, clientCfg.Port), opts...)
		if err != nil {
			log.Fatalf("grpc dial: %v\n", err)
		}
		defer conn.Close()

		echoCli := grpc.NewEchoServerClient(conn)
		res, err := echoCli.Echo(context.Background(), &grpc.Message{
			Data: msg,
		})
		if err != nil {
			log.Fatalf("receiving response: %v\n", err)
		}

		log.Infof("%s\n", res.Data)
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)

	clientCmd.Flags().StringVar(&serverAddr, "server", "127.0.0.1", "address of echo server")
	clientCmd.Flags().StringVar(&serverPort, "server-port", "9000", "port of echo server")
	clientCmd.Flags().StringVar(&msg, "msg", "", "message to send to echo server")
	clientCmd.Flags().StringVar(&clientCert, "client-cert", "cert.pem", "client tls certificate")

	viper.BindPFlags(clientCmd.Flags())
}
