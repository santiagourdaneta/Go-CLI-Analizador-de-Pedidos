📦 Go CLI Analizador de Pedidos
Una herramienta de línea de comandos (CLI) escrita en Go para procesar y analizar archivos de pedidos en formato CSV. Este programa lee un archivo pedidos.csv y genera un reporte completo con información valiosa sobre los pedidos.

🚀 Características
Análisis por zona: Cuenta el número total de pedidos por cada zona de entrega.

Productos populares: Identifica la cantidad de unidades vendidas de cada producto.

Tiempo de entrega promedio: Calcula el tiempo de entrega promedio para cada zona.

🛠️ Cómo Empezar
Estos pasos te guiarán para tener una copia de tu proyecto en funcionamiento en tu máquina local.

Prerrequisitos
Necesitas tener Go instalado en tu sistema. Puedes descargarlo desde la página oficial de Go.

Instalación
Clona el repositorio en tu máquina local:

git clone https://github.com/santiagourdaneta/Go-CLI-Analizador-de-Pedidos/ 
cd Go-CLI-Analizador-de-Pedidos

Asegúrate de tener un archivo llamado pedidos.csv en la misma carpeta que main.go. El archivo debe tener la siguiente estructura:

id_pedido,producto,zona,tiempo_entrega_min
101,pizza,centro,25
102,hamburguesa,norte,35
...

Uso
Para ejecutar el analizador de pedidos, simplemente usa el siguiente comando en tu terminal:

go run main.go

El programa imprimirá un reporte detallado directamente en la consola.

🧪 Ejecutando las Pruebas
El proyecto incluye pruebas para asegurar que todas las funciones de análisis de datos trabajan correctamente.

Pruebas de unidad e integración
Para ejecutar las pruebas unitarias y de integración, usa:

go test

Pruebas de estrés (Benchmark)
Para medir el rendimiento de la función de procesamiento de datos, usa:

go test -bench=.

🤝 Contribuciones
Las contribuciones son bienvenidas. Si encuentras un error o tienes una sugerencia de mejora, por favor, abre un "Issue" o envía un "Pull Request".

📄 Licencia
Este proyecto está bajo la Licencia MIT. Para más detalles, consulta el archivo LICENSE.md.

Labels (Etiquetas): go, golang, cli, data-analysis, csv, tool, report

Tags (Temas): go, cli, data-analysis, csv-parser

Hashtags: #golang #Go #CLI #DataAnalysis #CSV #OpenSource

Este video te puede dar más ideas sobre el análisis de datos con Go:
https://www.youtube.com/watch?v=h62mSkCjaEM
