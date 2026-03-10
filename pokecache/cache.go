package pokecache

import (
	"time"
	"sync"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type caughtPokemon struct {
	val []byte
}

type Cache struct {
	Entry map[string]cacheEntry
	PokemonCaught map[string]caughtPokemon
	mu *sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		Entry: make(map[string]cacheEntry),
		PokemonCaught: make(map[string]caughtPokemon),
		mu: &sync.Mutex{},
	}
	go c.reapLoop(interval)
	return &c
}

func (c Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Entry[key] = cacheEntry{
		createdAt: time.Now(),
		val: value,
	}
	return
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, ok := c.Entry[key]
	if ok {
		return value.val, true
	} else {
		return nil, false
	}
}

func (c Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mu.Lock()
		for key, value := range c.Entry {
			if time.Since(value.createdAt) > interval {
				delete(c.Entry, key)
			}
		}
		c.mu.Unlock()
	}
}

func (c Cache) CatchPokemon(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.PokemonCaught[key] = caughtPokemon{
		val: value,
	}
	return
}


func (c Cache) GetPokemon(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, ok := c.PokemonCaught[key]
	if ok {
		return value.val, true
	} else {
		return nil, false
	}
}
