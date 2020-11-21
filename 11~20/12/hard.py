class Solution(object):
    def countAndSay(self, n):
        
        if n == 1:
            return "1"
        res_str = self.countAndSay(n-1)
        return_str = ""
        count=0
        stack=res_str[0]
        for i in range(len(res_str)):
            if res_str[i] != stack:
                return_str += str(count) + str(stack)
                count = 0
            stack = res_str[i]
            count += 1
        return_str += str(count) + str(res_str[len(res_str)-1])
        return return_str
                
                