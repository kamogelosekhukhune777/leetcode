"""

621. Task Scheduler

You are given an array of CPU tasks, each labeled with a letter from A to Z, and a number n. Each CPU interval 
can be idle or allow the completion of one task. Tasks can be completed in any order, but there's a 
constraint: there has to be a gap of at least n intervals between two tasks with the same label.

Return the minimum number of CPU intervals required to complete all tasks.


Example 1:
    Input: tasks = ["A","A","A","B","B","B"], n = 2
    Output: 8
    Explanation: A possible sequence is: A -> B -> idle -> A -> B -> idle -> A -> B.
        After completing task A, you must wait two intervals before doing A again. The same applies to task B. 
        In the 3rd interval, neither A nor B can be done, so you idle. By the 4th interval, you can do A again 
        as 2 intervals have passed.

Example 2:
    Input: tasks = ["A","C","A","B","D","B"], n = 1
    Output: 6
    Explanation: A possible sequence is: A -> B -> C -> D -> A -> B.
        With a cooling interval of 1, you can repeat a task after just one other task.

Example 3:
    Input: tasks = ["A","A","A", "B","B","B"], n = 3
    Output: 10
    Explanation: A possible sequence is: A -> B -> idle -> idle -> A -> B -> idle -> idle -> A -> B.
        There are only two types of tasks, A and B, which need to be separated by 3 intervals. This leads 
        to idling twice between repetitions of these tasks.

 
Constraints:
    1 <= tasks.length <= 104
    tasks[i] is an uppercase English letter.
    0 <= n <= 100
"""
from collections import Counter, deque
import heapq

class Solution:
    def leastInterval(tasks, n):
        freq = Counter(tasks)
        max_heap = [-cnt for cnt in freq.values()]  # max-heap using min-heap
        heapq.heapify(max_heap)

        cooldown = deque()  # stores (available_time, freq)
        time = 0

        while max_heap or cooldown:
            time += 1

            # If task's cooldown ends, push back to heap
            if cooldown and cooldown[0][0] == time:
                _, task_freq = cooldown.popleft()
                heapq.heappush(max_heap, task_freq)

            if max_heap:
                task_freq = heapq.heappop(max_heap)
                if task_freq + 1 < 0:  # still has instances
                    cooldown.append((time + n + 1, task_freq + 1))
            # Else: CPU idles
        return time


print("Example: 1")
tasks = ["A","A","A","B","B","B"]
n = 2
print(Solution.leastInterval(tasks,n))#8

print("Example: 2")
tasks = ["A","C","A","B","D","B"]
n = 1
print(Solution.leastInterval(tasks,n))#6

print("Example: 3")
tasks = ["A","A","A", "B","B","B"]
n = 3
print(Solution.leastInterval(tasks,n))#10
