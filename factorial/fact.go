package main

import "fmt"

func Fact(a int) int{
	if a==1 || a==0 {
		return 1;
	}
	return a*Fact(a-1);
}

func main() {
	var a = 4
	fmt.Println("Factoriial of 4",Fact(a));
}