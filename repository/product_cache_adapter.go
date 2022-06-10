package repository

import (
	"encoding/json"

	"github.com/zipzoft/interview-k-donn/entity"
)

func NewProductWithCacheRepositoryAdapter(productRepository ProductRepository, cacheRepository CacheRepository) *ProductCacheRepositoryAdapter {
	return &ProductCacheRepositoryAdapter{
		product: productRepository,
		cache:   cacheRepository,
	}
}

var _ ProductRepository = (*ProductCacheRepositoryAdapter)(nil)

type ProductCacheRepositoryAdapter struct {
	cache   CacheRepository
	product ProductRepository
}

// All implements ProductRepository
func (repo *ProductCacheRepositoryAdapter) All() ([]*entity.Product, error) {
	// Get from cache
	jsonContent, err := repo.cache.Get("products")

	products := make([]*entity.Product, 0)

	// If found cache
	// Parse json string to products
	if err == nil {
		err = json.Unmarshal([]byte(jsonContent), &products)

		if err != nil {
			return nil, err
		}

		return products, nil
	}

	// If not found cache
	// Get from product repository
	products, err = repo.product.All()

	if err != nil {
		return nil, err
	}

	// Convert products to json string
	contents, err := json.Marshal(products)
	if err != nil {
		return nil, err
	}

	// Save to cache
	err = repo.cache.Set("products", string(contents))

	if err != nil {
		return nil, err
	}

	return products, nil
}
