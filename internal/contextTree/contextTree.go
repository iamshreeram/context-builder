package contextTree

import (
	"encoding/json"
	"fmt"
	"time"
)

type Node struct {
	ID         string   `json:"id"`
	Dependency []string `json:"dependency"`
	Timestamp  int64    `json:"timestamp"`
}

type ContextTree struct {
	Root      string          `json:"root"`
	Comment   string          `json:"comment"`
	Timestamp int64           `json:"timestamp"`
	Nodes     map[string]Node `json:"nodes"`
}

type ContextTreeManager struct {
	Tree ContextTree
}

func NewContextTreeManager() *ContextTreeManager {
	return &ContextTreeManager{}
}

func (m *ContextTreeManager) AddNode(id string, dependencies []string) {
	if m.Tree.Nodes == nil {
		m.Tree.Nodes = make(map[string]Node)
	}

	if m.Tree.Root != "" && m.Tree.Comment != "" {
		currentTime := time.Now().Unix()

		// Update timestamp for root node
		m.Tree.Timestamp = currentTime

		// Add new node with updated timestamp
		node := Node{ID: id, Dependency: dependencies, Timestamp: currentTime}
		m.Tree.Nodes[id] = node
	} else {
		fmt.Println("Cannot add node without root and comment")
	}
}

func (m *ContextTreeManager) PrintContextTree() {
	treeJson, _ := json.MarshalIndent(m.Tree, "", "  ")
	fmt.Println(string(treeJson))
}

func (m *ContextTreeManager) DeleteNode(id string) {
	delete(m.Tree.Nodes, id)
}

func (m *ContextTreeManager) AddDependency(id, dependency string) {
	if node, ok := m.Tree.Nodes[id]; ok {
		node.Dependency = append(node.Dependency, dependency)
	} else {
		fmt.Println("Node not found")
	}
}

func (m *ContextTreeManager) DeleteDependency(nodeID, dependency string) {
	if node, ok := m.Tree.Nodes[nodeID]; ok {
		for i, dep := range node.Dependency {
			if dep == dependency {
				node.Dependency = append(node.Dependency[:i], node.Dependency[i+1:]...)
				break
			}
		}
	} else {
		fmt.Println("Node not found")
	}
}
