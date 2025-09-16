package resources

import (
	"cloud-cli/models"
)

var Resources []models.Resource
var ResourceMap = make(map[int]*models.Resource)

func AddResource(id int, rType models.ResourceType, status models.ResourceStatus, specs map[models.ResourceSpecs]string) {
    resource := models.Resource{
        ID:     id,
        Type:   rType,
        Status: status,
        Specs:  specs,
    }
    Resources = append(Resources, resource)
    ResourceMap[id] = &Resources[len(Resources)-1]
}

func StartResource(id int) {
    if r, ok := ResourceMap[id]; ok {
        r.Status = models.Running
    }
}

func StopResource(id int) {
    if r, ok := ResourceMap[id]; ok {
        r.Status = models.Stopped
    }
}

func TerminateResource(id int) {
    if r, ok := ResourceMap[id]; ok {
        r.Status = models.Terminated
    }
}
