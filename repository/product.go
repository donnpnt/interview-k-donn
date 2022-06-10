package repository

import "github.com/zipzoft/interview-k-donn/entity"

type ProductRepository interface {
	All() ([]*entity.Product, error)
}
