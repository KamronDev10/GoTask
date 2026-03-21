// Son kiritiladi (1, 10) kiriltilgan sonni so’zlar orqali ifodalovchi dastur tuzing

package main

import "fmt"

func main() {

	var son int

	fmt.Print("son kiritng  ")
	fmt.Scan(&son)

	switch son {
	case 1:
		fmt.Println("bir")
	case 2:
		fmt.Println("ikki")
	case 3:
		fmt.Println("uch ")
	case 4:
		fmt.Println("to'rt")
	case 5:
		fmt.Println("besh")
	case 6:
		fmt.Println("olti")
	case 7:
		fmt.Println("yetti")
	case 8:
		fmt.Println("sakkiz")
	case 9:
		fmt.Println("to'qqiz")
	case 10:
		fmt.Println("o'n")

	}
}
