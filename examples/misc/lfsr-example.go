package main

import (
	"fmt"
	"strings"
)

type LFSR struct {
	state uint8
	taps  []int
	size  int
}

func NewLFSR(seed uint8, taps []int, size int) *LFSR {
	return &LFSR{
		state: seed,
		taps:  taps,
		size:  size,
	}
}

func (l *LFSR) Step() (newBit uint8, outBit uint8) {
	var feedback uint8 = 0
	for _, tap := range l.taps {
		feedback ^= (l.state >> tap) & 1
	}

	outBit = l.state & 1
	l.state = (l.state >> 1) | (feedback << (l.size - 1))
	return feedback, outBit
}

func formatBits(val uint8, size int) string {
	builder := strings.Builder{}
	for i := size - 1; i >= 0; i-- {
		bit := (val >> i) & 1
		builder.WriteByte('0' + bit)
	}
	return builder.String()
}

func main() {
	// LFSR de 4 bits com taps nas posições 3 e 0 (bit mais à esquerda e mais à direita)
	taps := []int{3, 0}
	size := 4
	seed := uint8(0b1001)

	lfsr := NewLFSR(seed, taps, size)

	fmt.Println("Simulação de LFSR com visualização:")
	fmt.Println("-------------------------------------")
	fmt.Printf("Estado inicial: %s\n", formatBits(lfsr.state, size))
	fmt.Printf("Taps: posições %v\n", taps)
	fmt.Println()

	// Simula 10 ciclos
	for step := 1; step <= 10; step++ {
		feedback, outBit := lfsr.Step()
		fmt.Printf("Etapa %2d:\n", step)
		fmt.Printf("  Estado atual:      %s\n", formatBits(lfsr.state, size))
		fmt.Printf("  Bit de saída:      %d (mais à direita antes do shift)\n", outBit)
		fmt.Printf("  Feedback (XOR taps): %d\n", feedback)
		fmt.Println()
	}
}
