func isPalindrome(x int) bool {
	isPalindrome := reverseInt(x) == x
	return isPalindrome
}

func reverseInt(n int) int {
	newInt := 0
	for n > 0 {
		remainder := n % 10
		newInt *= 10
		newInt += remainder
		n /= 10
	}
	return newInt
}