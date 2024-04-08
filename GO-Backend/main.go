package main

import (
	"fmt"
	"time"

	"github.com/go-chi/chi"
	"net/http"
)

type LRUCache struct {
	capacity int
	cache    map[string]CacheItem
}

type CacheItem struct {
	value      string
	expiration time.Time
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]CacheItem),
	}
}

func (lru *LRUCache) Set(key, value string, expiration time.Duration) {
	lru.cache[key] = CacheItem{
		value:      value,
		expiration: time.Now().Add(expiration),
	}
}

func (lru *LRUCache) Get(key string) (string, bool) {
	item, found := lru.cache[key]
	if !found || time.Now().After(item.expiration) {
		delete(lru.cache, key)
		return "", false
	}
	return item.value, true
}

func main() {
	lru := NewLRUCache(1024)

	r := chi.NewRouter()
	r.Post("/cache/set", func(w http.ResponseWriter, r *http.Request) {
		key := r.FormValue("key")
		value := r.FormValue("value")
		expiration, _ := time.ParseDuration(r.FormValue("expiration"))

		lru.Set(key, value, expiration)
		w.Write([]byte("Key/Value set in cache"))
	})

	r.Get("/cache/get/{key}", func(w http.ResponseWriter, r *http.Request) {
		key := chi.URLParam(r, "key")

		if val, found := lru.Get(key); found {
			w.Write([]byte(fmt.Sprintf("Value for key %s: %s", key, val)))
		} else {
			w.Write([]byte("Key not found in cache"))
		}
	})

	http.ListenAndServe(":5000", r)
}
