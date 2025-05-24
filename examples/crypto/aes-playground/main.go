package main

import (
	"fmt"
)

/*
	Polinômio irreducível de grau 8 utilizado no AES

	Representações equivalentes:

	- Representação polinomial (forma compacta):      x^8 + x^4 + x^3 + x + 1
	- Representação polinomial (forma longa):         1·x^8 + 0·x^7 + 0·x^6 + 0·x^5 + 1·x^4 + 1·x^3 + 0·x^2 + 1·x^1 + 1·x^0
	- Representação binária (bit 8 → bit 0):          1 0 0 0 1 1 0 1 1
	- Representação hexadecimal:                      0x11B
	- Representação decimal:                          283

	Detalhamento por posição de bit (do mais significativo ao menos significativo):

	| Potência (x^i) | Coeficiente | Bit posição |
	|----------------|-------------|-------------|
	| x^8            |      1      |     8       |
	| x^7            |      0      |     7       |
	| x^6            |      0      |     6       |
	| x^5            |      0      |     5       |
	| x^4            |      1      |     4       |
	| x^3            |      1      |     3       |
	| x^2            |      0      |     2       |
	| x^1            |      1      |     1       |
	| x^0            |      1      |     0       |

	Este polinômio é utilizado como módulo nas operações do campo finito GF(2^8) no AES.
	Ele define a aritmética de multiplicação e inversão no campo, pois é irreducível
	sobre GF(2) — ou seja, não pode ser fatorado, garantindo estrutura de corpo.

	Quando multiplicamos dois elementos do campo (bytes, ou polinômios de grau ≤ 7),
	o resultado pode ter grau até 14. Para manter o resultado dentro de GF(2^8),
	fazemos a redução módulo este polinômio:

		a(x) · b(x) mod m(x)

	Onde:
	- a(x), b(x) são polinômios binários representando os bytes
	- m(x) = x^8 + x^4 + x^3 + x + 1 é o polinômio redutor
	- Essa operação define a multiplicação no corpo GF(2^8)

	Importante: apesar dos elementos do campo terem no máximo 8 bits (grau 7),
	o polinômio de redução tem grau 8 (9 bits), pois é necessário para definir o corpo GF(2^8).
*/
const modulus = 0x11B

/*
	gfMultiply multiplica dois elementos (bytes) em GF(2^8),
	isto é, dois polinômios binários de grau ≤ 7, com coeficientes em GF(2),
	e reduz o resultado módulo m(x) = x^8 + x^4 + x^3 + x + 1 (0x11B), que define o campo.

	Observações:
	- Os bytes a e b são vistos como polinômios: ex: 0x57 → x^6 + x^4 + x^2 + x + 1
	- A multiplicação é feita como produto de polinômios binários
	- A redução (mod m(x)) é feita de forma implícita usando XOR com 0x1B
*/
func gfMultiply(a, b byte) byte {
	var result byte = 0

	for i := 0; i < 8; i++ {

		if (b & 1) != 0 {
			result ^= a
		}

		carry := (a & 0x80) != 0

		a <<= 1

		if carry {
			a ^= 0x1B
		}

		b >>= 1
	}

	return result
}
