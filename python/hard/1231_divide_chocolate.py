"""
Problem Link: https://leetcode.com/problems/divide-chocolate/

You have one chocolate bar that consists of some chunks. Each chunk has
its own sweetness given by the array sweetness.

You want to share the chocolate with your K friends so you start cutting
the chocolate bar into K+1 pieces using K cuts,
each piece consists of some consecutive chunks.

Being generous, you will eat the piece with the minimum total sweetness and
give the other pieces to your friends.

Find the maximum total sweetness of the piece you can get by cutting the
chocolate bar optimally.

Example 1:
    Input: sweetness = [1,2,3,4,5,6,7,8,9], K = 5
    Output: 6
    Explanation: You can divide the chocolate to [1,2,3], [4,5], [6], [7], [8], [9]

Example 2:
    Input: sweetness = [5,6,7,8,9,1,2,3,4], K = 8
    Output: 1
    Explanation: There is only one way to cut the bar into 9 pieces.

Example 3:
    Input: sweetness = [1,2,2,1,2,2,1,2,2], K = 2
    Output: 5
    Explanation: You can divide the chocolate to [1,2,2], [1,2,2], [1,2,2]

Constraints:
    0 <= K < sweetness.length <= 10^4
    1 <= sweetness[i] <= 10^5

"""


def maximizeSweetness(sweetness, K):
    def canDivide(mid):
        cuts = 0
        total = 0
        for sweet in sweetness:
            total += sweet
            if total >= mid:
                cuts += 1
                total = 0
        return cuts >= K + 1

    low, high = 1, sum(sweetness)
    result = 0

    while low <= high:
        mid = (low + high) // 2
        if canDivide(mid):
            result = mid  # Try a higher minimum sweetness
            low = mid + 1
        else:
            high = mid - 1

    return result


example = [1,2,3,4,5,6,7,8,9]
k = 5
print("example 1:")
print(maximizeSweetness(example, k))

example = [5, 6, 7, 8, 9, 1, 2, 3, 4]
k = 8
print("example 2:")
print(maximizeSweetness(example, k))

example = [1, 2, 2, 1, 2, 2, 1, 2, 2]
k = 2
print("example 3:")
print(maximizeSweetness(example, k))
