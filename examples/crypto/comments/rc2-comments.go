package main

import (
	"fmt"
)

type RC2 struct {
	subkeys [64]uint16
}

// piTable: Tabela pseudoaleatória usada na expansão da chave.
// Gerada com base nos dígitos de Pi. Essa é a tabela oficial usada por RC2.
var piTable = [256]byte{
	217, 120, 249, 196, 25, 221, 181, 237, 40, 231, 230, 173, 232, 96, 49, 222,
	179, 48, 65, 199, 251, 236, 139, 148, 245, 131, 160, 218, 251, 82, 206, 78,
	43, 250, 126, 186, 26, 187, 234, 242, 47, 238, 122, 169, 104, 121, 145, 21,
	178, 7, 63, 148, 194, 16, 137, 11, 34, 95, 33, 128, 127, 93, 154, 90,
	144, 50, 39, 53, 88, 98, 103, 30, 126, 6, 204, 24, 234, 236, 202, 241,
	18, 214, 227, 20, 101, 190, 70, 97, 122, 166, 148, 132, 123, 223, 203, 224,
	231, 174, 218, 35, 9, 168, 142, 23, 73, 158, 198, 133, 220, 249, 91, 17,
	149, 184, 115, 171, 99, 190, 52, 210, 105, 242, 60, 192, 175, 43, 252, 107,
	5, 211, 38, 141, 250, 13, 140, 153, 233, 212, 153, 197, 1, 193, 73, 209,
	76, 132, 187, 208, 89, 18, 169, 200, 196, 135, 130, 116, 188, 159, 86, 164,
	100, 109, 198, 173, 186, 3, 64, 52, 217, 226, 250, 124, 123, 222, 180, 247,
	244, 197, 225, 211, 138, 225, 114, 167, 40, 249, 193, 171, 239, 35, 114, 15,
	0, 203, 101, 191, 51, 238, 110, 200, 103, 119, 160, 46, 219, 56, 85, 66,
	94, 250, 109, 133, 203, 72, 66, 189, 226, 134, 59, 26, 192, 202, 254, 31,
	157, 130, 206, 247, 121, 245, 0, 70, 66, 205, 201, 164, 39, 184, 144, 90,
	23, 94, 11, 243, 140, 248, 135, 132, 59, 73, 41, 122, 10, 196, 153, 34,
}

// expandKey: Expande a chave de 128 bits para 64 subchaves de 16 bits.
// Parâmetros:
//   - userKey: []byte (128 bits) — chave original.
//   - T: número de bits efetivos da chave (geralmente 64)
// Retorno:
//   - [64]uint16: 64 subchaves de 16 bits.
//
// Descrição:
//   - Expande a chave de 128 bytes usando a tabela IP.
//   - Trunca conforme o parâmetro T (effective key bits)
//   - Constrói 64 palavras de 16 bits usando pares de bytes (little-endian)
func expandKey(userKey []byte, T int) [64]uint16 {
	var expandedKeyBytes [128]byte
	copy(expandedKeyBytes[:], userKey)

	userKeySize := len(userKey)

/*
	Se a chave for menor que 128 bits, expande a chave usando a tabela IP.

	O for vai de t até 128, pois já copiamos os primeiros t bytes da chave para L.
	Suponha que a chave original tem só 8 bytes, então t = 8.
	O for vai de 8 até 128, e L[i - 1] e L[i - t] são acessados.
	Logo, i - t vai de 0 até 120 e i -1 vai de 7 até 127.
	Ao somar ambos valores a cada iteração e forçar o intervalo de 0 a 255 com o operador %,
	estamos garantindo que o índice não saia do intervalo [0, 255] na IP
	Essa tabela foi gerada com valores pseudoaleatórios baseados nos dígitos do número Pi.

	Efetivamente: esse for preenche o restante do array L com bytes pseudoaleatórios extraídos da tabela IP,
	mas de forma determinística, permitindo que o mesmo vetor L seja reconstruído durante a decifragem
	desde que a chave original e a tabela IP sejam idênticas.
*/
	if userKeySize < 128 {
		for i := userKeySize; i < 128; i++ {

			a := expandedKeyBytes[i - 1]
			b := expandedKeyBytes[i - userKeySize]

			sum := int(a) + int(b)

			expandedKeyBytes[i] = piTable[sum % 256]
		}
	}


/*
	effectiveKeyInBits é o número de bits efetivos da chave, geralmente 64 bits.
	effectiveKeyInBytes é um arrendondamento para o múltiplo de 8 mais próximo de T
*/
	effectiveKeyInBits := T
	effectiveKeyInBytes := (T + 7) % 8

/*
	Controle de "bits efetivos" da chave, conhecido como parâmetro T no RFC 2268.

	O RC2 permite definir que apenas os T bits mais significativos da chave devem ser considerados "ativos".

	Etapas:
	1. Calcular o número mínimo de bytes para armazenar os bits.
	   Esse valor é arredondado para cima (ceil), para garantir que todos os bits caibam.

	2. Criar uma máscara para forçar o último byte da chave a obedecer ao número de bits efetivos.
	   - Se effectiveKeyInBits for múltiplo de 8, TM = 0xFF (sem truncamento)
	   - Se effectiveKeyInBits for, por exemplo, 61, então TM = 0x1F (últimos 3 bits desativados)
	   Isso garante que, mesmo com a expansão, a entropia total da chave não exceda os T bits permitidos.
*/
	trailingUnusedBits  := (8 * effectiveKeyInBytes - effectiveKeyInBits) & 7
	mask := byte(0xFF >> trailingUnusedBits )


	expandedKeyBytes[128 - effectiveKeyInBytes] = piTable[expandedKeyBytes[128 - effectiveKeyInBytes] & mask]

	for i := 127 - effectiveKeyInBytes; i >= 0; i-- {
		nextByte := expandedKeyBytes[i + 1]
		offsetByte := expandedKeyBytes[i + effectiveKeyInBytes]
		mixedIndex := nextByte ^ offsetByte
		expandedKeyBytes[i] = piTable[mixedIndex]
	}


	/* RFC 2268 na seção key expansion diz o seguinte:
		Desde que vamos lidar com operaçoes em nível de byte e também em nível de palavra,
		a key será acessada de duas formas:
		K[0], ..., K[63]; cada K[i] é uma palavra de 16-bits
		L[0], ..., L[127]; cada L[i] é um byte
		Onde tanto L quanto K acessam o mesmo array de bytes.

		A seguinte igualdade define a relação entre K e L:
		K[i] = L[2*i] + 256*L[2*i+1]

		 Ou seja, joga o segundo byte para parte alta da palavra (e.g., 0xAB se torna 0xAB00)
		 então faz um OR bit a bit com a parte baixa da palavra (e.g., 0xAB00 | 0x0010 se torna 0xAB10)
	*/
	var expandedKeyWords [64]uint16
	for i := 0; i < 64; i++ {
		lo := uint16(expandedKeyBytes[2*i])
		hi := uint16(expandedKeyBytes[2*i+1]) << 8
		expandedKeyWords[i] = hi | lo
	}

	return expandedKeyWords
}
func (rc2 *RC2) EncryptBlock(block []byte) []byte {
	if len(block) != 8 {
		panic("RC2: o bloco deve ter exatamente 8 bytes")
	}

	R := make([]uint16, 4)
	for i := 0; i < 4; i++ {
		lo := uint16(block[i*2])
		hi := uint16(block[i*2+1])
		R[i] = hi<<8 | lo
	}

	s := rc2.subkeys
	j := 0

	// Fase 1: 5 rodadas
	for i := 0; i < 5; i++ {
		R[0] += R[1] &^ R[3] ^ R[2] & R[3] + s[j]
		R[0] = R[0] << 1 | R[0] >> 15
		j++
		R[1] += R[2] &^ R[0] ^ R[3] & R[0] + s[j]
		R[1] = R[1] << 2 | R[1] >> 14
		j++
		R[2] += R[3] &^ R[1] ^ R[0] & R[1] + s[j]
		R[2] = R[2] << 3 | R[2] >> 13
		j++
		R[3] += R[0] &^ R[2] ^ R[1] & R[2] + s[j]
		R[3] = R[3] << 5 | R[3] >> 11
		j++
	}

	// Mix
	R[0] += s[R[3]&63]
	R[1] += s[R[0]&63]
	R[2] += s[R[1]&63]
	R[3] += s[R[2]&63]

	// Fase 2: 6 rodadas
	for i := 0; i < 6; i++ {
		R[0] += R[1] &^ R[3] ^ R[2] & R[3] + s[j]
		R[0] = R[0] << 1 | R[0] >> 15
		j++
		R[1] += R[2] &^ R[0] ^ R[3] & R[0] + s[j]
		R[1] = R[1] << 2 | R[1] >> 14
		j++
		R[2] += R[3] &^ R[1] ^ R[0] & R[1] + s[j]
		R[2] = R[2] << 3 | R[2] >> 13
		j++
		R[3] += R[0] &^ R[2] ^ R[1] & R[2] + s[j]
		R[3] = R[3] << 5 | R[3] >> 11
		j++
	}

	// Mix
	R[0] += s[R[3]&63]
	R[1] += s[R[0]&63]
	R[2] += s[R[1]&63]
	R[3] += s[R[2]&63]

	// Fase 3: 5 rodadas
	for i := 0; i < 5; i++ {
		R[0] += R[1] &^ R[3] ^ R[2] & R[3] + s[j]
		R[0] = R[0] << 1 | R[0] >> 15
		j++
		R[1] += R[2] &^ R[0] ^ R[3] & R[0] + s[j]
		R[1] = R[1] << 2 | R[1] >> 14
		j++
		R[2] += R[3] &^ R[1] ^ R[0] & R[1] + s[j]
		R[2] = R[2] << 3 | R[2] >> 13
		j++
		R[3] += R[0] &^ R[2] ^ R[1] & R[2] + s[j]
		R[3] = R[3] << 5 | R[3] >> 11
		j++
	}

	out := make([]byte, 8)
	for i := 0; i < 4; i++ {
		out[i*2] = byte(R[i] & 0xFF)
		out[i*2+1] = byte(R[i] >> 8)
	}
	return out
}

func main() {
	key := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF}
	plaintext := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF}

	rc2 := RC2{expandKey(key, 64)}
	ciphertext := rc2.EncryptBlock(plaintext)
	fmt.Printf("Ciphertext: %X\n", ciphertext)
}
