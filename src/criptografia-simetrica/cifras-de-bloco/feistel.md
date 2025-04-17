# A estrutura de Feistel

Cifras de bloco transformam blocos de dados de tamanho fixo (por exemplo, 64 bits) em blocos cifrados do mesmo tamanho. A `estrutura de Feistel` é um padrão de construção para cifras de bloco que permite aplicar os princípios de `confusão` e `difusão`, mesmo quando a função usada internamente não é reversível.

Proposta por `Horst Feistel` na década de 1970, essa estrutura permite construir cifras reversíveis a partir de funções simples. A ideia central é dividir o bloco de entrada em duas metades e, a cada rodada, aplicar uma função a uma das metades e combiná-la com a outra usando uma operação XOR. O resultado é uma cifra simétrica, ou seja, que pode ser `decifrada utilizando as mesmas operações em ordem inversa`.

## Estrutura Geral

Dado um bloco de entrada de tamanho $2n$, ele é dividido em duas partes iguais, $L_0$ (metade esquerda) e $R_0$ (metade direita).
A cada rodada $i$, aplica-se a seguinte transformação:

$$
\begin{aligned}
L_i &= R_{i-1} \\
R_i &= L_{i-1} \oplus f(R_{i-1}, k_i)
\end{aligned}
$$


Aqui, $f$ é uma função arbitrária que não precisa ser invertível, e $k_i$ é a subchave da rodada $i$.

Após todas as rodadas, o par $(L_{n}, R_{n})$ forma o bloco cifrado. Em algumas cifras, há uma `permutação final` ou troca de posição entre as mates antes da saída.
A `decifragem` é feita aplicando as mesmas operações, mas `com as subchaves na ordem inversa`.

### Exemplo Didático: `Rede de Feistel` com 2 rodadas

A seguir, apresentamos uma implementação simples em Go de uma rede de Feistel com blocos de 8 bits (divididos em duas metades de 4 bits) e duas rodadas com subchaves fixas.

```go
package main

import (
	"fmt"
)

func F(R, K byte) byte {
	return (R + K) % 16
}

func FeistelEncrypt(P byte, K []byte) byte {
	L := (P >> 4) & 0xF 
	R := P & 0xF

	for i := 0; i < len(K); i++ {
		newL := R
		R = L ^ F(R, K[i])
		L = newL
	}

	return (L << 4) | R
}

func FeistelDecrypt(C byte, K []byte) byte {
	L := (C >> 4) & 0xF
	R := C & 0xF
	
	for i := len(K) - 1; i >= 0; i-- {
		newR := L
		L = R ^ F(L, K[i])
		R = newR
	}
	return (L << 4) | R
}

func main() {
	P := byte('A')        // plaintext: bloco de 8 bits, caractere 'A' = 0x41 = 01000001
	K := []byte{7, 3}     // chaves: 2 chaves de 4 bits para 2 rodadas

	C := FeistelEncrypt(P, K)
	D := FeistelDecrypt(C, K)
	
	fmt.Println()
	fmt.Println("=== Demonstração da Cifra de Feistel (2 rodadas, 8 bits) ===")
	fmt.Println()
	fmt.Printf("%-12s %-10s %-10s %-10s %-10s\n", "Tipo", "Char", "Bin", "Hex", "Decimal")
	fmt.Println("------------------------------------------------------------")
	fmt.Printf("%-12s %-10q %08b   0x%02X       %-10d\n", "Plaintext", P, P, P, P)
	fmt.Printf("%-12s %-10q %08b   0x%02X       %-10d\n", "Ciphertext", C, C, C, C)
	fmt.Printf("%-12s %-10q %08b   0x%02X       %-10d\n", "Decrypted", D, D, D, D)
	fmt.Println()
}
```

Esse exemplo mostra como a estrutura de Feistel permite construir uma cifra reversível, mesmo que as operações internas (como a função F) não sejam reversíveis por si só.

## Algoritmos Baseados na Estrutura de Feistel

A seguir, listamos exemplos de cifras de bloco que usam a rede de Feistel, agrupadas por época ou aplicação:

### Clássicos

- **Lucifer** (IBM, 1973): primeira cifra baseada explicitamente em Feistel.
- **DES** (1977): 16 rodadas de Feistel, com S-boxes e permutações fixas.
- **3DES** (Triple DES): encadeamento de três aplicações do DES.

### Cifras Modernas Baseadas em Feistel

- **Blowfish** (1993): 16 rodadas, chave variável.
- **Twofish** (1998): estrutura similar a Feistel com difusão aprimorada.
- **CAST-128** / **CAST-256**: cifras com funções não lineares, também baseadas em Feistel.
- **Camellia** (2000): cifra com estrutura Feistel e S-boxes inspiradas no AES.
- **MARS** (IBM, finalista do AES): núcleo baseado em Feistel.
- **GOST 28147-89**: cifra soviética com 32 rodadas e S-boxes configuráveis.
- **RC5** / **RC6**: variantes com estrutura semelhante à de Feistel.

### Cifras Leves e Aplicações Específicas

- **ICE** (Information Concealment Engine)
- **FEAL** (Fast Data Encipherment Algorithm)
- **HIGHT**: projetada para dispositivos de baixo consumo.
- **LEA**: cifra leve coreana, com estrutura semelhante a Feistel


## Algoritmos que Não Utilizam Feistel

Algumas cifras modernas utilizam outras estruturas, como as redes de **substituição-permutação (SPN)** ou construções baseadas em permutações.

### SPN (Substitution–Permutation Network)

- **AES (Rijndael)**: usa uma rede SPN com substituições via S-boxes e permutações lineares; não é baseado em Feistel.
- **PRESENT**: cifra leve com estrutura SPN, amplamente usada em implementações de hardware.
- **LED**: projetada para dispositivos de recursos extremamente limitados; também segue a estrutura SPN.
- **NOEKEON**: cifra leve com estrutura SPN e foco em simplicidade algorítmica.
- **PRINCE**: cifra otimizada para baixa latência, com estrutura reflexiva baseada em SPN.

### Outras Estruturas

- **SIMON / SPECK (NSA)**: cifras leves com estruturas alternativas; SIMON usa operações bit a bit, SPECK opera com adição, rotação e XOR. Nenhuma delas segue o modelo Feistel tradicional.
- **ASCON**: cifra autenticada baseada em permutação, vencedora da competição NIST para criptografia leve (Lightweight Cryptography).
- **Khazad / Anubis**: cifradores de bloco com estruturas SPN, fortemente influenciados pelo AES.
- **Grain / Trivium**: cifras de fluxo, não baseadas em Feistel nem em SPN; operam em fluxo contínuo de bits.



