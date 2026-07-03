package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func main() {
	// Conectar a la base de datos
	db, err := sql.Open("sqlite", "./VentasTambo.db")
	if err != nil {
		log.Fatalf("Error al abrir la base de datos: %v", err)
	}
	defer db.Close()

	// Extraer datos
	query := "SELECT monto_recaudado FROM ventas_tambo"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Error al ejecutar la consulta: %v", err)
	}
	defer rows.Close()

	var datosTambo []float64
	for rows.Next() {
		var monto float64
		if err := rows.Scan(&monto); err != nil {
			log.Fatalf("Error al leer la fila: %v", err)
		}
		datosTambo = append(datosTambo, monto)
	}

	// Primero validamos que tengamos datos
	if len(datosTambo) == 0 {
		fmt.Println("No se encontraron registros.")
		return
	}

	tree := NewSegmentTree(datosTambo)

	fmt.Printf("¡Éxito! Árbol construido con %d registros.\n", len(datosTambo))

	inicio := 0
	fin := tree.N - 1
	fmt.Printf("Suma total: S/ %.2f\n", tree.QuerySum(inicio, fin))
	fmt.Printf("Mínimo: S/ %.2f\n", tree.QueryMin(inicio, fin))
	fmt.Printf("Máximo: S/ %.2f\n", tree.QueryMax(inicio, fin))

	fmt.Println("\n--- Aplicando ajuste masivo (Lazy Propagation) ---")

	tree.RangeUpdate(50, 100, 10.0)

	fmt.Println("Ajuste aplicado correctamente.")
	fmt.Printf("Nueva suma del rango [50,100]: S/ %.2f\n", tree.QuerySum(50, 100))
	fmt.Printf("Nueva suma total de la BD: S/ %.2f\n", tree.QuerySum(0, tree.N-1))
}
