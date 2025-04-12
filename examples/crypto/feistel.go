package main

import (
	"fmt"
)

func F(R, K byte) byte {
	return (R + K) % 16
}

func FeistelEncrypt(P byte, K []byte) byte {
	L := (P >> 4) & 0xF // 4 bits mais significativos (LSB)
	R := P & 0xF        // 4 bits menos significativos (MSB)

	for i := 0; i < len(K); i++ {
		newL := R
		R = L ^ F(R, K[i])
		L = newL
	}

	return (L << 4) | R
}

func FeistelDecrypt(C byte, K []byte) byte {
	L := (C >> 4) & 0xF // 4 bits mais significativos (LSB)
	R := C & 0xF        // 4 bits menos significativos (MSB)
	
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