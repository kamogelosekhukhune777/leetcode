"""
1167. Minimum Cost to Connect Sticks
You have some sticks with positive integer lengths.

You can connect any two sticks of lengths X and Y into one stick by paying a cost of X + Y.  
You perform this action until there is one stick remaining.

Return the minimum cost of connecting all the given sticks into one stick in this way.


Example 1:
    Input: sticks = [2,4,3]
    Output: 14

Example 2:
    Input: sticks = [1,8,3,5]
    Output: 30

Constraints:
    1 <= sticks.length <= 10⁴
    1 <= sticks[i] <= 10⁴

"""

import heapq

class Solution:
    def connectSticks(sticks):
        if len(sticks) <= 1:
            return 0
        
        heapq.heapify(sticks)
        total_cost = 0

        while len(sticks) > 1:
            first = heapq.heappop(sticks)
            second = heapq.heappop(sticks)
            cost = first + second
            total_cost += cost
            heapq.heappush(sticks, cost)

        return total_cost

print("Example 1: ")
sticks = [2,4,3]
print(Solution.connectSticks(sticks))

print("Example 1: ")
sticks = [1,8,3,5]
print(Solution.connectSticks(sticks))
