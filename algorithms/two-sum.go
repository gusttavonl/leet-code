func twoSum(nums []int, target int) []int {
	// Create a map to store the numbers and their indices
	numToIndex := make(map[int]int)

	// Iterate through the nums array
	for currentIndex, num := range nums {
		// Calculate the complement needed to reach the target
		complement := target - num

		// Check if the complement is already in the map
		if prevIndex, exists := numToIndex[complement]; exists {
			// Return the indices of the two numbers that add up to the target
			return []int{prevIndex, currentIndex}
		}

		// Store the current number and its index in the map
		numToIndex[num] = currentIndex
	}

	// If no solution is found, return nil
	return nil
}
