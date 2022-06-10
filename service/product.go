package service

import (
	"github.com/zipzoft/interview-k-donn/config"
	"github.com/zipzoft/interview-k-donn/entity"
	"github.com/zipzoft/interview-k-donn/repository"
)

type ProductService struct {
	repository repository.ProductRepository
}

func NewProductService() *ProductService {
	// Initial product repository
	var product repository.ProductRepository = repository.NewProductJsonFileRepository(
		config.ResolveBasePath("products.json"),
	)

	// Initial cache repository
	var cache repository.CacheRepository = repository.NewCacheFilesystemRepository("product")

	// Initial product repository adapter
	var adapter repository.ProductRepository = repository.NewProductWithCacheRepositoryAdapter(
		product, cache,
	)

	return &ProductService{
		repository: adapter,
	}
}

func (service *ProductService) All() ([]*entity.Product, error) {
	return service.repository.All()
}
