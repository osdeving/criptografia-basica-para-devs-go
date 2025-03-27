# Endianness - Organização de Bytes na Memória

## O que é Endianness?

Endianness se refere à ordem em que os bytes são armazenados na memória. Essa ordem pode variar entre sistemas diferentes e impacta como os dados binários são interpretados.

Existem dois tipos principais: Little-endian e Big-endian.

## Little-endian

Nesse modelo, o byte menos significativo (LSB - Least Significant Byte) vem primeiro. Vamos supor que temos o número 0x12345678 armazenado em 4 bytes. O byte menos significativo (LSB) é o 0x78, seguido pelo 0x56, 0x34 e 0x12. A representação em memória seria:

Endereço | Byte
--- | ---
0x1000 | 0x78
0x1001 | 0x56
0x1002 | 0x34
0x1003 | 0x12


## Big-endian

Neste modelo, o byte mais significativo (MSB - Most Significant Byte) vem primeiro. Novamente, usando o mesmo exemplo, o byte mais significativo (MSB) é o 0x12. A representação em memória seria:

Endereço | Byte
--- | ---
0x1000 | 0x12
0x1001 | 0x34
0x1002 | 0x56
0x1003 | 0x78

## Implementação em Go

Vamos ver como podemos implementar esses conceitos em Go.

### Identificando o Endianness

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
		fmt.Printf("Primeiro byte é %v, então este sistema é Little-endian (provavelmente um x86)\n", ptr[0])
	} else {
		fmt.Printf("Primeiro byte é %v, então este sistema é Big-endian\n", ptr[0])
	}

	b := make([]byte, 4)

	// Converte o inteiro para bytes no formato big-endian
	binary.BigEndian.PutUint32(b, uint32(i))

	fmt.Println("Convertido para Big-endian")

	// Verifica se o primeiro byte é 0x04 (LSB) ou 0x01 (MSB)
	if b[0] == 0x04 {
		fmt.Printf("Primeiro byte é %v, então é Little-endian\n", b[0])
	} else {
		fmt.Printf("Primeiro byte é %v, então é Big-endian\n", b[0])
	}
}

```

Vai imprimir:

```
Primeiro byte é 4, então é Little-endian
```







