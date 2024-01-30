func lengthOfLIS(nums []int) int {
	sub := []int{}

	for _, num := range nums {
		i := binarySearch(sub, num)

		if i == len(sub) {
			sub = append(sub, num)
		} else {
			sub[i] = num
		}
	}

	return len(sub)
}

func binarySearch(sub []int, target int) int {
	left, right := 0, len(sub)-1

	for left <= right {
		pivot := left + (right-left)/2

		if sub[pivot] == target {
			return pivot
		}

		if target < sub[pivot] {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}

	return left
}