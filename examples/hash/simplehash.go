package main

import (
    "fmt"
)

func simpleHash(in string) int {
    hash := 0
    for _, char := range in {
        hash += int(char)
    }
    return hash % 256
}

func main() {
	input1 := "willams sousa"
	input2 := "Go"
	input3 := "xyz"

    fmt.Printf("Hash de '%v':%v\n", input1, simpleHash(input1))
    fmt.Printf("Hash de '%v':%v\n", input2, simpleHash(input2))
    fmt.Printf("Hash de '%v':%v\n", input3, simpleHash(input3))
}