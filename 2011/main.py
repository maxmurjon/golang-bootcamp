class Solution(object):
    def finalValueAfterOperations(self, operations):
        x = 0
        for operation in operations:
            if "++" in operation:
                x += 1
            elif "--" in operation:
                x -= 1
        return x
solution = Solution()
result = solution.finalValueAfterOperations(["--X", "X++", "X++"])
print(result)