"""
You have n  tiles, where each tile has one letter tiles[i] 
printed on it.

Return the number of possible non-empty sequences of letters 
you can make using the letters printed on those tiles.

Example 1:
    Input: tiles = "AAB"
    Output: 8
    Explanation: The possible sequences are "A", "B", "AA", "AB", "BA", "AAB", "ABA", "BAA".

Example 2:
    Input: tiles = "AAABBC"
    Output: 188

Example 3:
    Input: tiles = "V"
    Output: 1
 

Constraints:
    1 <= tiles.length <= 7
    tiles consists of uppercase English letters.
"""

from collections import Counter

#Recursive Backtracking with Pruning
class Solution1(object):
    def numTilePossibilities(self, tiles):
        counter = Counter(tiles)

        def backtrack(counter):
            total = 0
            for ch in counter:
                if counter[ch] == 0:
                    continue
                # Choose this letter
                total += 1
                counter[ch] -= 1
                total += backtrack(counter)
                counter[ch] += 1  # Backtrack
            return total

        return backtrack(counter)

#Using a used set to avoid duplicates at same level
class Solution2(object):
    def numTilePossibilities(self,tiles):
        tiles = sorted(tiles)
        used = [False] * len(tiles)
        result = set()

        def backtrack(path):
            for i in range(len(tiles)):
                if used[i]:
                    continue
                if i > 0 and tiles[i] == tiles[i - 1] and not used[i - 1]:
                    continue
                used[i] = True
                result.add("".join(path + [tiles[i]]))
                backtrack(path + [tiles[i]])
                used[i] = False

        backtrack([])
        return len(result)
