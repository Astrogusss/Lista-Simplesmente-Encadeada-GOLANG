package funcauxiliadoras

import "fmt"

type lista struct {
	cabeca     *no
	quantidade int8
}

type no struct {
	prox *no
	dado int8
}

func CriarArvore() *lista {
	novo := &lista{nil, 0}
	return novo
}

func CriarNO(dado int8) *no {
	novo := &no{nil, dado}
	return novo
}

func InserirLista(l *lista, nNOVO *no) int8 {
	if l == nil || nNOVO == nil {
		//se der merda, entra nesse if para acabar
		return -1
	}
	//se a lista estiver vazia
	if l.quantidade == 0 || l.cabeca == nil {
		l.cabeca = nNOVO
	} else {
		//caso nao esteja vazia
		//declare uma variavel do tipo *no para receber o endereco de memoria da lista.cabeca
		var aux *no
		aux = l.cabeca
		//enquanto o proximo for diferente de nil (NULL), pode ir para frente
		for aux.prox != nil {
			aux = aux.prox
		}
		//atribui o prox de aux para o novoNO
		aux.prox = nNOVO
		nNOVO.prox = nil
	}
	l.quantidade++
	return 1
}

func RemoveLista(l *lista, dado int8) int8 {
	if l == nil {
		return -1
	}
	var aux *no
	aux = l.cabeca
	//vai parar ate achar o dado do proximo
	for aux.prox.dado != dado {
		aux = aux.prox
	}

	//aqui paramos antes de achar o no que temos que remover
	var remover *no
	remover = aux.prox
	aux.prox = remover.prox
	l.quantidade--
	return 1
}
func PercorrerLista(l *lista) {
	var aux *no = l.cabeca
	for aux != nil {
		fmt.Println(aux.dado)
		aux = aux.prox
	}
}
