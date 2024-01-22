// indexOf finds the index of the given element in the slice.
// If the element is not found, it returns -1.
func indexOf(element int, data []int) int {
	// Iterate over the elements in the slice with their index.
	for k, v := range data {
		// Check if the current element is equal to the target element.
		if element == v {
			// If a match is found, return the index of the element.
			return k
		}
	}
	// If the element is not found, return -1.
	return -1
}

// removeElement removes all occurrences of the specified value from the slice.
// It returns the new length of the slice after removal.
func removeElement(nums []int, val int) int {
	// Iterate over the elements in the slice.
	for i := 0; i < len(nums); i++ {
		// Check if the current element is equal to the target value.
		if nums[i] == val {
			// Find the index of the target value in the slice.
			index := indexOf(val, nums)
			// Check if the target value is found in the slice.
			if index != -1 {
				// Remove the element at the found index by slicing the slice.
				nums = append(nums[:index], nums[index+1:]...)
				// Decrement the loop counter to revisit the same index after removal.
				i--
			}
		}
	}

	// Return the new length of the slice after removal.
	return len(nums)
}
