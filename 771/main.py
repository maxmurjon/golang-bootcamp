class Solution:
    def numJewelsInStones(self, jewels: str, stones: str) -> int:
        counter=0
        for i in stones:
            if i in jewels:
                counter+=1
        return counter