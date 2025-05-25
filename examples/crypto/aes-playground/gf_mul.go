package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

// Converte um byte para forma polinomial, ex: 0x57 → x^6 + x^4 + x^2 + x + 1
func byteToPoly(x byte) string {
	if x == 0 {
		return "0"
	}
	terms := []string{}
	for i := 7; i >= 0; i-- {
		if (x>>i)&1 != 0 {
			if i == 0 {
				terms = append(terms, "1")
			} else if i == 1 {
				terms = append(terms, "x")
			} else {
				terms = append(terms, fmt.Sprintf("x^%d", i))
			}
		}
	}
	return strings.Join(terms, " + ")
}

// Imprime byte binário com bit ativo colorido, seguido da forma polinomial
func printByte(label string, b byte, activeBit int, comment string) {
	fmt.Print(label + " = ")
	for i := 7; i >= 0; i-- {
		bit := (b >> i) & 1
		if i == activeBit {
			color.New(color.FgHiGreen).Printf("%d", bit)
		} else {
			fmt.Printf("%d", bit)
		}
	}
	fmt.Printf("   %s", comment)
	fmt.Printf("   (%s)\n", byteToPoly(b))
}

func gfMultiplyVerbose(a0, b0 byte) byte {
	var result byte = 0
	a := a0
	b := b0

	fmt.Println()
	color.New(color.Bold).Printf("Teste de mesa para gfMultiply(a=0x%02X, b=0x%02X)\n", a, b)
	fmt.Println()

	for step := 1; step <= 8; step++ {
		color.New(color.FgHiWhite, color.Bold).Printf("Passo%d:\n", step)
		fmt.Println("Inicio:")

		printByte("b", b, 0, "")
		printByte("a", a, 7, "")

		fmt.Println("\nAlgoritmo:")

		if b&1 != 0 {
			color.Green("b & 1 ocorre, então result ^= a, agora result =")
			printByte("result", result^a, -1, "(resultado acumulado)")
			result ^= a
		} else {
			fmt.Println("b & 1 não ocorre, então segue")
		}

		if a&0x80 != 0 {
			color.Red("a & 10000000 ocorre, então haverá redução")
		} else {
			fmt.Println("a & 10000000 também não ocorre, então segue")
		}

		carry := (a & 0x80) != 0
		a <<= 1
		if carry {
			color.Red("a <<= 1 causou overflow (x^8), então aplicamos redução a ^= 00011011")
			a ^= 0x1B
			printByte("a", a, -1, "(após redução)")
		} else {
			printByte("a", a, -1, "a <<= 1 (bit andou para esquerda)")
		}

		b >>= 1
		printByte("b", b, 0, "b >>= 1 (bit andou para direita)")

		fmt.Println("\naqui termina a passada")
		fmt.Println(strings.Repeat("-", 64)) // separador
	}

	color.Cyan("Resultado final: 0x%02X   (%s)\n", result, byteToPoly(result))
	return result
}

func main() {
	a := byte(0x10) // x^4
	b := byte(0x02) // x^1

	gfMultiplyVerbose(a, b)
}
