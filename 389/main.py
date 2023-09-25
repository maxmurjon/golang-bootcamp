class Solution:
    def findTheDifference(self, s: str, t: str) -> str:
        sumS = 0
        sumT = 0
        for v in s:
            sumS += ord(v)
        for v in t:
            sumT += ord(v)
        return chr(sumT - sumS)