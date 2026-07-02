package main

import "fmt"

func main() {
	testCases := []string{
		"helloWorld",
		"hd",
		"h",
		".,",
		"",
		"hh",
		"wasdfdsaw",
		"wasdfds aw",
		"wasdfds    aw",
		"wa sdfdsaw",
		"wa  sdfd     saw",
		"w as dfd s aw",
		"wasd fdsaw",
		"wasd f dsaw",
		"wasdf dsaw",
		"A man, a plan, a canal: Panama",
	}

	for _, payload := range testCases {
		fmt.Printf("\"%s\", result: %v\n", payload, isPalindrome(payload))
	}
}

func matchLetter(a, b byte) bool {
	// Convert to lowercase by setting the 5th bit (a | 32)
	// Works perfectly for both A-Z and a-z in ASCII
	return (a | 32) == (b | 32)
}

func isAcceptableCharacter(c byte) bool {
	// Returns true if the byte is within uppercase or lowercase ranges
	return ('A' <= c && c <= 'Z') || ('a' <= c && c <= 'z') || ('0' <= c && c <= '9')
}

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	if len(s) == 0 {
		return true
	}

	if len(s) == 1 {
		return true
	}

	for left < right {
		if !isAcceptableCharacter(s[left]) {
			left++
			continue
		}

		if !isAcceptableCharacter(s[right]) {
			right--
			continue
		}

		if !matchLetter(s[left], s[right]) {
			return false
		}

		left++
		right--
	}

	return true
}
