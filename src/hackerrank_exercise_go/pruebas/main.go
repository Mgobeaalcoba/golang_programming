package main

import (
	"fmt"
)

func birthday(s []int32, d int32, m int32) int32 {
    // Write your code here
    var count int32 = 0
    var init int32 = 0
    var sum int32 = 0
    for init <= ( int32(len(s)) - m){
		//fmt.Println("init:", init)
        for i := init; i < init + m; i++ {
			//fmt.Println("i:", i)
			//fmt.Println("init + m:", init + m)
            sum += s[i]
			//fmt.Println("sum:", sum)
		}
		//fmt.Println("Comparo sum:",sum,"y d:",d)
        if sum == d {
            count += 1
			//fmt.Println("count queda como:", count)
        }
        init += 1
        sum = 0
    }
    return count
}

func main() {
	s := []int32{2, 2, 1, 3, 2}
	var d int32 = 4
	var m  int32 = 2

	fmt.Println(birthday(s, d, m))
}
