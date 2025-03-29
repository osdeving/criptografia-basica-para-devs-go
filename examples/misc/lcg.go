package main

/*
#include <stdlib.h>
#include <time.h>
*/
import "C"

import (
	"fmt"
	"math/big"
)

type PRNG struct {
    state *big.Int // Estado atual S_i
    A     *big.Int // Multiplicador
    B     *big.Int // Incremento
    M     *big.Int // Módulo
}

func (p *PRNG) rand() *big.Int {
    p.state.Mul(p.A, p.state)      // A * S_i
    p.state.Add(p.state, p.B)      // A * S_i + B
    p.state.Mod(p.state, p.M)      // (A * S_i + B) mod m
    return new(big.Int).Set(p.state)
}

func seed(A, B, M, s *big.Int) *PRNG {
    return &PRNG{A: A, B: B, M: M, state: s}
}

func main() {
	A := big.NewInt(1103515245)
	B := big.NewInt(12345)
	M := big.NewInt(1 << 32)
	state := big.NewInt(1) // <- igual ao srand(1)

	prng := PRNG{state, A, B, M}
	C.srand(1)

	fmt.Println("Comparando C.rand() e LCG manual:")
	for i := 0; i < 10; i++ {
		r := prng.rand()
		fmt.Printf("C: %5d | Go: %5d\n", C.rand(), r.Int64())
	}
}
