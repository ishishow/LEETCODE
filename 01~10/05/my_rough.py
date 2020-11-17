class Solution(object):
    def longestCommonPrefix(self, strs):
        if not strs or strs[0] == "":
            return ""
        str = strs[0]
        for i in range(len(strs)-1):
            j = 1
            while str[0:j] == strs[i+1][0:j] and j <= len(str):
                j += 1
            if j == 1:
                return ""
            str = str[0:j-1]
        return str
                