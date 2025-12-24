package main

import "fmt"

func LinearSearch(arr []int, target int) int {
	for i:=range(len(arr)) {
		if arr[i] == target {
			return i
		}
	}
	return -1
}

func main() {
	// Lets first define an array
	var arr = []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
    // Search with valid target of 10
	var b=LinearSearch(arr,10);
	// Search with invalid target of 11
	var c=LinearSearch(arr,11);
	fmt.Println(b); // Expected output: 9
	fmt.Println(c); // Expected output: -1
}