package repository

import (
	"encoding/base64"
	"io/ioutil"
	"os"

	"github.com/zipzoft/interview-k-donn/pkg/helpers"
)

func NewCacheFilesystemRepository(prefix string) *CacheFilesystemRepository {
	return &CacheFilesystemRepository{
		path: helpers.CachePath(prefix),
	}
}

var _ CacheRepository = (*CacheFilesystemRepository)(nil)

type CacheFilesystemRepository struct {
	path string
}

// Clear implements CacheRepository
func (cache *CacheFilesystemRepository) Clear() error {
	return os.RemoveAll(cache.path)
}

// Delete implements CacheRepository
func (cache *CacheFilesystemRepository) Delete(key string) error {
	return os.Remove(cache.filePath(cache.keyName(key)))
}

// Get implements CacheRepository
func (cache *CacheFilesystemRepository) Get(key string) (string, error) {
	// Open file
	file, err := os.Open(cache.filePath(cache.keyName(key)))

	if err != nil {
		return "", err
	}

	defer file.Close()

	value, err := ioutil.ReadAll(file)

	return string(value), err
}

// Set implements CacheRepository
func (cache *CacheFilesystemRepository) Set(key string, value string) error {

	// Create directory if not exists
	if err := os.MkdirAll(cache.path, os.ModePerm); err != nil {
		return err
	}

	// Save to file
	file, err := os.OpenFile(cache.filePath(cache.keyName(key)), os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		return err
	}

	defer file.Close()

	// Write value to file
	_, err = file.WriteString(value)

	return err
}

func (cache *CacheFilesystemRepository) keyName(key string) string {
	// encrypt key to base64
	key = base64.StdEncoding.EncodeToString([]byte(key))

	return key
}

func (cache *CacheFilesystemRepository) filePath(filename string) string {
	return cache.path + "/" + filename
}
