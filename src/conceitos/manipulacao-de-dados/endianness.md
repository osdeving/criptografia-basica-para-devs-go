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

## Implementação em Go

Vamos verificar a ordem dos bytes em memória e converter entre endiannesses.

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

## Considerações

Saber lidar com endianness é fundamental ao manipular dados binários em protocolos, arquivos binários, criptografia ou comunicação entre sistemas heterogêneos. Por exemplo:

* Algoritmos como **SHA** e **MD5** especificam explicitamente a ordem dos bytes em suas operações internas.

* Protocolos de rede, como **TCP/IP**, utilizam o modelo **big-endian**, também chamado de _network byte order_.


Arquiteturas modernas como x86 e x86-64 (Intel/AMD) adotam o modelo little-endian como padrão. Isso significa que o byte menos significativo de uma palavra é armazenado no menor endereço de memória. Essa decisão de projeto afeta diretamente a forma como inteiros são representados internamente e exige atenção ao manipular dados provenientes de ambientes que utilizam big-endian, como redes ou formatos binários padronizados.





