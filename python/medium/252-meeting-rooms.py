"""
252. Meeting Rooms

Given an array of meeting time intervals where intervals[i] = [starti, endi], 
determine if a person could attend all meetings.

Return true if a person can attend all meetings, otherwise return false.

Example 1:
    Input: [[0,30],[5,10],[15,20]]
    Output: false

Example 2:
    Input: [[7,10],[2,4]]
    Output: true

Constraints:
    0 <= intervals.length <= 10^4
    intervals[i].length == 2
    0 <= starti < endi <= 10^6
"""

class Solution:
    def canAttendMeetings(intervals):
        intervals.sort(key = lambda x :x[0])

        for i in  range(1, len(intervals)):
            prev_end = intervals[i - 1][1]
            curr_start = intervals[i][0]

            if curr_start < prev_end: 
                return False
            
        return True
