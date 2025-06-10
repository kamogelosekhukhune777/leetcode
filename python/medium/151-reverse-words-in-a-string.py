'''
Given an input string s, reverse the order of the words.

A word is defined as a sequence of non-space characters. The words 
in s will be separated by at least one space.

Return a string of the words in reverse order concatenated by a single 
space.

Note that s may contain leading or trailing spaces or multiple spaces between 
two words. The returned string should only have a single space separating the 
words. Do not include any extra spaces.

Example 1:
    Input: s = "the sky is blue"
    Output: "blue is sky the"

Example 2:
    Input: s = "  hello world  "
    Output: "world hello"
    Explanation: Your reversed string should not contain leading or 
        trailing spaces.

Example 3:
    Input: s = "a good   example"
    Output: "example good a"
    Explanation: You need to reduce multiple spaces between two words to a 
        single space in the reversed string.
 

Constraints:
    1 <= s.length <= 10^4
    s contains English letters (upper-case and lower-case), digits, and spaces ' '.
    There is at least one word in s.
 

Follow-up: If the string data type is mutable in your language, can you solve it in-place with O(1) extra space?
'''

class Solution():
    def reverseWords(self, s):
        result = []
        i = len(s) - 1

        while i >= 0:
            # Skip trailing spaces
            while i >= 0 and s[i] == ' ':
                i -= 1

            if i < 0:
                break

            # Find the end of the word
            end = i

            # Move left to the beginning of the word
            while i >= 0 and s[i] != ' ':
                i -= 1

            start = i + 1
            result.append(s[start:end+1])

        return ' '.join(result)
