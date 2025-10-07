package main

import (
	"fmt"
	"sync"
	"time"
)

type Bucket struct {
	token      float64
	lastRefill time.Time
}

type RateLimiter struct {
	maxTokens  float64
	refillRate float64
	buckets    map[string]*Bucket
	mu         sync.Mutex
}

func NewRateLimiter(maxTokens, refillRate float64) *RateLimiter {
	return &RateLimiter{
		maxTokens:  maxTokens,
		refillRate: refillRate,
		buckets:    make(map[string]*Bucket),
	}
}

func (rl *RateLimiter) Allow(clientId string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	bucket, exists := rl.buckets[clientId]
	if !exists {
		bucket = &Bucket{
			token:      rl.maxTokens,
			lastRefill: time.Now(),
		}
		rl.buckets[clientId] = bucket
	}

	now := time.Now()
	elapsed := now.Sub(bucket.lastRefill).Seconds()
	refilled := elapsed * rl.refillRate
	if refilled > 0 {
		bucket.token = min(rl.maxTokens, bucket.token+refilled)
		bucket.lastRefill = now
	}

	if bucket.token >= 1 {
		bucket.token -= 1
		return true
	}

	return false
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func hit(i, j float64, k int, client string) {
	rl := NewRateLimiter(i, j)
	for i := 1; i <= k; i++ {
		if rl.Allow(client) {
			fmt.Println(i, "✅ Allow- ", client)
		} else {
			fmt.Println(i, "❌ Limit- ", client)
		}
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	go hit(10, 1, 100, "user1")
	// go hit(10,2,20,"user2")
	// go hit(10,1,10,"user3")

	time.Sleep(time.Hour)
}
