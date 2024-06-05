package main

import (
	"log"

	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"

	"github.com/context-builder/internal/natsserver"
	"github.com/context-builder/simulator"
)

func main() {
	log.Println("Starting NATS server..")

	// Initializing NATS Server
	opts := &server.Options{}
	ns := natsserver.StartNatsServer(opts)
	// ns.InProcessConn() // Bypassing TCP for internal Connections
	defer ns.WaitForShutdown()

	// Validating NATS Server is up
	nc, err := nats.Connect(ns.ClientURL())
	if err != nil {
		log.Fatalf("Error connecting to NATS: %v", err)
	}
	defer nc.Close()

	// Running real world simulator that publishes events and deletes events
	go simulator.SimulateRealWorld()
}
