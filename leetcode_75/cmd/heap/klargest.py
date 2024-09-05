import heapq
from typing import List

class Solution:
    def findKthLargest(self, nums: List[int], k: int) -> int:
        nums = [-x for x in nums]
        heapq.heapify(nums)
        while k > 1:
            heapq.heappop(nums)
            k=k-1
        
        return -heapq.heappop(nums)


sol = Solution()
print(sol.findKthLargest([12,1,21,214,8,4],2))
