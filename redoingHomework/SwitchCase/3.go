// Son kiritiladi (1, 1000) kiriltilgan sonni so’zlar orqali ifodalovchi dastur tuzing

package main

import "fmt"

func main() {

	var son int

	fmt.Print("son kiritng :  ")
	fmt.Scan(&son)

	yuzlar := son / 100
	onlik := son / 10 % 10
	birlik := son % 10

	switch yuzlar {
	case 1:
		fmt.Println("yuz")
	case 2:
		fmt.Println("ikki yuz")
	case 3:
		fmt.Println("uch yuz ")
	case 4:
		fmt.Println("to'rt yuz")
	case 5:
		fmt.Println("besh yuz")
	case 6:
		fmt.Println("olti yuz ")
	case 7:
		fmt.Println("yetti yuz ")
	case 8:
		fmt.Println("sakkiz yuz")
	case 9:
		fmt.Println("to'qqiz yuz")
	case 10:
		fmt.Println("ming ")

	}

	switch onlik {
	case 1:
		fmt.Println("o'n")
	case 2:
		fmt.Println("yigirma")
	case 3:
		fmt.Println("o'ttiz ")
	case 4:
		fmt.Println("qirq")
	case 5:
		fmt.Println("ellik")
	case 6:
		fmt.Println("oltmish ")
	case 7:
		fmt.Println("tetmish ")
	case 8:
		fmt.Println("sakson")
	case 9:
		fmt.Println("to'qson")
	case 10:
		fmt.Println("yuz ")

	}

	switch birlik {
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
