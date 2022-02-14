# Maximum number of events that can be attended

## Challenge 
There are **N** events in Geek's city. You are given two arrays **start[]** and **end[]** denoting starting and ending day of the events respectively. Event **i** starts at **start[i]** and ends at **end[i]**.
You can attend an event **i** at any day **d** between **start[i]** and **end[i]** (start[i] ≤ d ≤ end[i]). But you can attend only one event in a day.
Find the maximum number of events you can attend.

### TEST CASE 1
````
Input:

N = 3

start[] = {1, 2, 1}

end[] = {1, 2, 2}

Output:
2

Explanation:
You can attend a maximum of two events.
You can attend 2 events by attending 1st event
at Day 1 and 2nd event at Day 2.
````

### TEST CASE 2
````
Input:
N = 3

start[i] = {1, 2, 3}

end[i] = {2, 3, 4} 

Output :
3

Explanation:
You can attend all events by attending event 1
at Day 1, event 2 at Day 2, and event 3 at Day 3.
````



### Solution applied (Java)

```java
class Solution {
    static class Pair{
        int startingDay;
        int endingDay;
        
        Pair(int startingDay, int endingDay){
            this.startingDay = startingDay;
            this.endingDay = endingDay; 
        }
    }
    
    static int maxEvents(int[] start, int[] end, int N) {
        Pair[] pairsArray = new Pair[N];
        int index = 0, eventsCount = 0, dayIndex= 0;
        
        for(int i = 0; i < N; i++){
            pairsArray[i] = new Pair(start[i], end[i]);
        }
        
        Arrays.sort(pairsArray, (a,b)-> (a.startingDay - b.startingDay));
        
        PriorityQueue<Integer> priorityQueue = new PriorityQueue<Integer>();
        
        while(!priorityQueue.isEmpty() || index < N){
            if(priorityQueue.isEmpty())
                dayIndex = pairsArray[index].startingDay;
            
            while(index<N && pairsArray[index].startingDay <=dayIndex){
                priorityQueue.offer(pairsArray[index++].endingDay);
            }
            priorityQueue.poll();
            ++eventsCount;
            ++dayIndex;
            while (!priorityQueue.isEmpty() && priorityQueue.peek() < dayIndex)
               priorityQueue.poll();
        }
        return eventsCount;
    }
};
```
