package tcp_conn

import (
	"gim/config"
	"gim/pkg/logger"
	"time"

	"github.com/alberliu/gn"
)

var server *gn.Server

func StartTCPServer() {
	var err error
	server, err = gn.NewServer(config.TCPConn.TCPListenAddr, &handler{}, 2, 254, 1024, 1000)
	if err != nil {
		logger.Sugar.Error(err)
		panic(err)
	}

	server.SetTimeout(5*time.Second, 11*time.Minute)
	server.Run()
}
