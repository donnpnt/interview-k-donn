package service

import (
	"github.com/zipzoft/interview-k-donn/config"
	"github.com/zipzoft/interview-k-donn/entity"
	"github.com/zipzoft/interview-k-donn/repository"
)

type ProductService struct {
	repository repository.ProductRepository
}

//var ctx = context.Background()

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

	//clientRedis := redis.NewClient(&redis.Options{
	//	Addr:     "127.0.0.1:6379", // host:port of the redis server
	//	Password: "",               // no password set
	//	DB:       0,                // use default DB
	//})
	//
	//fmt.Println(repository.GetRedis(clientRedis))

	repository.SetAdapter(adapter)

	var result repository.ProductRepository = repository.NewProductWithRedisRepositoryAdapter(
		product, cache,
	)

	result.All()

	return &ProductService{
		repository: result,
	}
}

func DBProductService(adapter repository.ProductRepository) {
	//clientRedis := redis.NewClient(&redis.Options{
	//	Addr:     "127.0.0.1:6379", // host:port of the redis server
	//	Password: "",               // no password set
	//	DB:       0,                // use default DB
	//})
	//
	//clientMongo := config.ConnectMongo()
	//dbClient := clientMongo.Database("test").Collection("product")
	//var products []*entity.Product
	//
	//find, err := dbClient.Find(ctx, bson.M{})
	//
	//if err = find.All(ctx, &products); err != nil {
	//	panic(err)
	//}
	//
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//if len(products) == 0 {
	//	all, err := adapter.All()
	//
	//	if err != nil {
	//		fmt.Println(err.Error())
	//		return
	//	}
	//	for _, value := range all {
	//		_, err := clientMongo.Database("test").Collection("product").InsertOne(ctx, value)
	//		if err != nil {
	//			fmt.Println(err.Error())
	//			return
	//		}
	//	}
	//	SetRedis(clientRedis, all)
	//	fmt.Println(GetRedis(clientRedis))
	//
	//}
}

//func GetRedis(clientRedis *redis.Client) string {
//
//	val, err := clientRedis.Get("product").Result()
//
//	if err != nil {
//		return err.Error()
//	}
//	return val
//}
//
//func SetRedis(clientRedis *redis.Client, value []*entity.Product) {
//	data, _ := json.Marshal(value)
//
//	err := clientRedis.Set("product", data, 0).Err()
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//}

func (service *ProductService) All() ([]*entity.Product, error) {
	return service.repository.All()
}
