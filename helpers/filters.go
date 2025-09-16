package helpers

import "cloud-cli/models"

func Filter(resources []models.Resource, predicate func(models.Resource) bool) []models.Resource {
    var result []models.Resource
    for _, r := range resources {
        if predicate(r) {
            result = append(result, r)
        }
    }
    return result
}
