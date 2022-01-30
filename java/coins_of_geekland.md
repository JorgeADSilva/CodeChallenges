# Coins of Geekland 

## Challenge 
In Geekland there is a grid of coins of size N x N. You have to find the maximum sum of coins in any sub-grid of size K x K.
Note: Coins of the negative denomination are also possible at Geekland.

### TEST CASE 1
````
Input: N = 5, K = 3 
mat[[]] = {1, 1, 1, 1, 1} 
          {2, 2, 2, 2, 2} 
          {3, 8, 6, 7, 3} 
          {4, 4, 4, 4, 4} 
          {5, 5, 5, 5, 5}
Output: 48
Explanation: {8, 6, 7}
             {4, 4, 4}
             {5, 5, 5}
has the maximum sum
````

### TEST CASE 2
````
Input: N = 1, K = 1
mat[[]] = {{4}} 
Output: 4
````



### Solution applied - Approach 1 - Aux Matrix (Java)
```java
class Solution
{
    public int Maximum_Sum(int mat[][],int N,int K){
        if(K>N) return 0;
        int[][] auxMatrix = new int[N][N];
        for(int i = 0;i<N;i++)
            for(int j = 0;j<N;j++){
                if(i == 0&&j==0) 
                    auxMatrix[i][j] = mat[i][j];
                else if(i==0) 
                    auxMatrix[i][j] = auxMatrix[i][j-1] + mat[i][j];
                else if(j==0) 
                    auxMatrix[i][j] = auxMatrix[i-1][j] + mat[i][j];
                else 
                    auxMatrix[i][j] = auxMatrix[i-1][j] + auxMatrix[i][j-1] + mat[i][j] -auxMatrix[i-1][j-1];
            }
        if(K==N) 
            return auxMatrix[N-1][N-1];
        int sum = 0;
        for(int i = K-1;i<N;i++)
            for(int j = K-1;j<N;j++){
                if(i==K-1 && j == K-1) 
                    sum = auxMatrix[K-1][K-1];
                else if(i == K-1) 
                    sum = Math.max(sum,auxMatrix[i][j]-auxMatrix[i][j-K]);
                else if(j == K-1) 
                    sum = Math.max(sum,auxMatrix[i][j]-auxMatrix[i-K][j]);
                else 
                    sum = Math.max(sum,auxMatrix[i][j]-auxMatrix[i-K][j]-auxMatrix[i][j-K]+auxMatrix[i-K][j-K]);
       }
       return sum;
    }
}
```


### Solution applied - Approach 2 - Moves Calculation (Java)
```java
class Solution
{
    public int Maximum_Sum(int mat[][],int N,int K){
        int max=0,sum=0;
        for(int moveHorizontally=0;moveHorizontally<=N-K;moveHorizontally++)
        for(int moveVertically=0;moveVertically<=N-K;moveVertically++){
            sum=0;
            for(int i=moveVertically;i<moveVertically+K;i++){
                for(int j=moveHorizontally;j<moveHorizontally+K;j++){
                    sum+=mat[i][j];
                }
                if(sum>max)
                    max=sum;
            }
       }
       return max;
   }
}
```
