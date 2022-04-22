package service

import (
	"context"
	"errors"
	"fmt"
	"golang-k8s/internal/domain/enitity"
	"golang-k8s/internal/domain/repository"
	"golang-k8s/internal/service/models"
	"golang-k8s/pkg/httpclient/category"
)

type ProductService interface {
	GetAll(ctx context.Context) ([]enitity.Product, error)
	GetById(ctx context.Context, id string) (*enitity.Product, error)
	GetProductDetailById(ctx context.Context, id string) (*models.ProductDetail, error)
	Insert(ctx context.Context, product *enitity.Product) error
	UpdateById(ctx context.Context, id string, product *enitity.Product) error
	DeleteById(ctx context.Context, id string) error
	Ping(ctx context.Context) error
}

type productService struct {
	productRepository repository.ProductRepository
	categoryClient    category.Client
}

func (self *productService) Ping(ctx context.Context) error {
	return self.productRepository.Ping(ctx)
}

func (self *productService) GetAll(ctx context.Context) (categories []enitity.Product, err error) {

	if categories, err = self.productRepository.GetAll(ctx); err != nil {
		return nil, err
	}

	return categories, nil

}

func (self *productService) GetById(ctx context.Context, id string) (product *enitity.Product, err error) {

	if product, err = self.productRepository.GetById(ctx, id); err != nil {
		return nil, err
	}

	return product, err
}

func (self *productService) GetProductDetailById(ctx context.Context, id string) (*models.ProductDetail, error) {

	var (
		product          *enitity.Product
		categoryResponse *category.GetCategoryByIdResponse
		err              error
	)

	if product, err = self.GetById(ctx, id); err != nil {
		return nil, err
	}

	if categoryResponse, err = self.categoryClient.GetCategoryById(product.CategoryId); err != nil {
		return nil, err
	}

	if categoryResponse == nil {
		return nil, errors.New("product category not found")
	}

	response := models.ProductDetail{
		Name:  product.Name,
		Price: product.Price,
		Category: models.CategoryItem{
			Id:   categoryResponse.Id,
			Name: categoryResponse.Name,
		},
	}

	return &response, err
}

func (self *productService) Insert(ctx context.Context, product *enitity.Product) (err error) {
	return self.productRepository.Insert(ctx, product)
}

func (self *productService) UpdateById(ctx context.Context, id string, updateEntity *enitity.Product) error {

	var (
		product *enitity.Product
		err     error
	)

	if product, err = self.GetById(ctx, id); err != nil {
		return err
	}

	if product == nil {
		return errors.New(fmt.Sprintf("%s Product not found.", id))
	}

	return self.productRepository.UpdateById(ctx, id, updateEntity)
}

func (self *productService) DeleteById(ctx context.Context, id string) error {

	var (
		product *enitity.Product
		err     error
	)

	if product, err = self.GetById(ctx, id); err != nil {
		return err
	}

	if product == nil {
		return errors.New(fmt.Sprintf("%s Product not found.", id))
	}

	return self.productRepository.RemoveById(ctx, id)
}

func NewProductService(productRepository repository.ProductRepository, categoryClient category.Client) ProductService {
	return &productService{productRepository: productRepository, categoryClient: categoryClient}
}
