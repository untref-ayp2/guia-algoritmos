package utiles

import "fmt"

func InvertirArreglo(arreglo []int) []int {
	n := len(arreglo)
	arregloInvertido := make([]int, n)
	for i := 0; i < n; i++ {
		arregloInvertido[i] = arreglo[n-i-1]
	}
	return arregloInvertido
}

func InvertirArreglo2(arreglo []int) {
	n := len(arreglo)
	for i := range n / 2 {
		arreglo[i], arreglo[n-i-1] = arreglo[n-i-1], arreglo[i]
	}
}

func ProductoMatrices(matrizA, matrizB [][]int) ([][]int, error) {
	if len(matrizA[0]) != len(matrizB) {
		return nil, fmt.Errorf("las matrices no se pueden multiplicar")
	}

	filasA := len(matrizA)
	columnasA := len(matrizA[0])
	columnasB := len(matrizB[0])

	matrizProducto := make([][]int, filasA)
	for i := range matrizProducto {
		matrizProducto[i] = make([]int, columnasB)
	}

	for i := 0; i < filasA; i++ {
		for j := 0; j < columnasB; j++ {
			for k := 0; k < columnasA; k++ {
				matrizProducto[i][j] += matrizA[i][k] * matrizB[k][j]
			}
		}
	}

	return matrizProducto, nil
}

func BusquedaTernaria(arreglo []int, valor int) int {
	izquierda := 0
	derecha := len(arreglo) - 1

	for izquierda <= derecha {
		tercio1 := izquierda + (derecha-izquierda)/3
		tercio2 := derecha - (derecha-izquierda)/3

		if arreglo[tercio1] == valor {
			return tercio1
		}
		if arreglo[tercio2] == valor {
			return tercio2
		}

		if valor < arreglo[tercio1] {
			derecha = tercio1 - 1
		} else if valor > arreglo[tercio2] {
			izquierda = tercio2 + 1
		} else {
			izquierda = tercio1 + 1
			derecha = tercio2 - 1
		}
	}

	return -1
}
