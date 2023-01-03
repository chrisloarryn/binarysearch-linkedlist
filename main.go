package main

import (
	"fmt"
	listmanager "github.com/chrisloarryn/listmanager"
)

func main() {
	// 1. Escribe una funcion que realice una buqueda binaria de un ID `n`
	// sobre el siguiente slice:
	productIDs := []int{2, 5, 7, 9, 23, 28, 36, 39, 56, 68, 75, 76, 77, 88, 89, 99}
	// y retorne la posición donde se encontro y el numero de iteraciones
	// para encontrarlo.

	pos, it := BinarySearch(99, productIDs)
	println("pos:", pos, "it:", it)

	//	2. Escribe una libreria para poder manejar Listas Enlazadas, y crea
	//  un ejemplo basico de uso.

	l := listmanager.LinkedList{}

	fmt.Println("\n************* Insert *************")
	l.Insert(12)
	l.Insert(13)
	l.Insert(14)
	l.Insert(15)
	fmt.Println("************* Print *************")
	l.Print()

	fmt.Println("\n************* Search *************")
	fmt.Println("Position of 12 value: ", l.Search(12))
	fmt.Println("Position of 14 value: ", l.Search(14))
	fmt.Println("Position of 15 value: ", l.Search(15))
	fmt.Println("Position of 100 value: ", l.Search(100))

	fmt.Println("\n************* GetAt *************")
	fmt.Println("Get At 1st position: ", l.GetAt(0))
	fmt.Println("Get At 3rd position: ", l.GetAt(2))
	fmt.Println("Get At 4th position: ", l.GetAt(3))
	fmt.Println("Get At -5 position (head is returned): ", l.GetAt(-5))

	fmt.Println("\n************* DeleteAt *************")
	l.DeleteAt(3)
	fmt.Println("************* Print *************")
	l.Print()
	// scatalan@deuna.com
}

func BinarySearch(n int, productIDs []int) (position int, iterations int) {
	// your code here
	low := 0
	high := len(productIDs) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := productIDs[mid]
		if guess == n {
			return mid, iterations
		}
		if guess > n {
			high = mid - 1
		} else {
			low = mid + 1
		}
		iterations++
	}

	return -1, iterations
}
