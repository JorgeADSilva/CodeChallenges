# Implementing Dijkstra Algorithm

## Challenge 
Given a weighted, undirected and connected graph of V vertices and E edges, Find the shortest distance of all the vertex's from the source vertex S.
Note: The Graph doesn't contain any negative weight cycle.

### TEST CASE 1
````
Input: 
1
2 1
0 1 9
0

Output: 
0 9 
````

### TEST CASE 2
````
Input: 
1
3 4
0 1 1
0 2 6
1 2 3
2 1 3
2

Output:
4 3 0
````

### TEST CASE 3
````
Input: 
1
3 4
0 1 1
0 2 6
1 2 3
2 1 3
0

Output:
0 1 4
````


### Solution applied - Approach 1 - Aux Matrix (Java)

**Algorithm**:

1) Create a set sptSet (shortest path tree set) that keeps track of vertices included in shortest path tree, i.e., whose minimum distance from source is calculated and finalized. Initially, this set is empty. 
2) Assign a distance value to all vertices in the input graph. Initialize all distance values as INFINITE. Assign distance value as 0 for the source vertex so that it is picked first. 
3) While sptSet doesn’t include all vertices 

Pick a vertex u which is not there in sptSet and has minimum distance value. 
Include u to sptSet. 
Update distance value of all adjacent vertices of u. To update the distance values, iterate through all adjacent vertices. For every adjacent vertex v, if sum of distance value of u (from source) and weight of edge u-v, is less than the distance value of v, then update the distance value of v. 


```java

class Solution
{
    static int[] dijkstra(int V, ArrayList<ArrayList<ArrayList<Integer>>> adj, int S)
    {
        PriorityQueue<Node> pq = new PriorityQueue<Node>(new Comparator<Node> (){
            public int compare(Node n1, Node n2){
                return n1.value-n2.value;
            }
        });
        
        boolean [] vis = new boolean [V];
        int [] dist = new int[V];
        for(int i = 0; i < V ; i++){
            dist[i] = Integer.MAX_VALUE;
        }
        dist[S] = 0;
        pq.add(new Node(S,0));
        
        
        while(!pq.isEmpty()){
            Node curr = pq.poll();
            vis[curr.id] = true;
            
            for(ArrayList<Integer> i : adj.get(curr.id)){
                int end = i.get(0);
                int w = i.get(1);
              
                if(vis[end]) continue;
                
                int tempDist = dist[curr.id]+ w;
                
                if(dist[end] > tempDist){
                    dist[end] = tempDist;
                    pq.add(new Node(end, dist[end]));
                }
            }
        }
		return dist;
    }
    
public static class Node{
    int id;
    int value;
    
    public Node(int id, int value){
        this.id = id;
        this.value = value;
    }
}
}

```

