"""
Given a string s, you can transform every letter individually 
to be lowercase or uppercase to create another string.

Return a list of all possible strings we could create. Return 
the output in any order.

Example 1:
    Input: s = "a1b2"
    Output: ["a1b2","a1B2","A1b2","A1B2"]

Example 2:
    Input: s = "3z4"
    Output: ["3z4","3Z4"]
 

Constraints:
    1 <= s.length <= 12
    s consists of lowercase English letters, uppercase English letters, and digits.

"""

#Iterative (Subset-like Expansion)
class Solution1(object):
    def letterCasePermutation(self,s):
        result = [""]

        for char in s:
            if char.isalpha():
                temp = []
                for item in result:
                    temp.append(item + char.lower())
                    temp.append(item + char.upper())
                result = temp
            else:
                result = [item + char for item in result]

        return result

#Recursive Backtracking (DFS)
class Solution2(object):
    def letterCasePermutation(self,s):
        result = []

        def backtrack(index, path):
            if index == len(s):
                result.append("".join(path))
                return

            if s[index].isalpha():
                # Lowercase
                path.append(s[index].lower())
                backtrack(index + 1, path)
                path.pop()

                # Uppercase
                path.append(s[index].upper())
                backtrack(index + 1, path)
                path.pop()
            else:
                path.append(s[index])
                backtrack(index + 1, path)
                path.pop()

        backtrack(0, [])
        return result
