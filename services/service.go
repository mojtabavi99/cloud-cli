package services

import "cloud-cli/models"

type Service interface {
    Start(resource *models.Resource) error
    Stop(resource *models.Resource) error
    Terminate(resource *models.Resource) error
}
