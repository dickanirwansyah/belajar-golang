package main

import "fmt"

func main(){

	var a [2]string
	a[0] = "Hello"
	a[1] = "go !"

	//100 adalah banyaknya kasus
	primes := [100]int {1,2,3,4,5,6,7}

	fmt.Println(a[0], a[1])
	fmt.Println(a)
	fmt.Println(primes)
}