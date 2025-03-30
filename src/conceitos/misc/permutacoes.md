# Permutações e Trocas

Em cifras simétricas, a dispersão de dependências locais — fundamental para a difusão — pode ser alcançada por meio de operações lineares, como multiplicações matriciais, ou estruturais, como permutações e trocas de posição. Este tipo de transformação reorganiza os dados sem alterar seu conteúdo individual, mas com profundo impacto estatístico na propagação de alterações pela estrutura do algoritmo.

## Permutações como mecanismo de difusão

Uma permutação é uma bijeção sobre o conjunto de posições de bits ou bytes. Aplicada a um vetor de entrada, ela redistribui os elementos segundo uma ordem fixa ou pseudoaleatória. Embora não introduza não-linearidade, a permutação é essencial para ampliar o alcance das alterações provocadas pelas S-Boxes ou outras funções não lineares.

Na arquitetura de cifras modernas, as permutações são utilizadas para:

* Propagar a influência de bits da entrada por várias S-Boxes na rodada seguinte.

* Impedir que padrões locais na entrada gerem padrões locais na saída.

* Compor difusão com custo computacional baixo.

## Exemplo em SPNs

No AES (FIPS-197), duas camadas implementam difusão com base em permutação:

* **ShiftRows**: uma permutação de bytes por rotação de linhas.
* **MixColumns**: uma multiplicação linear que atua em colunas, mas combinada com ShiftRows produz difusão bidimensional.

Essas operações não introduzem confusão, mas são vitais para garantir o efeito avalanche ao longo das rodadas.

## Exemplo em Feistel

Na estrutura de Feistel, usada no DES e em muitas cifras derivadas, a difusão é construída ao longo das rodadas, intercalando trocas de metades do bloco com aplicações de funções não lineares (geralmente baseadas em S-Boxes).

Permutações também aparecem de forma explícita:

* Initial Permutation (IP) e Final Permutation (FP) no DES são mapeamentos fixos de 64 bits, historicamente motivados por eficiência em hardware, mas com efeito colateral de difusão inicial e final.

* Permutação P após a aplicação das S-Boxes na função F do DES serve para espalhar os resultados locais das substituições.

## Implementação ilustrativa em Go

Abaixo, um exemplo de permutação simples de bytes em um vetor de 16 posições:

```go
package main

import "fmt"

// Permutação fixa (exemplo artificial)
var permutation = [16]int{2, 0, 3, 1, 6, 4, 7, 5, 10, 8, 11, 9, 14, 12, 15, 13}

func permute(input []byte, p [16]int) []byte {
	out := make([]byte, len(input))
	for i := range p {
		out[i] = input[p[i]]
	}
	return out
}

func main() {
	entrada := []byte("abcdefghijklmnop")
	saida := permute(entrada, permutation)
	fmt.Printf("Original : %s\n", entrada)
	fmt.Printf("Permutada: %s\n", saida)
}
```

Esse tipo de permutação, apesar de simples, é o mecanismo base por trás de transformações como ShiftRows. A escolha da ordem é crítica: ela deve maximizar a dispersão estatística sem introduzir simetrias.

## Considerações Finais

Permutações e trocas estruturadas são ferramentas indispensáveis na construção de cifras robustas. Embora não introduzam entropia ou não-linearidade, são essenciais para distribuir os efeitos locais das S-Boxes e garantir que os princípios de difusão definidos por Shannon sejam plenamente realizados.

Além disso, entender como essas operações interagem com as camadas de substituição permite projetar SPNs eficientes e compreender como redes de Feistel obtêm segurança mesmo com funções internas relativamente simples.
