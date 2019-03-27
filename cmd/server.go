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
	"github.com/jharshman/simple-client-server/config"
	"github.com/jharshman/simple-client-server/pkg/server"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfg    config.Server
	port   string
	loglvl string
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the server component.",
	Long:  `Run the echo server`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := viper.Unmarshal(&cfg); err != nil {
			log.Fatalf("unmarshal server config: %v\n", err)
		}
		server.Start(&cfg)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	rootCmd.Flags().StringVar(&port, "port", "9000", "grpc bind address")
	rootCmd.Flags().StringVar(&loglvl, "loglvl", "info", "log level")

	viper.BindPFlags(rootCmd.Flags())
}
