package simulator

import (
	"log"
	"time"

	"github.com/context-builder/internal/contextTree"
	"github.com/context-builder/internal/pubsub"
)

func SimulateRealWorld() {
	// Initialize pubsub client for publishing
	natsPubSub, err := pubsub.NewNatsPubSub()
	if err != nil {
		panic(err)
	}

	// manager := contextTree.NewContextTreeManager()
	manager := contextTree.NewContextTreeManager(natsPubSub)

	// Add root and comment before adding nodes
	log.Println("INIT : Initializing context-builder app..")
	time.Sleep(1 * time.Second) // adding delay to simulate real time

	log.Println("TASK : New incoming event")
	manager.Tree.Root = "pharma"
	manager.Tree.Comment = "pharma down | inc123 | z:9867"
	manager.Tree.Timestamp = time.Now().Unix()
	manager.PrintContextTree()
	time.Sleep(1 * time.Second) // adding delay to simulate real time

	// Add nodes with dependencies
	log.Println("TASK : New incoming task for event")
	time.Sleep(1 * time.Second)
	manager.AddNode("opm", []string{})
	manager.PrintContextTree()

	// Add nodes with dependencies
	log.Println("TASK : Adding new dependency")
	time.Sleep(1 * time.Second)
	manager.AddDependency("opm", "errors")
	manager.PrintContextTree()

	// Add nodes with dependencies
	log.Println("TASK : Adding new dependency")
	time.Sleep(1 * time.Second)
	manager.AddDependency("opm", "connection")
	manager.PrintContextTree()

	// Add nodes with dependencies
	log.Println("TASK : New incoming task")
	time.Sleep(1 * time.Second)
	manager.AddNode("users", []string{"someuser"})
	manager.PrintContextTree()

	// Add nodes with dependencies
	log.Println("TASK : New incoming task")
	time.Sleep(1 * time.Second)
	manager.AddNode("infra", []string{"disk", "network"})
	manager.PrintContextTree()

	// Add nodes with dependencies
	log.Println("TASK : New incoming task")
	time.Sleep(1 * time.Second)
	manager.AddNode("network", []string{"dns"})
	manager.PrintContextTree()

	// Add specific dependency for network
	// Add nodes with dependencies
	log.Println("TASK : Adding new dependency")
	time.Sleep(1 * time.Second)
	manager.AddDependency("network", "dnspoison")
	manager.PrintContextTree()

	// Delete nodes with dependencies
	log.Println("TASK : Deleting a node")
	time.Sleep(1 * time.Second)
	manager.DeleteNode("users")
	manager.PrintContextTree()

	// Delete new dependency
	log.Println("TASK : Deleting a dependency")
	time.Sleep(1 * time.Second)
	manager.DeleteDependency("infra", "disk")
	manager.PrintContextTree()
}

/*
Error :
Though i have `import github.com/iamshreeram/context-builder/internal/contextTree`, go is throwing error that says, "No required module provides package". how to fix this? but, this import `import github.com/context-builder/internal/contextTree` is working well without error.
*/
