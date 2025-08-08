package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

// Helper para capturar la salida de la consola y poder testearla.
func capturarSalida(f func()) string {
	// Creamos un "tubo" para capturar la salida.
	r, w, _ := os.Pipe()
	originalStdout := os.Stdout
	os.Stdout = w
	
	// Aseguramos que la salida vuelva a la normalidad al final.
	defer func() {
		os.Stdout = originalStdout
		w.Close()
	}()
	
	// Corremos la función que queremos testear.
	f()
	w.Close()
	
	// Leemos todo lo que salió por el tubo.
	out, _ := io.ReadAll(r)
	return string(out)
}

// Helper para crear un archivo de prueba.
func crearArchivoDePrueba(nombreArchivo string, contenido string) error {
    return os.WriteFile(nombreArchivo, []byte(contenido), 0644)
}

// Test de Integración/Unidad: verifica el reporte completo con datos de prueba.
func TestReporteDePedidos(t *testing.T) {
	// Contenido del archivo de prueba
	csvData := `id_pedido,producto,zona,tiempo_entrega_min
101,pizza,centro,25
102,hamburguesa,norte,35
103,pizza,sur,15
104,sushi,centro,30
105,hamburguesa,norte,40
106,pizza,centro,20
107,pizza,centro,10
108,sushi,norte,30`

	// Creamos el archivo temporalmente.
	crearArchivoDePrueba("pedidos.csv", csvData)
	defer os.Remove("pedidos.csv") // Aseguramos que se borre al final.

	// Capturamos la salida de la función main.
	output := capturarSalida(main)

	// Verificamos que la salida contenga los datos correctos.
	if !strings.Contains(output, "Pedidos por zona:") {
		t.Error("La salida no contiene el encabezado 'Pedidos por zona:'")
	}
	if !strings.Contains(output, "Productos más populares:") {
		t.Error("La salida no contiene el encabezado 'Productos más populares:'")
	}
	if !strings.Contains(output, "Tiempo promedio de entrega por zona:") {
		t.Error("La salida no contiene el encabezado 'Tiempo promedio de entrega por zona:'")
	}
	if !strings.Contains(output, "- centro: 4 pedidos") {
		t.Error("Reporte de pedidos por zona incorrecto para 'centro'")
	}
	if !strings.Contains(output, "- norte: 3 pedidos") {
		t.Error("Reporte de pedidos por zona incorrecto para 'norte'")
	}
	if !strings.Contains(output, "- pizza: 4 unidades") {
		t.Error("Reporte de productos populares incorrecto para 'pizza'")
	}
	if !strings.Contains(output, "- norte: 35.00 minutos") {
		t.Error("Reporte de tiempo promedio incorrecto para 'norte'")
	}
	if !strings.Contains(output, "- centro: 21.25 minutos") {
		t.Error("Reporte de tiempo promedio incorrecto para 'centro'")
	}
}

// Prueba E2E: simula la ejecución real del programa.
// Esta prueba usa un archivo de datos diferente para demostrar el E2E.
func TestE2EFlow(t *testing.T) {
	// Creamos un archivo de prueba.
	csvData := `id_pedido,producto,zona,tiempo_entrega_min
101,pizza,ciudad,10
102,hamburguesa,ciudad,15
103,sushi,campo,25`
	crearArchivoDePrueba("pedidos.csv", csvData)
	defer os.Remove("pedidos.csv")

	// Capturamos la salida de la función main.
	output := capturarSalida(main)

	// Verificamos los resultados de la prueba E2E.
	if !strings.Contains(output, "ciudad: 2 pedidos") {
		t.Error("E2E failed: incorrect count for 'ciudad'")
	}
	if !strings.Contains(output, "campo: 1 pedidos") {
		t.Error("E2E failed: incorrect count for 'campo'")
	}
	if !strings.Contains(output, "ciudad: 12.50 minutos") {
		t.Error("E2E failed: incorrect average time for 'ciudad'")
	}
}

// Prueba de Estrés (Benchmark): mide el rendimiento del procesamiento de datos.
func BenchmarkProcesarDatos(b *testing.B) {
	// Simulamos un archivo CSV muy grande.
	records := [][]string{
		{"id_pedido", "producto", "zona", "tiempo_entrega_min"},
	}
	for i := 0; i < 10000; i++ {
		records = append(records, []string{
			fmt.Sprintf("%d", i),
			"producto",
			"zona",
			"20",
		})
	}

	// Creamos un archivo de prueba grande.
	archivo, _ := os.Create("pedidos_benchmark.csv")
	writer := csv.NewWriter(archivo)
	writer.WriteAll(records)
	writer.Flush()
	archivo.Close()
	defer os.Remove("pedidos_benchmark.csv")

	// El benchmark se ejecutará b.N veces para medir el rendimiento.
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Aquí está la lógica de procesamiento que queremos medir.
		archivo, _ := os.Open("pedidos_benchmark.csv")
		lector := csv.NewReader(archivo)
		registros, _ := lector.ReadAll()
		archivo.Close()

		pedidosPorZona := make(map[string]int)
		pedidosPorProducto := make(map[string]int)
		tiemposDeEntrega := make(map[string][]int)

		for j, registro := range registros {
			if j == 0 {
				continue
			}
			producto := registro[1]
			zona := registro[2]
			tiempoStr := registro[3]

			tiempo, _ := strconv.Atoi(tiempoStr)

			pedidosPorZona[zona]++
			pedidosPorProducto[producto]++
			tiemposDeEntrega[zona] = append(tiemposDeEntrega[zona], tiempo)
		}
	}
}
