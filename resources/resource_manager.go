package resources

import (
	"cloud-cli/models"
	"errors"
	"fmt"
)

// نگهداری منابع هر سرویس
var VMMap = make(map[int]*models.Resource)
var DBMap = make(map[int]*models.Resource)
var StorageMap = make(map[int]*models.Resource)

// AddVMResource ایجاد یک VM جدید
func AddVMResource(id int, specs map[models.ResourceSpecs]string, status models.ResourceStatus) (*models.Resource, error) {
	if _, exists := VMMap[id]; exists {
		return nil, errors.New("VM with this ID already exists")
	}

	vm := &models.Resource{
		ID:     id,
		Type:   models.Server,
		Status: status,
		Specs:  specs,
	}

	VMMap[id] = vm
	fmt.Printf("VM %d created successfully \n", id)
	return vm, nil
}

// AddDBResource ایجاد یک Database جدید
func AddDBResource(id int, specs map[models.ResourceSpecs]string, status models.ResourceStatus) (*models.Resource, error) {
	if _, exists := DBMap[id]; exists {
		return nil, errors.New("Database with this ID already exists")
	}

	db := &models.Resource{
		ID:     id,
		Type:   models.Database,
		Status: status,
		Specs:  specs,
	}

	DBMap[id] = db
	fmt.Printf("Database %d created successfully ✅\n", id)
	return db, nil
}

// AddStorageResource ایجاد یک Storage جدید
func AddStorageResource(id int, specs map[models.ResourceSpecs]string, status models.ResourceStatus) (*models.Resource, error) {
	if _, exists := StorageMap[id]; exists {
		return nil, errors.New("Storage with this ID already exists")
	}

	storage := &models.Resource{
		ID:     id,
		Type:   models.Storage,
		Status: status,
		Specs:  specs,
	}

	StorageMap[id] = storage
	fmt.Printf("Storage %d created successfully ✅\n", id)
	return storage, nil
}
