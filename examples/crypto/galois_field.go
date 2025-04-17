package main

import (
	"fmt"
	"strings"
)

// Converte um inteiro binário para representação polinomial
func formatPolynomial(poly uint16) string {
	if poly == 0 {
		return "0"
	}
	var terms []string
	for i := 15; i >= 0; i-- {
		if poly&(1<<i) != 0 {
			switch i {
			case 0:
				terms = append(terms, "1")
			case 1:
				terms = append(terms, "x")
			default:
				terms = append(terms, fmt.Sprintf("x^%d", i))
			}
		}
	}
	return strings.Join(terms, " + ")
}

// Gera todos os polinômios de grau exato n sobre F_2[x]
func generatePolynomialsOfDegree(n int) []uint16 {
	start := 1 << n         // 2^n
	end := (1 << (n + 1)) - 1 // 2^(n+1) - 1
	var polys []uint16
	for i := start; i <= end; i++ {
		polys = append(polys, uint16(i))
	}
	return polys
}

func polyAdd(a, b uint16) uint16 {
	return a ^ b // Soma binária em F_2 é XOR
}

func polyMultiply(a, b uint16) uint16 {
	var result uint16 = 0
	for i := 0; i < 16; i++ {
		if (b & (1 << i)) != 0 {
			result ^= a << i
		}
	}
	return result
}

func polyMod(p, mod uint16) uint16 {
	modDeg := degree(mod)
	for degree(p) >= modDeg {
		shift := degree(p) - modDeg
		p ^= mod << shift
	}
	return p
}

// Calcula o grau do polinômio
func degree(p uint16) int {
	for i := 15; i >= 0; i-- {
		if p&(1<<i) != 0 {
			return i
		}
	}
	return -1 // grau de zero
}

func addF3(a, b int) int {
	return (a + b) % 3
}

func subF3(a, b int) int {
	return (a - b + 3) % 3 // evita negativo
}

func polyAddArray(p, q []int) []int {
	n := max(len(p), len(q))
	result := make([]int, n)

	for i := 0; i < n; i++ {
		var pi, qi int
		if i < len(p) {
			pi = p[i]
		}
		if i < len(q) {
			qi = q[i]
		}
		result[i] = pi + qi
	}
	return result
}

func polySubArray(p, q []int) []int {
	n := max(len(p), len(q))
	result := make([]int, n)

	for i := 0; i < n; i++ {
		var pi, qi int
		if i < len(p) {
			pi = p[i]
		}
		if i < len(q) {
			qi = q[i]
		}
		result[i] = pi - qi
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func formatPolyArray(p []int) string {
	result := ""
	for i := len(p) - 1; i >= 0; i-- {
		coef := p[i]
		if coef == 0 {
			continue
		}
		if result != "" && coef > 0 {
			result += " + "
		} else if coef < 0 {
			result += " - "
			coef = -coef
		}
		if i == 0 {
			result += fmt.Sprintf("%d", coef)
		} else if i == 1 {
			result += fmt.Sprintf("%dx", coef)
		} else {
			result += fmt.Sprintf("%dx^%d", coef, i)
		}
	}
	if result == "" {
		return "0"
	}
	return result
}

func testPolynomialOperations() {
	p := []int{1, 4, 3} // 3x^2 + 4x + 1
	q := []int{5, 0, 2} // 2x^2 + 5

	sum := polyAddArray(p, q)
	sub := polySubArray(p, q)

	fmt.Println("P(x) =", formatPolyArray(p))
	fmt.Println("Q(x) =", formatPolyArray(q))
	fmt.Println("P + Q =", formatPolyArray(sum))
	fmt.Println("P - Q =", formatPolyArray(sub))
}

func main() {
	n := 4
	fmt.Printf("Polinômios binários de grau exato %d:\n", n)
	polys := generatePolynomialsOfDegree(n)
	for _, p := range polys {
		fmt.Printf("0x%X = %s\n", p, formatPolynomial(p))
	}

	a := uint16(0b1011) // x^3 + x + 1
	b := uint16(0b0111) // x^2 + x + 1
	m := uint16(0x11B)  // AES irreduzível

	sum := polyAdd(a, b)
	mul := polyMultiply(a, b)
	mod := polyMod(mul, m)

	fmt.Println()
	fmt.Println("Exemplos de operações com polinômios em F_2[x]")
	fmt.Println()

	fmt.Printf("a(x) = %s\n", formatPolynomial(a))
	fmt.Printf("b(x) = %s\n", formatPolynomial(b))
	fmt.Printf("a + b = %s\n", formatPolynomial(sum))
	fmt.Printf("a * b = %s\n", formatPolynomial(mul))
	fmt.Printf("(a * b) mod m = %s\n", formatPolynomial(mod))

	fmt.Println()
	fmt.Print("Exemplos de operações")
	fmt.Println()

	testPolynomialOperations()
}
