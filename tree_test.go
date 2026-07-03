package main

import "testing"

// Verifica que el árbol se construya correctamente
func TestBuild(t *testing.T) {
	data := []float64{10, 20, 30, 40, 50}
	tree := NewSegmentTree(data)

	if tree.Tree[1].Sum != 150 {
		t.Errorf("Sum esperado 150, obtenido %.2f", tree.Tree[1].Sum)
	}
	if tree.Tree[1].Min != 10 {
		t.Errorf("Min esperado 10, obtenido %.2f", tree.Tree[1].Min)
	}
	if tree.Tree[1].Max != 50 {
		t.Errorf("Max esperado 50, obtenido %.2f", tree.Tree[1].Max)
	}
}

// Verifica la consulta de suma sobre un rango.
func TestQuerySum(t *testing.T) {
	data := []float64{10, 20, 30, 40, 50}
	tree := NewSegmentTree(data)
	result := tree.QuerySum(1, 3)

	if result != 90 {
		t.Errorf("Esperado 90, obtenido %.2f", result)
	}
}

// Verifica la consulta del valor mínimo en un rango.
func TestQueryMin(t *testing.T) {
	data := []float64{10, 20, 30, 40, 50}
	tree := NewSegmentTree(data)
	result := tree.QueryMin(1, 3)

	if result != 20 {
		t.Errorf("Esperado 20, obtenido %.2f", result)
	}
}

// Verifica la consulta del valor máximo en un rango.
func TestQueryMax(t *testing.T) {
	data := []float64{10, 20, 30, 40, 50}
	tree := NewSegmentTree(data)
	result := tree.QueryMax(1, 3)

	if result != 40 {
		t.Errorf("Esperado 40, obtenido %.2f", result)
	}
}

// Verifica que una actualización puntual modifique correctamente el árbol.
func TestUpdate(t *testing.T) {
	data := []float64{10, 20, 30, 40, 50}
	tree := NewSegmentTree(data)
	tree.Update(2, 100)

	if tree.QuerySum(1, 3) != 160 {
		t.Errorf("Esperado 160, obtenido %.2f", tree.QuerySum(1, 3))
	}
	if tree.QueryMax(1, 3) != 100 {
		t.Errorf("Esperado 100, obtenido %.2f", tree.QueryMax(1, 3))
	}
}

// Verifica el funcionamiento de Lazy Propagation mediante una actualización sobre un rango.
func TestRangeUpdate(t *testing.T) {
	data := []float64{10, 20, 30, 40, 50}
	tree := NewSegmentTree(data)
	tree.RangeUpdate(1, 4, 10)

	if tree.QuerySum(1, 4) != 180 {
		t.Errorf("Esperado 180, obtenido %.2f", tree.QuerySum(1, 4))
	}
	if tree.QueryMin(1, 4) != 30 {
		t.Errorf("Esperado 30, obtenido %.2f", tree.QueryMin(1, 4))
	}
	if tree.QueryMax(1, 4) != 60 {
		t.Errorf("Esperado 60, obtenido %.2f", tree.QueryMax(1, 4))
	}
}
