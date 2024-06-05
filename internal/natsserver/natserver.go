package natsserver

import (
	"time"

	"github.com/nats-io/nats-server/v2/server"
)

type Server interface {
	StartNatsServer() error
	StopNatsServer() error
}

func StartNatsServer(opts *server.Options) *server.Server {
	// Initialize new server with options
	ns, err := server.NewServer(opts)
	if err != nil {
		panic(err)
	}

	// Start the server via goroutine
	go ns.Start()

	// Wait for server to be ready for connections
	if !ns.ReadyForConnections(4 * time.Second) {
		panic("not ready for connection")
	}
	return ns
}

func StopNatsServer(ns *server.Server) *server.Server {
	// Shutdown the server (optional)
	ns.Shutdown()
	return ns
}
