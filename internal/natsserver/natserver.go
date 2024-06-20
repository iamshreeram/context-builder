package natsserver

import (
	"time"

	"github.com/nats-io/nats-server/v2/server"
)

type natsServer struct {
	server *server.Server
}

type Server interface {
	StartNatsServer(opts *server.Options) *server.Server
	StopNatsServer() *server.Server
}

func NewNatsServer() Server {
	return &natsServer{}
}

func (ns *natsServer) StartNatsServer(opts *server.Options) *server.Server {
	// Initialize new server with options
	ns.server, _ = server.NewServer(opts)
	// Start the server via goroutine
	go ns.server.Start()
	// server.FlagSnapshot.Debug = true

	// Wait for server to be ready for connections
	if !ns.server.ReadyForConnections(10 * time.Second) {
		panic("not ready for connection")
	}
	return ns.server
}

func (ns *natsServer) StopNatsServer() *server.Server {
	// Shutdown the server (optional)
	ns.server.Shutdown()
	return ns.server
}
