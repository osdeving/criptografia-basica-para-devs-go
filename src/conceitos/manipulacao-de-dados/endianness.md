# Endianness - Organização de Bytes na Memória

## O que é Endianness?

**Endianness** refere-se à ordem na qual os bytes de uma palavra multibyte (como `int32`, `uint64`, etc.) são armazenados na memória. Essa ordem pode variar entre arquiteturas e afeta diretamente a forma como os dados binários são interpretados em operações de leitura, escrita, serialização e comunicação entre sistemas.

Existem dois modelos principais: **Little-endian** e **Big-endian**.

## Little-endian

Nesse modelo, o **byte menos significativo** (*LSB – Least Significant Byte*) é armazenado primeiro, no menor endereço de memória. Suponha que o valor `0x12345678` (em hexadecimal) seja armazenado em 4 bytes. A ordem de armazenamento será:

| Endereço | Byte  |
|----------|-------|
| 0x1000   | 0x78  |
| 0x1001   | 0x56  |
| 0x1002   | 0x34  |
| 0x1003   | 0x12  |


## Big-endian

Neste modelo, o **byte mais significativo** (*MSB – Most Significant Byte*) é armazenado primeiro. Para o mesmo valor `0x12345678`, teríamos:

| Endereço | Byte  |
|----------|-------|
| 0x1000   | 0x12  |
| 0x1001   | 0x34  |
| 0x1002   | 0x56  |
| 0x1003   | 0x78  |


## Endianess na Prática

A ordem dos bytes na memória varia entre arquiteturas. A seguir, veremos como isso afeta a manipulação de dados binários e como determinar a ordem de bytes de um sistema.

### Detectando o endianness do sistema

```go
package main

import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

func main() {
	var i int32 = 0x01020304
	ptr := (*[4]byte)(unsafe.Pointer(&i))

	if ptr[0] == 0x04 {
		fmt.Printf("Primeiro byte é %v → sistema é Little-endian (ex: x86)\n", ptr[0])
	} else {
		fmt.Printf("Primeiro byte é %v → sistema é Big-endian\n", ptr[0])
	}
}
```

Saída esperada:

```
Primeiro byte é 4 → sistema é Little-endian (ex: x86)
```

### Conversão explícita usando o pacote encoding/binary

```go
package main

import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

func main() {
	var i uint32 = 0x01020304
	b := make([]byte, 4)

	binary.BigEndian.PutUint32(b, i)
	fmt.Printf("Big-endian: % x\n", b)

	binary.LittleEndian.PutUint32(b, i)
	fmt.Printf("Little-endian: % x\n", b)
}
```

Saída esperada:

```
Big-endian:    01 02 03 04
Little-endian: 04 03 02 01

```

### Manipulação de bytes individuais de acordo com o endianness

Aqui está um exemplo que mostra como o endianness influencia diretamente a forma como blocos de dados são interpretados em cifras de bloco como o RC2 (ou mesmo DES). Vamos simular a codificação de um bloco de 64 bits em que a ordem dos bytes impacta o valor numérico processado:

```go
package main

import (
	"encoding/binary"
	"fmt"
)

// Simula um processamento de bloco que depende da ordem dos bytes.
// Ex: RC2/DES usam blocos de 64 bits interpretados como dois inteiros de 32 bits.
func processBlock(b []byte, order binary.ByteOrder) {
	if len(b) != 8 {
		panic("Bloco precisa ter 8 bytes")
	}

	left := order.Uint32(b[:4])
	right := order.Uint32(b[4:])

	fmt.Printf("Bloco interpretado como:\n")
	fmt.Printf("  L = 0x%08X\n", left)
	fmt.Printf("  R = 0x%08X\n", right)
}

func main() {
	// Exemplo de bloco de 64 bits (8 bytes) com valor fixo
	block := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF}

	fmt.Println("== Interpretação Big-endian (protocolos de rede, SHA, etc.)")
	processBlock(block, binary.BigEndian)

	fmt.Println("\n== Interpretação Little-endian (como em x86 e x86_64)")
	processBlock(block, binary.LittleEndian)
}

```

Saída esperada:

```
== Interpretação Big-endian (protocolos de rede, SHA, etc.)
Bloco interpretado como:
  L = 0x01234567
  R = 0x89ABCDEF

== Interpretação Little-endian (como em x86 e xx86_64)
Bloco interpretado como:
  L = 0x67452301
  R = 0xEFCDAB89
```

Esse exemplo mostra que o mesmo bloco de bytes resulta em valores completamente diferentes para L e R, dependendo do endianness. Esse fator afeta diretamente muitos algoritmos de cifra — especialmente cifras de bloco — que tratam blocos de N bits como inteiros de 16 ou 32 bits, sobre os quais aplicam operações lógicas, rotações e substituições.

## Considerações

Endianness será um fator a ser considerado em diversas situações:

* Algoritmos como **SHA** e **MD5** especificam explicitamente a ordem dos bytes em suas operações internas.

* Protocolos de rede, como **TCP/IP**, utilizam o modelo **big-endian**, também chamado de _network byte order_.


Arquiteturas modernas como `x86` e `x86-64` (Intel/AMD) adotam o modelo `little-endian` como padrão. Isso significa que o byte menos significativo de uma palavra é armazenado no menor endereço de memória. Essa decisão de projeto afeta diretamente a forma como inteiros são representados internamente e exige atenção ao manipular dados provenientes de ambientes que utilizam big-endian, como redes ou formatos binários padronizados.





