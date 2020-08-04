package depth

import (
	"testing"

	"github.com/aidancy/LeetCodeInGo/struct/binarytree"
)

func TestMaxDepth(t *testing.T)  {
	a := binarytree.New(3)
	b := binarytree.New(9)
	c := binarytree.New(20)
	d := binarytree.New(15)
	e := binarytree.New(17)
	a.Lchild = b
	a.Rchild = c
	c.Lchild = d
	c.Rchild = e
	ans := MaxDepth(a)
	ans2 := MaxDepth2(a)
	ans3 := MaxDepth3(a)
	if ans != 3 {
		t.Errorf("MaxDepth() fail, got %v, expect %v", ans, 3)
	}
	if ans2 != 3 {
		t.Errorf("MaxDepth2() fail, got %v, expect %v", ans2, 3)
	}
	if ans3 != 3 {
		t.Errorf("MaxDepth3() fail, got %v, expect %v", ans3, 3)
	}
}