package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	arrayCadena := strings.Split(s, ":")
	horario := strings.Trim(arrayCadena[len(arrayCadena)-1], "0123456789")
	arrayCadena[len(arrayCadena)-1] = strings.ReplaceAll(arrayCadena[len(arrayCadena)-1], "PM", "")
	arrayCadena[len(arrayCadena)-1] = strings.ReplaceAll(arrayCadena[len(arrayCadena)-1], "AM", "")

	if horario == "AM" {
		switch arrayCadena[0] {
		case "12":
			arrayCadena[0] = "00"
		}
	}
	if horario == "PM" {
		switch arrayCadena[0] {
		case "01":
			arrayCadena[0] = "13"
		case "02":
			arrayCadena[0] = "14"
		case "03":
			arrayCadena[0] = "15"
		case "04":
			arrayCadena[0] = "16"
		case "05":
			arrayCadena[0] = "17"
		case "06	":
			arrayCadena[0] = "18"
		case "07":
			arrayCadena[0] = "19"
		case "08":
			arrayCadena[0] = "20"
		case "09":
			arrayCadena[0] = "21"
		case "10":
			arrayCadena[0] = "22"
		case "11":
			arrayCadena[0] = "23"
		}
	}

	return strings.Join(arrayCadena, ":")
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

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
