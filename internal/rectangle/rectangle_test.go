package rectangle

import (
	"testing"
)

var hasOverlapTests = []struct {
	name           string
	rect1          Rectangle
	rect2          Rectangle
	expectedResult bool
}{
	{"rect1_rect2_equal", Rectangle{1, 1, 5, 6}, Rectangle{1, 1, 5, 6}, true},
	{"rect2_right_to_rect1_has_overlap", Rectangle{1, 1, 5, 6}, Rectangle{5, 2, 7, 3}, true},
	{"rect2_right_rect1_no_overlap", Rectangle{12, 5, 13, 7}, Rectangle{14, 5, 17, 9}, false},
	{"rect2_top_rect1_no_overlap", Rectangle{20, 15, 25, 22}, Rectangle{20, 23, 25, 28}, false},
}

func TestRectangle_HasOverlap(t *testing.T) {
	for _, tt := range hasOverlapTests {
		res := tt.rect1.HasOverlap(tt.rect2)
		if res != tt.expectedResult {
			t.Fatalf("%s result: %v, expected: %v", tt.name, res, tt.expectedResult)
		}
	}
}

func BenchmarkRectangle_HasOverlap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hasOverlapTests[0].rect1.HasOverlap(hasOverlapTests[0].rect2)
	}
}