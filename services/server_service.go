package services

import (
	"cloud-cli/models"
	"cloud-cli/repository"
	"errors"
	"fmt"
)

type ServerService struct {
	Repo *repository.ResourceRepository
}

func NewServerService(repo *repository.ResourceRepository) *ServerService {
	return &ServerService{
		Repo: repo,
	}
}

func (v *ServerService) Start(id int) error {
	resource, err := v.Repo.GetResourceByID(id)
	if err != nil {
		return err
	}
	if resource.Status == models.Running {
		return errors.New("Server is already running")
	}
	resource.Status = models.Running
	fmt.Printf("Server %d started successfully \n", resource.ID)
	return nil
}

func (v *ServerService) Stop(id int) error {
	resource, err := v.Repo.GetResourceByID(id)
	if err != nil {
		return err
	}
	if resource.Status != models.Running {
		return errors.New("Server is not running")
	}
	resource.Status = models.Stopped
	fmt.Printf("Server %d stopped successfully \n", resource.ID)
	return nil
}

func (v *ServerService) Terminate(id int) error {
	resource, err := v.Repo.GetResourceByID(id)
	if err != nil {
		return err
	}
	if resource.Status == models.Terminated {
		return errors.New("Server is already terminated")
	}
	resource.Status = models.Terminated
	fmt.Printf("Server %d terminated successfully \n", resource.ID)
	return nil
}

func (v *ServerService) Restart(id int) error {
	resource, err := v.Repo.GetResourceByID(id)
	if err != nil {
		return err
	}
	if resource.Status == models.Terminated {
		return errors.New("cannot restart a terminated Server")
	}

	if resource.Status == models.Running {
		fmt.Printf("Server %d is running, stopping before restart...\n", resource.ID)
		resource.Status = models.Stopped
	}

	resource.Status = models.Running
	fmt.Printf("Server %d restarted successfully \n", resource.ID)
	return nil
}

func (v *ServerService) StatusCheck(id int) error {
	resource, err := v.Repo.GetResourceByID(id)
	if err != nil {
		return err
	}
	fmt.Printf("Server %d current status: %s\n", resource.ID, resource.Status)
	return nil
}
