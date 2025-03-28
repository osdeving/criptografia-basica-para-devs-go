package main

import (
	"bufio"
	"fmt"
	"os"
)

func xorEncrypt(plaintext, key string) string {
	ciphertext := make([]byte, len(plaintext))
	for i := 0; i < len(plaintext); i++ {
		ciphertext[i] = plaintext[i] ^ key[i%len(key)]
	}
	return string(ciphertext)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println()
	fmt.Println("🔒 Encriptador XOR 🔒")
	fmt.Println()

	fmt.Print("📄 Digite a mensagem a ser encriptada: ")
	scanner.Scan()
	plaintext := scanner.Text()

	fmt.Print("🔑 Digite a chave de encriptação:")
	scanner.Scan()
	key := scanner.Text()

	ciphertext := xorEncrypt(plaintext, key)

	fmt.Println()

	fmt.Printf("🔒 Mensagem encriptada: %v\n", ciphertext)

	plainTextAgain := xorEncrypt(ciphertext, key)

	fmt.Printf("🔓 Mensagem decriptada: %v\n", plainTextAgain)
}
