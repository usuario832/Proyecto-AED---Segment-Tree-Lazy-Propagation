package main

// Un "paso" representa un nodo del arbol que fue visitado
// durante una consulta o una actualizacion. Esto es lo que
// el frontend en Vue va a usar para dibujar la animacion.
type PasoVisitado struct {
	Rango     [2]int  `json:"rango"`
	EsLazy    bool    `json:"esLazy"`
	ValorLazy float64 `json:"valorLazy"`
}

// Hace lo mismo que QuerySum, pero ademas va guardando
// en una lista cada nodo por el que pasa
func (st *SegmentTree) QuerySumConPasos(l, r int) (float64, []PasoVisitado) {
	var pasos []PasoVisitado
	suma := st.querySumPasos(1, 0, st.N-1, l, r, &pasos)
	return suma, pasos
}

func (st *SegmentTree) querySumPasos(node, left, right, qLeft, qRight int, pasos *[]PasoVisitado) float64 {
	if qRight < left || qLeft > right {
		return 0
	}

	paso := PasoVisitado{
		Rango:     [2]int{left, right},
		EsLazy:    st.Tree[node].Lazy != 0,
		ValorLazy: st.Tree[node].Lazy,
	}
	*pasos = append(*pasos, paso)

	if qLeft <= left && right <= qRight {
		return st.Tree[node].Sum
	}

	st.push(node, left, right)
	mid := (left + right) / 2
	sumaIzq := st.querySumPasos(node*2, left, mid, qLeft, qRight, pasos)
	sumaDer := st.querySumPasos(node*2+1, mid+1, right, qLeft, qRight, pasos)
	return sumaIzq + sumaDer
}

// Hace lo mismo que RangeUpdate, pero ademas va guardando
// en una lista cada nodo por el que pasa
func (st *SegmentTree) RangeUpdateConPasos(l, r int, delta float64) []PasoVisitado {
	var pasos []PasoVisitado
	st.rangeUpdatePasos(1, 0, st.N-1, l, r, delta, &pasos)
	return pasos
}

func (st *SegmentTree) rangeUpdatePasos(node, left, right, qLeft, qRight int, delta float64, pasos *[]PasoVisitado) {
	if qRight < left || qLeft > right {
		return
	}

	// si el rango del nodo entra completo dentro del rango pedido,
	// aca es exactamente donde se aplica el "lazy": no seguimos
	// bajando a los hijos, solo marcamos este nodo como pendiente
	entraCompleto := qLeft <= left && right <= qRight

	if entraCompleto {
		st.apply(node, left, right, delta)
		*pasos = append(*pasos, PasoVisitado{
			Rango:     [2]int{left, right},
			EsLazy:    true,
			ValorLazy: st.Tree[node].Lazy,
		})
		return
	}

	*pasos = append(*pasos, PasoVisitado{
		Rango:  [2]int{left, right},
		EsLazy: false,
	})

	st.push(node, left, right)
	mid := (left + right) / 2
	st.rangeUpdatePasos(node*2, left, mid, qLeft, qRight, delta, pasos)
	st.rangeUpdatePasos(node*2+1, mid+1, right, qLeft, qRight, delta, pasos)

	leftNode := st.Tree[node*2]
	rightNode := st.Tree[node*2+1]
	st.Tree[node].Sum = leftNode.Sum + rightNode.Sum
	st.Tree[node].Min = min(leftNode.Min, rightNode.Min)
	st.Tree[node].Max = max(leftNode.Max, rightNode.Max)
}
