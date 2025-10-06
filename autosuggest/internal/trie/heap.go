package trie

import (
	"container/heap"
	"sort"
)

type TopItem struct {
	PhraseId int64
	Score    float64
}

type TopHeap struct {
	items []TopItem
	limit int
	index map[int64]int
}

func NewTopHeap(limit int) *TopHeap {
	return &TopHeap{
		items: make([]TopItem, 0, limit),
		limit: limit,
		index: make(map[int64]int),
	}
}

func (h *TopHeap) Len() int {
	return len(h.items)
}

func (h *TopHeap) Less(i, j int) bool {
	return h.items[i].Score < h.items[j].Score
}

func (h *TopHeap) Swap(i, j int) {
	h.index[h.items[i].PhraseId], h.index[h.items[j].PhraseId] = j, i
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *TopHeap) Push(x interface{}) {
	it := x.(TopItem)
	h.index[it.PhraseId] = len(h.items)
	h.items = append(h.items, it)
}

func (h *TopHeap) Pop() interface{} {
	n := len(h.items)
	it := h.items[n-1]
	delete(h.index, it.PhraseId)
	h.items = h.items[:n-1]
	return it
}

func (h *TopHeap) Upsert(it TopItem) {
	if idx, ok := h.index[it.PhraseId]; ok {
		h.items[idx].Score = it.Score
		heap.Fix(h, idx)
		return
	}
	if h.Len() < h.limit {
		heap.Push(h, it)
		return
	}
	if h.items[0].Score >= it.Score {
		return
	}
	heap.Pop(h)
	heap.Push(h, it)
}

func (h *TopHeap) TopK(k int) []TopItem {
	n := h.Len()
	if k > n {
		k = n
	}
	cp := make([]TopItem, n)
	copy(cp, h.items)
	sort.Slice(cp,func(i,j int) bool {
		return cp[i].Score > cp[j].Score
	})
	return cp[:k]
}


