class Solution(object):
    def reverse(self, x):
        if x < 0:
            return (self.roop_func(-x) * -1)
        else:
            return self.roop_func(x)
        
    def roop_func(self,x):
        ans = 0
        while x > 0:
            ans = ans*10 + x % 10
            x /= 10
        if -(ans) < -(2**31) or ans > 2**31 - 1:
            return 0
        return ans