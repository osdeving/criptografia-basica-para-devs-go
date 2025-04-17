/*
	Este arquivo é um "artigo interativo" em forma de código Go, focado em aritmética sobre campos finitos,
	em especial o campo com 256 elementos, construído sobre o corpo F₂, notado formalmente como F_{2^8}.

	Ele aborda conceitos fundamentais da álgebra moderna que aparecem na prática da computação, especialmente
	em criptografia, correção de erros e aritmética binária avançada.

	Tópicos cobertos:
	- Aritmética modular: operações realizadas "com resto" após divisão por um número fixo.
	  Ex: 7 mod 5 = 2. Em F₂, tudo é feito mod 2: 1 + 1 = 0.

	- Corpo (field): estrutura algébrica com adição, multiplicação, identidade e inverso para ambas (exceto o zero).
	  Ex: F₂ = {0, 1}. Já F_{2^8} é um corpo estendido construído com polinômios sobre F₂, reduzidos por um polinômio irredutível.

	- Anel (ring): estrutura parecida com campo, mas sem inverso multiplicativo obrigatório.

	- Polinômios sobre F₂[x]: usados para representar os elementos do campo F_{2^8}. Cada byte (8 bits) é um polinômio
	  de grau até 7 com coeficientes 0 ou 1. Ex: 0x57 = x^6 + x^4 + x^2 + x + 1

	- Redução modular com polinômios: usada para "fechar" o conjunto de polinômios em um campo finito
	  através da operação de módulo com um polinômio irredutível. No AES, esse polinômio é x⁸ + x⁴ + x³ + x + 1 (0x11B).

	- Representação binária de polinômios: útil para simular toda essa aritmética usando operadores bit a bit (XOR, SHIFT).

	Este programa também mostra a aplicação prática desses conceitos, por exemplo na operação MixColumns do AES,
	que depende da multiplicação matricial sobre F_{2^8}.
*/

// galois_field_polynomials.go
// ------------------------------------------------------------
// Um "artigo interativo" em forma de código Go
// Este programa explora a estrutura do campo finito F_{2^8},
// mostrando como bytes podem ser interpretados como polinômios
// binários sobre F_2[x], e como fazer operações como soma,
// multiplicação, redução e construção de tabelas completas de
// elementos do campo, incluindo inversos, logs e anti-logs.
// Esta versão também implementa multiplicação vetorial no estilo
// MixColumns do AES, para mostrar aplicações práticas.
// ------------------------------------------------------------
package main

import (
	"fmt"
	"strings"
)

// formatPolynomial imprime um polinômio no formato tradicional
func formatPolynomial(poly uint16) string {
	if poly == 0 {
		return "0"
	}
	var terms []string
	for i := 15; i >= 0; i-- {
		if poly&(1<<i) != 0 {
			switch i {
			case 0:
				terms = append(terms, "1")
			case 1:
				terms = append(terms, "x")
			default:
				terms = append(terms, fmt.Sprintf("x^%d", i))
			}
		}
	}
	return strings.Join(terms, " + ")
}

func degree(p uint16) int {
	for i := 15; i >= 0; i-- {
		if p&(1<<i) != 0 {
			return i
		}
	}
	return -1
}

/*
	polyAdd realiza a soma de dois polinômios sobre F₂, ou seja, polinômios com coeficientes binários.

	No campo F₂, os coeficientes possíveis são apenas 0 e 1, e a operação de soma é definida como:
	  0 + 0 = 0
	  0 + 1 = 1
	  1 + 0 = 1
	  1 + 1 = 0 (porque 2 ≡ 0 mod 2)

	Ou seja, a soma em F₂ corresponde exatamente à operação lógica XOR.

	Neste código, cada polinômio é representado como um byte, onde cada bit indica a presença (1) ou ausência (0) do termo x^i.
	Exemplo:
	  0b1011 representa o polinômio x³ + x + 1
	  0b0110 representa o polinômio x² + x

	A soma dos dois:
	  0b1011 ^ 0b0110 = 0b1101 → x³ + x² + 1

	Portanto, esta função faz a soma bit a bit usando XOR.
*/
func polyAdd(a, b uint8) uint8 {
	return a ^ b
}

/*
	polyMultiply realiza a multiplicação de dois polinômios com coeficientes binários sobre F₂[x].

	Contexto matemático:
	Na álgebra polinomial, multiplicar dois polinômios consiste em aplicar a distributiva entre todos os termos.
	Como estamos em F₂, a operação de soma de coeficientes é XOR, e a multiplicação continua sendo 1*1 = 1, com 0 anulando o termo.

	Dado dois polinômios:
	  A(x) = a₀ + a₁x + a₂x² + ...
	  B(x) = b₀ + b₁x + b₂x² + ...
	O produto é:
	  C(x) = Σₖ (Σᵢ aᵢ * bₖ₋ᵢ) xᵏ

	Como estamos representando os polinômios como bytes, cada bit indica a presença ou ausência de um termo x^i.
	Para cada bit 1 no polinômio B (o segundo fator), fazemos o shift de A para a esquerda (multiplicação por x^i),
	e acumulamos esse termo no resultado com XOR.

	Exemplo:
	  A(x) = 0b1011 → x³ + x + 1
	  B(x) = 0b0110 → x² + x

	  Multiplicação:
	    (x³ + x + 1) * x  = x⁴ + x² + x
	    (x³ + x + 1) * x² = x⁵ + x³ + x²
	    Resultado final = (x⁴ + x² + x) ⊕ (x⁵ + x³ + x²)
	                      = x⁵ + x⁴ + x³ + x

	A função abaixo implementa exatamente isso, usando XOR para acumular os termos
	e shift para multiplicar os termos por potências de x.
*/
func polyMultiply(a, b uint8) uint16 {
	var result uint16
	for i := 0; i < 8; i++ {
		if b&(1<<i) != 0 {
			result ^= uint16(a) << i
		}
	}
	return result
}

/*
	polyMod realiza a operação de redução de um polinômio binário módulo um polinômio irredutível.

	Essa operação é fundamental para "fechar" o conjunto de polinômios sobre F₂[x] em um corpo finito,
	isto é, garantir que após multiplicações o resultado permaneça dentro de um conjunto limitado —
	especificamente, o campo F_{2^8}.

	Matematicamente, ao realizar uma multiplicação A(x) * B(x), obtemos um polinômio C(x) de grau até 14.
	Mas os elementos de F_{2^8} são representados como polinômios de grau até 7.
	Logo, fazemos C(x) mod m(x), onde m(x) é um polinômio irredutível de grau 8.

	A operação de "mod" com polinômios é análoga à divisão com resto:
	- Subtrai-se (com XOR) o polinômio m(x) alinhado com o maior grau de C(x)
	- Repete-se até que o grau do resto seja menor que o grau de m(x)

	Exemplo:
	  C(x) = x^9 + x^4 + x + 1 → 0b10010011_000
	  m(x) = x^8 + x^4 + x^3 + x + 1 → 0x11B

	  Alinha-se m(x) com x^9 → shift de 1
	  XOR para subtrair → novo resto
	  Repete-se até grau final < 8

	Essa função executa esse processo usando operações bit a bit:
	- degree(p): encontra o termo de maior grau
	- shift: alinha m(x) com p
	- XOR: subtrai polinômios binários (soma módulo 2)

	No final, o polinômio resultante tem grau ≤ 7, e pode ser convertido de volta em uint8.
*/
func polyMod(p, mod uint16) uint8 {
	for degree(p) >= degree(mod) {
		shift := degree(p) - degree(mod)
		p ^= mod << shift
	}
	return uint8(p)
}

/*
	multiplyInF2_8 realiza a multiplicação de dois elementos no campo finito F_{2^8},
	utilizando a representação polinomial sobre F₂[x] com redução módulo um polinômio irredutível.

	Conceito:
	O campo F_{2^8} é construído como o conjunto de todos os polinômios de grau ≤ 7 com coeficientes em F₂,
	com a operação de multiplicação sendo definida como a multiplicação usual de polinômios
	seguida da redução módulo um polinômio irredutível de grau 8 (como x⁸ + x⁴ + x³ + x + 1, usado no AES).

	Multiplicar dois elementos em F_{2^8} envolve:
	1. Interpretar os bytes como polinômios binários (bit i representa x^i)
	2. Multiplicar esses polinômios sobre F₂ usando distributiva (função polyMultiply)
	3. Reduzir o resultado módulo m(x) para que o grau final seja ≤ 7 (função polyMod)

	O valor retornado é o resultado da multiplicação dentro do campo, ou seja,
	o produto polinomial reduzido que também pode ser representado por um byte.
*/
func multiplyInF2_8(a, b uint8) uint8 {
	raw := polyMultiply(a, b)
	return polyMod(raw, 0x11B)
}

/*
	multiplicativeInverse encontra o inverso multiplicativo de um elemento no campo finito F_{2^8}.

	Conceito:
	Em qualquer corpo (field), todo elemento diferente de zero possui um único inverso multiplicativo,
	ou seja, um elemento b tal que: a * b ≡ 1

	No caso de F_{2^8}, os elementos são polinômios binários de grau ≤ 7, e a multiplicação é definida
	com redução módulo um polinômio irredutível (como 0x11B).

	Esta função percorre todos os elementos possíveis de F_{2^8} (exceto zero), e testa se o produto
	entre a e o candidato é igual a 1.

	Embora ineficiente, essa abordagem serve como demonstração prática do conceito de inverso multiplicativo,
	sendo útil para validação, testes e aprendizado.

	Nota:
	- O zero não possui inverso (por definição do corpo)
	- Em aplicações reais (ex: AES), o inverso é obtido via tabelas (log/antilog) ou Extended Euclidean Algorithm
*/
func multiplicativeInverse(a uint8) uint8 {
	if a == 0 {
		return 0
	}
	for i := uint8(1); i < 255; i++ {
		if multiplyInF2_8(a, i) == 1 {
			return i
		}
	}
	return 0
}

func generateLogTables() ([]uint8, []uint8) {
	log := make([]uint8, 256)
	alog := make([]uint8, 256)
	alog[0] = 1
	for i := 1; i < 255; i++ {
		alog[i] = multiplyInF2_8(alog[i-1], 0x03)
		log[alog[i]] = uint8(i)
	}
	return log, alog
}

// mixColumnsStyleMultiplication realiza multiplicação vetorial
// como na operação MixColumns do AES. Cada vetor é multiplicado
// por uma matriz fixa sobre F_{2^8}.
func mixColumnsStyleMultiplication(col [4]uint8) [4]uint8 {
	// Matriz fixa do AES (para MixColumns):
	// [2 3 1 1]
	// [1 2 3 1]
	// [1 1 2 3]
	// [3 1 1 2]

	mul := func(a, b uint8) uint8 {
		return multiplyInF2_8(a, b)
	}

	var result [4]uint8
	result[0] = mul(0x02, col[0]) ^ mul(0x03, col[1]) ^ col[2] ^ col[3]
	result[1] = col[0] ^ mul(0x02, col[1]) ^ mul(0x03, col[2]) ^ col[3]
	result[2] = col[0] ^ col[1] ^ mul(0x02, col[2]) ^ mul(0x03, col[3])
	result[3] = mul(0x03, col[0]) ^ col[1] ^ col[2] ^ mul(0x02, col[3])
	return result
}

func printColumnAsPolys(label string, col [4]uint8) {
	fmt.Printf("%s:\n", label)
	for i := 0; i < 4; i++ {
		fmt.Printf("  [%d] = 0x%02X = %s\n", i, col[i], formatPolynomial(uint16(col[i])))
	}
}

func main() {
	a := uint8(0x57)
	b := uint8(0x13)

	fmt.Println("========== DEMONSTRAÇÃO PRINCIPAL ==========")
	fmt.Printf("a = 0x%X = %s\n", a, formatPolynomial(uint16(a)))
	fmt.Printf("b = 0x%X = %s\n", b, formatPolynomial(uint16(b)))

	sum := polyAdd(a, b)
	fmt.Printf("\nSoma (a + b):\n0x%X = %s\n", sum, formatPolynomial(uint16(sum)))

	raw := polyMultiply(a, b)
	fmt.Printf("\nMultiplicação binária (sem mod):\n0x%X = %s\n", raw, formatPolynomial(raw))

	mod := polyMod(raw, 0x11B)
	fmt.Printf("\nRedução módulo m(x):\n0x%X = %s\n", mod, formatPolynomial(uint16(mod)))

	fieldResult := multiplyInF2_8(a, b)
	fmt.Printf("\nMultiplicação completa em F_{2^8}:\n0x%X = %s\n", fieldResult, formatPolynomial(uint16(fieldResult)))

	inv := multiplicativeInverse(a)
	fmt.Printf("\nInverso multiplicativo de a (0x%X):\n0x%X = %s\n", a, inv, formatPolynomial(uint16(inv)))

	fmt.Println("\n========== TABELA DE LOG E ANTILOG ==========")
	log, alog := generateLogTables()
	fmt.Println("Primeiros 20 elementos da tabela de log:")
	for i := 1; i <= 20; i++ {
		fmt.Printf("log(0x%02X) = %3d\n", alog[i], log[alog[i]])
	}
	fmt.Println("\nPrimeiros 20 elementos da tabela de antilog:")
	for i := 0; i < 20; i++ {
		fmt.Printf("alog[%3d] = 0x%02X = %s\n", i, alog[i], formatPolynomial(uint16(alog[i])))
	}

	fmt.Println("\n========== MIXCOLUMNS DEMO ==========")
	inputCol := [4]uint8{0xdb, 0x13, 0x53, 0x45} // Exemplo famoso do AES
	printColumnAsPolys("Coluna original", inputCol)
	outputCol := mixColumnsStyleMultiplication(inputCol)
	printColumnAsPolys("Coluna transformada (MixColumns)", outputCol)
}
