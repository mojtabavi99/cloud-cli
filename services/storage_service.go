package services

import (
	"cloud-cli/models"
	"cloud-cli/repository"
	"errors"
	"fmt"
)

type StorageService struct {
	Repo *repository.ResourceRepository
}

func NewStorageService(repo *repository.ResourceRepository) *StorageService {
	return &StorageService{
		Repo: repo,
	}
}

// Start برای Storage معمولاً به Mount شدن اشاره دارد
func (s *StorageService) Start(id int) error {
	resource, err := s.Repo.GetResourceByID(id)
	if err != nil {
		return err
	}
	if resource.Status == models.Running {
		return errors.New("Storage is already mounted")
	}
	resource.Status = models.Running
	fmt.Printf("Storage %d mounted successfully \n", resource.ID)
	return nil
}

// Stop برای Storage به Unmount شدن اشاره دارد
func (s *StorageService) Stop(id int) error {
	resource, err := s.Repo.GetResourceByID(id)
	if err != nil {
		return err
	}
	if resource.Status != models.Running {
		return errors.New("Storage is not mounted")
	}
	resource.Status = models.Stopped
	fmt.Printf("Storage %d unmounted successfully \n", resource.ID)
	return nil
}

// Terminate حذف کامل Storage
func (s *StorageService) Terminate(id int) error {
	resource, err := s.Repo.GetResourceByID(id)
	if err != nil {
		return err
	}
	if resource.Status == models.Terminated {
		return errors.New("Storage is already terminated")
	}
	resource.Status = models.Terminated
	fmt.Printf("Storage %d terminated successfully \n", resource.ID)
	return nil
}

// Restart در Storage می‌تواند Unmount و Mount مجدد باشد
func (s *StorageService) Restart(id int) error {
	resource, err := s.Repo.GetResourceByID(id)
	if err != nil {
		return err
	}
	if resource.Status == models.Terminated {
		return errors.New("cannot restart a terminated Storage")
	}

	if resource.Status == models.Running {
		fmt.Printf("Storage %d is mounted, unmounting before restart...\n", resource.ID)
		resource.Status = models.Stopped
	}

	resource.Status = models.Running
	fmt.Printf("Storage %d remounted successfully \n", resource.ID)
	return nil
}

// StatusCheck وضعیت فعلی Storage را نمایش می‌دهد
func (s *StorageService) StatusCheck(id int) error {
	resource, err := s.Repo.GetResourceByID(id)
	if err != nil {
		return err
	}
	fmt.Printf("Storage %d current status: %s\n", resource.ID, resource.Status)
	return nil
}
