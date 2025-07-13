package main

import "testing"

func TestSboxCorrectness(t *testing.T) {
	known := [256]byte{
		0x63, 0x7c, 0x77, 0x7b, 0xf2, 0x6b, 0x6f, 0xc5,
		0x30, 0x01, 0x67, 0x2b, 0xfe, 0xd7, 0xab, 0x76,
		0xca, 0x82, 0xc9, 0x7d, 0xfa, 0x59, 0x47, 0xf0,
		0xad, 0xd4, 0xa2, 0xaf, 0x9c, 0xa4, 0x72, 0xc0,
		// ... completar até 256 se desejar
	}
	gen := generateSbox()
	for i := 0; i < len(known); i++ {
		if known[i] != gen[i] {
			t.Fatalf("S-box[%02X]: esperado %02X, obtido %02X", i, known[i], gen[i])
		}
	}
}

func TestSboxInverse(t *testing.T) {
	s := generateSbox()
	inv := generateInvSbox()
	for i := 0; i < 256; i++ {
		res := inv[s[i]]
		if byte(i) != res {
			t.Fatalf("InvSbox[Sbox[%02X]] != %02X (deu %02X)", i, i, res)
		}
	}
}
