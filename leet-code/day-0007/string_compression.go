package leetcode

/*
Given an array of characters chars, compress it using the following algorithm:

Begin with an empty string s. For each group of consecutive repeating characters in chars:

If the group's length is 1, append the character to s.
Otherwise, append the character followed by the group's length.
The compressed string s should not be returned separately, but instead, be stored in the input character array chars. Note that group lengths that are 10 or longer will be split into multiple characters in chars.

After you are done modifying the input array, return the new length of the array.

You must write an algorithm that uses only constant extra space.

Example 1:

Input: chars = ["a","a","b","b","c","c","c"]
Output: Return 6, and the first 6 characters of the input array should be: ["a","2","b","2","c","3"]
Explanation: The groups are "aa", "bb", and "ccc". This compresses to "a2b2c3".
Example 2:

Input: chars = ["a"]
Output: Return 1, and the first character of the input array should be: ["a"]
Explanation: The only group is "a", which remains uncompressed since it's a single character.
Example 3:

Input: chars = ["a","b","b","b","b","b","b","b","b","b","b","b","b"]
Output: Return 4, and the first 4 characters of the input array should be: ["a","b","1","2"].
Explanation: The groups are "a" and "bbbbbbbbbbbb". This compresses to "ab12".
*/

import "strconv"

func compress(chars []byte) int {
	if len(chars) == 1 {
		return 1
	}
	writeIdx := 0
	count := 1
	lastChar := chars[0]
	for i := 1; i < len(chars); i++ {
		if chars[i] == lastChar {
			count++
		} else {
			// need to write out the char with the length
			chars[writeIdx] = lastChar
			writeIdx++
			if count > 1 {
				str := strconv.Itoa(count)
				for j := 0; j < len(str); j++ {
					chars[writeIdx] = str[j]
					writeIdx++
				}
			}
			// reset count
			lastChar = chars[i]
			count = 1
		}
	}
	chars[writeIdx] = lastChar
	writeIdx++
	if count > 1 {
		str := strconv.Itoa(count)
		for j := 0; j < len(str); j++ {
			chars[writeIdx] = str[j]
			writeIdx++
		}
	}
	return writeIdx
}

func leetcodeSolution(chars []byte) int {
	i := 0
	res := 0
	for i < len(chars) {
		groupLength := 1
		for i+groupLength < len(chars) && chars[i+groupLength] == chars[i] {
			groupLength++
		}
		chars[res] = chars[i]
		res++
		if groupLength > 1 {
			for _, c := range strconv.Itoa(groupLength) {
				chars[res] = byte(c)
				res++
			}
		}
		i += groupLength
	}
	return res
}
