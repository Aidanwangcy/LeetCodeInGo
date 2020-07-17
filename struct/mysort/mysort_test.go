package mysort_test

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/Aidanwangcy/leetcode/struct/mysort"
)

func Testquicksort1(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	mysort.QuickSort(values, 0, len(values)-1)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 || values[4] != 5 {
		t.Error("quicksort() failed, Got ", values, " Expect 1 2 3 4 5")
	}
}

func Testquicksort2(t *testing.T) {
	values := []int{5, 5, 3, 2, 1}
	mysort.QuickSort(values, 0, len(values)-1)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 5 || values[4] != 5 {
		t.Error("quicksort() failed, Got ", values, " Expect 1 2 3 5 5")
	}
}

func Testquicksort3(t *testing.T) {
	values := []int{5}
	mysort.QuickSort(values, 0, len(values)-1)
	if values[0] != 5 {
		t.Error("quicksort() failed, Got ", values, " Expect 5")
	}
}

func Testquciksort_go1(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	done := make(chan struct{})
	mysort.QuickSortInGo(values, 0, len(values)-1, done, 5)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 || values[4] != 5 {
		t.Error("quicksort() failed, Got ", values, " Expect 1 2 3 4 5")
	}
}

func Testquciksort_go2(t *testing.T) {
	values := []int{5, 5, 3, 2, 1}
	done := make(chan struct{})
	mysort.QuickSortInGo(values, 0, len(values)-1, done, 5)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 5 || values[4] != 5 {
		t.Error("quicksort() failed, Got ", values, " Expect 1 2 3 5 5")
	}
}
func Testquciksort_go3(t *testing.T) {
	values := []int{5}
	done := make(chan struct{})
	mysort.QuickSortInGo(values, 0, len(values)-1, done, 5)
	if values[0] != 5 {
		t.Error("quicksort() failed, Got ", values, " Expect 5")
	}
}

func TimeTest() {
	rand.Seed(time.Now().UnixNano())
	testData1, testData2, testData3 := make([]int, 0, 100000000), make([]int, 0, 100000000), make([]int, 0, 100000000)
	times := 100000000
	for i := 0; i < times; i++ {
		val := rand.Intn(20000000)
		testData1 = append(testData1, val)
		testData2 = append(testData2, val)
		testData3 = append(testData3, val)
	}

	start := time.Now()
	mysort.QuickSort(testData1, 0, len(testData1)-1)
	fmt.Println("singal goroutine:", time.Now().Sub(start))
	if sort.IntsAreSorted(testData1) {
		fmt.Println("wrong quick_sort implementation")
	}

	done := make(chan struct{})
	start = time.Now()
	go mysort.QuickSortInGo(testData2, 0, len(testData2)-1, done, 5)
	<-done
	fmt.Println("multiple goroutine: ", time.Now().Sub(start))
	if !sort.IntsAreSorted(testData2) {
		fmt.Println("wrong quickSort_go implementation")
	}
	start = time.Now()
	sort.Ints(testData3)
	fmt.Println("std lib: ", time.Now().Sub(start))
	if !sort.IntsAreSorted(testData3) {
		fmt.Println("wrong std lib implementation")
	}
}
