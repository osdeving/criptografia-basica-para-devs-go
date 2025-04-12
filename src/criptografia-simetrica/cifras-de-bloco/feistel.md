# A estrutura de Feistel

Vimos que cifras de bloco transformam blocos de dados de tamanho fixo (e.g. 64 bits) em blocos cifrados de mesmo tamanho. A `estrutura de Feistel` é um padrão de construção de cifras de bloco que permite aplicar `confusão` e `difusão`. 

A estrutura de Feistel (chamada também de `Feistel Network`), proposta por Horst Feistel em 1970, é uma abordagem muito utilizada para construir cifras de bloco a partir de funções não necessariamente invertíveis. A ideia central dessa estrutura é dividir o bloco de entrada em duas partes iguais, aplicar uma função não invertível a uma das partes e combinar o resultado com a outra parte usando uma operação `XOR`.

## Estrutura Geral

Dado um bloco de entrada de tamanho $2n$, ele é dividido em duas partes iguais, $L_0$ e $R_0$.

$$
\begin{aligned}
L_0 &\leftarrow \text{parte esquerda do bloco} \\
R_0 &\leftarrow \text{parte direita do bloco}
\end{aligned}
$$

Cada rodada $i$ aplica a seguinte transformação:

$$
\begin{aligned}
L_i &= R_{i-1} \\
R_i &= L_{i-1} \oplus f(R_{i-1}, k_i)
\end{aligned}
$$


Aqui, $f$ é uma função que pode ser não invertível, e $k_i$ é a subchave da rodada $i$.

Ao final das rodadas, o par $(L_{n}, R_{n})$ é combinado e (em alguns casos) passado por uma permutação final.

Uma característica importante da estrutura de Feistel é que a decifração é realizada aplicando as mesmas operações em ordem inversa, com as subchaves utilizadas na ordem inversa.

### Exemplo simplificado (Rede de Feistel de 2 rodadas)

Para ilustrar, considere uma rede de Feistel com 2 rodadas e blocos de 8 bits, divididos em duas metades de 4 bits.

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

Esse exemplo usa operações simples para fins didáticos. A ideia é mostrar como, mesmo com funções unilaterais simples, a estrutura de Feistel permite a construção de cifras reversíveis.

## Algoritmos que usam a estrutura de Feistel

Esses algoritmos seguem a estrutura clássica ou variantes da rede de Feistel:

Clássicos / Históricos
Lucifer (IBM, 1973): primeiro a usar Feistel explicitamente.

DES (1977): 16 rodadas de Feistel com S-boxes e permutações fixas.

3DES (Triple-DES): composição de DES, também baseado em Feistel.

Modernos (até anos 2000)
Blowfish (Schneier, 1993): Feistel de 16 rodadas com chave variável.

Twofish (Schneier et al., 1998): usa uma estrutura pseudo-Feistel.

CAST-128/CAST-256: cifras de Feistel com funções não lineares.

Camellia (NTT, Mitsubishi, 2000): estrutura Feistel com S-boxes inspiradas no AES.

MARS (IBM, candidato AES): usa um núcleo do tipo Feistel.

GOST 28147-89 (União Soviética): 32 rodadas de Feistel com S-boxes configuráveis.

RC5 / RC6: estrutura Feistel modificada (dependendo da parametrização).

Outros exemplos notáveis
ICE (Information Concealment Engine)

FEAL (Fast Data Encipherment Algorithm)

HIGHT (cifra leve baseada em Feistel, usada em IoT)

LEA (cifra leve coreana, com estrutura semelhante a Feistel)

## Algoritmos que não usam rede de Feistel

Esses usam redes de substituição-permutação (SPN) ou outras estruturas:

AES e derivados
AES (Rijndael): usa uma rede SPN (Substituição–Permutação Network), não é Feistel.

PRESENT: cifra leve SPN, amplamente usada em hardware.

LED (Lightweight Encryption Device): também SPN.

Modernos e pós-quantum
SIMON / SPECK (NSA): SIMON usa estrutura bitwise alternada, não exatamente Feistel.

ASCON (padrão NIST para cifragem autenticada): SPN com permutação.

NOEKEON: cifra leve, rede SPN.

PRINCE: estrutura reflexiva, não Feistel.

Outros exemplos
Khazad / Anubis: SPN, baseados em ideias do AES.

Grain / Trivium: cifras de fluxo (não são cifras de bloco, mas também não usam Feistel).

