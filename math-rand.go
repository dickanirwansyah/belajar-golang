package main

import(
	"fmt"
	"math"
	"math/rand"
)

//seperti ini
func tambahkan(x int, y int, z int) int {
	return x + y - z
}

//bisa juga seperti ini
func tambah(x, y, z int) int {
	return x + y - z
}

//fungtion swap
func swap(x, y string) (string, string){
	return x, y
}

func main(){
	fmt.Println("Bilangan favorit : ",rand.Intn(9))
	fmt.Println("Bilang %g lagi: ",math.Sqrt(7))
	fmt.Println("Hasil dari x+y-z: ",tambah(100, 100, 7))
	fmt.Println("Hasil dari x+y-z: ",tambahkan(100,100,7))

	//swap
	a, b := swap("belajar", "golang basic")
	fmt.Println("hasil dari swap", a, b)
}