package main

// https://leetcode.com/problems/longest-substring-without-repeating-characters/
// Time: O(n) | Space: O(1)
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	m := make(map[byte]bool, n)
	ans := 0

	for left, right := 0, 0; right < n; left++ {
		for right < n && !m[s[right]] {
			m[s[right]] = true
			right++
		}

		subLength := right - left
		if ans < subLength {
			ans = subLength
		}

		delete(m, s[left])
	}

	return ans
}
