package hackerrank

// problem https://www.hackerrank.com/challenges/challenging-palindromes/problem?isFullScreen=true
func Palindrome(a string) bool {
	reverse := []rune(a)
	for i, j := 0, len(reverse)-1; i < j; i, j = i+1, j-1 {
		reverse[i], reverse[j] = reverse[j], reverse[i]
	}
	if a == string(reverse) {
		return true
	}
	return false

}
