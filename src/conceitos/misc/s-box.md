# S-Boxes: Substituições Não Lineares


As S-Boxes (Substitution Boxes) são funções fundamentais em cifras simétricas modernas. Elas realizam substituições não lineares, introduzindo confusão no sistema criptográfico, conforme definido por Shannon [1]. Em redes de substituição-permutação (SPN) e cifras de Feistel, as S-Boxes são os principais elementos responsáveis por tornar complexa a relação entre a chave e o texto cifrado.

## Definição Formal

Uma S-Box é uma função

$S: \{0, 1\}^n \rightarrow \{0, 1\}^m$

que mapeia uma entrada binária de $n$ bits para uma saída de $m$ bits. Quando $n = m$, a S-Box é dita bijetiva (ou uma permutação) se cada saída for única.

Exemplo típico:

* No AES, a S-Box é uma permutação de 8 bits para 8 bits, isto é, $S: \{0, 1\}^8 \rightarrow \{0, 1\}^8$.

Propriedades desejáveis:

* Não-linearidade alta (distância de Hamming em relação a todas funções lineares).

* Avalanche: uma mudança de 1 bit na entrada afeta múltiplos bits da saída.

* Resistência a diferenciais e lineares: saída deve ser estatisticamente imprevisível sob diferenças ou aproximações lineares na entrada.


Métodos de construção:

A criação de S-Boxes pode ser feita de diferenes maneiras, dependendo do contexto e objetivos de segurança:

1. Baseadas em estruturas algébricas

* AES (Rijndael) utiliza a inversa multiplicativa no corpo finito $\mathbb{F}_{2^8}$, seguida de uma transformação afim:
    * Inverso: $x \mapsto x^{-1}$ (com $0 \mapsto 0$)
    * Transformação afim: multiplicação e soma sobre $\mathbb{F}_2$

Essa construção é baseada no padrão FIPS 197 [2]. e foi escolhida por sua alta não-linearidade e resistência a ataques diferenciais.

2. Baseadas em constantes (look-up fixo)

* MD2 define uma S-Box de 256 bytes fixos, supostamente derivada dos dígitos de $\pi$. A origem exata nunca foi publicada por Rivest, mas acredita-se que a sequência tenha sido escolhida de forma a evitar padrões evidentes [3].

* DES usa oito S-Boxes diferentes com mapeamentos de 6 bits para 4 bits, cuidadosamente projetadas para resistir a criptoanálise diferencial (ver Coppersmith [4]).

3. Pseudoaleatórias derivadas de chave

* Em cifras modernas como Blowfish, a S-Box é preenchida dinamicamente durante a expansão da chave. Isso dificulta análise e torna a estrutura dependente da chave secreta.


## Exemplo prático em Go: Construindo uma S-Box simples

Aqui, mostramos duas formas distintas de definir S-Boxes em Go:

1. AES-like (algébrica) - inverso multiplicativo seguido de transformação afim (simplificada).
2. MD2-like (constante fixa) - look-up fixo.

```go
package main

import (
	"fmt"
	"math/bits"
)

// Exemplo simplificado de S-Box: inverso em GF(2^8) + afim (AES-like simplificada)
func simpleAffineSBox() [256]byte {
	var sbox [256]byte
	for i := 0; i < 256; i++ {
		inv := inverseByte(byte(i))
		sbox[i] = inv ^ bits.RotateLeft8(inv, 1) ^ 0x63 // afim simplificada
	}
	return sbox
}

// Inverso multiplicativo em GF(2^8) com polinômio redutor x^8 + x^4 + x^3 + x + 1
func inverseByte(x byte) byte {
	if x == 0 {
		return 0
	}
	var inv byte = 1
	for i := 1; i < 256; i++ {
		if byteMultiply(x, inv) == 1 {
			return inv
		}
		inv++
	}
	return 0
}

// Multiplicação em GF(2^8)
func byteMultiply(a, b byte) byte {
	var p byte = 0
	for b != 0 {
		if b&1 != 0 {
			p ^= a
		}
		hi := a & 0x80
		a <<= 1
		if hi != 0 {
			a ^= 0x1B // AES irreducible polynomial
		}
		b >>= 1
	}
	return p
}

// Exemplo de S-Box constante (MD2-like)
func md2SBox() [256]byte {
	return [256]byte{
		0x29, 0x2E, 0x43, 0x32, 0x8B, 0x1C, 0x1A, 0x06, 0x3F, 0x3C, 0x7F, 0x8A, 0x0E, 0x19, 0x4F, 0x43,
		// ... complete até 256 valores reais conforme tabela do MD2
	}
}

func main() {
	sbox := simpleAffineSBox()
	fmt.Printf("Exemplo de S-Box (simplificada):\n")
	for i := 0; i < 16; i++ {
		fmt.Printf("%02x ", sbox[i])
	}
	fmt.Println()
}
```

## Considerações finais

As S-Boxes são blocos fundamentais na construção de cifras seguras. Sua estrutura influencia diretamente a resistência contra ataques diferenciais, lineares e por análise estatística. S-Boxes bem projetadas devem apresentar não-linearidade elevada e boa dispersão da entrada na saída, contribuindo para o efeito avalanche.

Embora possam ser fixas ou derivadas de chave, é essencial que sejam analisadas formalmente quanto às suas propriedades de segurança — algo que o projeto do AES exemplifica de forma rigorosa. Em contrapartida, designs mais antigos, como MD2, revelam escolhas empíricas, ainda que eficazes em seu tempo.

## Referências

[1] C. E. Shannon, Communication Theory of Secrecy Systems, Bell System Technical Journal, 1949.

[2] National Institute of Standards and Technology, FIPS-197: Advanced Encryption Standard (AES), 2001.

[3] R. Rivest, The MD2 Message-Digest Algorithm, RFC 1319, 1992.

[4] D. Coppersmith, The Data Encryption Standard (DES) and Its Strength Against Attacks, IBM Journal of Research and Development, 1994.

[5] D. E. Knuth, The Art of Computer Programming, Vol. 2, Addison-Wesley, 1997.

