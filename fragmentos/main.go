package main

import (
	"fmt"
	"untref-ayp2/guia-algoritmos/fragmentos/utiles"
)

func main() {
	arreglo := []int{1, 2, 3, 4, 5}
	fmt.Println("Arreglo Original: ", arreglo)

	arregloInvertido := utiles.InvertirArreglo(arreglo)
	fmt.Println("Arreglo Invertido: ", arregloInvertido)

	utiles.InvertirArreglo2(arreglo)
	fmt.Println("Arreglo Invertido 2: ", arreglo)

	matrizA := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	matrizB := [][]int{
		{7, 8},
		{9, 10},
		{11, 12},
	}
	matrizProducto, err := utiles.ProductoMatrices(matrizA, matrizB)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Producto de Matrices:")
		for _, fila := range matrizProducto {
			fmt.Println(fila)
		}
	}

	arregloTernario := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	buscar := 5
	indice := utiles.BusquedaTernaria(arregloTernario, buscar)

	if indice != -1 {
		fmt.Printf("El valor %d se encuentra en el índice %d\n", buscar, indice)
	} else {
		fmt.Printf("El valor %d no se encuentra en el arreglo\n", buscar)
	}
}
