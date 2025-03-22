package main

import (
	"fmt"
	"strings"
)


type LFSR struct {
	state uint8
	taps []int
}

func NewLFSR(seed uint8, taps []int) *LFSR {
	return &LFSR{state: seed, taps: taps}
}


func (l *LFSR) Step() uint8 {
	var feedback uint8 = 0

	for _, tap := range l.taps {
		feedback ^= (l.state >> tap) & 1
	}
	
	out := l.state & 1
	l.state = (l.state >> 1) | (feedback << 7)
	return out
}

func (l *LFSR) Generate(n int) []uint8 {
	seq := make([]uint8, n)

	for i := 0; i < n; i++ {
		seq[i] = l.Step()
	}

	return seq
}


func main() {
	seed := uint8(0b11001010)
	taps := []int{7, 5}

	lfsr := NewLFSR(seed, taps)

	fmt.Println("Gerando 32 bits de keystream:")
	bits := lfsr.Generate(32)

	builder := strings.Builder{}

	for i, bit := range bits {
		builder.WriteByte('0' + bit)

		if (i + 1) % 8 == 0 {
			builder.WriteByte('\n')
		}
	}

	fmt.Println(builder.String())
}