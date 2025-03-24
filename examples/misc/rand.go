// go:build ignore
package main

import (
	"fmt"
)

type PRNG struct {
	state uint8
	taps  []int
}

func srand(seed uint8) *PRNG {
	// para simplificar, taps serão sempre 7 e 5
	return &PRNG{state: seed, taps: []int{7, 5}}
}

// gera um bit pseudo-aleatório semelhante a rand.Intn(2) ou rand() % 2 do C
func (p *PRNG) nextBit() uint8 {
	var feedback uint8

	// equivalente a:
	// feedback = 0 XOR bit5 (pega o primeito tap, que sabemos que é 5)
	// feedback = bit5 XOR bit7 (pega o segundo tap, que sabemos que é 7, e faz XOR com o resultado anterior, que sabemos que é bit5)
	for _, tap := range p.taps {
		feedback ^= (p.state >> tap) & 1
	}

	// obtém o bit menos significativo do estado
	out := p.state & 1

	// desloca o estado para a direita e insere o feedback no bit mais significativo
	// p.ex.: 
	// 10000001 >> 1 = 01000000 (desloca o estado para a direita, perdendo o bit menos significativo e ganhando um 0 no bit mais significativo)
	// 00000001 << 7 = 10000000 (o feedback é sempre apenas 0 ou 1 e não precisamos fazer AND com 1, o deslocamento joga esse bit para o bit mais significativo)
	// 01000000 | 10000000 = 11000000 (faz um OR bit a bit entre o estado deslocado e o feedback deslocado, e obtém o novo estado)
	p.state = (p.state >> 1) | (feedback << 7)
	return out
}

func (p *PRNG) rand() uint8 {
	var b uint8

	// obtém 8 bits pseudo-aleatórios
	// o OR é feito sempre entre os bits obtidos e o valor de b, que é inicializado com 0
	// cada bit é deslocado para sua posição correta (0 a 7)
	for i := 0; i < 8; i++ {
		b |= p.nextBit() << i
	}
	return b
}

func main() {
	// seed (semelhante a srand do C)
	prng := srand(0b11001010)

	fmt.Println("Gerando 10 números pseudo-aleatórios com LFSR:")
	for i := 0; i < 10; i++ {
		fmt.Printf("rand() = %3d (0b%08b)\n", prng.rand(), prng.rand())
	}
}
