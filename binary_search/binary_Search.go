package main

import "fmt"

func Helper(arr []int,low int,high int,target int) int {
	if low > high {
		return -1;
	}
	var mid = (low + high)/2
	if arr[mid] == target {
		return  mid;
	}else if arr[mid]>target {
		return Helper(arr,low,mid-1,target);
	}
	return Helper(arr,mid+1,high,target)
}

func BinarySearch(arr []int,target int) int {
	return Helper(arr,0,len(arr)-1,target);
}

func main() {
	var arr=[]int{1,2,3,4,5};
	var b=BinarySearch(arr,4);
	fmt.Println("Found element at index",b);
	var c=BinarySearch(arr,0);
	fmt.Println("0 found at index",c);
}