package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */
 
 func roundToSixDecimalPlaces(num float64) float64 {

    return math.Round(num*1000000) / 1000000
	
}

func plusMinus(arr []int32) {

    len_arr := len(arr)
    var positive_count int
    var negative_count int
    var cero_count int
    
    for i := 0; i < len(arr); i++ {
        if arr[i] < 0 {
            negative_count += 1
        } else if arr[i] > 0 {
            positive_count += 1
        } else {
            cero_count += 1
        }
    }
    
    //fmt.Println(positive_count)
    //fmt.Println(negative_count)
    //fmt.Println(cero_count)
    //fmt.Println(len_arr)
    
    //fmt.Println()
    
    positive_ratio := float64(positive_count) / float64(len_arr)
    negative_ratio := float64(negative_count) / float64(len_arr)
    cero_ratio := float64(cero_count) / float64(len_arr)
    
    //fmt.Println(positive_ratio)
    //fmt.Println(negative_ratio)
    //fmt.Println(cero_ratio)
    
    fmt.Printf("%f\n",roundToSixDecimalPlaces(float64(positive_ratio)))
    fmt.Printf("%f\n",roundToSixDecimalPlaces(float64(negative_ratio)))
    fmt.Printf("%f\n",roundToSixDecimalPlaces(float64(cero_ratio)))
    
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    plusMinus(arr)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
