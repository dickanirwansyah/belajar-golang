package main

import (
	"fmt"
	"math"
	"runtime"
)


func sqrt(x float64) string {
	if x < 0{
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}


func main() {

	fmt.Println("Go berjalan pada OS ")
	switch os := runtime.GOOS; os {

	case "solaris os":
		fmt.Println("SOLARIS")
	case "linux":
		fmt.Println("Linux")
	case "darwin":
		fmt.Println("OS X")
	default:
		fmt.Println("%s.", os)
	}

	sum := 1
	for ; sum < 1000; {
		sum += sum
	}

	sums := 0
	for i:=0; i<10; i++{
		sums += i
	}

	fmt.Println("hasil dari sums :",sums)
	fmt.Println("hasil dari sum :",sum)

	//sqrt
	fmt.Println(sqrt(2), sqrt(-4))
}