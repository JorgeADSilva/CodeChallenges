'''
Online status
The aim of this challenge is, given a dictionary 
of people's online status, to count the number 
of people who are online.

For example, consider the following dictionary:

statuses = {
    "Alice": "online",
    "Bob": "offline",
    "Eve": "online",
}

In this case, the number of people online is 2.

Write a function named online_count that takes one 
parameter. The parameter is a dictionary that maps 
from strings of names to the string "online" or 
"offline", as seen above.

Your function should return the number of people 
who are online.
'''

# Solution 1
def online_count(dictionary):
    online_status_counter = 0
    for person_name, status in dictionary.keys():
        if status == "online":
            online_status_counter += 1
    return online_status_counter

# Solution 2
def online_count(people):
    return len([p for p in people if people[p] == "online"])