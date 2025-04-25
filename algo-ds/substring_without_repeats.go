package main

import (
	"fmt"
)

// lengthOfLongestSubstring finds the length of the longest substring without repeating characters.
func lengthOfLongestSubstring(s string) int {
	// Create a map to store the last index of each character
	lastIndex := make(map[rune]int)
	maxLength := 0
	start := 0

	for i, char := range s {
		// If the character is already in the map and its last index is greater than or equal to the start index
		if lastIdx, found := lastIndex[char]; found && lastIdx >= start {
			// Move the start index to the right of the last occurrence of the character
			start = lastIdx + 1
		}
		// Update the last index of the character
		lastIndex[char] = i
		// Calculate the length of the current substring and update maxLength if it's greater
		if length := i - start + 1; length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}

func main() {
	// Test cases
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // Output: 3
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // Output: 1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // Output: 3
	fmt.Println(lengthOfLongestSubstring(""))         // Output: 0
	fmt.Println(lengthOfLongestSubstring(" "))        // Output: 1
}