package main

import (
	"fmt"
	"sort"
	"time"
	"untref-ayp2/guia-algoritmos/mediciones/busquedas"
	"untref-ayp2/guia-algoritmos/mediciones/utiles"
)

func main() {
	arreglo := utiles.GenerarArreglo(10, 100000)
	buscado := -1

	//fmt.Println(arreglo)

	inicio := time.Now()
	// Busqueda Lineal
	busquedas.BusquedaLineal(arreglo, buscado)
	fmt.Println("Busqueda Lineal: ", time.Since(inicio))

	inicio = time.Now()
	// Ordenar el arreglo para la busqueda binaria
	sort.Ints(arreglo)
	fmt.Println("Ordenamiento:    ", time.Since(inicio))
	//fmt.Println(arreglo)

	inicio = time.Now()
	// Busqueda Binaria
	busquedas.BusquedaBinaria(arreglo, buscado)
	fmt.Println("Busqueda Binaria:", time.Since(inicio))
}
