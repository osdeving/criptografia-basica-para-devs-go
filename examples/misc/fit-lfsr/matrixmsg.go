package main

import (
	"fmt"
	"time"
)

func typewriter(s string, delay time.Duration) {
	for _, ch := range s {
		fmt.Printf("%c", ch)
		time.Sleep(delay)
	}
	fmt.Println()
}

func main() {
	typewriter("Knock, knock, Neo.", 200*time.Millisecond)
	time.Sleep(1 * time.Second)
	typewriter("Follow the white rabbit.", 200*time.Millisecond)

}
