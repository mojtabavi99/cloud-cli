package repository

import (
	"cloud-cli/models"
	"errors"
)

type ResourceRepository struct {
	resources []*models.Resource
	idCounter int
}

// ایجاد repository جدید
func NewResourceRepository() *ResourceRepository {
	return &ResourceRepository{
		resources: []*models.Resource{},
		idCounter: 0,
	}
}

// افزودن Resource
func (r *ResourceRepository) AddResource(resource *models.Resource) error {
	for _, res := range r.resources {
		if res.ID == resource.ID {
			return errors.New("resource with this ID already exists")
		}
	}
	r.resources = append(r.resources, resource)
	return nil
}

// دریافت Resource بر اساس ID
func (r *ResourceRepository) GetResourceByID(id int) (*models.Resource, error) {
	for _, res := range r.resources {
		if res.ID == id {
			return res, nil
		}
	}
	return nil, errors.New("resource not found")
}

// گرفتن تمام منابع
func (r *ResourceRepository) GetAllResources() []*models.Resource {
	return r.resources
}

// تولید ID یکتا برای Resource جدید
func (r *ResourceRepository) NextID() int {
	r.idCounter++
	return r.idCounter
}
