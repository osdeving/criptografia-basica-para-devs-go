# Operadores Bitwise

## O que são operadores bit a bit?

Operadores bitwise (bit a bit) são operadores que atuam diretamente sobre os bits de inteiros binários. São fundamentais em implementações de algoritmos criptográficos, protocolos binários, manipulação de flags e otimizações de espaço e tempo.

Ao contrário dos operadores aritméticos tradicionais `(+, -, *, /)`, operadores bitwise operam posição a posição, diretamente sobre os bits das representações binárias dos operandos.

## Operadores Fundamentais

| Operador | Nome | Descrição |
| -------- | ---- | --------- |
| &        | AND  | Retorna 1 se ambos os bits forem 1. |
| \|       | OR   | Retorna 1 se pelo menos um dos bits é 1. |
| ^        | XOR  | Exclusivo: Retorna 1 se exatamente um dos bits é 1. |
| ~        | NOT  | Inverte todos os bits. |
| &^        | AND NOT (Go)  | Zera os bits do primeiro operando que forem 1 no segundo operando. |
| <<       | Deslocamento à esquerda | Desloca os bits n posições à esquerda. (equivalente a multiplicar por $2^n$) |
| >>       | Deslocamento à direita | Desloca os bits n posições à direita. (equivalente a dividir por $2^n$) |


A tabela a seguir ilustra como os operadores bitwise atuam sobre valores binários. Os exemplos utilizam inteiros de 8 bits para facilitar a visualização:

| Expressão         | Operação         | Resultado Binário | Resultado Decimal |
|------------------|------------------|-------------------|-------------------|
| `0b1100 & 0b1010` | AND              | `0b1000`          | `8`               |
| `0b1100 \| 0b1010`| OR               | `0b1110`          | `14`              |
| `0b1100 ^ 0b1010` | XOR              | `0b0110`          | `6`               |
| `^0b1100`         | NOT (unário)     | `0b...11110011`¹  | depende do tipo   |
| `0b0001 << 2`     | Shift à esquerda | `0b0100`          | `4`               |
| `0b1000 >> 2`     | Shift à direita  | `0b0010`          | `2`               |
| `0b1111 &^ 0b0101`| AND NOT (Go)     | `0b1010`²          | `10`              |

¹ Em Go, o operador `^` representa o NOT unário. O resultado depende do tipo da variável (`uint8`, `int`, etc.). Por exemplo, `^uint8(0b1100)` resulta em `0b11110011` (243 decimal).

² O operador `&^` é específico da linguagem Go e realiza uma operação conhecida como "bit clear". Ele equivale a `a & (~b)`, ou seja, **faz um NOT bit a bit do segundo operando e aplica AND com o primeiro**. Em C, essa operação seria expressa como `a & (~b)`.

## Exemplo prático 1: isolando bits com máscaras (Base64, SHA, etc)

Uma das aplicações mais comuns do operador AND (`&`) é a extração de porções específicas de um número binário, por meio de máscaras. Isso é amplamente usado em codificações como Base64, compressão, protocolos e algoritmos criptográficos.

Suponha que você queira extrair os 6 bits menos significativos de um byte. Para isso, podemos usar uma máscara `0b00111111`, que equivale a `0x3F` em hexadecimal. Veja:

```
BYTE qualquer:            01010010  (82 em decimal)
Máscara de 6 bits:        00111111  (0x3F em hexadecimal)
Resultado após AND:       00010010  (18 em decimal)

```

O operador `&` preserva apenas os bits em que ambos os operandos são 1. Portanto, ele serve como uma forma seletiva de “manter” ou “zerar” bits.

**Tabela verdade para `AND`**

| A | AND | B | Resultado |
|---|-----|---|-----------|
| 1 | &   | 1 | 1         |
| 1 | &   | 0 | 0         |
| 0 | &   | 1 | 0         |
| 0 | &   | 0 | 0         |

Com isso, chegamos à conclusão de que:

* Uma máscara 0b00000000 apaga todos os bits.

* Uma máscara 0b00000001 testa apenas o bit menos significativo.

* Para extrair bits específicos, basta construir uma máscara com 1s nas posições desejadas.

Por exemplo, podemos extrair os bits 3 e 4 de um byte da seguinte forma:

```
BYTE qualquer:            01010010  (82)
Máscara:                  00011000  (0x18)
Resultado após AND:       00010000  (16)
```


## Exemplo prático 2: usando bits como flags

Uma aplicação prática de operadores bitwise é o uso de flags, onde cada bit representa uma característica binária (ligado ou desligado). Esse padrão é comum em linguagens como C, Go e Rust, e em estruturas compactas de controle, como permissões ou estados.

Abaixo, usaremos o contexto do anime Naruto para ilustrar como os operadores bitwise podem ser usados para representar e manipular características das personagens.

### Definindo estilos de um ninja

Abaixo, cada constante representa uma posição de bit única:

```go
type NinjaRank uint8

const (
	Genin     NinjaRank = 1 << iota // 00000001
	Chunin                          // 00000010
	Jounin                          // 00000100
	Ambu                            // 00001000
	Sannin                          // 00010000
	Kage                            // 00100000
	Sage                            // 01000000
	Jinchuriki                      // 10000000
)
```

#### Combinando características com OR (`|`)

```go
naruto  := Genin  | Sage   | Jinchuriki       // 10100001
tsunade := Jounin | Sannin | Kage             // 00110100
kakashi := Jounin | Ambu                      // 00001100
```

**Tabela verdade para `OR`**

Para entender como a combinação funciona bit a bit, veja a tabela verdade da operação OR:

| A | OR | B | Resultado |
|---|----|---|-----------|
| 1 | \| | 1 | 1         |
| 1 | \| | 0 | 1         |
| 0 | \| | 1 | 1         |
| 0 | \| | 0 | 0         |

O operador `|` retorna `1` sempre que pelo menos um dos bits for `1`. Por isso ele é ideal para adicionar flags a uma variável sem interferir nos bits já definidos.

#### Testando características com AND (`&`)

```go
if naruto & Genin == Genin {
	fmt.Println("Naruto é Genin.")
}
```

#### Adicionando características com OR (`|`)

```go
naruto |= Jounin // promove Naruto para Jounin
```

#### Removendo características com AND NOT (`&^`)

```go
kakashi &^= Kage // remove o status de Kage de Kakashi
```

**Tabela verdade para `NOT`**

O operador NOT (`~`) inverte cada bit individualmente:

| A | NOT | Resultado |
|---|-----|-----------|
| 1 | ~   | 0         |
| 0 | ~   | 1         |


## Exemplo prático 3: usando `XOR` em PRNGs e criptografia

O operador XOR (`^`) é um dos mais utilizados em criptografia moderna, principalmente por suas propriedades de reversibilidade e difusão controlada. Ele também aparece como base de muitos geradores de números pseudoaleatórios (PRNGs), como a família **Xorshift**.

A propriedade central do `XOR` é:

```
A ^ A = 0        (auto-cancelamento)
A ^ 0 = A        (identidade)
A ^ B ^ B = A    (reversível)
```
Essas propriedades tornam o `XOR` ideal para cifragem, embaralhamento e geração determinística de entropia.

### Xorshift simplificado

A seguir, um PRNG de 32 bits usando apenas `XOR` e `shifts`:

```go
package main

import "fmt"

func xorshift32(x uint32) uint32 {
	x ^= x << 13
	x ^= x >> 17
	x ^= x << 5
	return x
}

func main() {
	seed := uint32(123456789)
	for i := 0; i < 5; i++ {
		seed = xorshift32(seed)
		fmt.Printf("Valor %d: %032b\n", i+1, seed)
	}
}
```

Esse gerador, embora simples, é capaz de produzir uma sequência pseudoaleatória com boa dispersão de bits. Ele é construído exclusivamente com operadores `^`, `<<` e `>>`.

**Tabela verdade para `XOR`**

| A | XOR | B | Resultado |
|---|-----|---|-----------|
| 1 | ^   | 1 | 0         |
| 0 | ^   | 0 | 0         |
| 1 | ^   | 0 | 1         |
| 0 | ^   | 1 | 1         |

A operação retorna `1` apenas quando os bits são diferentes. Por isso, é chamada de *OU exclusivo* (exclusive OR).


## Exemplo prático 4: deslocamento de bits (`<<` e `>>`)

Os operadores de deslocamento `<<` (shift à esquerda) e `>>` (shift à direita) são amplamente utilizados em operações de dispersão, ajustes de bits e estruturação de dados binários. Eles são comuns em geradores pseudoaleatórios, cifras e funções de hash.

A operação `x << n` desloca os bits de `x` para a esquerda em `n` posições, preenchendo com `0` à direita. Já `x >> n` desloca os bits para a direita, preenchendo com `0` à esquerda (em inteiros sem sinal).

### Aplicação prática: passo de um LFSR (Linear Feedback Shift Register)

Sem entrar nos detalhes do algoritmo agora, podemos observar um exemplo de uso de deslocamento para atualizar um valor de estado:

```go
package main

import "fmt"

func stepLFSR(state uint8) uint8 {
	feedback := ((state >> 0) ^ (state >> 2)) & 1
	state = (state >> 1) | (feedback << 7)
	return state
}

func main() {
	state := uint8(0b10010010)
	for i := 0; i < 5; i++ {
		state = stepLFSR(state)
		fmt.Printf("Estado %d: %08b\n", i+1, state)
	}
}
```

Neste exemplo:

* Um bit de realimentação é calculado via `XOR` de dois bits do estado.

* O estado é deslocado para a direita `(>> 1)`, descartando o LSB.

* O novo bit de feedback é colocado na posição mais à esquerda via feedback `<< 7`.

Embora não exista uma tabela verdade para o deslocamento, segue um exemplo visual do deslocamento:

| Operação | Entrada     | Resultado    | Observação |
|----------|-------------|--------------|------------|
| `x << 1` | <span style="color:#ff6666;">0</span>1010010 | 1010010<span style="color:#66ccff;">0</span> | Perde o MSB, ganha 0 no LSB |
| `x >> 1` | 0101001<span style="color:#ff6666;">0</span> | <span style="color:#66ccff;">0</span>0101001 | Perde o LSB, ganha 0 no MSB |
| `x >> 3` | 111<span style="color:#ff6666;">00000</span> | <span style="color:#66ccff;">000</span>11100 | Bits deslizam 3 casas à direita |


Esse comportamento é determinístico, barato computacionalmente e essencial para muitas transformações em cifras e PRNGs.


## Importância em Criptografia

Operadores bitwise são usados constantemente em algoritmos criptográficos. Exemplos:

* O `XOR` (`^`) é a operação mais usada em cifras de fluxo e blocos — por exemplo, no modo CTR ou no cálculo de paridade.

* Cifras de fluxo: cada byte da mensagem é combinado com um byte do keystream usando `XOR`.

* Funções hash: o `XOR` é usado para compressão, mistura e difusão de bits internos.

* `AND`, (`&`) `OR` (`|`), e `NOT` (`~`) são usados em compressão de mensagens, máscaras, e manipulação de estruturas internas como as de **SHA-2**.

* Shifts (`<<, >>`) implementam rotações e misturas, fundamentais em **S-Boxes** e funções de dispersão.


## Considerações

Dominar os operadores bitwise é essencial para compreender o funcionamento interno de muitos algoritmos criptográficos. Além disso, esse tipo de manipulação é eficiente, compacta e amplamente utilizada em implementações de baixo nível, seja em bibliotecas de segurança, protocolos ou microcontroladores.

Nos próximos tópicos, aplicaremos esses conceitos na manipulação de blocos de dados, operações estruturais e transformação binária direta.