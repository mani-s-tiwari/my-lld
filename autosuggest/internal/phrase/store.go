package phrase

import (
	"math"
	"sync"
	"time"
)

type Phrase struct {
	ID       int64
	Text     string
	Freq     int64
	LastSeen int64
	Score    float64
}

type Store struct {
	mu     sync.RWMutex
	nextID int64
	byText map[string]*Phrase
	byId   map[int64]*Phrase
}

func NewStore() *Store {
	return &Store{
		nextID: 1,
		byText: make(map[string]*Phrase),
		byId:   make(map[int64]*Phrase),
	}
}

func (s *Store) GetByText(text string) (*Phrase, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	p, ok := s.byText[text]
	return p, ok
}

func (s *Store) GetById(Id int64) (*Phrase, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	p, ok := s.byId[Id]
	return p, ok
}

func (s *Store) Upsert(text string, freqDelta int64, lastSeen int64, computeScore func(*Phrase)) *Phrase {
	s.mu.Lock()
	defer s.mu.Unlock()
	p, ok := s.byText[text]
	if !ok {
		p = &Phrase{
			ID:       s.nextID,
			Text:     text,
			Freq:     0,
			LastSeen: time.Now().Unix(),
			Score:    0,
		}
		s.byText[text] = p
		s.byId[p.ID] = p
		s.nextID++
	}
	p.Freq += freqDelta
	if lastSeen != 0 {
		p.LastSeen = lastSeen
	}
	computeScore(p)
	return p
}

func (s *Store) Delete(text string) (*Phrase, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	p, ok := s.byText[text]
	if !ok {
		return nil, false
	}
	delete(s.byText, text)
	delete(s.byId, p.ID)

	return p, true
}

func Score(freq int64, lastSeen int64, wFreq, wRec float64, tau time.Duration) float64 {
	now := time.Now().UnixMilli()
	fScore := math.Log(1+float64(freq))
	recency := math.Exp(-float64(now-lastSeen)/float64(tau.Milliseconds()))

	return wFreq*fScore +wRec*recency
}
