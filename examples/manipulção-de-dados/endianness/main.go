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
