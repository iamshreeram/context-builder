package main

import (
	"time"

	"github.com/context-builder/internal/contextTree"
)

func logger(msg string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	println(timestamp + " : " + msg)
}

func main() {
	manager := contextTree.NewContextTreeManager()

	// Add root and comment before adding nodes

	logger("INIT : Initializing context-builder app..")
	time.Sleep(1 * time.Second) // adding delay to simulate real time

	logger("TASK : New incoming event")

	manager.Tree.Root = "pharma"
	manager.Tree.Comment = "pharma down | inc123 | z:9867"
	manager.Tree.Timestamp = time.Now().Unix()
	manager.PrintContextTree()
	time.Sleep(1 * time.Second) // adding delay to simulate real time

	// Add nodes with dependencies
	logger("TASK : New incoming task for event")
	time.Sleep(1 * time.Second)
	manager.AddNode("opm", []string{})
	manager.PrintContextTree()

	// Add nodes with dependencies
	logger("TASK : Adding new dependency")
	time.Sleep(1 * time.Second)
	manager.AddDependency("opm", "errors")
	manager.PrintContextTree()

	// Add nodes with dependencies
	logger("TASK : Adding new dependency")
	time.Sleep(1 * time.Second)
	manager.AddDependency("opm", "connection")
	manager.PrintContextTree()

	// Add nodes with dependencies
	logger("TASK : New incoming task")
	time.Sleep(1 * time.Second)
	manager.AddNode("users", []string{"someuser"})
	manager.PrintContextTree()

	// Add nodes with dependencies
	logger("TASK : New incoming task")
	time.Sleep(1 * time.Second)
	manager.AddNode("infra", []string{"disk", "network"})
	manager.PrintContextTree()

	// Add nodes with dependencies
	logger("TASK : New incoming task")
	time.Sleep(1 * time.Second)
	manager.AddNode("network", []string{"dns"})
	manager.PrintContextTree()

	// Add specific dependency for network
	// Add nodes with dependencies
	logger("TASK : Adding new dependency")
	time.Sleep(1 * time.Second)
	manager.AddDependency("network", "dnspoison")
	manager.PrintContextTree()

	// Delete nodes with dependencies
	logger("TASK : Deleting a node")
	time.Sleep(1 * time.Second)
	manager.DeleteNode("users")
	manager.PrintContextTree()

	// Delete new dependency
	logger("TASK : Deleting a dependency")
	time.Sleep(1 * time.Second)
	manager.DeleteDependency("infra", "disk")
	manager.PrintContextTree()
}
