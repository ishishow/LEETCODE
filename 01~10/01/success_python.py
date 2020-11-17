class Solution(object):
    def twoSum(self, nums, target):
        d = {} 
        for i in range(len(nums)):
            if (target - nums[i]) in d:
                return [d[target - nums[i]],i]
            else:
                d[nums[i]] = i
        return 0