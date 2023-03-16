# Guía: Análisis de Algoritmos

## Mediciones

En la carpeta [`mediciones/busquedas`](./mediciones/busquedas) se encuentran algunos algoritmos de búsqueda conocidos.

En [`mediciones/main.go`](./mediciones/main.go), se encuentra un ejemplo de como medir el tiempo que demora la ejecución de una función.

Se pide:

1. Tomar mediciones para $N = 1.000, 10.000, \dots, 10.000.000$ de los algoritmos de búsquedas.

   - Guardar los resultados en una plantilla de cálculo y graficar ambas tablas en un mismo gráfico.
   - Verificar si las curvas obtenidas se aproximan a las curvas teóricas. **El o los archivos deben estar ubicados dentro de la carpeta ya creada `mediciones/resultados`**

2. Implementar el algoritmo de ordenamiento de burbujeo, escribir tests y tomar las mismas mediciones que el punto anterior.

## Subsecuencia de Suma Máxima

El problema de subsecuencia de suma máxima consiste en encontrar una secuencia (en posiciones consecutivas) cuya suma sea máxima dentro de un arrego original.

Por ejemplo: en el arreglo `[-1, 6, -2, 5, -1, 4, 3, -4, 3]` la subsecuencia de suma máxima es `[6, -2, 5, -1, 4, 3]`, cuya suma es `15`.

1. Analizar el orden de la función `SubsecuenciaSumaMaxima`.
2. Encontrar otra solución que sea $O(n)$

## El mito de la máquina superpoderosa

1. Mirar el siguiente video, analizar y sacar tus propias conclusiones
   - [El Problema del Viajante de Comercio](https://www.youtube.com/watch?v=oSPkod-M6Gc)
2. Clonar la siguiente
   - [Planilla de cálculo](https://docs.google.com/spreadsheets/d/1i1jjbYdJ63bpxtnJy2ZiMYXDjYz4bIOcJqGDrb1qE1g/copy)

Completar los cálculos, probar con distintas velocidades de CPU y obtener tus propias conclusiones

## Análisis de la complejidad de fragmentos de código

En la carpeta [`fragmentos/utiles`](./fragmentos/utiles) se encuentran fragmentos de código que implementan distintas operaciones sobre arreglos y matrices

1. Analizar el orden de las funciones `InvertirArreglo` e `InvertirArreglo2`. ¿Cuál de las dos es más eficiente? ¿Por qué?

2. Analizar el orden de la función `BusquedaTernaria` y comparla con la función `BusquedaBinaria`. ¿Cuál de las dos es más eficiente? ¿Por qué?

3. Analizar el orden de la llamada `matrizProducto, err := utiles.ProductoMatrices(matrizA, matrizB)` de la línea 31 de [`fragmentos/main.go`](./fragmentos/main.go) si la `matrizA` es de NxM y la `matrizB` es de MxP.

---

## Como trabajar con la guía

Una vez clonado el repositorio primero debemos asegurarnos de que las dependencias de la guía están instaladas.

```bash
go mod tidy
```

Para asegurarnos de que nuestro código no tiene errores de compilación ejecutamos:

```bash
go build ./...
```

Para ejecutar todos los tests de la guía ejecutamos:

```bash
go test ./...
```

Si queremos ejecutar un test en particular, podemos hacerlo de la siguiente manera:

```bash
go test ./... -run ^TestResultados$
```

donde `TestResultados` es el nombre del test.
