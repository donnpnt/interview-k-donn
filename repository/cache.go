package repository

type CacheRepository interface {
	// Get by key
	Get(key string) (string, error)

	// Set by key
	Set(key string, value string) error

	// Delete removes the key from the cache
	Delete(key string) error

	// Clear clears the cache
	Clear() error
}
