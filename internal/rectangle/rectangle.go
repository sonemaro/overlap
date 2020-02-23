package rectangle

import "github.com/golang/geo/r2"

// Rectangle consists of 4 x and y for 2 points
type Rectangle struct {
	// x and y of 2 points of a rectangle
	X1 float64 `json:"x1"`
	Y1 float64 `json:"y1"`
	X2 float64 `json:"x2"`
	Y2 float64 `json:"y2"`
}

// Equal returns true if r == rect2
func (r *Rectangle) Equal(rect2 Rectangle) bool {
	return r.X1 == rect2.X1 && r.X2 == rect2.X2 && r.Y1 == rect2.Y1 && r.Y2 == rect2.Y2
}

// HasOverlap is a wrapper on native golang rect library
// It's easier to use and make this project more readable
func (rect1 *Rectangle) HasOverlap(other Rectangle) bool {
	// for sake of not reinventing the wheel, I used official golang rectangle library
	// for more info check their great test file: https://github.com/golang/geo/blob/master/r2/rect_test.go#L36
	rr := r2.RectFromPoints(r2.Point{rect1.X1, rect1.Y1}, r2.Point{rect1.X2, rect1.Y2})
	rr2 := r2.RectFromPoints(r2.Point{other.X1, other.Y1}, r2.Point{other.X2, other.Y2})
	return rr.Intersects(rr2)
}

func GetOverlapInArray(main *Rectangle, others []Rectangle) []Rectangle {
	overlapRects := make([]Rectangle, 0, len(others))
	for _, rect := range others {
		if main.HasOverlap(rect) {
			overlapRects = append(overlapRects, rect)
		}
	}
	return overlapRects
}