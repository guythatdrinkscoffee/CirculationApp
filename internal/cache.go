package internal

import (
	"errors"
	"github.com/ReneKroon/ttlcache/v2"
	"github.com/guythatdrinkscoffee/CirculationApp/models"
	"time"
)

type TTLCache struct {
	cache ttlcache.SimpleCache
}

func NewStorage() TTLCache {
	c := ttlcache.NewCache()
	return TTLCache{
		cache: c,
	}
}

func (s TTLCache) Get(code string) (models.CurrencyRates, error) {
	val, err := s.cache.Get(code)

	if err != nil {
		return nil, err
	}

	return val.(models.CurrencyRates), nil
}

func (s TTLCache) Set(code string, value interface{}) error {
	if len(code) == 0 {
		err := errors.New("Code must not be empty")
		return err
	}

	err := s.cache.SetWithTTL(code, value, time.Hour)

	if err != nil {
		return err
	}

	return nil
}
