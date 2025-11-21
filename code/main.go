package main

import "fmt"

/*
1.	Longest Palindromic Substring
Given a string s, the goal is to find the longest substring that is a palindrome.
If there are multiple valid answers, return the first appearing substring.
Examples:
✅ Input: s = "racecarapple"
left = 0
right = 1

var largest string

left == right
0 , 6 => racecar
'r' == 'r'

O(2 ^ n)

🔹 Output: "racecar"


✅ Input: s = "banana"
🔹 Output: "anana"
✅ Input: s = "abc"
🔹 Output: "a" (Since no longer palindromes exist, return the first character)
✅ Input: s = ""
🔹 Output: "" (An empty string should return an empty result)

*/

func isPalindromic(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		if s[left] == s[right] {
			left++
			right--
		}
	}
	return true
}

func LongestPalindromicSubstring(s string) string {
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return s
	}
	var largest string
	for left := 0; left < len(s); left++ {
		for right := left + 1; right < len(s); right++ {
			if s[left] == s[right] {
				sub := s[left:right] + string(s[right])
				fmt.Println(sub, " ", isPalindromic(sub))
				if isPalindromic(sub) {
					if len(sub) > len(largest) {
						largest = sub
					}
				}
			}
		}
	}

	if len(largest) == 0 {
		return string(s[0])
	}
	return largest
}

func main() {
	str := "" //"abc" //"banana" //"racecarapple"
	result := LongestPalindromicSubstring(str)
	fmt.Println(result)

}
