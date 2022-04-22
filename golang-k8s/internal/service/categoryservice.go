package service

import (
	"context"
	"errors"
	"fmt"
	"golang-k8s/internal/domain/enitity"
	"golang-k8s/internal/domain/repository"
)

type CategoryService interface {
	GetAll(ctx context.Context) ([]enitity.Category, error)
	GetById(ctx context.Context, id string) (*enitity.Category, error)
	Insert(ctx context.Context, category *enitity.Category) error
	UpdateById(ctx context.Context, id string, category *enitity.Category) error
	DeleteById(ctx context.Context, id string) error
}

type categoryService struct {
	categoryRepository repository.CategoryRepository
}

func (self *categoryService) GetAll(ctx context.Context) (categories []enitity.Category, err error) {

	if categories, err = self.categoryRepository.GetAll(ctx); err != nil {
		return nil, err
	}

	return categories, nil

}

func (self *categoryService) GetById(ctx context.Context, id string) (category *enitity.Category, err error) {

	if category, err = self.categoryRepository.GetById(ctx, id); err != nil {
		return nil, err
	}

	return category, err
}

func (self *categoryService) Insert(ctx context.Context, category *enitity.Category) (err error) {
	return self.categoryRepository.Insert(ctx, category)
}

func (self *categoryService) UpdateById(ctx context.Context, id string, updateEntity *enitity.Category) error {

	var (
		category *enitity.Category
		err      error
	)

	if category, err = self.GetById(ctx, id); err != nil {
		return err
	}

	if category == nil {
		return errors.New(fmt.Sprintf("%s Category not found.", id))
	}

	return self.categoryRepository.UpdateById(ctx, id, updateEntity)
}

func (self *categoryService) DeleteById(ctx context.Context, id string) error {

	var (
		category *enitity.Category
		err      error
	)

	if category, err = self.GetById(ctx, id); err != nil {
		return err
	}

	if category == nil {
		return errors.New(fmt.Sprintf("%s Category not found.", id))
	}

	return self.categoryRepository.RemoveById(ctx, id)
}

func NewCategoryService(categoryRepository repository.CategoryRepository) CategoryService {
	return &categoryService{categoryRepository: categoryRepository}
}
