# LFSR – Linear Feedback Shift Register

O LFSR (Registrador de Deslocamento com Realimentação Linear em português) é uma estrutura fundamental na geração de números pseudoaleatórios (PRNGs), além de desempenhar papel central em técnicas de criptografia de fluxo. Nesta seção vamos entender como o LFSR funciona e como ele é usado para gerar números aleatórios. Ao final, implementaremos um LFSR em Go para ilustrar seu funcionamento simulando as funções srand e rand tais como são definidas na biblioteca padrão de C. O LFSR não é um algorítmo seguro, mas é uma ferramenta útil para entender como os PRNGs funcionam.

## Funcionamento

Um LFSR é composto por um vetor de bits que representa o estado interno. A cada iteração, os bits do registrador são deslocados à direita, e um novo bit entra pela extremidade esquerda. Esse novo bit, chamado de feedback, é calculado como o XOR de certos bits do estado atual. Os bits utilizados nesse cálculo são denominados taps (ou pontos de realimentação).

A escolha dos taps não é arbitrária. Eles devem corresponder a um polinômio primitivo sobre o corpo finito $\mathbb{F}_2$, garantindo que o LFSR atinja seu período máximo, ou seja, percorra todos os $2^n - 1$ estados possíveis (exceto o estado nulo).

## Taps e Polinômios

Para um LFSR de $n$ bits, os taps definem um **polinômio binário** da forma:

$$
P(x) = x^n + a_{n-1}x^{n-1} + \dots + a_1x + 1
$$

Por exemplo, o polinômio:

$$
P(x) = x^4 + x + 1
$$

representa um LFSR de 4 bits com taps nas posições 4 e 1 (i.e. índices 3 e 0, respectivamente).

Se esse polinômio for **primitivo**, o LFSR terá **período máximo**, ou seja, irá percorrer todos os $2^n - 1$ estados possíveis distintos (exceto o estado 0000, que leva ao ciclo trivial de zeros).


## Operação Passo a Passo

A seguir, demonstramos o funcionamento de um LFSR de 4 bits com estado inicial 1001 e taps nas posições 4 e 1 (índices 3 e 0):

- Passo 1

   Estado: `1001`.

   Feedback: `1 ⊕ 1 = 0`.

   Novo estado: `0100`.

- Passo 2

   Estado: `0100`.

   Feedback: `0 ⊕ 0 = 0`.

   Novo estado: `0010`.


- Passo 3

   Estado: `0010`.

   Feedback: `0 ⊕ 0 = 0`.

   Novo estado: `0001`.


- Passo 4

   Estado: `0001`.

   Feedback: `0 ⊕ 1 = 1`.

   Novo estado: `1000`.


- Passo 5

   Estado: `1000`.

   Feedback: `1 ⊕ 0 = 1`.

   Novo estado: `1100`.


E assim por diante.   

### Tabela de Transiçãoo de Estados

| Passo | Estado   | Feedback     | Novo Estado |
|-------|----------|--------------|-------------|
| 1     | 1001     | 1 ⊕ 1 = 0    | 0100        |
| 2     | 0100     | 0 ⊕ 0 = 0    | 0010        |
| 3     | 0010     | 0 ⊕ 0 = 0    | 0001        |
| 4     | 0001     | 0 ⊕ 1 = 1    | 1000        |
| 5     | 1000     | 1 ⊕ 0 = 1    | 1100        |
| ...   | ...      | ...          | ...         |

A operação em cada passo pode ser descrita da seguinte forma:

1. O feedback é calculado aplicando XOR entre os bits nas posições definidas pelos taps. (p.ex.: 1001 os taps serão os bits 3 e 0, então 1 ⊕ 1 = 0)

2. O estado é deslocado uma posição à direita. (p.ex.: o estado 1001 ao descolar para direita se torna 0100)

3. O feedback é inserido como o novo bit mais à esquerda. (p.ex.: o estado 0100, já deslocado à direita, vai receber o feedback 0)

Esse processo é iterado indefinidamente, e a sequência de estados dependerá da escolha inicial da semente (estado) e dos taps. Quando os taps não forem escolhidos adequadamente, o LFSR pode entrar em ciclos curtos, ou até mesmo repetir rapidamente os estados, falhando em percorrer todos os $2^n - 1$ estados possíveis.

## LFSRs de Tamanhos Diferentes

| Tamanho | Taps        | Polinômio                  | Período máximo     |
|---------|-------------|----------------------------|--------------------|
| 4 bits  | [3, 0]      | $x^4 + x + 1$              | $2^4 - 1 = 15$     |
| 8 bits  | [7, 5]      | $x^8 + x^6 + 1$            | $255$              |
| 16 bits | [15, 13]    | $x^{16} + x^{14} + 1$      | $2^{16} - 1$       |
| 32 bits | [31, 21]    | $x^{32} + x^{22} + 1$      | $2^{32} - 1$       |


## Importância dos Taps

Taps mal escolhidos podem resultar em:

- Períodos muito curtos
- Ciclos triviais (repetição precoce)
- Saídas estatisticamente fracas

Por isso, os taps devem ser derivados de polinômios **primitivos** cuidadosamente estudados. Um polinômio primitivo é um polinômio irredutível sobre o corpo finito $\mathbb{F}_2$, o que significa que não pode ser fatorado em polinômios de menor grau sobre $\mathbb{F}_2$.


## Fórmula Geral
<br>


<style>
/* Estilo base para caixas */
.box {
  border: 1px solid var(--fg);
  border-left: 6px solid var(--fg);
  padding: 1rem;
  margin: 1.5rem 0;
  background-color: var(--bg);
  box-shadow: 2px 2px 6px var(--shadow);
  border-radius: 6px;
  font-size: 1.5rem;
}

/* Definição (azul adaptado) */
.box.definition {
  border-left-color: var(--primary);
}

.box.definition::before {
  content: "Definição";
  font-weight: bold;
  display: block;
  margin-bottom: 0.5rem;
  color: var(--primary);
  font-size: 1.5rem;
}

  
  </style>

<div class="box definition">

O texto puro, o texto cifrado e o key stream consiste de bits individuais:

$$
x_i, y_i, s_i \in \{0, 1\}
$$

**Cifragem:**
$$
y_i = e_{s_i}(x_i) \equiv x_i + s_i \mod 2
$$

**Decifragem:**
$$
x_i = d_{s_i}(y_i) \equiv y_i + s_i \mod 2
$$

</div>

## Implementações

### Gerador de números pseudo-aleatórios (srand e rand)

```go
package main

import (
	"fmt"
)

type PRNG struct {
	state uint8
	taps  []int
}

// inicializa com uma semente
func srand(seed uint8) *PRNG {
	return &PRNG{state: seed, taps: []int{7, 5}}
}

// gera um único bit pseudo-aleatório
// é como jogar uma moeda para escolher 0 ou 1, só que deterrminístico
func (p *PRNG) nextBit() uint8 {
	var feedback uint8
	
	for _, tap := range p.taps {
		feedback ^= (p.state >> tap) & 1
	}

	out := p.state & 1

	p.state = (p.state >> 1) | (feedback << 7)
	return out
}

// gera um número pseudo-aleatório de 8 bits
// obtido a partir de 8 chamadas de nextBit()
func (p *PRNG) rand() uint8 {
	var b uint8
	
	for i := 0; i < 8; i++ {
		b |= p.nextBit() << i
	}
	return b
}

func main() {

   // inicializa o gerador com uma semente
	prng := srand(0b11001010)

	fmt.Println("Gerando 10 números pseudo-aleatórios com LFSR:")
	for i := 0; i < 10; i++ {
		fmt.Printf("rand() = %3d (0b%08b)\n", prng.rand(), prng.rand())
	}
}

```
### Explicação


No loop onde percorrremos os taps, o feedback é calculado aplicando XOR entre os bits nas posições definidas pelos taps. É equivalente a:

feedback = 0 XOR bit5 (pega o primeito tap, que sabemos que é 5)

feedback = bit5 XOR bit7 (pega o segundo tap, que sabemos que é 7, e faz XOR com o resultado anterior, que sabemos que é bit5)

O estado é atualizando da seguinte forma:

Desloca o estado para a direita e insere o feedback no bit mais significativo. P.ex.: 

10000001 >> 1 = 01000000 (desloca o estado para a direita, perdendo o bit menos significativo e ganhando um 0 no bit mais significativo)

00000001 << 7 = 10000000 (o feedback é sempre apenas 0 ou 1 e não precisamos fazer AND com 1, o deslocamento joga esse bit para o bit mais significativo)

01000000 | 10000000 = 11000000 (faz um OR bit a bit entre o estado deslocado e o feedback deslocado, e obtém o novo estado)

Na função rand(), o loop percorre os 8 bits do número pseudo-aleatório e obtém cada bit usando a função nextBit().

O OR é feito sempre entre os bits obtidos e o valor de b, que é inicializado com 0 e cada bit é deslocado para sua posição correta (0 a 7)
