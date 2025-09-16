package services

import (
	"cloud-cli/models"
	"errors"
	"fmt"
)

type VMService struct{}

func (v *VMService) Start(resource *models.Resource) error {
	if resource.Status == models.Running {
		return errors.New("VM is already running")
	}
	resource.Status = models.Running
	fmt.Printf("VM %d started successfully \n", resource.ID)
	return nil
}

func (v *VMService) Stop(resource *models.Resource) error {
	if resource.Status != models.Running {
		return errors.New("VM is not running")
	}
	resource.Status = models.Stopped
	fmt.Printf("VM %d stopped successfully \n", resource.ID)
	return nil
}

func (v *VMService) Terminate(resource *models.Resource) error {
	if resource.Status == models.Terminated {
		return errors.New("VM is already terminated")
	}
	resource.Status = models.Terminated
	fmt.Printf("VM %d terminated successfully \n", resource.ID)
	return nil
}

func (v *VMService) Restart(resource *models.Resource) error {
	if resource.Status == models.Terminated {
		return errors.New("cannot restart a terminated VM")
	}

	if resource.Status == models.Running {
		fmt.Printf("VM %d is running, stopping before restart...\n", resource.ID)
		resource.Status = models.Stopped
	}

	resource.Status = models.Running
	fmt.Printf("VM %d restarted successfully \n", resource.ID)
	return nil
}

func (v *VMService) StatusCheck(resource *models.Resource) {
	fmt.Printf("VM %d current status: %s\n", resource.ID, resource.Status)
}


