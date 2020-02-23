package rectangle

import "testing"

var rectExamples = []Rectangle {
	{1, 1, 5, 6},
	{5, 2, 7, 3},
	{20, 23, 25, 28},
}

func TestRectangleMemoryStorage_GetFirst(t *testing.T) {
	st := NewStorage()
	for _, r := range rectExamples {
		st.Save(r)
	}
	if rm, isNull := st.GetFirst(); isNull {
		t.Fatalf("expected: %#v recieved null", rectExamples[0])
	} else {
		if !rm.Rect.Equal(rectExamples[0]) {
			t.Fatalf("expected: %#v received: %#v", rectExamples[0], rm.Rect)
		}
	}
}

func BenchmarkMemoryStorage_Save(b *testing.B) {
	st := NewStorage()
	for i := 0; i < b.N; i++ {
		st.Save(rectExamples[2])
	}
}