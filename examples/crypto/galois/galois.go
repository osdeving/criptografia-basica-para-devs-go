package main

import (
	"fmt"
)

// Polinômio irredutível m(x) = x^8 + x^4 + x^3 + x + 1
// Representado como uint16 porque tem grau 8 (bit 8 = x^8)
const modulus uint16 = 0x11B

// Soma ou subtração em F_{2^8} é XOR
func gfAdd(a, b byte) byte {
	return a ^ b
}

// Multiplicação em F_{2^8} com redução módulo m(x)
func gfMul(a, b byte) byte {
	var result byte = 0
	for i := 0; i < 8; i++ {
		// Se o bit menos significativo de b estiver ativo
		if b&1 != 0 {
			result ^= a // Soma = XOR
		}

		// Verifica overflow antes de deslocar
		hiBit := a & 0x80
		a <<= 1
		if hiBit != 0 {
			// Redução por m(x) — só a parte baixa do polinômio irreduzível (sem x^8)
			a ^= byte(modulus & 0xFF)
		}

		b >>= 1
	}
	return result
}

// Calcula o grau de um polinômio binário (uint16)
func degree(p uint16) int {
	for i := 15; i >= 0; i-- {
		if p&(1<<i) != 0 {
			return i
		}
	}
	return -1
}

// Divisão polinomial binária: retorna quociente
func polyDiv(dividend, divisor uint16) uint16 {
	if divisor == 0 {
		panic("divisão por zero")
	}

	var quotient uint16 = 0
	for degree(dividend) >= degree(divisor) {
		shift := degree(dividend) - degree(divisor)
		quotient ^= 1 << shift
		dividend ^= divisor << shift
	}
	return quotient
}

// Multiplica um polinômio (uint16) por um byte (interpreta o byte como polinômio)
func polyMul(a byte, b uint16) uint16 {
	var result uint16 = 0
	for i := 0; i < 8; i++ {
		if a&(1<<i) != 0 {
			result ^= b << i
		}
	}
	return result
}

// Retorna o inverso multiplicativo de a em F_{2^8}
func gfInverse(a byte) byte {
	if a == 0 {
		panic("zero não tem inverso em um corpo")
	}

	var r0, r1 uint16 = modulus, uint16(a)
	var t0, t1 uint16 = 0, 1

	for r1 != 0 {
		q := polyDiv(r0, r1)
		r0, r1 = r1, r0^polyMul(byte(q), r1)
		t0, t1 = t1, t0^polyMul(byte(q), t1)
	}

	// O inverso está em t0 — ele é garantidamente < 256
	return byte(t0)
}

func main() {
	a := byte(0x57) // x^6 + x^4 + x^2 + x + 1
	b := byte(0x83) // x^7 + x + 1

	fmt.Printf("a        = 0x%02X\n", a)
	fmt.Printf("b        = 0x%02X\n", b)

	sum := gfAdd(a, b)
	fmt.Printf("a + b    = 0x%02X\n", sum)

	sub := gfAdd(a, b) // mesmo que soma
	fmt.Printf("a - b    = 0x%02X\n", sub)

	mul := gfMul(a, b)
	fmt.Printf("a * b    = 0x%02X\n", mul)

	inv := gfInverse(a)
	fmt.Printf("a^-1     = 0x%02X\n", inv)

	check := gfMul(a, inv)
	fmt.Printf("a * a^-1 = 0x%02X\n", check) // deve dar 0x01
}
