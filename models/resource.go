package models

import "fmt"

type ResourceType string
const (
    Server   ResourceType = "Server"
    Database ResourceType = "Database"
    Storage  ResourceType = "Storage"
)

type ResourceStatus string
const (
    Running    ResourceStatus = "Running"
    Stopped    ResourceStatus = "Stopped"
    Terminated ResourceStatus = "Terminated"
)

type ResourceSpecs string
const (
    CPU     ResourceSpecs = "CPU"
    Memory  ResourceSpecs = "Memory"
    Stotage ResourceSpecs = "Storage"
)

type Resource struct {
    ID     int
    Type   ResourceType
    Status ResourceStatus
    Specs  map[ResourceSpecs]string
}

func (r Resource) Display() string {
    return fmt.Sprintf("ID: %d | Type: %s | Status: %s | Specs: %v",
        r.ID, r.Type, r.Status, r.Specs)
}
