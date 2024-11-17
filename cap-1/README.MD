# Pesquisa binaria

 A pesquisa binária é um algoritmo. Sua entrada é uma lista ordenada de elementos (explicarei mais tarde por que motivo a lista precisa ser ordenada). Se o elemento que você está buscando está na lista, a pesquisa binária retorna a sua localização. Caso contrário, a pesquisa binária retorna None.

Com a pesquisa binária, você chuta um número intermediário e elimina a metade dos números restantes a cada vez.


# EXERCÍCIOS

## 1.1 Número máximo de etapas para lista com 128 nomes

**Pergunta**: Suponha que você tenha uma lista com 128 nomes e esteja fazendo uma pesquisa binária. Qual seria o número máximo de etapas que você levaria para encontrar o nome desejado?

**Solução**:
A pesquisa binária divide a lista ao meio em cada etapa, reduzindo o número de elementos pela metade. O número máximo de etapas para uma lista de n elementos é [log₂(n)].

### Cálculo de log₂(128):
Para resolver log₂(128), perguntamos: "2 elevado a qual potência é igual a 128?"

2ⁿ = 128

Decomposição:
* 2¹ = 2
* 2² = 4
* 2³ = 8
* 2⁴ = 16
* 2⁵ = 32
* 2⁶ = 64
* 2⁷ = 128

Portanto, o número máximo de etapas é **7**.

## 1.2 Número máximo de etapas para lista com 256 nomes

**Pergunta**: Suponha que você duplique o tamanho da lista. Qual seria o número máximo de etapas agora?

**Solução**:
Aplicando a mesma lógica para 256 nomes:
* 2⁸ = 256

Portanto, o número máximo de etapas seria **8**.



## Tempo de execução


