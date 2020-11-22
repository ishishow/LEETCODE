class Solution:
    def climbStairs(self, n):
        a, b = 1, 1
        for i in range(n):
            tmp = b
            b = a + b
            a = tmp
            
        return a