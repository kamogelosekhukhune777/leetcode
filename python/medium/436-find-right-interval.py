"""
436. Find Right Interval

You are given an array of intervals, where intervals[i] = [starti, endi] and each starti is unique.

The right interval for an interval i is an interval j such that startj >= endi and startj is minimized. Note 
that i may equal j.

Return an array of right interval indices for each interval i. If no right interval exists for interval i, then 
put -1 at index i.

Example 1:
    Input: intervals = [[1,2]]
    Output: [-1]
    Explanation: There is only one interval in the collection, so it outputs -1.

Example 2:
    Input: intervals = [[3,4],[2,3],[1,2]]
    Output: [-1,0,1]
    Explanation: There is no right interval for [3,4].
        The right interval for [2,3] is [3,4] since start0 = 3 is the smallest start that is >= end1 = 3.
        The right interval for [1,2] is [2,3] since start1 = 2 is the smallest start that is >= end2 = 2.

Example 3:
    Input: intervals = [[1,4],[2,3],[3,4]]
    Output: [-1,2,-1]
    Explanation: There is no right interval for [1,4] and [3,4].
        The right interval for [2,3] is [3,4] since start2 = 3 is the smallest start that is >= end1 = 3.

Constraints:
    1 <= intervals.length <= 2 * 10^4
    intervals[i].length == 2
    -10^6 <= starti <= endi <= 10^6
    The start point of each interval is unique.
"""

import heapq
import bisect

#Binary Search
class Solution1(object):
    def findRightInterval(self, intervals):
        n = len(intervals)
        starts = sorted((s, i) for i, (s, e) in enumerate(intervals))
        result = [-1] * n
        
        for i, (s, e) in enumerate(intervals):
            idx = bisect.bisect_left(starts, (e,))
            if idx < n:
                result[i] = starts[idx][1]
        return result


#Min-Heap Simulation
#Time Limit Exceeded -> O(n^2 * log n)
class Solution2(object):
    def findRightInterval(self, intervals):
        n = len(intervals)
        result = [-1] * n

        start_heap = [(s, i) for i, (s,e) in enumerate(intervals)]
        heapq.heapify(start_heap)

        for i, (s, e) in enumerate(intervals):
            temp = list(start_heap)
            while temp and temp[0][0] <  e:
                heapq.heappop(temp)
            if temp:
                result[i]= temp[0][1]
        return result