package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// main es la función principal de nuestro robot.
func main() {
	// 1. Decirle al robot dónde están los papeles.
	fmt.Println("🤖 ¡Hola! Soy el robot de análisis de pedidos.")
	fmt.Println("📦 Leyendo el archivo de pedidos...")

	// Aquí abrimos el archivo de papeles (CSV).
	archivo, err := os.Open("pedidos.csv")
	if err != nil {
		fmt.Println("¡Ups! No pude encontrar el archivo pedidos.csv.")
		return
	}
	// Cuando terminemos, cerramos el archivo.
	defer archivo.Close()

	// 2. Leer todos los papeles.
	lector := csv.NewReader(archivo)
	registros, err := lector.ReadAll()
	if err != nil {
		fmt.Println("¡Ups! No pude leer los datos del archivo.")
		return
	}

	// 3. Contar y organizar los datos.
	// Estas son las "cajas" donde el robot guardará la información.
	pedidosPorZona := make(map[string]int)
	pedidosPorProducto := make(map[string]int)
	tiemposDeEntrega := make(map[string][]int)

	// Recorrer cada papel (registro) en la pila.
	for i, registro := range registros {
		if i == 0 {
			// El primer papel es solo el título, no lo usamos.
			continue
		}

		// Tomar la información de cada papel.
		producto := registro[1]
		zona := registro[2]
		tiempoStr := registro[3]

		// Convertir el tiempo de texto a número.
		tiempo, _ := strconv.Atoi(tiempoStr)

		// Guardar la información en las cajas.
		pedidosPorZona[zona]++
		pedidosPorProducto[producto]++
		tiemposDeEntrega[zona] = append(tiemposDeEntrega[zona], tiempo)
	}

	// 4. Decirle a la computadora qué mostrar en la pantalla.
	fmt.Println("\n--- 📊 Reporte de Pedidos ---")

	fmt.Println("\nPedidos por zona:")
	for zona, cantidad := range pedidosPorZona {
		fmt.Printf("- %s: %d pedidos\n", zona, cantidad)
	}

	fmt.Println("\nProductos más populares:")
	for producto, cantidad := range pedidosPorProducto {
		fmt.Printf("- %s: %d unidades\n", producto, cantidad)
	}

	fmt.Println("\nTiempo promedio de entrega por zona:")
	for zona, tiempos := range tiemposDeEntrega {
		totalTiempo := 0
		for _, t := range tiempos {
			totalTiempo += t
		}
		promedio := float64(totalTiempo) / float64(len(tiempos))
		fmt.Printf("- %s: %.2f minutos\n", zona, promedio)
	}

	fmt.Println("\n¡Análisis completo!")
}