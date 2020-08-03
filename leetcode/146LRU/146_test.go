package lru

import "testing"

func Test(t *testing.T)  {
	cache := Constructor(2)  
	cache.Put(1, 1)
	cache.Put(2, 2);
	ans1 := cache.Get(1);       // 返回  1
	cache.Put(3, 3);    // 该操作会使得密钥 2 作废
	ans2 := cache.Get(2);       // 返回 -1 (未找到)
	cache.Put(4, 4);    // 该操作会使得密钥 1 作废
	ans3 := cache.Get(1);       // 返回 -1 (未找到)
	ans4 := cache.Get(3);       // 返回  3
	ans5 := cache.Get(4);       // 返回  4
	if ans1 != 1 {
		t.Errorf("Get fial, got %v, expect %v", ans1, 1)
	}
	if ans2 != -1 {
		t.Errorf("Get fial, got %v, expect %v", ans2, -1)
	}
	if ans3 != -1 {
		t.Errorf("Get fial, got %v, expect %v", ans3, -1)
	}
	if ans4 != 3 {
		t.Errorf("Get fial, got %v, expect %v", ans4, 3)
	}
	if ans5 != 4 {
		t.Errorf("Get fial, got %v, expect %v", ans5, 4)
	}
}