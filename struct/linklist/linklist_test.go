package linklist

import (
	"testing"
)

func TestLen(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	if ans := l.Len(); ans != 3 {
		t.Errorf("len must be 3, but %d got", ans)
	}
}
