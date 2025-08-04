package service

import (
	"math/rand"
	"sync"
	"time"
)

type Url struct {
	Original string
	ShortUrl string
}

type UrlService struct {
	mu     sync.RWMutex
	urls   map[string]*Url   // original URL -> Url
	urlmap map[string]string // short URL -> original URL
}

func NewUrlService() *UrlService {
	return &UrlService{
		urls:   make(map[string]*Url),
		urlmap: make(map[string]string),
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (us *UrlService) randomUrlGenerator() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for {
		url := ""
		for i := 0; i < 6; i++ {
			url += string(charset[rand.Intn(len(charset))])
		}

		us.mu.RLock()
		_, exists := us.urlmap[url]
		us.mu.RUnlock()

		if !exists {
			return url
		}
	}
}

func (us *UrlService) MakeUrl(orgUrl string) Url {
	short := us.randomUrlGenerator()

	us.mu.Lock()
	defer us.mu.Unlock()

	url := Url{
		Original: orgUrl,
		ShortUrl: short,
	}
	us.urls[orgUrl] = &url
	us.urlmap[short] = orgUrl

	return url
}

func (us *UrlService) RedirectUrl(shortUrl string) string {
	us.mu.RLock()
	defer us.mu.RUnlock()

	return us.urlmap[shortUrl]
}

func (us *UrlService) EditUrl(shortUrl, newOriginal string) Url {
	us.mu.Lock()
	defer us.mu.Unlock()

	// Delete old mapping (if exists)
	if oldOriginal, exists := us.urlmap[shortUrl]; exists {
		delete(us.urls, oldOriginal)
		delete(us.urlmap, shortUrl)
	}

	url := Url{
		Original: newOriginal,
		ShortUrl: shortUrl,
	}
	us.urls[newOriginal] = &url
	us.urlmap[shortUrl] = newOriginal

	return url
}

func (us *UrlService) DeleteUrl(original string) bool {
	us.mu.Lock()
	defer us.mu.Unlock()

	urlEntry, exists := us.urls[original]
	if !exists {
		return false
	}

	delete(us.urlmap, urlEntry.ShortUrl)
	delete(us.urls, original)
	return true
}
