package rectangle

import (
	"sync"
	"time"
)

var DB Storage

type Storage interface {
	GetFirst() (*Model, bool)
	GetAll() []*Model
	Save(Rectangle) *Model
}

type MemoryStorage struct {
	Rects map[RectID]*Model
	LastID RectID
	mu sync.Mutex
}

func NewStorage() Storage {
	return &MemoryStorage{Rects: make(map[RectID]*Model)}
}

func (rm *MemoryStorage) GetFirst() (rectModel *Model, isNull bool) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	rectModel, ok := rm.Rects[1]
	return rectModel, !ok
}

func (rm *MemoryStorage) GetAll() []*Model {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	rects := make([]*Model, 0, len(rm.Rects))
	for _, r := range rm.Rects {
		rects = append(rects, r)
	}
	return rects
}

func (rm *MemoryStorage) Save(r Rectangle) *Model {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	id := rm.LastID + 1
	rectModel := &Model{
		ID:        id,
		Rect:      r,
		CreatedAt: time.Now(),
	}
	rm.Rects[id] = rectModel
	rm.LastID++
	return rectModel
}

