'''
Middle letter
Write a function named mid that takes a 
string as its parameter. Your function 
should extract and return the middle letter. 
If there is no middle letter, your function 
should return the empty string.

For example, 

mid("abc") should return "b" 

and 

mid("aaaa") should return "".
'''

#Solution 1
def mid(parameter):
    halfsize = (len(parameter) // 2)
    if halfsize != (len(parameter) - halfsize):
        return parameter[halfsize]
    else:
        return ""

#Solution 2
def mid(string):
    if len(string) % 2 == 0:
        return ""
    return string[len(string)//2]