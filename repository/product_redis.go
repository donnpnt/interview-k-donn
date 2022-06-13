package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/zipzoft/interview-k-donn/config"
	"github.com/zipzoft/interview-k-donn/entity"
	"go.mongodb.org/mongo-driver/bson"
)

var products ProductRepository

var ctx = context.Background()

func NewProductWithRedisRepositoryAdapter(productRepository ProductRepository, cacheRepository CacheRepository) *ProductRedisRepositoryAdapter {
	return &ProductRedisRepositoryAdapter{
		product: productRepository,
		cache:   cacheRepository,
	}
}

var _ ProductRepository = (*ProductRedisRepositoryAdapter)(nil)

type ProductRedisRepositoryAdapter struct {
	cache   CacheRepository
	product ProductRepository
}

func (p ProductRedisRepositoryAdapter) All() ([]*entity.Product, error) {
	adapter := products

	clientRedis := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // host:port of the redis server
		Password: "",               // no password set
		DB:       0,                // use default DB
	})
	clientRedis.Del("product")
	result, _ := GetRedis(clientRedis)

	if result == nil {
		clientMongo := config.ConnectMongo()
		dbClient := clientMongo.Database("test").Collection("product")

		var products []*entity.Product

		find, err := dbClient.Find(ctx, bson.M{})

		if err = find.All(ctx, &products); err != nil {
			fmt.Println(err.Error() + "test")
		}

		if err != nil {
			fmt.Println(err.Error() + "test2")
			return nil, err
		}

		all, err := adapter.All()

		SetRedis(clientRedis, all)

		if len(products) == 0 {

			if err != nil {
				fmt.Println(err.Error() + "test3")
				return nil, err
			}
			for _, value := range all {
				_, err := clientMongo.Database("test").Collection("product").InsertOne(ctx, value)
				if err != nil {
					fmt.Println(err.Error() + "test4")
					return nil, err
				}
			}
		}
		return products, nil
	}

	return result, nil
}

func SetAdapter(adapter ProductRepository) {
	products = adapter
}

func GetRedis(clientRedis *redis.Client) ([]*entity.Product, error) {
	val, err := clientRedis.Get("product").Result()

	if err != nil {
		return nil, err
	}

	var products []*entity.Product

	err = json.Unmarshal([]byte(val), &products)

	if err != nil {
		return nil, err
	}

	return products, nil
}
func SetRedis(clientRedis *redis.Client, value []*entity.Product) {
	data, _ := json.Marshal(value)

	err := clientRedis.Set("product", data, 0).Err()
	if err != nil {
		fmt.Println(err.Error())
	}
}
