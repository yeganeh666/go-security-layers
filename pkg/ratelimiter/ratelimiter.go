package ratelimiter

import (
	"sync"
)

const rateLimit = 5

var (
	mu            sync.Mutex
	requestCounts = make(map[string]int)
)

func Check(remoteAddr string) bool {
	mu.Lock()
	defer mu.Unlock()

	count, ok := requestCounts[remoteAddr]
	if !ok {
		requestCounts[remoteAddr] = 1
		return true
	}

	if count >= rateLimit {
		return false
	}

	requestCounts[remoteAddr]++
	return true
}
