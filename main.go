package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "modernc.org/sqlite"
)

var miArbol *SegmentTree

func main() {

	var datos []float64

	db, err := sql.Open("sqlite", "./VentasTambo.db")
	if err != nil {
		log.Fatal("Error al abrir BD:", err)
	}
	defer db.Close() //cerrar la conexion justo antes que termine la func, aunque termine en error

	aux, err := db.Query("SELECT monto_recaudado FROM ventas_tambo")
	if err != nil {
		log.Fatal("Error al consultar la tabla:", err)
	}

	for aux.Next() {
		var val float64
		aux.Scan(&val)
		datos = append(datos, val)
	}
	aux.Close()

	miArbol = NewSegmentTree(datos)
	fmt.Println("Datos cargados de la BD y Árbol construido.")

	// API: Consultar suma en un rango
	http.HandleFunc("/api/consultar", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		l, _ := strconv.Atoi(r.URL.Query().Get("l"))
		r_idx, _ := strconv.Atoi(r.URL.Query().Get("r"))

		suma, pasos := miArbol.QuerySumConPasos(l, r_idx)

		respuesta := map[string]interface{}{
			"suma":           suma,
			"nodosVisitados": pasos,
		}
		json.NewEncoder(w).Encode(respuesta)
	})

	// API: Actualizar un rango (Lazy Propagation)
	http.HandleFunc("/api/actualizar", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		l, _ := strconv.Atoi(r.URL.Query().Get("l"))
		r_idx, _ := strconv.Atoi(r.URL.Query().Get("r"))
		val, _ := strconv.ParseFloat(r.URL.Query().Get("valor"), 64)

		pasos := miArbol.RangeUpdateConPasos(l, r_idx, val)

		respuesta := map[string]interface{}{
			"mensaje":        "OK",
			"nodosVisitados": pasos,
		}
		json.NewEncoder(w).Encode(respuesta)
	})

	fmt.Println("\n--- Iniciando servidor para el frontend ---")
	fmt.Println("Escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
