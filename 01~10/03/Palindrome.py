class Solution(object):
    def isPalindrome(self, x):
        if x >= 0:
            if str(x) == str(x)[::-1]:
                return True
        return False
            
        