class Solution(object):
    def removeDuplicates(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        
        nw_index = 0
        for i in range(len(nums)):
            if nums[nw_index] != nums[i]:
                nw_index +=1
                nums[nw_index] = nums[i]
        return (nw_index + 1)