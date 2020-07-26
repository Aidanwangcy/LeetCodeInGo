package myallsort

//快速排序
//partition 分区操作， pivot为基准
func partition(nums []int, lo, hi int) int {
	i, j := lo, hi
	//基准为nums[lo]
	p := lo
	//基准放入temp
	temp := nums[p]

	//左右同时开始循环
	for i < j {
		//当j的值大于基准，且j大于p时，不用动，j-1
		for j >= p && nums[j] >= temp {
			j--
		}
		//j停在值小于基准的位置，这时候j如果仍然>=p那么交换位置
		if j >= p {
			nums[p] = nums[j]
			p = j
		}
		//如果i的值小于基准，且i小于p不用动，i+1
		for i <= p && nums[i] <= temp {
			i++
		}
		//i停在值大于基准的位置，这是胡如果仍然<=p那么交换位置
		if i <= p {
			nums[p] = nums[i]
			p = i
		}
	}
	//最后把基准值放到相应位置
	nums[p] = temp
	//返回基准位置
	return p
}

//QuickSort 单线程快速排序
func QuickSort(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	//先进行一次划分
	p := partition(a, lo, hi)
	QuickSort(a, lo, p-1)
	QuickSort(a, p+1, hi)
}

//QuickSortInGo go多线程快速排序
func QuickSortInGo(a []int, lo, hi int, done chan struct{}, depth int) {
	if lo >= hi {
		done <- struct{}{}
		return
	}
	depth--
	p := partition(a, lo, hi)
	if depth > 0 {
		childDone := make(chan struct{}, 2)
		go QuickSortInGo(a, lo, p-1, childDone, depth)
		go QuickSortInGo(a, p+1, hi, childDone, depth)
		<-childDone
		<-childDone
	} else {
		QuickSort(a, lo, p-1)
		QuickSort(a, p+1, hi)
	}
	done <- struct{}{}
}

//SimpleSelectSort 简单选择排序
//从未排序的数组序列中，选择最大或者最小元素添加入已排序数组 最开始已排序数组为空
func SimpleSelectSort(nums []int)  {
	for i := 0; i < len(nums); i++ {
		min := i
		for j := i + 1; j < len(nums); j++ {
			if nums[min] > nums[j] {
				min = j
			}
		}
		if min != i {
			tmp := nums[min]
            nums[min] = nums[i]
            nums[i] = tmp
		}
	}
}

//InsertSort 直接插入排序
//从未排序的数组元素中，取数据填写到以排序的数组中，并且形成已排序数组，数组长度为1的数组为以排序数组
func InsertSort(nums []int) {
    for i := 1; i < len(nums); i++ {
        for j := 0; j < i; j++ {
            if nums[i] < nums[j] {
                tmp := nums[i]
                nums[i] = nums[j]
                nums[j] = tmp
            }
        }
    }
}

//BubbleSort 冒泡排序
//相邻两个元素进行对比，顺序相反则互换。经过一轮的对比后，最大或者最小的元素“浮”到顶端，经过len(nums)-1轮后，nums变为有序数组
func bubbleSort(nums []int) {
    for i := 1; i < len(nums); i++ { //经过了len(nums)-1轮冒泡
        for j := 0; j < len(nums)-i; j++ {
            if nums[j] > nums[j+1] {
                tmp := nums[j]
                nums[j] = nums[j+1]
                nums[j+1] = tmp
            }
        }
    }
}

//MergeSort 归并排序
//将数组对分至若干有序数组（数组元素为一个的数组为有序数组)，将有序数组逐次两两合并为有序数组，直至合并为一个数组
func MergeSort(nums []int, start int, boundce int, end int) {
    if start >= end { //结束边界
        return
    }
    MergeSort(nums, start, start+(boundce-start)/2, boundce) //数组拆分前半部分
    MergeSort(nums, boundce+1, boundce+(end-boundce)/2, end) //数组拆分后半部分
    var tmp = make([]int, end-start+1, end-start+1)
    i := 0
    //拷贝数组
    for i <= end-start {
        tmp[i] = nums[i+start]
        i = i + 1
    }
    //数组合并过程
    i = 0                                      //左半部分数组的起始地址
    j := boundce - start + 1                   //右半部分的起始地址
    ii := start                                //终选数组的起始地址
    for i <= boundce-start && j <= end-start { //结束条件
        /*在左半部分选取较小元素填充目标数组*/
        for i <= boundce-start && tmp[i] <= tmp[j] {
            nums[ii] = tmp[i]
            ii = ii + 1
            i = i + 1
        }
        //here --> tmp[i] > tmp[j]
        /*在右半部分选取较小数据填充目标数组*/
        for j <= end-start && tmp[j] <= tmp[i] {
            nums[ii] = tmp[j]
            ii = ii + 1
            j = j + 1
        }
    }
    //结束判断
    //左数组先结束
    if i == boundce-start+1 {
        for j <= end-start {
            nums[ii] = tmp[j]
            ii = ii + 1
            j = j + 1
        }
    }
    //右数组先结束
    if j == end-start+1 {
        for i <= boundce-start {
            nums[ii] = tmp[i]
            ii = ii + 1
            i = i + 1
        }
    }
}

//HeapSort 堆排序
//1. 建堆 2. 堆调整 3. 堆排序
func HeapSort(nums []int) {
    tmp := len(nums)
    if tmp == 0 {
        return
    }
    buildHeap(nums)
    for i := 0; i < len(nums); i++ {
        nums[tmp-1], nums[0] = nums[0], nums[tmp-1] //提取堆顶
        tmp--
        adjustHeap(nums, 0, tmp) //数组尾部是已经排序部分，不再做调整
    }
}

//buildHeap 建堆
func buildHeap(nums []int) {
    len := len(nums)
    for i := len/2 - 1; i >= 0; i-- { //可以理解为先为子树构建“堆”，再在子树“堆”上构建父堆， len/2 为第一个非叶子节点，叶子节点肯定是“堆”
        adjustHeap(nums, i, len) //每次都为全局调整
    }
}

//adjustHeap 调整方式：当父节点的数据变化时，子节点的数据必须做相应调整
//2*n + 1 左子节点， 2*n + 2 右子节点
func adjustHeap(nums []int, pos int, len int) {
    tmp := pos
    j := 2*pos + 1
    for j < len {
        //寻找较大子节点
        if j+1 < len {
            if nums[j+1] > nums[j] {
                j++
            }
            if nums[tmp] < nums[j] {
                nums[tmp], nums[j] = nums[j], nums[tmp]
            }
        } else {
            if nums[tmp] < nums[j] {
                nums[tmp], nums[j] = nums[j], nums[tmp]
            }
        }
        //遍历此节点的子节点，调整
        tmp = j
        j = j*2 + 1
    }
}