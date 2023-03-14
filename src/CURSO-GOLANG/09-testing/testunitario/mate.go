package testunitario

func Suma(a, b int) int {
	return a + b
}

func GetMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

/*
Si solo testeo Suma no estoy testeando el 100% de mi aplicación. Para saber que % de mi app estoy testeando puedo usar:

go test -cover

PASS
        github.com/Mgobeaalcoba/golang_programming/tree/main/src/CURSO-GOLANG/testing/testunitario      coverage: 25.0% of statements
ok      github.com/Mgobeaalcoba/golang_programming/tree/main/src/CURSO-GOLANG/testing/testunitario      0.002s

Si quiero puedo exportar los resultados de mi testeo de la siguiente forma:

go test -coverprofile=coverage.out

        github.com/Mgobeaalcoba/golang_programming/tree/main/src/CURSO-GOLANG/testing/testunitario  coverage: 25.0% of statements
ok      github.com/Mgobeaalcoba/golang_programming/tree/main/src/CURSO-GOLANG/testing/testunitario  0.002s

Y generó un file coverage.out en mi carpeta de testunitario.

Este archivo lo voy a usar para saber cuales de mis funciones ya están testeadas y en que porcentaje así:

go tool cover -func=coverage.out

github.com/Mgobeaalcoba/golang_programming/tree/main/src/CURSO-GOLANG/testing/testunitario/mate.go:3:       Suma            100.0%
github.com/Mgobeaalcoba/golang_programming/tree/main/src/CURSO-GOLANG/testing/testunitario/mate.go:7:       GetMax          0.0%
total:                                                                     (statements)     25.0%

Tmb podría ver los resultados de mis testeos de forma visualmente simple con el siguiente comando:

go tool cover -html=coverage.out

Me va a abrir una pagina donde lo testeado aparecerá en verde y lo no testeado en rojo.

-------------------------------------------

En ocasiones, una de nuestras funciones, por ejemplo fibonacci, puede ser mucho mas lenta que las demas. Por tal motivo, puedo observar su performance de la siguiente forma:

go test -cpuprofile=cou.out

Crea un archivo cou.out y otro archivo llamado testunitario.test

luego debo correr el siguiente comando:

go tool pprof cou.out

Abre una interaccion con el archivo y escribimos "top"

y nos mostrará donde está ocurriendo la demora significativa. En este caso en "fibonaci"

14:53:22 👽 with 🤖 mgobea 🐶 in CURSO-GOLANG/09-testing/testunitario on  main [!] via 🐹 v1.20 took 1m 0.4s …
➜ go tool pprof cou.out
File: testunitario.test
Type: cpu
Time: Mar 14, 2023 at 2:52pm (-03)
Duration: 60.03s, Total samples = 59.72s (99.48%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 59.60s, 99.80% of 59.72s total
Dropped 29 nodes (cum <= 0.30s)
      flat  flat%   sum%        cum   cum%
    59.60s 99.80% 99.80%     59.65s 99.88%  github.com/Mgobeaalcoba/golang_programming/tree/main/src/CURSO-GOLANG/testing/testunitario.Fibonacci
         0     0% 99.80%     59.65s 99.88%  github.com/Mgobeaalcoba/golang_programming/tree/main/src/CURSO-GOLANG/testing/testunitario.TestFibo
         0     0% 99.80%     59.65s 99.88%  testing.tRunner
(pprof)

Puedo ver en que parte del codigo se está demorando agregando al "pprof" el pedido de "list Fibonacci"

(pprof) list Fibonacci
Total: 59.72s
ROUTINE ======================== github.com/Mgobeaalcoba/golang_programming/tree/main/src/CURSO-GOLANG/testing/testunitario.Fibonacci in /home/mgobea/develop/go/src/CURSO-GOLANG/09-testing/testunitario/mate.go
    59.60s     95.24s (flat, cum) 159.48% of Total
    24.85s     24.86s     14:func Fibonacci(n int) int {
     2.57s      2.57s     15:   if n <= 1 {
     8.16s      8.16s     16:           return n
         .          .     17:   }
         .          .     18:
    24.02s     59.65s     19:   return Fibonacci(n-1) + Fibonacci(n-2)
         .          .     20:}
         .          .     21:
         .          .     22:/*
         .          .     23:Si solo testeo Suma no estoy testeando el 100% de mi aplicación. Para saber que % de mi app estoy testeando puedo usar:
         .          .     24:
(pprof)

Y ahí veo que la demora se está originando en la recursión del return.

Si escribo en "pprof" "web" me va a abrir un grafico donde veré de forma "mas visual" donde está la demora.

Con "quit" en "pprof" salgo del tool "pprof" de golang y puedo continuar usando la consola.

Escribiendo dentro del "pprof" "pdf" voy a ver el mismo reporte grafico svg que vi con "web" en un reporte formato "pdf" que quedará guardada en la carpeta de testeo.

Luego en Linux puedo visualizarlo con el programa GIMP y el comando

gimp "nombre.pdf"
*/
