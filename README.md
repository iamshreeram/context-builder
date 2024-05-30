# Context Builder

## Summary:
Context Builder is a simple Golang application that allows users to build, manage, and visualize context trees for any given scenario. It enables users to create a hierarchical structure of contexts and their relationships, stored as JSON in memory. This tool is designed to provide a flexible and user-friendly way to organize and understand complex relationships between different elements.

## Description:
The Context Builder allows users to construct a context tree by defining various contexts and establishing relationships between them. Users can easily add, edit, delete, and view contexts and their dependencies. This tool is particularly useful for troubleshooting, categorizing information, and organizing data in a structured manner.

## Use Cases:
**Troubleshooting**: When diagnosing issues or problems, users can create a context tree to map out potential causes and their relationships to find the root cause efficiently.

**Knowledge Management**: Organizations can utilize this tool to organize and categorize information, making it easier to retrieve and understand complex relationships between different concepts.

**Process Mapping**: By creating a context tree for a specific process or workflow, users can visualize the steps, dependencies, and interactions involved in a clear and concise manner.

**Project Management**: Team members can collaborate on building context trees to plan, track, and communicate project requirements, dependencies, and tasks effectively.

### Run the app:

To start the context-builder example, use the following command in the terminal:
```
go run cmd/main.go
```

**Context-tree Payload** :

```json
{
  "root": "pharma",
  "comment": "pharma down | inc123 | z:9867",
  "timestamp": 1717037126,
  "nodes": [
    {
      "id": "opm",
      "dependency": [
        "errors",
        "connection"
      ],
      "timestamp": 1717037122
    },
    {
      "id": "users",
      "dependency": [
        "someuser"
      ],
      "timestamp": 1717037125
    },
    {
      "id": "infra",
      "dependency": [
        "disk",
        "network"
      ],
      "timestamp": 1717037126
    }
  ]
}

```

### Current features
1. Agents can create a new event in the context tree.
2. Agent are able to add or remove tasks associated with the event.
3. Agent can add new dependencies to the tasks in the context tree.

### Upcoming feature:
1. Agents will have the ability to subscribe to receive notifications about changes in the context tree.