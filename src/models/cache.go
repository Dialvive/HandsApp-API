package models

import (
	"time"
)

// Cache stores an object of any type, recording it's stored timestamp, and assigns a time to live value
type Cache struct {
	initTime time.Time
	ttl      time.Duration
	renewed  uint
	storage  interface{}
}

// NewCache is a constructor for a Cache of any type
func NewCache(s interface{}, timeToLive string) (Cache, error) {
	if ttlaux, err := time.ParseDuration(timeToLive); err != nil {
		return Cache{}, err
	} else {
		cache := Cache{
			initTime: time.Now(),
			ttl:      ttlaux,
			renewed:  0,
			storage:  s,
		}
		return cache, nil
	}
}

// InitTime returns the time the Cache was created
func (c Cache) InitTime() time.Time {
	return c.initTime
}

// TTL returns the time to live of the cache
func (c Cache) TTL() time.Duration {
	return c.ttl
}

// Retrieve returns the stored atribute of the cache
func (c Cache) Retrieve() interface{} {
	return c.storage
}

// Expired returns true if Now is greater than initTime plus the time to live
func (c Cache) Expired() bool {
	return c.InitTime().Add(c.ttl).After(time.Now())
}

// Update sets the storage of the cache to a new one, and resets the init value to Now
func (c Cache) Update(s interface{}) {
	c.storage = s
	c.initTime = time.Now()
	c.renewed = c.renewed + 1
}

// Renew renews the Cache, without changing the stored object
func (c Cache) Renew() {
	c.initTime = time.Now()
	c.renewed = c.renewed + 1
}
