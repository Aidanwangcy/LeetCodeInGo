package linklist_test

import (
	"testing"

	"github.com/Aidanwangcy/LeetCodeInGo/struct/linklist"
)

func TestLen(t *testing.T) {
	l := linklist.New()
	l.Append(1)
	l.Append(2)
	if ans := l.Len(); ans != 3 {
		t.Errorf("len must be 3, but %d got", ans)
	}
}
