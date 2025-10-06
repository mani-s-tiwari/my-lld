package app

import (
	"autosuggest/internal/phrase"
	"autosuggest/internal/trie"
	"time"
)

type Service struct {
	Store *phrase.Store
	Trie  *trie.Trie
	WFreq float64
	WRec  float64
	Tau   time.Duration
}

func NewService() *Service {
	store := phrase.NewStore()
	tr := trie.NewTrie(50, store)
	return &Service{
		Store: store,
		Trie:  tr,
		WFreq: 1.0,
		WRec:  1.0,
		Tau:   24 * time.Hour,
	}
}

func (s *Service) AddOrUpdate(text string, freq int64) {
	p := s.Store.Upsert(text, freq, time.Now().UnixMilli(), func(p *phrase.Phrase) {
		p.Score = phrase.Score(p.Freq, p.LastSeen, s.WFreq, s.WRec, s.Tau)
	})
	s.Trie.InsertOrUpdate(p)
}

func (s *Service) Suggest(prefix string, k int) []*phrase.Phrase {
	return s.Trie.Suggest(prefix, k)
}
