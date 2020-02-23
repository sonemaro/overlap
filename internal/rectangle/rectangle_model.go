package rectangle

import "time"

type RectID uint8

type Model struct {
	ID RectID `json:"id"`
	Rect Rectangle `json:"rect"`
	CreatedAt time.Time `json:"created_at"`
}
