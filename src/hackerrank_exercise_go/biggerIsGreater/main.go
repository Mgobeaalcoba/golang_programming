package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'biggerIsGreater' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING w as parameter.
 */

func biggerIsGreater(w string) string {

	arrayString := strings.Split(w, "")
	copiaArray := make([]string, len(arrayString)) // Armo un array vacio independiente del anterior
	copy(copiaArray, arrayString)                  // Copio mi array original en el segundo para luego compararlos
	var init int = len(arrayString)

	for i := len(arrayString) - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			count1j := 1
			if arrayString[i] > arrayString[j] {
				aux := arrayString[j]
				arrayString[j] = arrayString[i]
				arrayString[i] = aux
				init = j + 1
				fmt.Println("J vale:", j, ".I vale:", i)
				fmt.Println("Entro en este primer ciclo:", count1j, "veces")
				fmt.Println("Así está el string recibido en el primer ciclo:", arrayString)
				count1j += 1
				break
			}
		}
		fmt.Println("copiaArray:", copiaArray)
		fmt.Println("arrayString:", arrayString)
		if strings.Join(copiaArray, "") != strings.Join(arrayString, "") {
			break
		}
	}
	if init < len(arrayString)-2 {
		for i := init; i < len(arrayString); i++ {
			//count2i := 1
			for j := i + 1; j < len(arrayString); j++ {
				count1j := 1
				if arrayString[i] > arrayString[j] {
					aux := arrayString[j]
					arrayString[j] = arrayString[i]
					arrayString[i] = aux
					fmt.Println("Entro en este segundo ciclo:", count1j, "veces")
					fmt.Println("Así está el string recibido:", arrayString)
					count1j += 1
				}
			}
		}
	}
	if strings.Join(copiaArray, "") != strings.Join(arrayString, "") {
		return strings.Join(arrayString, "")
	} else {
		return "no answer"
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	TTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	T := int32(TTemp)

	for TItr := 0; TItr < int(T); TItr++ {
		w := readLine(reader)

		result := biggerIsGreater(w)

		fmt.Fprintf(writer, "%s\n", result)
	}

	writer.Flush()
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
