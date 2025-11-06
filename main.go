package main

import funcauxiliadoras "ListaDinamicaGO/funcAuxiliadoras"

func main() {
	listaNOVA := funcauxiliadoras.CriarArvore()

	no1 := funcauxiliadoras.CriarNO(3)
	no2 := funcauxiliadoras.CriarNO(10)
	no3 := funcauxiliadoras.CriarNO(15)

	funcauxiliadoras.InserirLista(listaNOVA, no1)
	funcauxiliadoras.InserirLista(listaNOVA, no2)
	funcauxiliadoras.InserirLista(listaNOVA, no3)

	funcauxiliadoras.PercorrerLista(listaNOVA)
}
