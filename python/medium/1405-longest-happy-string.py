"""
1405. Longest Happy String

A string s is called happy if it satisfies the following conditions:

    s only contains the letters 'a', 'b', and 'c'.
    s does not contain any of "aaa", "bbb", or "ccc" as a substring.
    s contains at most a occurrences of the letter 'a'.
    s contains at most b occurrences of the letter 'b'.
    s contains at most c occurrences of the letter 'c'.

Given three integers a, b, and c, return the longest possible happy string. If there 
are multiple longest happy strings, return any of them. If there is no such string, 
return the empty string "".

A substring is a contiguous sequence of characters within a string.

 

Example 1:
    Input: a = 1, b = 1, c = 7
    Output: "ccaccbcc"
    Explanation: "ccbccacc" would also be a correct answer.

Example 2:
    Input: a = 7, b = 1, c = 0
    Output: "aabaa"
    Explanation: It is the only correct answer in this case.
 

Constraints:
    0 <= a, b, c <= 100
    a + b + c > 0
"""

import heapq

class Solution():
    def longestDiverseString(a, b, c):
        max_heap = []
        for count, char in [(-a, 'a'), (-b, 'b'), (-c, 'c')]:
            if count != 0:
                heapq.heappush(max_heap, (count, char))

        result = []

        while max_heap:
            count1, char1 = heapq.heappop(max_heap)

            # Check if last two are same
            if len(result) >= 2 and result[-1] == result[-2] == char1:
                if not max_heap:
                    break  # No alternative, can't proceed
                count2, char2 = heapq.heappop(max_heap)
                result.append(char2)
                count2 += 1  # we used one (count is negative)
                if count2 != 0:
                    heapq.heappush(max_heap, (count2, char2))
                heapq.heappush(max_heap, (count1, char1))  # push back first
            else:
                result.append(char1)
                count1 += 1
                if count1 != 0:
                    heapq.heappush(max_heap, (count1, char1))

        return ''.join(result)


print("Example 1:")
a = 1 
b = 1
c = 7
print(Solution.longestDiverseString(a, b, c))#Output: "ccaccbcc"
   

print("Example 2:")
a = 7
b = 1
c = 0
print(Solution.longestDiverseString(a, b, c)) #Output: "aabaa"