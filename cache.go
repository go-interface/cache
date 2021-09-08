package cache

/*Cacher defined a cache interface for common*/
type Cacher interface {
	// Get a value
	Get(key string) ([]byte, error)
	// Set a value to cache
	Set(key string, val []byte) error
	// SetWithTTL set a value with ttl
	SetWithTTL(key string, val []byte, ttl int64) error
	// Has check the value is exist
	Has(key string) (bool, error)
	// Delete delete a value
	Delete(key string) error
	// Clear all the values
	Clear() error
	// SetMultiple set multi values with [key]value
	SetMultiple(values map[string][]byte) error
	// DeleteMultiple delete multi values
	DeleteMultiple(keys ...string) error
}

type DefaultCacher interface {
	Cacher
	// GetD get a value from cache with default
	GetD(key string, v []byte) []byte
}
