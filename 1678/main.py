# First Solution
class Solution(object):
    def interpret(self, command):
        """
        :type command: str
        :rtype: str
        """
        command=command.replace('()','o')
        command=command.replace('(al)','al')
        return command

# Second Solution
class SolutionTwo(object):
    def interpret(self, command):
        """
        :type command: str
        :rtype: str
        """
        text=""
        i=0
        while i<len(command):
            if command[i]=="G":
                text+="G"
                i+=1
            elif command[i:i+2]=="()":
                text+="o"
                i+=2
            elif command[i:i+4]=="(al)":
                text+="al"
                i+=4
        return text