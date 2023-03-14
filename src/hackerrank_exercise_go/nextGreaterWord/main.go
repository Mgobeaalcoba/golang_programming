package main

import "fmt"

func nextGreaterWord(word string) string {
	// Convert the string to a slice of bytes for easier manipulation
	chars := []byte(word)

	// Find the largest index i such that chars[i] < chars[i+1]
	i := len(chars) - 2
	for ; i >= 0; i-- {
		if chars[i] < chars[i+1] {
			break
		}
	}

	// If no such index exists, the word is already the largest permutation
	if i == -1 {
		return "no answer"
	}

	// Find the largest index j greater than i such that chars[j] > chars[i]
	j := len(chars) - 1
	for ; j > i; j-- {
		if chars[j] > chars[i] {
			break
		}
	}

	// Swap the characters at indices i and j
	chars[i], chars[j] = chars[j], chars[i]

	// Reverse the substring after index i
	reverse(chars[i+1:])

	return string(chars)
}

func reverse(chars []byte) {
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
}

func main() {
	nextWord := nextGreaterWord("abcd")
	fmt.Println(nextWord) // Outputs "abdc"
}
