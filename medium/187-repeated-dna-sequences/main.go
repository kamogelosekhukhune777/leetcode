package main

import "fmt"

/*
The DNA sequence is composed of a series of nucleotides abbreviated as 'A', 'C', 'G', and 'T'.

For example, "ACGAATTCCG" is a DNA sequence.
When studying DNA, it is useful to identify repeated sequences within the DNA.

Given a string s that represents a DNA sequence, return all the 10-letter-long sequences (substrings)
that occur more than once in a DNA molecule. You may return the answer in any order.


Example 1:
	Input: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	Output: ["AAAAACCCCC","CCCCCAAAAA"]

Example 2:
	Input: s = "AAAAAAAAAAAAA"
	Output: ["AAAAAAAAAA"]


Constraints:
	1 <= s.length <= 10^5
	s[i] is either 'A', 'C', 'G', or 'T'.
*/

func findRepeatedDnaSequence(s string) []string {
	//length of the DNA sequence must be atleat 10 to have any 10-letter-long sequences
	if len(s) <= 10 {
		return []string{}
	}

	//to store the counts of each 10-letter-long sequence
	sequenceMap := make(map[string]int)

	//to store the results of the repeated sequences
	result := []string{}

	//iterate through the DNA sequence using a sliding window of size 10.
	for i := 0; i < len(s)-10; i++ {
		//extract the 10-letter-long sequence
		sequence := s[i : i+10]

		//increment the count of the sequence
		sequenceMap[sequence]++

		//if the sequence count becomes 2, it means it's repeated so add to the result
		if sequenceMap[sequence] == 2 {
			result = append(result, sequence)
		}
	}

	return result
}

func main() {
	s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	fmt.Println(findRepeatedDnaSequence(s)) //Output: ["AAAAACCCCC","CCCCCAAAAA"]

	//Example 2:
	s = "AAAAAAAAAAAAA"
	fmt.Println(findRepeatedDnaSequence(s)) //Output: ["AAAAAAAAAA"]
}
