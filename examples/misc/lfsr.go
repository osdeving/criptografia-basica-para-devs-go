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

	// Ataque
	fmt.Println("Ataque ao LFSR:")

	secretSeed := uint8(0b10101010)
	
	alice := NewLFSR(secretSeed, taps)
	bob := alice.Generate(8)

	fmt.Println("Alice gerou 8 bits de keystream:", bob)

	fmt.Println("Eve observa 8 primeiros bits da saída de Alice para Bob tenta recuperar o estado:", bob)

	recovered := solveLFSR2(bob, taps)
	fmt.Printf("Eve recupera o estado: %08b\n", recovered)

	eve := NewLFSR(recovered, taps)
	recreated := eve.Generate(8)

	fmt.Println("Eve recria os 8 bits de keystream:", recreated)
}

// ---------- ATAQUE ----------
func solveLFSR(outputs []uint8, taps []int) uint8 {
	// Cria matriz 8x9 (8 equações, 8 variáveis + 1 resultado)
	matrix := make([][]uint8, 8)
	for i := range matrix {
		matrix[i] = make([]uint8, 9)
	}

	// Para cada bit de saída, monte equações sobre os bits do estado
	for i := 0; i < 8; i++ {
		// Simula 8 steps com incógnitas s0...s7
		for j := 0; j < 8; j++ {
			matrix[i][j] = uint8(((1 << j) >> i) & 1)
		}

		// Aplica taps
		for _, tap := range taps {
			if tap == 7 {
				continue // bit que acabou de sair, já está nas equações
			}
			for j := 0; j < 8; j++ {
				if ((1 << tap) >> j & 1) == 1 {
					matrix[i][j] ^= 1
				}
			}
		}

		// Termo independente = bit de saída
		matrix[i][8] = outputs[i]
	}

	// Aplica eliminação de Gauss em módulo 2 (XOR)
	for i := 0; i < 8; i++ {
		if matrix[i][i] == 0 {
			for j := i + 1; j < 8; j++ {
				if matrix[j][i] == 1 {
					matrix[i], matrix[j] = matrix[j], matrix[i]
					break
				}
			}
		}
		for j := 0; j < 8; j++ {
			if i != j && matrix[j][i] == 1 {
				for k := 0; k < 9; k++ {
					matrix[j][k] ^= matrix[i][k]
				}
			}
		}
	}

	// Monta o estado como uint8
	var recovered uint8 = 0
	for i := 0; i < 8; i++ {
		if matrix[i][8] == 1 {
			recovered |= (1 << i)
		}
	}
	return recovered
}

func solveLFSR2(observed []uint8, taps []int) uint8 {
	// Construção formal
	T := buildTransitionMatrix(taps)
	L := [N]uint8{1, 0, 0, 0, 0, 0, 0, 0}

	var M [N][N]uint8
	for i := 0; i < N; i++ {
		Ti := matrixPow(T, i)
		M[i] = mulRowMat(L, Ti)
	}

	var y [N]uint8
	copy(y[:], observed[:N])
	
	stateVec := solveSystemMod2(M, y)
	return bitsToUint8(stateVec)
}
// Tamanho fixo do LFSR
const N = 8

// Constrói a matriz de transição T a partir dos taps
func buildTransitionMatrix(taps []int) [N][N]uint8 {
	var T [N][N]uint8

	// Shift: linha i copia coluna i+1 (exceto última)
	for i := 0; i < N-1; i++ {
		T[i][i+1] = 1
	}

	// Última linha: feedback linear
	for _, tap := range taps {
		T[N-1][tap] = 1
	}
	return T
}

// Multiplica vetor linha (1xN) por matriz NxN → retorna vetor linha 1xN
func mulRowMat(row [N]uint8, mat [N][N]uint8) [N]uint8 {
	var result [N]uint8
	for j := 0; j < N; j++ {
		for k := 0; k < N; k++ {
			result[j] ^= row[k] * mat[k][j]
		}
	}
	return result
}

// Eleva matriz ao expoente (mod 2)
func matrixPow(mat [N][N]uint8, exp int) [N][N]uint8 {
	// identidade
	var result [N][N]uint8
	for i := 0; i < N; i++ {
		result[i][i] = 1
	}
	for exp > 0 {
		if exp%2 == 1 {
			result = matrixMul(result, mat)
		}
		mat = matrixMul(mat, mat)
		exp /= 2
	}
	return result
}

// Multiplica duas matrizes NxN (mod 2)
func matrixMul(a, b [N][N]uint8) [N][N]uint8 {
	var result [N][N]uint8
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			for k := 0; k < N; k++ {
				result[i][j] ^= a[i][k] * b[k][j]
			}
		}
	}
	return result
}

// Resolve sistema M * x = y (mod 2)
func solveSystemMod2(M [N][N]uint8, y [N]uint8) [N]uint8 {
	// Matriz aumentada
	var mat [N][N + 1]uint8
	for i := 0; i < N; i++ {
		copy(mat[i][:N], M[i][:])
		mat[i][N] = y[i]
	}

	// Eliminação de Gauss
	for i := 0; i < N; i++ {
		if mat[i][i] == 0 {
			for j := i + 1; j < N; j++ {
				if mat[j][i] == 1 {
					mat[i], mat[j] = mat[j], mat[i]
					break
				}
			}
		}
		for j := 0; j < N; j++ {
			if i != j && mat[j][i] == 1 {
				for k := 0; k <= N; k++ {
					mat[j][k] ^= mat[i][k]
				}
			}
		}
	}

	// Extrai solução
	var result [N]uint8
	for i := 0; i < N; i++ {
		result[i] = mat[i][N]
	}
	return result
}

// Converte vetor de bits para uint8
func bitsToUint8(bits [N]uint8) uint8 {
	var val uint8 = 0
	for i := 0; i < N; i++ {
		if bits[i] == 1 {
			val |= 1 << i
		}
	}
	return val
}


/*

Alice tinha um estado secreto de 0b10101010
Ela rodou o LFSR com taps 7 e 5 e gerou 8 bits de keystream: [1 0 1 0 1 0 1 0]
Eve observou os 8 primeiros bits de keystream: [1 0 1 0 1 0 1 0] usada para cibrar a mensagem 0b10101010.
Eve sabia os taps e a estrutura do LFSR e tentou recuperar o estado de Alice.
Com isso, Eve:
- Montou um sistema de equações lineares com os bits observados e os taps conhecidos.
- Montou a matriz de transição T a partir dos taps.
- Montou o sistema M * x = y (mod 2) e resolveu para x.
- Resolveu o sistema linear em Z2 com eliminação de Gauss.
- Obeteve o estado inicial do LFSR
Eve então rodou o LFSR com o estado recuperado e gerou 8 bits de keystream: [1 0 1 0 1 0 1 0] que estava sendo enviado para Bob.

*/

/* 
Alice usa um LFSR com taps fixos e uma semente para gerar a keystream. Ela usa essa keystream para XOR com a mensagem e gerar o ciphertext.
Eve intercepta o ciphertext, e como ela conhece parte do conteúdo esperado da mensagem (por exemplo, "Email: "), ela consegue recuperar parte da keystream.
Com esses bits da keystream, Eve resolve um sistema de equações lineares para recuperar o estado inicial do LFSR, podendo assim decifrar o restante da mensagem.




(a xor b) xor b = a


*/
