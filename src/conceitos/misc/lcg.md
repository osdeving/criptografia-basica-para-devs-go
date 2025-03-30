# LCG - Linear Congruential Generator

O gerador congruencial linear (LCG) é um Pseudo-Random Number Generator (PRNG) que produz uma sequência de números inteiros por meio de uma **fórmula de recorrência**, que define cada novo número com base no anterior.

$S_i+1 = A \cdot S_i + B \mod m$

Esta fórmula diz que, para obter o próximo valor da sequência $(S_i+1)$, pegamos o valor atual $(S_i)$, multiplicamos por um número fixo $A$, somamos um descolamento $B$, e aplicamos o módulo $m$ para que o resultado fique dentro de um intervalo limitado.

O que significa recorrência? **Recorrência** é um termo matemático para uma definição que depende do próprio valor anterior. É como uma função que chamamos com os resultados anteriores, p.ex.:

```go
S0 := 123456  // semente inicial
S1 := (A * S0 + B) % m
S2 := (A * S1 + B) % m
S3 := (A * S2 + B) % m
// ...
```
A sequência inteira $S_0, S_1, S_2, \dots$ é construída recursivamente a partir da semente inicial $S_0$.

Esse tipo de estrutura é comum em geradores determinísticos, pois com os mesmos parâmetros iniciais e a mesma semente, a sequência será sempre igual.

Os valores $S_i$ são números inteiros, mas o tamanho pode variar dependendo da aplicação. Em teoria, $S_i \in Z_m$, ou seja, um número de $0$ até $m-1$. Em código, os valores são frequentemente armazenados em um tipo de dado inteiro de 32 ou 64 bits para casos simples ou em um `big.Int` para suportar números maiores (e.g. um valor com 100 bits). Às vezes, o número inteiro é dividido em bits, caso a aplicação exigir uma saída bit a bit (como no caso de construir uma key stream para cifra de fluxo).

Os parâmetros fixos: $A, B, m$

* $A$ é chamado de multiplicador e controla o "rítmo" de variação da sequência.
* $B$ é chamado de incremento e adiciona uma constante a cada passo.
* $m$ é chamado de módulo e define o tamanho do espaço onde os valores vivem (e.g.: $m=2^{32}$ ou $m=2^{64}$ ou $m=2^{100}$).

O módulo $m$, além de manter os valores dentro de um intervalo finito (e.g. entre $0$ e $2^{32}-1$) também permite operar em espaços cíclicos, como nos corpos finitos em criptografia.

No código, podemos implementar uma função, convenientemente chamada rand, da seguinte forma:

```go
package main

type PRNG struct {
    state *big.Int // Estado atual S_i
    A     *big.Int // Multiplicador
    B     *big.Int // Incremento
    M     *big.Int // Módulo
}

func (p *PRNG) rand() *big.Int {
    p.state.Mul(p.A, p.state)      // A * S_i
    p.state.Add(p.state, p.B)      // A * S_i + B
    p.state.Mod(p.state, p.M)      // (A * S_i + B) mod m
    return new(big.Int).Set(p.state)
}

func main() {
    // Inicialização do PRNG
    A := big.NewInt(1103515245)
    B := big.NewInt(12345)
    M := big.NewInt(1 << 32)

    state := big.NewInt(123456)
    prng := PRNG{state, A, B, M}

    for i := 0; i < 10; i++ {
        fmt.Println(prng.rand())
    }
}
```

### Teorema de Hull-Dobell (para LCGs modulares)

Como veremos também em LFSR, é importante que um PRNG tenha um período máximo. Ou seja, depois de um certo número de iterações, a sequência deve começar a se repetir e o número de iterações deve ser o maior possível, nesse caso $m$. O **Teorema de Hull-Dobell** estabelece as condições necessárias e suficientes para um LCG modular ter um período máximo. Assim, para um LCG modular ter um período máximo, é necessário que:

* $m$ e $B$ são coprimos (i.e. $mdc(m, B) = 1$)
* $a - 1$ é divisível por todos os fatores primos de $m$
* se $m$ é múltiplo de 4, $a - 1$ também deve ser múltiplo de 4

A seguinte, os valores clássicos em C (glibc):

```c
A = 1103515245
B = 12345
M = 1 << 32
```
Para criptografia, esses critérios são insuficientes, porque o LCG é linear e, como veremos, pode ser atacado mesmo que atacante conheça apenas algumas poucas saídas. Para maior segurança precisamos de:

* Tornar a função não linear
* Ou esconder a saída (i.e. aplicar uma função hash sobre $S_i$)
* Ou usar um gerador de números aleatórios criptograficamente seguro (CSPRNG)


## Variações e Referências Clássicas (Knuth e outros)

O Gerador Congruencial Linear foi extensivamente estudado e documentado por Donald Knuth em sua obra fundamental The Art of Computer Programming, Volume 2: Seminumerical Algorithms [5]. Knuth fornece uma análise detalhada dos critérios estatísticos e estruturais que um bom gerador deve seguir, além de discutir limitações importantes como o ciclo curto, a baixa entropia de bits menos significativos e o comportamento previsível.

Ao longo dos anos, várias variações do LCG foram propostas para mitigar alguns desses problemas ou adaptar o algoritmo a restrições computacionais específicas. Algumas dessas variações incluem:

Gerador de Lehmer: também chamado de Multiplicative Congruential Generator, onde $B = 0$, ou seja, sem incremento. O foco é na escolha precisa de $A$ e $m$.

Método de Park-Miller: um caso específico do gerador de Lehmer com $A = 16807$, $m = 2^{31} - 1$. Bastante usado historicamente por seu bom desempenho e simplicidade.

Algoritmo de Schrage: técnica usada para evitar overflow em implementações do método de Park-Miller usando aritmética inteira. Divide o cálculo de $A \cdot S_i \mod m$ em partes menores.

LCGs combinados (ex: Wichmann-Hill): combinam múltiplos LCGs com diferentes parâmetros e somam as saídas para aumentar o período e melhorar propriedades estatísticas.

Miller’s LCG: variante pouco comum, mas citada por Knuth, com foco em evitar correlações entre bits.

Apesar das melhorias, todos esses métodos compartilham a característica de linearidade sobre $\mathbb{Z}_m$, o que os torna inadequados para aplicações criptográficas. Essa linearidade permite a reconstrução de parâmetros internos ou do estado com relativamente pouca informação (ex: alguns termos consecutivos da sequência). Por esse motivo, LCGs são aceitáveis apenas em contextos não criptográficos, como simulações estatísticas ou algoritmos probabilísticos onde segurança não é uma exigência.


### Caso de uso clássico em Assembly

```asm
Obs.: Essa implementação foi usada em um jogo 
escrito em assembly é uma variação do método de
Park-Miller (LCG multiplicativo), com constantes
específicas (16807 e 2836) e otimização inspirada 
no algoritmo de Schrage para evitar overflow.                                        
                                                                                                
Esteja à vontade para converter para Go
ou outra linguagem de alto nível. 


;---------------------------------------------------------;
; Função Random.                                          ;
; Parâmetros: DWORD teto.                                 ;
; Retorno:    Retorna o némero aleatério em EAX.          ;
; Descrição:  Gera um némero aleatério.                   ;
;---------------------------------------------------------;
Random:
    PUSH EBP
    MOV EBP, ESP
    SUB ESP, 4
	
    PUSH EDX
	
    CALL [GetTickCount]
    MOV DWORD [EBP - 4], EAX
	
    XOR EDX, EDX        
    PUSH 127773
    DIV DWORD [ESP]
    PUSH EAX
    MOV EAX, 16807    
    MUL EDX
    POP EDX
    PUSH EAX        
    MOV EAX, 2836              
    MUL EDX
    POP EDX                
    SUB EDX, EAX
    MOV EAX, EDX
    MOV DWORD [EBP - 4], EDX
    PUSH DWORD [EBP + 8]
    MOV EDX, 0
    DIV DWORD [ESP]
    ADD ESP, 8
    MOV EAX, EDX
	
    POP EDX
	
    MOV ESP, EBP
    POP EBP
    RET	4
```











