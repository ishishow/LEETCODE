class Solution(object):
    def lengthOfLastWord(self, s):
        strs = s.strip().split(" ")
        return len(strs[-1])
        