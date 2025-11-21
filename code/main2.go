package main

import "fmt"

/*
2.	Check if a given string is a rotation of a palindrome
Given a string s, determine whether it is a rotation of a palindrome.
A string is considered a rotation of a palindrome if there exists some rotation of it that forms a valid palindrome.
Examples:
✅ Input: s = "aab"
🔹 Output: 1
💡 Explanation: "aab" is a rotation of "aba", which is a palindrome.

"aab"

solution 1:
count := map[rune] int

a, a, b

'a' = 2
'b' = 1

one of them is even

solution 2:
"aab"
1. only 2 letters

var letter1, letter2 rune
var count1, count2 int

2. one of the letter should be repeated in even number "aaab" 2

letter1 = 'a'
count1 = 1

letter1 = 'a'
count1 = 2

letter2 = 'b'
count2 = 1

- if we see third letter we will return false

count := max(count1, count2)

count%2 == 0{
return true


✅ Input: s = "aaaad"
🔹 Output: 1
💡 Explanation: "aaaad" is a rotation of "aadaa", which is a palindrome.

Input: "aaab"
output: ? 0

Input: "aaaabb"
output: ? 1

Input: "aaaabbb"
output: 1
💡 Explanation: "aabbbaa"

Input: "aaaabbbb"
output: 1
💡 Explanation: "aabbbaa"

Input: "aaaabc"
output: ? 0




✅ Input: s = "abcd"
🔹 Output: 0
💡 Explanation: "abcd" is not a rotation of any palindrome

*/

func isRotationPalindrome(s string) bool {
	var (
		letter1 rune
		count1  int
		letter2 rune
		count2  int
	)
	if len(s) == 0 {
		fmt.Println("0", count1, letter1, count2, letter2)

		return false
	}
	if len(s) == 1 {
		fmt.Println("1", count1, letter1, count2, letter2)
		return false
	}
	for _, r := range s {
		fmt.Println(string(r))
		if r == letter1 {
			count1++
		} else if r == letter2 {
			count2++
		} else if letter1 == 0 {
			letter1 = r
			count1 = 1
		} else if letter2 == 0 {
			letter2 = r
			count2 = 1
		} else {
			fmt.Println("else", count1, letter1, count2, letter2)
			return false
		}

	}
	fmt.Println("final", count1, letter1, count2, letter2)
	maxCount := max(count1, count2)
	if (maxCount % 2) == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(isRotationPalindrome("aaaabc")) // false
}
