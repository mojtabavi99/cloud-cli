package services

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
