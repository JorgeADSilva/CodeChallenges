# Handshakes in table

## Challenge 
We have N persons sitting on a round table. Any person can do a handshake with any other person.
```
     1
2         3
     4
```
Handshake with 2-3 and 1-4 will cause cross.

In how many ways these N people can make handshakes so that no two handshakes cross each other. N would be even.


### TEST CASE 1
````
Input:
N = 2

Output:
1

Explanation:
{1,2} handshake is
possible.
````

### TEST CASE 2
````
Input:
N = 4
Output:
2
Explanation:
{{1, 2}, {3, 4}} are the
two handshakes possible.
{{1, 3}, {2, 4}} is another
set of handshakes possible.
````



### Solution applied (Java)
Idea

- Suppose person 1 handshakes with person k (2≤k≤N), then the table is divided into 2 parts. 
- Part 1 with persons from 2 to <k and Part 2 with persons from k+1 to N.
- Now, no person from Part 1 can handshake with a person from Part 2 as this will cause cross handshaking.
- Thus, each part must have even number of persons and so k must be even.
- Hence, the problem can be divided into two subproblems.
- So, count(N) = count(k-2)*count(N-k) for all even k from 2 to N.

Code: 
```java
static int count(int N) 
{ 
    if(N%2==1) return 0;
   
    else if(N==0)
    return 1;
    int res=0;
 
    for(int i=0;i<N;i+=2)
    res+=count(i)*count(N-2-i);
    
    return res;
    
}
```
