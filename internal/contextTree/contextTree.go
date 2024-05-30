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
	Root      string `json:"root"`
	Comment   string `json:"comment"`
	Timestamp int64  `json:"timestamp"`
	Nodes     []Node `json:"nodes"`
}

type ContextTreeManager struct {
	Tree ContextTree
}

func NewContextTreeManager() *ContextTreeManager {
	return &ContextTreeManager{}
}

/*
func (m *ContextTreeManager) AddNode(id string, dependencies []string) {
	if m.Tree.Root != "" && m.Tree.Comment != "" {
		node := Node{ID: id, Dependency: dependencies, Timestamp: time.Now().Unix()}
		m.Tree.Nodes = append(m.Tree.Nodes, node)
	} else {
		fmt.Println("Cannot add node without root and comment")
	}
}
*/

func (m *ContextTreeManager) AddNode(id string, dependencies []string) {
	if m.Tree.Root != "" && m.Tree.Comment != "" {
		currentTime := time.Now().Unix()

		// Update timestamp for root node
		m.Tree.Timestamp = currentTime

		// Add new node with updated timestamp
		node := Node{ID: id, Dependency: dependencies, Timestamp: currentTime}
		m.Tree.Nodes = append(m.Tree.Nodes, node)
	} else {
		fmt.Println("Cannot add node without root and comment")
	}
}

func (m *ContextTreeManager) DeleteNode(id string) {
	for i, node := range m.Tree.Nodes {
		if node.ID == id {
			m.Tree.Nodes = append(m.Tree.Nodes[:i], m.Tree.Nodes[i+1:]...)
			break
		}
	}
}

func (m *ContextTreeManager) PrintContextTree() {
	treeJson, _ := json.MarshalIndent(m.Tree, "", "  ")
	fmt.Println(string(treeJson))
}

func (m *ContextTreeManager) AddDependency(id, dependency string) {
	for i, node := range m.Tree.Nodes {
		if node.ID == id {
			m.Tree.Nodes[i].Dependency = append(m.Tree.Nodes[i].Dependency, dependency)
			break
		}
	}
}

func (m *ContextTreeManager) DeleteDependency(nodeID, dependency string) {
	for i, node := range m.Tree.Nodes {
		if node.ID == nodeID {
			for j, dep := range m.Tree.Nodes[i].Dependency {
				if dep == dependency {
					m.Tree.Nodes[i].Dependency = append(m.Tree.Nodes[i].Dependency[:j], m.Tree.Nodes[i].Dependency[j+1:]...)
					break
				}
			}
		}
	}
}
