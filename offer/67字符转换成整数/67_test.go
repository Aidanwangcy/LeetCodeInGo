package strtoint

import "testing"

func TestStrTInt(t *testing.T)  {
	str1 := "42"
	str2 := "   -42"
	str3 := "4193 with words"
	str4 := "words and 987"
	str5 := "-91283472332"
	ans1 := StrToInt(str1)
	ans2 := StrToInt(str2)
	ans3 := StrToInt(str3)
	ans4 := StrToInt(str4)
	ans5 := StrToInt(str5)
	if ans1 != 42 {
		t.Errorf("StrToInt() fail, got %v, expect %v.", ans1, 42)
	}
	if ans2 != -42 {
		t.Errorf("StrToInt() fail, got %v, expect %v.", ans2, -42)
	}
	if ans3 != 4193 {
		t.Errorf("StrToInt() fail, got %v, expect %v.", ans3, 4193)
	}
	if ans4 != 0 {
		t.Errorf("StrToInt() fail, got %v, expect %v.", ans4, 0)
	}
	if ans5 != -2147483648 {
		t.Errorf("StrToInt() fail, got %v, expect %v.", ans5, -2147483648)
	}
}