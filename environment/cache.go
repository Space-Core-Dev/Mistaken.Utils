package env

import (
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

func NewExperationCache(expiration, cleanupInterval time.Duration, values map[string]string) *cache.Cache {
	return cache.New(expiration, cleanupInterval)
}

func NewPermanentCache(values map[string]string) *sync.Map {
	return &sync.Map{}
}
