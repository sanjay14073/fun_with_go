package main

func Prime(n int) bool {
	if n<=1 {
		return false
	}
	for i:=2; i<n; i++ {
		if n%i==0 {
			return false
		}
	}
	return true
}

func main() {
	var n=6
	println(Prime(n))
}