class Solution:
    def mergeAlternately(self, word1: str, word2: str) -> str:
        text=""
        if len(word1)>=len(word2):
            for i in range(len(word2)):
                text+=word1[i]+word2[i]
            if len(word1)>len(word2):
                text+=word1[len(word2):]
        else:
            for i in range(len(word1)):
                text+=word1[i]+word2[i]
            if len(word2)>len(word1):
                text+=word2[len(word1):]
        return text

