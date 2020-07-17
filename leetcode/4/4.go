package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 4, 5, 6, 7}
	nums2 := []int{8, 9, 10, 11, 12, 13, 14, 15, 16}
	n := findMedianSortedArrays(nums1, nums2)
	fmt.Println(n)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		return findMedianSortedArrays(nums2, nums1)
	}
	iMin, iMax, laftLen := 0, m, (n+m+1)/2
	for iMin <= iMax {
		i := iMin + (iMax-iMin)/2
		j := laftLen - i
		if i > iMin && nums1[i-1] > nums2[j] {
			iMax = i - 1
		} else if i < iMax && nums2[j-1] > nums1[i] {
			iMin = i + 1
		} else {
			leftMax := 0
			if i == 0 {
				leftMax = nums2[j-1]
			} else if j == 0 {
				leftMax = nums1[i-1]
			} else {
				leftMax = max(nums1[i-1], nums2[j-1])
			}
			if (m+n)%2 != 0 {
				return float64(leftMax)
			}
			rightMin := 0
			if i == m {
				rightMin = nums2[j]
			} else if j == n {
				rightMin = nums1[i]
			} else {
				rightMin = min(nums1[i], nums2[j])
			}
			return float64(leftMax+rightMin) / 2.0
		}
	}
	return 0.0
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
