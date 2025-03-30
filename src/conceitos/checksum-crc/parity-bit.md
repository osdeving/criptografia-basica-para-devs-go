# Bit de Paridade (Parity Bit)

O **bit de paridade** é um dos mecanismos mais simples de detecção de erros. Consiste em adicionar um único bit à mensagem original para que o total de bits com valor 1 na transmissão seja par (*paridade par*) ou ímpar (*paridade ímpar*).

Esse bit adicional permite detectar se ocorreu a inversão de um único bit durante a transmissão ou armazenamento. Contudo, não permite localizar o erro nem corrigi-lo, e não detecta alterações que envolvam um número par de bits.

---

## Paridade Par

Garante que o número total de bits 1 na mensagem, incluindo o bit de paridade, seja par.

Exemplo:

```
Mensagem original: 10101110 (8 bits)
Número de bits "1": 5 (impar)
Bit de paridade: 1  → Total de bits "1" passa a ser 6 (par)
Mensagem transmitida: 101011101
```

## Paridade Ímpar

Garante que o total de bits 1 seja ímpar.

```
Mensagem original: 10101110
Número de bits "1": 5 (impar)
Bit de paridade: 0  → Total de bits "1" continua ímpar
Mensagem transmitida: 101011100
```

---

## Verificação

O receptor recalcula a paridade da mensagem recebida (incluindo o bit de paridade) e verifica se o valor é coerente com a paridade esperada (par ou ímpar).

Exemplo de erro detectado:

```
Transmissão esperada: 101011101 (paridade par)
Erro no terceiro bit:
Recebido: 101001101
Número de bits "1": 4 (paridade inesperada) → erro detectado
```

Erro não detectado (dois bits invertidos):

```
Erro nos bits 2 e 6:
Original: 101011101
Recebido: 111001101
Número de bits "1": 6 (mesma paridade) → erro não detectado
```

---

## Implementação em Go

```go
package main

import (
	"fmt"
	"math/bits"
)

// Calcula paridade par (retorna 0 ou 1)
func parityBit(data byte) byte {
	if bits.OnesCount8(data)%2 == 0 {
		return 0 // já é par
	}
	return 1
}

// Verifica a paridade de um byte + bit extra
func checkParity(data byte, parity byte) bool {
	return parityBit(data) == parity
}

func main() {
	data := byte(0b10101110) // 5 bits 1
	parity := parityBit(data)
	fmt.Printf("Paridade calculada: %d\n", parity)

	valid := checkParity(data, parity)
	fmt.Printf("Verificação: %v\n", valid)

	// Simulando erro
	corrupted := data ^ 0b00000100 // inverte o bit 2
	valid = checkParity(corrupted, parity)
	fmt.Printf("Verificação após erro: %v\n", valid)
}
```

---

## Limitações e Aplicabilidades

- Detecta apenas **um erro de bit**.
- Não localiza nem corrige o erro.
- Falha em detectar erros com número par de bits invertidos.
- É adequado para ambientes com ruído limitado e simplicidade de hardware.

### Exemplos de uso:
- Memórias RAM com detecção simples
- Protocolos seriais (UART, RS-232)
- Dispositivos com restrição de custo e complexidade
- Parte de códigos de correção como Hamming

---

## Considerações finais

O bit de paridade ilustra o conceito central da detecção de erros com custo mínimo de redundância. Apesar de suas limitações, é historicamente importante e ainda aplicável em contextos restritos. Os próximos métodos que estudaremos, **Checksums** e **CRC**, vão superar essas limitações oferecendo maior poder de detecção sem aumento significativo de complexidade.