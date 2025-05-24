package main

import (
	"fmt"

	"github.com/fatih/color"
)

// Nomes das "pessoas"
var names = []string{"A", "B", "C", "D", "E"}

// Permutação: índices indicam quem vira quem
// A → C, B → A, C → E, D → D, E → B
var perm = []int{2, 0, 4, 3, 1}

// Decompõe a permutação em ciclos disjuntos
func decomposeCycles(p []int) [][]int {
	n := len(p)
	visited := make([]bool, n)
	var cycles [][]int

	for i := 0; i < n; i++ {
		if !visited[i] {
			var cycle []int
			x := i
			for !visited[x] {
				cycle = append(cycle, x)
				visited[x] = true
				x = p[x]
			}
			if len(cycle) > 1 {
				cycles = append(cycles, cycle)
			}
		}
	}
	return cycles
}

// Imprime os ciclos disjuntos com estilo
func printCycles(cycles [][]int, names []string) {
	title := color.New(color.FgCyan, color.Bold)
	title.Println("\n🔁 Ciclos disjuntos:")
	cycleColor := color.New(color.FgHiYellow)
	for _, cycle := range cycles {
		cycleColor.Print("  (")
		for i, idx := range cycle {
			if i > 0 {
				cycleColor.Print(" → ")
			}
			cycleColor.Print(names[idx])
		}
		cycleColor.Println(")")
	}
}

// Imprime o mapeamento direto
func printMapping(p []int, names []string) {
	title := color.New(color.FgCyan, color.Bold)
	title.Println("🔀 Mapeamento (quem vira quem):")
	fromColor := color.New(color.FgGreen)
	toColor := color.New(color.FgHiMagenta)

	for i, to := range p {
		fromColor.Printf("  %s ", names[i])
		fmt.Print("→ ")
		toColor.Printf("%s\n", names[to])
	}
}

// Calcula a inversa da permutação
func invertPermutation(p []int) []int {
	n := len(p)
	inv := make([]int, n)
	for i := 0; i < n; i++ {
		inv[p[i]] = i
	}
	return inv
}

// Imprime a inversa corretamente (quem virou quem originalmente)
func printInverse(inv []int, names []string) {
	title := color.New(color.FgCyan, color.Bold)
	title.Println("\n🔁 Inversa da permutação (reverte o mapeamento):")
	fromColor := color.New(color.FgHiMagenta)
	toColor := color.New(color.FgGreen)

	for i, from := range inv {
		fromColor.Printf("  %s ", names[i]) // imagem → origem
		fmt.Print("→ ")
		toColor.Printf("%s\n", names[from])
	}
}

// Funções auxiliares: mdc e mmc
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

// Calcula a ordem da permutação via MMC dos tamanhos dos ciclos
func orderOfPermutation(p []int) int {
	cycles := decomposeCycles(p)
	order := 1
	for _, cycle := range cycles {
		order = lcm(order, len(cycle))
	}
	return order
}

func main() {
	printMapping(perm, names)

	cycles := decomposeCycles(perm)
	printCycles(cycles, names)

	inv := invertPermutation(perm)
	printInverse(inv, names)

	order := orderOfPermutation(perm)
	color.New(color.FgHiBlue, color.Bold).Printf("\n⏳ Ordem da permutação: %d\n", order)
}
