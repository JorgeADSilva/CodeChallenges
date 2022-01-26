# Gadgets Of Doraland

## Challenge 
In Doraland, people have unique Identity Numbers called D-id. Doraemon owns the most popular gadget shop in Doraland. Since his gadgets are in high demand and he has only K gadgets left he has decided to sell his gadgets to his most frequent customers only. N customers visit his shop and D-id of each customer is given in an array array[ ]. In case two or more people have visited his shop the same number of time he gives priority to the customer with higher D-id. Find the D-ids of people he sells his K gadgets to.

### TEST CASE 1
````
Input: 
1
6
1 1 1 2 2 3
2

Expected Output: 
1 2 
````

### TEST CASE 2
````
1
6
1 1 2 2 2 3 3
1 

Expected Output: 
2 1 3
````



### Solution applied (Java)
```java
class Solution
{
   public class Pair implements Comparable<Pair>{
       int id;
       int value;
       Pair(int id,int value){
           this.id=id;
           this.val=value;
       }
       public int compareTo(Pair otherPair){
           if(this.value==otherPair.value)
           return this.id-otherPair.id;
           return this.value-otherPair.value;
       }
   }
   ArrayList<Integer> TopK(ArrayList<Integer> array, int k)
   {
       ArrayList<Integer> finalArrayList=new ArrayList<>();
       PriorityQueue<Pair> priorityQueue=new PriorityQueue<>(Collections.reverseOrder());
       HashMap<Integer,Integer> hashMap=new HashMap<>();
       for(int arrayElement:array){
           if(hashMap.containsKey(arrayElement)==false)
           hashMap.put(arrayElement,1);
           else
           hashMap.put(arrayElement,hashMap.get(arrayElement)+1);
       }
       for(int key:hashMap.keySet()){
           priorityQueue.add(new Pair(key,hashMap.get(key)));
       }
       
       for(int i=0;i<k;i++){
           Pair rp=priorityQueue.remove();
           finalArrayList.add(rp.id);
       }
       return finalArrayList;
   }
}
```
