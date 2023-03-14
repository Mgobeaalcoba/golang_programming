package testunitario

import "testing"

// Modulo para testing.

/* Primera forma de testeo:
func TestSuma(t *testing.T) {
	total := Suma(5,5)

	// Condicional que expresa que ha habido un error
	if total != 11 {
		t.Errorf("Suma incorrecta, tiene %d se esperaba %d", total, 11)
	}
}
*/

/* Ejecuto mi codigo por consola con "$ go test" desde "testunitario"
Consola y Output si no falla:

C:\Users\mgobea\go\src\CURSO-GOLANG\09-testing\testunitario(main -> origin)
λ go test
PASS
ok      github.com/Mgobeaalcoba/golang_programming/tree/main/src/CURSO-GOLANG/testing/testunitario      0.768s

Consola y Output si falla:

C:\Users\mgobea\go\src\CURSO-GOLANG\09-testing\testunitario(main -> origin)
λ go test
--- FAIL: TestSuma (0.00s)
    mate_test.go:10: Suma incorrecta, tiene 10 se esperaba 11
FAIL
exit status 1
FAIL    github.com/Mgobeaalcoba/golang_programming/tree/main/src/CURSO-GOLANG/testing/testunitario      0.716s
*/

// Testing convencional:
func TestSuma(t *testing.T) {
	// Slicen de structs. Tabla de testeo convencional
	tabla := []struct {
		a int
		b int
		c int
		// Para cada struct de abajo, "a" va a ser el primer valor, "b" el segundo valor y "c" el tercer valor
	}{
		{1, 2, 3},    // Se lee: 1 y 2 debería resultar en 3
		{2, 2, 4},    // 2 y 2 debería resultar 4
		{25, 25, 50}, // 25 y 25 debería resultar 50
	}

	// Aqui entonces voy a testear mi función "Suma" con mi "tabla"
	for _, item := range tabla {
		total := Suma(item.a, item.b)

		if total != item.c {
			t.Errorf("Suma incorrecta, tiene %d se esperaba %d", total, item.c) // Gracias a este msj se que en que prueba está fallando
		}
	}
}

// Desafio: Crear una función para testear GetMax que aún no está hecho:

func TestGetMax(t *testing.T) {
	tabla := []struct {
		a int
		b int
		c int
	}{
		{2, 4, 4}, // Importante: enviar un ejemplo donde a sea el mayor y otro donde b sea el mayor para testear el 100% de nuestra función.
		{3, 1, 3},
		{5, 17, 17},
	}

	for _, item := range tabla {
		maximo := GetMax(item.a, item.b)

		if maximo != item.c {
			t.Errorf("Maximo incorrecto, devuelve %d se esperaba %d", maximo, item.c)
		}
	}
}

func TestFibo(t *testing.T) {
	tabla := []struct {
		a int
		b int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, item := range tabla {
		myFibo := Fibonacci(item.a)

		if myFibo != item.b {
			t.Errorf("Fibonacci incorrecto, devuelve %d se esperaba %d", myFibo, item.b)
		}
	}
}

/* Consola y Output si no falla:

C:\Users\mgobea\go\src\CURSO-GOLANG\09-testing\testunitario(main -> origin)
λ go test
PASS
ok      github.com/Mgobeaalcoba/golang_programming/tree/main/src/CURSO-GOLANG/testing/testunitario      0.725s

Consola y Output si falla:

C:\Users\mgobea\go\src\CURSO-GOLANG\09-testing\testunitario(main -> origin)
λ go test
--- FAIL: TestSuma (0.00s)
    mate_test.go:55: Suma incorrecta, tiene 3 se esperaba 4
FAIL
exit status 1
FAIL    github.com/Mgobeaalcoba/golang_programming/tree/main/src/CURSO-GOLANG/testing/testunitario      0.866s
*/
