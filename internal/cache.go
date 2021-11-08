package internal

import (
	"errors"
	"github.com/ReneKroon/ttlcache/v2"
	"github.com/guythatdrinkscoffee/CirculationApp/models"
	"log"
	"time"
)

type TTLCache struct {
	Cache ttlcache.SimpleCache
}

func NewTLLCache() TTLCache {
	c := ttlcache.NewCache()
	return TTLCache{
		Cache: c,
	}
}

func (s TTLCache) Get(code string) (*models.APIResponse, error) {
	val, err := s.Cache.Get(code)

	if err != nil {
		return nil, ttlcache.ErrNotFound
	}

	log.Println("Response found in Cache.")

	return val.(*models.APIResponse), nil
}

func (s TTLCache) Set(code string, value interface{}) error {
	if len(code) == 0 {
		err := errors.New("Code must not be empty")
		return err
	}

	err := s.Cache.SetWithTTL(code, value, time.Hour)

	if err != nil {
		return err
	}

	log.Println("Inserted into Cache")

	return nil
}
