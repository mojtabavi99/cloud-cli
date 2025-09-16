package services

import (
	"cloud-cli/models"
	"errors"
	"fmt"
)

type DatabaseService struct{}

func (d *DatabaseService) Start(resource *models.Resource) error {
	if resource.Status == models.Running {
		return errors.New("Database is already running")
	}
	resource.Status = models.Running
	fmt.Printf("Database %d started successfully \n", resource.ID)
	return nil
}

func (d *DatabaseService) Stop(resource *models.Resource) error {
	if resource.Status != models.Running {
		return errors.New("Database is not running")
	}
	resource.Status = models.Stopped
	fmt.Printf("Database %d stopped successfully \n", resource.ID)
	return nil
}

func (d *DatabaseService) Terminate(resource *models.Resource) error {
	if resource.Status == models.Terminated {
		return errors.New("Database is already terminated")
	}
	resource.Status = models.Terminated
	fmt.Printf("Database %d terminated successfully \n", resource.ID)
	return nil
}

func (d *DatabaseService) Restart(resource *models.Resource) error {
	if resource.Status == models.Terminated {
		return errors.New("cannot restart a terminated Database")
	}

	if resource.Status == models.Running {
		fmt.Printf("Database %d is running, stopping before restart...\n", resource.ID)
		resource.Status = models.Stopped
	}

	resource.Status = models.Running
	fmt.Printf("Database %d restarted successfully \n", resource.ID)
	return nil
}

func (d *DatabaseService) StatusCheck(resource *models.Resource) {
	fmt.Printf("Database %d current status: %s\n", resource.ID, resource.Status)
}
