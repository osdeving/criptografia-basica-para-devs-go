/*
	Estruturas usadas:

	Permutações (bit a bit, via tabela)
		- IP    (64 -> 64):  Permutação inicial
		- IPInv (64 -> 64):  Permutação final
		- PC1   (64 -> 56):  Permuted Choice 1
		- PC2   (56 -> 48):  Permuted Choice 2
		- EBox  (32 -> 48):  Expansão da parte direita
		- PBox  (32 -> 32):  Permutação após S-boxes

	Substituções Não Lineares (S-boxes)
		- SBoxes  (8 * 64 entradas):  Cada S-box mapeia 6 bits para 4 bits

	Controle de rotações no key schedule
		- Rotations (16 elementos):  Número de rotações para cada rodada

*/

package main

import (
	"fmt"
)

const (
	BlockSize = 8
	KeySize   = 8
	NumRounds = 16
)

type DES struct {
	key     [KeySize]byte
	subkeys [NumRounds][6]byte
}

// Tabela de permutação inicial
// Veja a explicação detalhada no comentário da função permute
var IP = [64]uint8{
	58, 50, 42, 34, 26, 18, 10, 2,
	60, 52, 44, 36, 28, 20, 12, 4,
	62, 54, 46, 38, 30, 22, 14, 6,
	64, 56, 48, 40, 32, 24, 16, 8,
	57, 49, 41, 33, 25, 17, 9, 1,
	59, 51, 43, 35, 27, 19, 11, 3,
	61, 53, 45, 37, 29, 21, 13, 5,
	63, 55, 47, 39, 31, 23, 15, 7,
}

// Tabela de permutação final, desfaz a permutação feita com IP
// Veja a explicação detalhada no comentário da função permute
var IPInv = [64]uint8{
	40, 8, 48, 16, 56, 24, 64, 32,
	39, 7, 47, 15, 55, 23, 63, 31,
	38, 6, 46, 14, 54, 22, 62, 30,
	37, 5, 45, 13, 53, 21, 61, 29,
	36, 4, 44, 12, 52, 20, 60, 28,
	35, 3, 43, 11, 51, 19, 59, 27,
	34, 2, 42, 10, 50, 18, 58, 26,
	33, 1, 41, 9, 49, 17, 57, 25,
}

/*
E (EBox): Tabela de expansão usada na função f

Descrição:
  - Expande um bloco de 32 bits para 48 bits, duplicando alguns bits conforme a tabela.
  - É aplicada sobre a metade direita (R) do bloco durante cada rodada.
  - A expansão cria sobreposição entre blocos de 4 bits vizinhos, o que aumenta a difusão e permite mistura com a subchave de 48 bits via XOR.

Formato:
  - Entrada: 32 bits
  - Tabela: 48 posições, com índices de bits da entrada no intervalo [1,32]
  - Saída: 48 bits

Exemplo simplificado (com 6 bits para 8 bits):

	Entrada: [A, B, C, D, E, F]
	Tabela: [6, 1, 2, 3, 4, 5, 4, 5]
	Saída:  [F, A, B, C, D, E, D, E]  (alguns bits se repetem)

Obs: veja mais detalhes na descrição da função permute
*/
var E = [48]uint8{
	32, 1, 2, 3, 4, 5,
	4, 5, 6, 7, 8, 9,
	8, 9, 10, 11, 12, 13,
	12, 13, 14, 15, 16, 17,
	16, 17, 18, 19, 20, 21,
	20, 21, 22, 23, 24, 25,
	24, 25, 26, 27, 28, 29,
	28, 29, 30, 31, 32, 1,
}

/*
P (PBox): Permutação aplicada após as substituições nas S-boxes, dentro da função f

Descrição:
  - Permuta os 32 bits resultantes da saída das 8 S-boxes.
  - A permutação é fixa e tem como objetivo espalhar (difusão) os bits substituídos.
  - Essa difusão aumenta a complexidade da relação entre o plaintext e o ciphertext.

Formato:
  - Entrada: 32 bits
  - Tabela: 32 posições, com índices no intervalo [1,32]
  - Saída: 32 bits permutados
*/
var P = [32]uint8{
	16, 7, 20, 21,
	29, 12, 28, 17,
	1, 15, 23, 26,
	5, 18, 31, 10,
	2, 8, 24, 14,
	32, 27, 3, 9,
	19, 13, 30, 6,
	22, 11, 4, 25,
}

/*
PC1 (Permuted Choice 1): Seleção inicial de bits da chave

Descrição:
  - Aplica a primeira permutação sobre a chave de 64 bits.
  - Remove os 8 bits de paridade (1 por byte), resultando em 56 bits úteis.
  - Os bits resultantes são usados como base para o key schedule (geração das subchaves).

Formato:
  - Entrada: 64 bits (com paridade)
  - Tabela: 56 posições, com índices no intervalo [1,64]
  - Saída: 56 bits (sem paridade), divididos em blocos C e D
*/
var PC1 = [56]uint8{
	57, 49, 41, 33, 25, 17, 9,
	1, 58, 50, 42, 34, 26, 18,
	10, 2, 59, 51, 43, 35, 27,
	19, 11, 3, 60, 52, 44, 36,
	63, 55, 47, 39, 31, 23, 15,
	7, 62, 54, 46, 38, 30, 22,
	14, 6, 61, 53, 45, 37, 29,
	21, 13, 5, 28, 20, 12, 4,
}

/*
PC2 (Permuted Choice 2): Seleção e permutação final dos bits da subchave

Descrição:
  - Aplica uma permutação sobre os blocos C e D (já rotacionados), que juntos têm 56 bits.
  - Seleciona 48 bits específicos e os reordena para formar a subchave da rodada.
  - Essa subchave é usada no XOR com a expansão da metade direita (R) dentro da função f.

Formato:
  - Entrada: 56 bits (C || D)
  - Tabela: 48 posições, com índices no intervalo [1,56]
  - Saída: 48 bits (subchave da rodada)
*/
var PC2 = [48]uint8{
	14, 17, 11, 24, 1, 5,
	3, 28, 15, 6, 21, 10,
	23, 19, 12, 4, 26, 8,
	16, 7, 27, 20, 13, 2,
	41, 52, 31, 37, 47, 55,
	30, 40, 51, 45, 33, 48,
	44, 49, 39, 56, 34, 53,
	46, 42, 50, 36, 29, 32,
}

/*
S (SBoxes): Substituições não lineares aplicadas na função f

Descrição:
  - 8 S-boxes, cada uma com 64 entradas.
  - Cada S-box recebe 6 bits de entrada e gera 4 bits de saída.
  - A entrada de 48 bits (resultado da expansão + XOR com subchave) é dividida em 8 blocos de 6 bits.
  - Cada bloco de 6 bits é processado por uma das 8 S-boxes, produzindo 32 bits no total.

Endereçamento:
  - Os bits de entrada são indexados da seguinte forma:
    bits 1 e 6: formam o número da linha (entre 0 e 3)
    bits 2 a 5: formam o número da coluna (entre 0 e 15)
  - O valor de S[i][linha * 16 + coluna] fornece os 4 bits de saída da S-box i.

Formato:
  - Entrada total: 48 bits (8 blocos de 6 bits)
  - Saída total:   32 bits (8 blocos de 4 bits)
  - Cada S-box tem 64 entradas (4 linhas × 16 colunas)

Objetivo:
  - Introduzir não linearidade na cifra, essencial para segurança contra ataques lineares e diferenciais.
*/
var S = [8][64]uint8{
	{ // S1
		14, 4, 13, 1, 2, 15, 11, 8, 3, 10, 6, 12, 5, 9, 0, 7,
		0, 15, 7, 4, 14, 2, 13, 1, 10, 6, 12, 11, 9, 5, 3, 8,
		4, 1, 14, 8, 13, 6, 2, 11, 15, 12, 9, 7, 3, 10, 5, 0,
		15, 12, 8, 2, 4, 9, 1, 7, 5, 11, 3, 14, 10, 0, 6, 13,
	},
	{ // S2
		15, 1, 8, 14, 6, 11, 3, 4, 9, 7, 2, 13, 12, 0, 5, 10,
		3, 13, 4, 7, 15, 2, 8, 14, 12, 0, 1, 10, 6, 9, 11, 5,
		0, 14, 7, 11, 10, 4, 13, 1, 5, 8, 12, 6, 9, 3, 2, 15,
		13, 8, 10, 1, 3, 15, 4, 2, 11, 6, 7, 12, 0, 5, 14, 9,
	},
	{ // S3
		10, 0, 9, 14, 6, 3, 15, 5, 1, 13, 12, 7, 11, 4, 2, 8,
		13, 7, 0, 9, 3, 4, 6, 10, 2, 8, 5, 14, 12, 11, 15, 1,
		13, 6, 4, 9, 8, 15, 3, 0, 11, 1, 2, 12, 5, 10, 14, 7,
		1, 10, 13, 0, 6, 9, 8, 7, 4, 15, 14, 3, 11, 5, 2, 12,
	},
	{ // S4
		7, 13, 14, 3, 0, 6, 9, 10, 1, 2, 8, 5, 11, 12, 4, 15,
		13, 8, 11, 5, 6, 15, 0, 3, 4, 7, 2, 12, 1, 10, 14, 9,
		10, 6, 9, 0, 12, 11, 7, 13, 15, 1, 3, 14, 5, 2, 8, 4,
		3, 15, 0, 6, 10, 1, 13, 8, 9, 4, 5, 11, 12, 7, 2, 14,
	},
	{ // S5
		2, 12, 4, 1, 7, 10, 11, 6, 8, 5, 3, 15, 13, 0, 14, 9,
		14, 11, 2, 12, 4, 7, 13, 1, 5, 0, 15, 10, 3, 9, 8, 6,
		4, 2, 1, 11, 10, 13, 7, 8, 15, 9, 12, 5, 6, 3, 0, 14,
		11, 8, 12, 7, 1, 14, 2, 13, 6, 15, 0, 9, 10, 4, 5, 3,
	},
	{ // S6
		12, 1, 10, 15, 9, 2, 6, 8, 0, 13, 3, 4, 14, 7, 5, 11,
		10, 15, 4, 2, 7, 12, 9, 5, 6, 1, 13, 14, 0, 11, 3, 8,
		9, 14, 15, 5, 2, 8, 12, 3, 7, 0, 4, 10, 1, 13, 11, 6,
		4, 3, 2, 12, 9, 5, 15, 10, 11, 14, 1, 7, 6, 0, 8, 13,
	},
	{ // S7
		4, 11, 2, 14, 15, 0, 8, 13, 3, 12, 9, 7, 5, 10, 6, 1,
		13, 0, 11, 7, 4, 9, 1, 10, 14, 3, 5, 12, 2, 15, 8, 6,
		1, 4, 11, 13, 12, 3, 7, 14, 10, 15, 6, 8, 0, 5, 9, 2,
		6, 11, 13, 8, 1, 4, 10, 7, 9, 5, 0, 15, 14, 2, 3, 12,
	},
	{ // S8
		13, 2, 8, 4, 6, 15, 11, 1, 10, 9, 3, 14, 5, 0, 12, 7,
		1, 15, 13, 8, 10, 3, 7, 4, 12, 5, 6, 11, 0, 14, 9, 2,
		7, 11, 4, 1, 9, 12, 14, 2, 0, 6, 10, 13, 15, 3, 5, 8,
		2, 1, 14, 7, 4, 10, 8, 13, 15, 12, 9, 0, 3, 5, 6, 11,
	},
}

/*
Rotations (16 elementos):  Número de rotações para cada rodada

Descrição:

	No DES, a chave principal de 64 bits é reduzida para 56 bits (removendo bits de paridade) usando a tabela PC1.
	Esses 56 bits são divididos em dois blocos de 28 bits: C e D.

	A cada rodada, esses blocos são rotacionados à esquerda de forma circular.
	O número de rotações varia conforme a rodada, e está definido na tabela Rotations.

	Após rotacionar, os blocos C e D são concatenados (C || D) e passam pela tabela PC2 (é permutado com ela) para gerar uma subchave de 48 bits que será usada na rodada correspondente da cifra.

Quando ocorre:

	As rotações ocorrem antes da geração de cada uma das 16 subchaves, ou seja, são aplicadas 16 vezes.

Quantas rotações:

	A quantidade de bits a rotacionar é especificada para cada rodada:

	Rodada  1:               1 rotação
	Rodada  2:               1 rotação
	Rodada  3 até Rodada 8:  2 rotações por rodada
	Rodada  9:               1 rotação
	Rodada 10 até 15:        2 rotações por rodada
	Rodada 16:               1 rotação

Exemplo com 6 bits em vez de 28 (para simplificar a visualização):

		Rotação 1 à esquerda:
	 		Entrada: [A, B, C, D, E, F]   (6 bits, cada letra é {0,1})
			Saída:   [B, C, D, E, F, A]   (6 bits, rotacionou 1 à esquerda)

		Rotação 2 à esquerda:
			Entrada: [A, B, C, D, E, F]   (6 bits, cada letra é {0,1})
			Saída:   [C, D, E, F, A, B]   (6 bits, rotacionou 2 à esquerda)
*/
var Rotations = [16]uint8{
	1, 1, 2, 2, 2, 2, 2, 2,
	1, 2, 2, 2, 2, 2, 2, 1,
}

/*
New: Inicializa uma nova instância da cifra DES com base em uma chave de 64 bits.

Parâmetros:
  - key: [8]byte (64 bits) contendo a chave original com bits de paridade.

Retorna:
  - *DES: ponteiro para a instância da cifra DES com subchaves geradas.
  - error: erro caso a geração de subchaves falhe (ex: tamanho incorreto após permutação).

Etapas:
 1. Aplica a permutação PC1 na chave original descartando os 8 bits de paridade.
 2. Realiza rotações controladas sobre os blocos C e D (28 bits cada) com 16 rodadas.
 3. Em cada rodada, gera uma subchave de 48 bits aplicando PC2 sobre (C || D).
 4. Armazena todas as 16 subchaves na estrutura DES para uso nas rodadas da cifra.

Obs:

	A função assume que a chave já está no formato exigido pelo padrão DES (8 bytes).
*/
func New(key [KeySize]byte) (*DES, error) {
	d := &DES{}

	err := d.generateSubkeys(key)
	if err != nil {
		return nil, err
	}
	return d, nil
}

/*
		Essa função realiza a operação de expansão e permutação.

		Explicação:
			Ela recebe um slice de N bytes e uma tabela de permutação de N bytes e retorna um slice de N bytes.

			A tabela é usada para determinar a ordem dos bits no slice de entrada e no slice de saída.
			podemos ter valores repetidos, o que significa que a mesma posição de bit pode aparecer em várias posições na saída.
			Isso permite expandir a saída para um tamanho maior do que a entrada.

		Exemplos:
			Simples sem expansão:
				Entrada: [A, B, C, D] (4 bits, cada letra equivale a {0,1})
				Tabela: [3, 1, 2, 4] (posições de bits)
				Saída: [C, A, B, D] (4 bits, cada letra equivale a {0,1})

			Com expansão:
				Entrada: [A, B, C, D] (4 bits, cada letra equivale a {0,1})
				Tabela: [3, 1, 2, 4, 1, 2] (posições de bits)
				Saída: [C, A, B, D, A, B] (6 bits, cada letra equivale a {0,1}, repete o A e o B)

		Restrições:
			Os valores na tabela devem estar no intervalo [1, N], onde N é o tamanho do slice de entrada.
			Isso garante que não tente acessar posições inválidas no slice de entrada.

		Funções Auxiliares:
			getBit: Essa função recebe um slice de bytes e uma posição (0-indexed) e retorna o bit correspondente.
			setBit: Essa função recebe um slice de bytes, uma posição (0-indexed) e um valor (0 ou 1) e "seta" o bit correspondente.

		+----------------+-----------------+-------------------+------------------+------------+--------------------------------------+
	 	| Nome da tabela | Bits de entrada | Tamanho da tabela | Bits de saída    | Tipo       | Uso                                  |
	 	+----------------+-----------------+-------------------+------------------+------------+--------------------------------------+
	 	|      IP        |       64        |          64       |        64        | permutação | Permutação inicial                   |
	 	+----------------+-----------------+-------------------+------------------+------------+--------------------------------------+
		|      IPInv     |       64        |          64       |        64        | permutação | Permutação final                     |
		+----------------+-----------------+-------------------+------------------+------------+--------------------------------------+
		|      PC1       |       64        |          56       |        56        | seleção    | Redução da chave (remove paridade)   |
		+----------------+-----------------+-------------------+------------------+------------+--------------------------------------+
		|      PC2       |       56        |          48       |        48        | seleção    | Geração das subchaves                |
		+----------------+-----------------+-------------------+------------------+------------+--------------------------------------+
		|      E (EBox)  |       32        |          48       |        48        | expansão   | Parte da função f                    |
		+----------------+-----------------+-------------------+------------------+------------+--------------------------------------+
		|      P (PBox)  |       32        |          32       |        32        | permutação | Após as S-boxes, dentro de f         |
		+----------------+-----------------+-------------------+------------------+------------+--------------------------------------+
*/
func permute(input, table []byte) []byte {
	n := len(table)
	output := make([]byte, (n+7)/8)

	for i, pos := range table {
		bit := getBit(input, int(pos)-1)
		setBit(output, i, bit)
	}
	return output
}

func getBit(data []byte, pos int) byte {
	byteIndex := pos / 8
	bitIndex := 7 - (pos % 8)
	return (data[byteIndex] >> bitIndex) & 1
}

func setBit(data []byte, pos int, value byte) {
	byteIndex := pos / 8
	bitIndex := 7 - (pos % 8)

	if value == 1 {
		data[byteIndex] |= (1 << bitIndex)
	} else {
		data[byteIndex] &^= (1 << bitIndex)
	}
}

/*
f: Função usada em cada rodada do DES

Descrição:
  - Recebe os 32 bits da metade direita R (4 bytes) e a subchave da rodada (48 bits = 6 bytes).
  - Aplica a expansão E para transformar R em 48 bits.
  - Realiza XOR entre o resultado e a subchave.
  - Divide em 8 blocos de 6 bits e aplica as 8 S-boxes.
  - A saída (32 bits) é permutada com a tabela P e retornada.

Etapas:

	R (32) -> expansão E (48) -> XOR com subchave -> S-boxes (48 -> 32) -> permutação P (32) -> saída
*/
func f(R []byte, subkey []byte) []byte {
	// Etapa 1 Expande R de 32 para 48 bits
	R48 := permute(R, E[:]) // 6 bytes

	// Etapa 2 XOR com subchave
	x := xor(R48, subkey)

	// Etapa 3 Aplicar S-boxes
	sboxOut := make([]byte, 4) // 32 bits = 4 bytes
	for i := 0; i < 8; i++ {
		offset := i * 6
		val := uint8(0)
		for j := 0; j < 6; j++ {
			bit := getBit(x, offset+j)
			val |= bit << (5 - j)
		}

		row := ((val & 0x20) >> 4) | (val & 0x01) // bits 1 e 6
		col := (val >> 1) & 0x0F                  // bits 2 a 5
		sboxVal := S[i][row*16+col]

		// Inserir 4 bits resultantes no sboxOut
		for j := 0; j < 4; j++ {
			bit := (sboxVal >> (3 - j)) & 1
			setBit(sboxOut, i*4+j, bit)
		}
	}

	// Etapa 4 Permutação final (PBox)
	return permute(sboxOut, P[:])
}

// Realiza XOR bit a bit entre dois slices de mesmo comprimento.
func xor(a, b []byte) []byte {
	if len(a) != len(b) {
		panic("tamanhos diferentes no xor")
	}
	out := make([]byte, len(a))
	for i := range a {
		out[i] = a[i] ^ b[i]
	}
	return out
}

/*
generateSubkeys: Gera as 16 subchaves de 48 bits a partir da chave original de 64 bits.

Descrição:
  - Aplica PC1 à chave original (64 bits) para obter 56 bits sem paridade.
  - Converte esses bits em slices puros de bits (0 ou 1).
  - Divide em dois blocos de 28 bits: C (esquerda) e D (direita).
  - Executa rotações circulares definidas pela tabela Rotations.
  - Concatena C || D e aplica PC2 para gerar a subchave de 48 bits da rodada.
  - Armazena cada subchave no campo `d.subkeys`.

Parâmetros:
  - key: [8]byte — chave original de 64 bits (com bits de paridade).

Retorno:
  - error: retorna erro se houver falha inesperada (não esperado com chave válida).
*/
func (d *DES) generateSubkeys(key [KeySize]byte) error {
	// Aplica PC1 -> 56 bits
	key56 := permute(key[:], PC1[:]) // 7 bytes = 56 bits

	// Converte key56 para slice de bits (cada bit é 0 ou 1)
	bits := bytesToBits(key56)

	// Divide bits em C e D (28 bits cada)
	C := make([]byte, 28)
	D := make([]byte, 28)
	copy(C, bits[:28])
	copy(D, bits[28:])

	for i := 0; i < NumRounds; i++ {
		// Rotações circulares à esquerda
		C = rotateLeft(C, int(Rotations[i]))
		D = rotateLeft(D, int(Rotations[i]))

		// Concatena C || D
		CD := append(C, D...)

		// Aplica PC2 -> 48 bits
		subkeyBits := permuteBits(CD, PC2[:]) // 48 bits
		subkey := bitsToBytes(subkeyBits)     // 6 bytes

		copy(d.subkeys[i][:], subkey)
	}
	return nil
}

/*
bytesToBits: Converte um slice de bytes em um slice de bits.

Descrição:
  - Cada byte é decomposto em seus 8 bits (MSB primeiro).
  - O resultado é um slice com bits individuais representados como 0 ou 1 (byte).

Parâmetros:
  - data: slice de bytes.

Retorno:
  - Slice de bits com tamanho len(data) * 8.
*/
func bytesToBits(data []byte) []byte {
	bits := make([]byte, len(data)*8)
	for i := 0; i < len(bits); i++ {
		bits[i] = getBit(data, i)
	}
	return bits
}

/*
bitsToBytes: Converte um slice de bits (0 ou 1) em um slice de bytes.

Descrição:
  - Agrupa os bits a cada 8 posições para formar bytes.
  - Bits são tratados no formato MSB primeiro.
  - Se a quantidade de bits não for múltiplo de 8, o último byte é preenchido à esquerda.

Parâmetros:
  - bits: slice de bits (0 ou 1).

Retorno:
  - Slice de bytes correspondente.
*/
func bitsToBytes(bits []byte) []byte {
	n := (len(bits) + 7) / 8
	data := make([]byte, n)
	for i := 0; i < len(bits); i++ {
		setBit(data, i, bits[i])
	}
	return data
}

/*
permuteBits: Realiza uma permutação direta sobre um slice de bits.

Descrição:
  - Usa a tabela de permutação (1-based) para reorganizar os bits do slice.
  - Cada valor da tabela especifica a posição do bit original a ser copiado.

Parâmetros:
  - input: slice de bits (0 ou 1).
  - table: tabela de permutação (valores entre 1 e len(input)).

Retorno:
  - Slice de bits permutado.
*/
func permuteBits(input []byte, table []uint8) []byte {
	out := make([]byte, len(table))
	for i, pos := range table {
		out[i] = input[pos-1]
	}
	return out
}

/*
rotateLeft: Executa uma rotação circular à esquerda em um slice de bits.

Descrição:
  - Roda n posições os bits de entrada para a esquerda.
  - Bits deslocados à esquerda são recolocados no fim do slice.

Parâmetros:
  - input: slice de bits (0 ou 1).
  - n: número de posições a rotacionar.

Retorno:
  - Novo slice com bits rotacionados.
*/
func rotateLeft(input []byte, n int) []byte {
	n %= len(input)
	return append(input[n:], input[:n]...)
}

/*
Encrypt: Cifra um bloco de 64 bits usando o algoritmo DES.

Descrição:
  - Aplica a permutação inicial (IP) sobre o bloco de entrada.
  - Divide o bloco em L0 e R0 (32 bits cada).
  - Executa 16 rodadas de Feistel, onde:
  - Li = Ri-1
  - Ri = Li-1 XOR f(Ri-1, subkey_i)
  - Após as rodadas, concatena R16 || L16 (note a inversão).
  - Aplica a permutação final (IPInv) para gerar o bloco cifrado.

Parâmetros:
  - plaintext: [8]byte (64 bits) — bloco de entrada.

Retorno:
  - [8]byte: bloco cifrado.
  - error: reservado para casos futuros (atualmente sempre nil).
*/
func (d *DES) Encrypt(plaintext [BlockSize]byte) ([BlockSize]byte, error) {
	fmt.Printf("Plaintext (hex): %x\n", plaintext)
	fmt.Printf("Plaintext (bits): %08b\n", plaintext)

	permuted := permute(plaintext[:], IP[:])
	fmt.Printf("After IP: %x\n", permuted)

	L, R := permuted[:4], permuted[4:]
	fmt.Printf("L0: %x\n", L)
	fmt.Printf("R0: %x\n", R)

	for i := 0; i < NumRounds; i++ {
		fOut := f(R, d.subkeys[i][:])
		fmt.Printf("Round %2d - Subkey: %x\n", i+1, d.subkeys[i])
		fmt.Printf("Round %2d - f(R, K): %x\n", i+1, fOut)

		Li := R
		Ri := xor(L, fOut)
		fmt.Printf("Round %2d - L: %x\n", i+1, Li)
		fmt.Printf("Round %2d - R: %x\n", i+1, Ri)
		L, R = Li, Ri
	}

	fmt.Printf("Preoutput (R || L): %x%x\n", R, L)
	preoutput := append(R, L...)
	finalOutput := permute(preoutput, IPInv[:])
	fmt.Printf("Final Output: %x\n", finalOutput)

	var out [BlockSize]byte
	copy(out[:], finalOutput)
	return out, nil
}

/*
	Decrypt: Decifra um bloco de 64 bits usando o algoritmo DES.

	Descrição:
		- Aplica a permutação inicial (IP) sobre o bloco cifrado.
		- Divide em L0 e R0 (32 bits cada).
		- Executa 16 rodadas de Feistel no sentido inverso (chaves na ordem reversa):
			- Li = Ri-1
			- Ri = Li-1 XOR f(Ri-1, subkey_i)
		- Após as rodadas, concatena R16 || L16 (inversão).
		- Aplica a permutação final (IPInv) para obter o plaintext original.

	Parâmetros:
		- ciphertext: [8]byte (64 bits) — bloco cifrado.

	Retorno:
		- [8]byte: bloco decifrado.
		- error: reservado para casos futuros (atualmente sempre nil).
*/

func (d *DES) Decrypt(ciphertext [BlockSize]byte) ([BlockSize]byte, error) {
	permuted := permute(ciphertext[:], IP[:])
	L, R := permuted[:4], permuted[4:]

	for i := NumRounds - 1; i >= 0; i-- {
		Li := R
		Ri := xor(L, f(R, d.subkeys[i][:]))
		L, R = Li, Ri
	}

	preoutput := append(R, L...)
	finalOutput := permute(preoutput, IPInv[:])
	var out [BlockSize]byte
	copy(out[:], finalOutput)
	return out, nil
}

/*
	main: Função de teste e demonstração (driver) do algoritmo DES.

	Descrição:
		- Define chave e bloco de teste conforme especificação do padrão FIPS PUB 81.
		- Instancia o algoritmo DES com a chave dada.
		- Cifra o bloco de 64 bits e imprime o resultado em hexadecimal.
		- Decifra o resultado e verifica se o texto recuperado é idêntico ao original.
		- Se tudo estiver correto, imprime sucesso.

	Teste FIPS:
		- Chave:    133457799BBCDFF1
		- Bloco:    0123456789ABCDEF
		- Esperado: 85E813540F0AB405 (cifrado)

	Objetivo:
		- Verificar se a implementação segue o comportamento padronizado do DES.
*/

func main() {
	// Chave e bloco de teste (padrão FIPS - Federal Information Processing Standards)
	var chave = [8]byte{0x13, 0x34, 0x57, 0x79, 0x9B, 0xBC, 0xDF, 0xF1}
	var bloco = [8]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF}

	// Instancia o DES com a chave
	des, err := New(chave)
	if err != nil {
		panic(err)
	}

	// Cifra o bloco
	ciphertext, err := des.Encrypt(bloco)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Texto cifrado:   %x\n", ciphertext)

	// Decifra o resultado
	plaintext, err := des.Decrypt(ciphertext)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Texto decifrado: %x\n", plaintext)

	// Verificação de correção
	if plaintext != bloco {
		panic("Erro no processo de decifragem: valores não coincidem.")
	}
	fmt.Println("Sucesso! Decifragem correta.")
}

/*
================================================================================
Referências
================================================================================

(conceitos e definições)
[1] C. E. Shannon, "Communication Theory of Secrecy Systems," Bell System Technical Journal, vol. 28, pp. 656–715, Oct. 1949.

(algoritmo DES)
[2] National Bureau of Standards (NBS), "Data Encryption Standard (DES)," FIPS PUB 46, U.S. Department of Commerce, Jan. 1977.

(topzera demais!)
[3] C. Paar and J. Pelzl, *Understanding Cryptography: A Textbook for Students and Practitioners*, Springer, 2010.

(topzera demais!)
[4] B. Schneier, *Applied Cryptography: Protocols, Algorithms, and Source Code in C*, 2nd ed., Wiley, 1996. (Edição comemorativa de 20 anos disponível em 2015).

[5] https://chatgpt.com/ (lol)
================================================================================
*/
