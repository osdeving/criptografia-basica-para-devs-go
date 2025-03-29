// fitting_gen.go
package main

import (
	"fmt"
	"os"
	"strings"
)

type LFSRSegment struct {
	Seed   uint64
	Taps   []int
	Length int
}

func bytesToBits(data []byte) []uint8 {
	var bits []uint8
	for _, b := range data {
		for i := 7; i >= 0; i-- {
			bits = append(bits, (b>>i)&1)
		}
	}
	return bits
}

func simulateLFSR(seed uint64, taps []int, n int, size int) []uint8 {
	state := seed
	var out []uint8
	for i := 0; i < n; i++ {
		bit := state & 1
		out = append(out, uint8(bit))
		var feedback uint64
		for _, t := range taps {
			feedback ^= (state >> t) & 1
		}
		state = (state >> 1) | (feedback << (size - 1))
	}
	return out
}

func matchBits(a, b []uint8) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func isAllZero(bits []uint8) bool {
	for _, b := range bits {
		if b != 0 {
			return false
		}
	}
	return true
}

func fitLFSR(bits []uint8) []LFSRSegment {
	var segments []LFSRSegment
	pos := 0

	trySizes := []int{16, 12, 8, 6, 4}
	for pos < len(bits) {
		fitted := false
		for _, size := range trySizes {
			if pos+size > len(bits) {
				continue
			}
			window := bits[pos : pos+size]

			// Atalho para zero puro
			if isAllZero(window) {
				segments = append(segments, LFSRSegment{
					Seed:   0,
					Taps:   []int{0, 1},
					Length: size,
				})
				fmt.Printf("⚠️  Segmento zero puro no offset %d: seed=000...000 taps=[0,1] size=%d\n", pos, size)
				pos += size
				fitted = true
				break
			}

			for seed := uint64(1); seed < (1 << size); seed++ {
				tapCombos := generateTapCombos(size)
				for _, taps := range tapCombos {
					gen := simulateLFSR(seed, taps, size, size)
					if matchBits(gen, window) {
						segments = append(segments, LFSRSegment{
							Seed:   seed,
							Taps:   taps,
							Length: size,
						})
						fmt.Printf("✔️  Segmento encontrado no offset %d: seed=%0*b taps=%v size=%d\n", pos, size, seed, taps, size)
						pos += size
						fitted = true
						goto nextBlock
					}
				}
			}
		nextBlock:
			if fitted {
				break
			}
		}
		if !fitted {
			fmt.Printf("❌ Falha no offset %d com janela: %v\n", pos, bits[pos:pos+16])
			panic(fmt.Sprintf("Não foi possível encaixar LFSR no offset %d", pos))
		}
	}
	return segments
}

func generateTapCombos(size int) [][]int {
	var combos [][]int
	for t1 := 0; t1 < size; t1++ {
		for t2 := t1 + 1; t2 < size; t2++ {
			combos = append(combos, []int{t1, t2})
			for t3 := t2 + 1; t3 < size; t3++ {
				combos = append(combos, []int{t1, t2, t3})
			}
		}
	}
	return combos
}

func writeGoFile(segments []LFSRSegment, output string) {
	var b strings.Builder
	b.WriteString("package main\n\n")
	b.WriteString("var KeyStream = []LFSRSegment{\n")
	for _, seg := range segments {
		b.WriteString(fmt.Sprintf("  {Seed: 0b%0*b, Taps: []int{%d, %d}, Length: %d},\n",
			seg.Length, seg.Seed, seg.Taps[0], seg.Taps[1], seg.Length))
	}
	b.WriteString("}\n")
	os.WriteFile(output, []byte(b.String()), 0644)
}

func main() {
	data, err := os.ReadFile("matrixmsg")
	if err != nil {
		panic(err)
	}
	bits := bytesToBits(data)
	segments := fitLFSR(bits)
	writeGoFile(segments, "matrix_data.go")
	fmt.Println("✅ Gerado com sucesso: matrix_data.go")
}
