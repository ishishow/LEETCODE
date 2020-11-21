def removeDuplicates(self, nums):
    nums[:] = sorted(set(nums))
    return len(nums)


# https://leetcode.com/problems/remove-duplicates-from-sorted-array/discuss/11880/Python-2-liner-O(N)
# 
