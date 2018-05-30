package main

import(
	"fmt"
	"time"
)

func main(){

	//defer function --> peng-eksekusian funsi sampai funsi yang pertama kali berhasil di eksekusi
	//maka lanjut ke fungsi berikutnya
	defer fmt.Println(1)
	fmt.Println(2)

	fmt.Println("kapan hari sabtu ?")

	today := time.Now().Weekday()

	switch time.Saturday{

	case today + 0:
		fmt.Println("sekarang adalah hari sabtu")
	case today + 1:
		fmt.Println("Besok adalah hari sabtu")
	case today + 2:
		fmt.Println("dua hari lagi hari sabtu")
	default:
		fmt.Println("hari sabtu masih jauh")

	}

	//switch condition on go
	t := time.Now()
	switch{
	case t.Hour() < 12:
		fmt.Println("selamat pagi !")
	case t.Hour() < 17:
		fmt.Println("selamat sore !")
	default:
		fmt.Println("selamat malam !")
	}
}