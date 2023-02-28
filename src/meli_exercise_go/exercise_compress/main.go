/*
Given an array of characters chars, compress it using the following algorithm:

Begin with an empty string s. For each group of consecutive repeating characters in chars:

If the group's length is 1, append the character to s.
Otherwise, append the character followed by the group's length.
The compressed string s should not be returned separately, but instead, be stored in the input character array chars. Note that group lengths that are 10 or longer will be split into multiple characters in chars.

After you are done modifying the input array, return the new length of the array.

You must write an algorithm that uses only constant extra space.
*/

package main

import "fmt"

func compress(chars *[]byte) int {
	new_length := 0

	mapa := make(map[byte]int) // Estructura map vacia. Es como un diccionario para Python. Clave byte y valor int

	/*
	Luego, se utiliza un bucle for para iterar sobre cada carácter en la cadena de texto str1. Dentro del bucle, se utiliza un condicional if para verificar si el carácter actual ya existe en el mapa mapa. Si el carácter aún no se ha encontrado en el mapa, se agrega como una nueva clave en el mapa y se establece su valor en 1. Si el carácter ya existe en el mapa, se incrementa su valor en 1.
	*/
	for _, v := range *chars {
		if _, ok := mapa[v]; !ok { 
			mapa[v] = 1
		} else {
			mapa[v] += 1
		}
	}
	
	new_arr := []byte{}

	// Uso un for each para recorrer clave-valor mi mapa
	// Uso el metodo string para transformar la letra en unicode (numero) en letra e imprimirlo transformado
	

	for k, v := range mapa {
		//fmt.Printf("%s:%d ",  string(k), v) 
		//fmt.Println()
		new_arr = append(new_arr, k)
		if v == 1 {
			continue
		}else if v >= 10 {
			div := v / 10
			mod := v % 10
			new_arr = append(new_arr, byte(div))
			new_arr = append(new_arr, byte(mod))
		} else {
			new_arr = append(new_arr, byte(v))
		}
	}

	//fmt.Println(new_arr)

	/* Intento por convertir los unicodes de golang a strings: 
	for i := 0; i < len(new_arr); i++ {
		if i % 2 == 0 {
			fmt.Print(string(new_arr[i]),":")
		} else {
			fmt.Println(new_arr[i])
		}
	}
	*/

	new_length = len(new_arr)

	*chars = new_arr

	return new_length
}

func main() {
	arr1 := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	arr2 := []byte{'a'}
	arr3 := []byte{'a','b','b','b','b','b','b','b','b','b','b','b','b'}

	dev := compress(&arr1)
	//fmt.Printf("El tipo de arr1 es %T\n", arr1)
	fmt.Printf("El largo del array devuelto es de %d\n", dev)
	fmt.Println(arr1)
	fmt.Println("Finished")
	dev = compress(&arr2)
	//fmt.Printf("El tipo de arr2 es %T\n", arr2)
	fmt.Printf("El largo del array devuelto es de %d\n", dev)
	fmt.Println(arr2)
	fmt.Println("Finished")
	dev = compress(&arr3)
	//fmt.Printf("El tipo de arr3 es %T\n", arr3)
	fmt.Printf("El largo del array devuelto es de %d\n", dev,)
	fmt.Println(arr3)
	fmt.Println("Finished")

}

/*

Input: chars = ["a","a","b","b","c","c","c"]
Output: Return 6, and the first 6 characters of the input array should be: ["a","2","b","2","c","3"]
Explanation: The groups are "aa", "bb", and "ccc". This compresses to "a2b2c3".

Input: chars = ["a"]
Output: Return 1, and the first character of the input array should be: ["a"]
Explanation: The only group is "a", which remains uncompressed since it's a single character.

Input: chars = ["a","b","b","b","b","b","b","b","b","b","b","b","b"]
Output: Return 4, and the first 4 characters of the input array should be: ["a","b","1","2"].
Explanation: The groups are "a" and "bbbbbbbbbbbb". This compresses to "ab12".

Constraints:

1 <= chars.length <= 2000
chars[i] is a lowercase English letter, uppercase English letter, digit, or symbol.

*/