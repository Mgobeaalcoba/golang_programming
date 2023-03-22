package main

import (
	"fmt"
)

func getTotalX(a []int32, b []int32) int32 {
	var count int32 = 0
	myList := []int32{}
	var validate bool = true

	for i := a[0]; i <= b[0]; i++ {
		for _, value := range a {
			//fmt.Println("value:", value)
			//fmt.Println("i:", i)
			if i%value == 0 {
				continue
			} else {
				validate = false
				break
			}
		}
		if validate {
			//fmt.Println("Agregado a la lista:", i)
			myList = append(myList, int32(i))
		}
		validate = true
	}

	//fmt.Println("Imprimo mi lista a mitad de proceso:")
	//fmt.Println(myList)

	validate = true

	for _, valueMyList := range myList {
		for _, valueB := range b {
			//fmt.Println("valueB:", valueB)
			//fmt.Println("valueMyList:", valueMyList)
			if valueB%valueMyList == 0 {
				continue
			} else {
				validate = false
				break
			}
		}
		if validate {
			count += 1
		}
		validate = true
	}

	return count

}

func main() {
	a := []int32{2, 4}
	b := []int32{16, 32, 96}

	fmt.Println(getTotalX(a, b))
}
