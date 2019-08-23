package chop

func Chop(target int, arr []int) (ret int) {
	ret = bisearch1(target, arr)
	return ret
}

// 非递归
func bisearch1(target int, arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	left, mid, right := 0, 0, len(arr)-1
	for {
		// 中位数计算，下位中位数
		mid = left + (right-left)/2
		if arr[mid] > target {
			right = mid - 1
		} else if arr[mid] < target {
			left = mid + 1
		} else if arr[mid] == target {
			// 找到结果终止
			return mid
		}
		// 没有查到 mid设置-1，终止
		if left > right {
			mid = -1
			break
		}
	}

	return mid
}

// TODO 递归
func bisearchRecursive() {}

func bsr() {}
