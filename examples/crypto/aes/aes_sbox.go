package main

import (
	"fmt"
)
var irreducible int = 0x11B

func gfMul(a, b byte) byte {
	var res byte = 0
	for b > 0 {
		if b&1 != 0 {
			res ^= a
		}
		if a&0x80 != 0 {
			a = a<<1 ^ byte(irreducible)
		} else {
			a <<= 1
		}
		b >>= 1
	}
	return res
}

func gfInv(a byte) byte {
	if a == 0 {
		return 0
	}

	r0, r1 := irreducible, int(a)
	t0, t1 := 0, 1

	for r1 != 0 {
		q := r0 / r1
		r0, r1 = r1, r0%r1

		// Corrigir multiplicação: q é int, t1 é int, convertemos para byte
		t0, t1 = t1, t0^int(gfMul(byte(q), byte(t1)))
	}
	return byte(t0)
}


func affineTransform(x byte) byte {
	var result byte = 0
	for i := 0; i < 8; i++ {
		bit := ((x >> i) & 1) ^
			((x >> ((i + 4) % 8)) & 1) ^
			((x >> ((i + 5) % 8)) & 1) ^
			((x >> ((i + 6) % 8)) & 1) ^
			((x >> ((i + 7) % 8)) & 1) ^
			((0x63 >> i) & 1)
		result |= bit << i
	}
	return result
}

func inverseAffineTransform(y byte) byte {
	var result byte = 0
	for i := 0; i < 8; i++ {
		bit := ((y >> ((i + 2) % 8)) & 1) ^
			((y >> ((i + 5) % 8)) & 1) ^
			((y >> ((i + 7) % 8)) & 1) ^
			((y >> i) & 1) ^
			((0x05 >> i) & 1)
		result |= bit << i
	}
	return result
}

func sbox(x byte) byte {
	inv := gfInv(x)
	return affineTransform(inv)
}

func invSbox(y byte) byte {
	pre := inverseAffineTransform(y)
	return gfInv(pre)
}

func generateSbox() [256]byte {
	var table [256]byte
	for i := 0; i < 256; i++ {
		table[i] = sbox(byte(i))
	}
	return table
}

func generateInvSbox() [256]byte {
	var table [256]byte
	for i := 0; i < 256; i++ {
		table[i] = invSbox(byte(i))
	}
	return table
}

func main() {
	s := generateSbox()
	i := generateInvSbox()

	fmt.Printf("S-box[0x53] = 0x%02X\n", s[0x53]) // Esperado: 0xED
	fmt.Printf("InvS-box[0xED] = 0x%02X\n", i[0xED]) // Esperado: 0x53

	fmt.Println("Tabela S-box gerada:")
	for i := 0; i < 256; i++ {
		fmt.Printf("0x%02X,", s[i])
		if (i+1)%16 == 0 {
			fmt.Println()
		}
	}
}

