package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/zipzoft/interview-k-donn/entity"
)

func NewProductJsonFileRepository(filePath string) *ProductJsonFileRepository {
	return &ProductJsonFileRepository{
		filePath: filePath,
	}
}

var _ ProductRepository = (*ProductJsonFileRepository)(nil)

type ProductJsonFileRepository struct {
	filePath string
}

// All implements ProductRepository
func (repo *ProductJsonFileRepository) All() ([]*entity.Product, error) {
	// Pause for 5 second for simulate loading data from database
	// Don't remove this line
	time.Sleep(5 * time.Second)

	// Open file
	file, err := os.Open(repo.filePath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	fmt.Println("Reading file:", repo.filePath)

	// Read value from file
	var products []*entity.Product

	// Map json from file to products
	err = json.NewDecoder(file).Decode(&products)

	if err != nil {
		return nil, err
	}

	return products, err
}
