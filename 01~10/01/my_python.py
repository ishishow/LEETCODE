class Solution(object):
    def twoSum(self, nums, target):
        a = len(nums)
        for i in range(a):
            j = i+1
            while j!=a:
                if nums[i] + nums[j] == target:
                    return [i,j]
                j +=1
        return 0
                