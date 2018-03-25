package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := []int{1, 0, 3, 9, 2}
	fmt.Printf(strconv.Itoa(maxSum(arr)))
}

func maxSum(arr []int) int {
	sum := make([]int, len(arr))
	//A[i] = Max(A[i - 2] + watcA[i], A[i - 1]) ; Bottom-up.
	sum[0] = arr[0]
	sum[1] = Max(sum[0], arr[1])
	for i := 2; i <= len(arr)-1; i++ {
		sum[i] = Max(sum[i-2]+arr[i], sum[i-1])
	}

	return sum[len(arr)-1]
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*

int dp[2] = {a[0],a[1]};

for(int i=2;i<a.size();i++)
{
    int temp = dp[1];
    dp[1] = dp[0]+a[i];
    dp[0] = max(dp[0],temp);
}

int ans = max(dp[0],dp[1]);

----

int sum(int *arr,int size) {
        int *sum = malloc(sizeof(int) * size);
        int i;

        for(i=0;i<size;i++) {
                if(i==0) {
                        sum[0] = arr[0];
                }else if(i==1) {
                        sum[1] = max(sum[0],arr[1]);
                }else{
                        sum[i] = max(sum[i-2]+arr[i],sum[i-1]);
                }
        }
        return sum[size-1];
}
 */
