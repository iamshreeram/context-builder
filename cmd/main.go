package main

import (
	"log"

	"github.com/context-builder/internal/contextTree"
	"github.com/context-builder/internal/natsserver"
	"github.com/context-builder/internal/pubsub"
	"github.com/context-builder/simulator"

	"github.com/nats-io/nats-server/v2/server"
)

func main() {

	// Logging the start of NATS server
	log.Println("Starting NATS server..")

	// Initializing NATS Server
	natsServer := natsserver.NewNatsServer()
	ns := natsServer.StartNatsServer(&server.Options{})

	// ns.InProcessConn() // Bypassing TCP for internal Connections (future scope)

	// Waiting for the server to shutdown
	defer ns.WaitForShutdown()

	// Initialize nats pubsub client for context-tree to publish events
	natsPubSub, err := pubsub.NewNatsPubSub()
	if err != nil {
		panic(err)
	}

	// Creating a new ContextTreeManager with the initialized natsPubSub
	manager := contextTree.NewContextTreeManager(natsPubSub)

	// Running real world simulator that publishes events and deletes events
	go simulator.SimulateRealWorld(manager)
}
