package services

import (
	"cloud-cli/models"
	"cloud-cli/repository"
	"errors"
	"fmt"
)

type DatabaseService struct {
	Repo *repository.ResourceRepository
}

func NewDatabaseService(repo *repository.ResourceRepository) *DatabaseService {
	return &DatabaseService{
		Repo: repo,
	}
}

func (d *DatabaseService) Start(id int) error {
	resource, err := d.Repo.GetResourceByID(id)
	if err != nil {
		return err
	}
	if resource.Status == models.Running {
		return errors.New("Database is already running")
	}
	resource.Status = models.Running
	fmt.Printf("Database %d started successfully \n", resource.ID)
	return nil
}

func (d *DatabaseService) Stop(id int) error {
	resource, err := d.Repo.GetResourceByID(id)
	if err != nil {
		return err
	}
	if resource.Status != models.Running {
		return errors.New("Database is not running")
	}
	resource.Status = models.Stopped
	fmt.Printf("Database %d stopped successfully \n", resource.ID)
	return nil
}

func (d *DatabaseService) Terminate(id int) error {
	resource, err := d.Repo.GetResourceByID(id)
	if err != nil {
		return err
	}
	if resource.Status == models.Terminated {
		return errors.New("Database is already terminated")
	}
	resource.Status = models.Terminated
	fmt.Printf("Database %d terminated successfully \n", resource.ID)
	return nil
}

func (d *DatabaseService) Restart(id int) error {
	resource, err := d.Repo.GetResourceByID(id)
	if err != nil {
		return err
	}
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

func (d *DatabaseService) StatusCheck(id int) error {
	resource, err := d.Repo.GetResourceByID(id)
	if err != nil {
		return err
	}
	fmt.Printf("Database %d current status: %s\n", resource.ID, resource.Status)
	return nil
}
