<script setup>
import { ref } from 'vue'

// URL base de la API en Go
const API = 'http://localhost:8080/api'

// Valores de los inputs del usuario
const l = ref(0)
const r = ref(0)
const valor = ref(0)

// Resultado que se muestra en pantalla
const resultado = ref(0)
const mensaje = ref('')

// Guardamos los nodos que el backend dice que visitó
// para poder dibujarlos y mostrarlos paso a paso
const nodos = ref([])
const pasoActual = ref(0)
const hayAnimacion = ref(false)

// Pide al backend la suma de un rango
async function consultar() {
  const url = API + '/consultar?l=' + l.value + '&r=' + r.value
  const res = await fetch(url)
  const data = await res.json()

  resultado.value = data.suma
  mensaje.value = 'Consulta: suma en el rango [' + l.value + ', ' + r.value + ']'

  // Si el backend nos manda los nodos recorridos, armamos la animacion
  if (data.nodosVisitados) {
    nodos.value = data.nodosVisitados
    pasoActual.value = 0
    hayAnimacion.value = true
  }
}

// Pide al backend que sume "valor" a todo el rango [l, r]
async function actualizar() {
  const url = API + '/actualizar?l=' + l.value + '&r=' + r.value + '&valor=' + valor.value
  const res = await fetch(url)
  const data = await res.json()

  mensaje.value = 'Actualizacion: se sumo ' + valor.value + ' al rango [' + l.value + ', ' + r.value + ']'

  if (data.nodosVisitados) {
    nodos.value = data.nodosVisitados
    pasoActual.value = 0
    hayAnimacion.value = true
  }

  // despues de actualizar, consultamos de nuevo para ver el cambio
  consultar()
}

// Avanza un paso en la animacion (muestra un nodo mas)
function siguientePaso() {
  if (pasoActual.value < nodos.value.length) {
    pasoActual.value = pasoActual.value + 1
  }
}

// Vuelve a empezar la animacion desde cero
function reiniciarAnimacion() {
  pasoActual.value = 0
}

// Le dice al template si un nodo ya se debe ver o todavia no
function seVeElNodo(indice) {
  return indice < pasoActual.value
}
</script>

<template>
  <div class="caja">
    <h1>Demo Segment Tree (Lazy Propagation)</h1>

    <div class="panel">
      <label>L: <input v-model.number="l" type="number" ></label>
      <label>R: <input v-model.number="r" type="number" ></label>
      <label>Valor a sumar: <input v-model.number="valor" type="number" ></label>
    </div>

    <div class="botones">
      <button @click="consultar">Consultar Rango</button>
      <button @click="actualizar" class="verde">Actualizar Rango (Lazy)</button>
    </div>

    <h3>{{ mensaje }}</h3>
    <h2 class="resultado">Resultado: {{ resultado }}</h2>

    <!-- Esta parte solo aparece despues de hacer una consulta o actualizacion -->
    <div v-if="hayAnimacion" class="zona-arbol">
      <h3>Nodos recorridos (paso {{ pasoActual }} de {{ nodos.length }})</h3>

      <div class="botones">
        <button @click="siguientePaso">Siguiente paso</button>
        <button @click="reiniciarAnimacion">Reiniciar</button>
      </div>

      <div class="fila">
        <div
          v-for="(nodo, indice) in nodos"
          :key="indice"
          class="nodo"
          :class="{ activo: seVeElNodo(indice), lazy: nodo.esLazy }"
        >
          <p>[{{ nodo.rango[0] }}, {{ nodo.rango[1] }}]</p>
          <p v-if="nodo.esLazy">pendiente: {{ nodo.valorLazy }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.caja {
  font-family: sans-serif;
  max-width: 700px;
  margin: 40px auto;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 10px;
}

.panel {
  border: 2px solid #ffffff;  
  background-color: #00000000;    
  border-radius: 10px;         
  padding: 15px;                
  max-width: 400px;   
  margin-bottom: 15px;                         
}

.panel input {
  width: 50px;
  margin-left: 5px;
  margin-right: 15px;
}

.botones {
  display: flex;
  gap: 10px;
  margin-bottom: 10px;
}

button {
  padding: 10px;
  cursor: pointer;
}

.verde {
  background: #4CAF50;
  color: white;
  border: none;
}

.resultado {
  color: blue;
}

.zona-arbol {
  margin-top: 20px;
  padding: 15px;
  background: #ffffff;
  border: 2px solid #999;
  border-radius: 10px;
  color: #222;
}

.fila {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}

.nodo {
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 6px;
  opacity: 0.2;
  color: #222;
}

.nodo.activo {
  opacity: 1;
  background: #e3f2fd;
}

.nodo.activo.lazy {
  background: #fff3cd;
  border-color: #ffc107;
}
</style>