package trie

import (
	"autosuggest/internal/phrase"
	"strings"
	"sync"
)

type Node struct {
	Children map[rune]*Node
	Top      *TopHeap
	IsWord   bool
	PhraseId int64
}

func NewNode(limit int) *Node {
	return &Node{Children: make(map[rune]*Node), Top: NewTopHeap(limit)}
}

type Trie struct {
	root     *Node
	store    *phrase.Store
	topLimit int
	rw       sync.Mutex
}

func NewTrie(limit int, store *phrase.Store) *Trie {
	return &Trie{root: NewNode(limit), topLimit: limit, store: store}
}

func (t *Trie) InsertOrUpdate(p *phrase.Phrase) {
	runes := []rune(strings.ToLower(p.Text))
	t.rw.Lock()
	defer t.rw.Unlock()
	node := t.root
	node.Top.Upsert(TopItem{p.ID, p.Score})
	for _, r := range runes {
		if _, ok := node.Children[r]; !ok {
			node.Children[r] = NewNode(t.topLimit)
		}
		node = node.Children[r]
		node.Top.Upsert(TopItem{p.ID, p.Score})
	}
	node.IsWord = true
	node.PhraseId = p.ID
}

func (t *Trie) Suggest(prefix string, k int) []*phrase.Phrase {
	runes := []rune(strings.ToLower(prefix))
	t.rw.Lock()
	node := t.root
	for _, r := range runes {
		next, ok := node.Children[r]
		if !ok {
			t.rw.Unlock()
			return nil
		}
		node = next
	}
	items := node.Top.TopK(k)
	t.rw.Unlock()

	results := make([]*phrase.Phrase, 0, len(items))
	for _, it := range items {
		if ph, ok := t.store.GetById(it.PhraseId); ok {
			results = append(results, ph)
		}
	}
	return results
}
