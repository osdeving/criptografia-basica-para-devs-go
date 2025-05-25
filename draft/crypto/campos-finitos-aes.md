# Operando em ${F}_{2^8}$

Praticamente todos os algoritmos de criptografia envolvem operações aritiméticas em inteiros, e se uma dessas operações é divisão, então precisamos de um corpo finito, visto que na divisão precisamos que todo inteiro diferente de zero tenha um inverso multiplicativo.

É conveniente trabalhar com inteiros que se encaixam em um número fixo de bits, por exemplo 8 bits, 16 bits, 32 bits, 64 bits, etc. Isso nos permite otimizar a implementação em hardware e facilitar a implementação em software.

Infelizmente, ao trabalhar com $Z_{2^n}$, não temos um corpo finito, pois nem todo inteiro tem um inverso multiplicativo. Por exemplo, $2$ não tem inverso multiplicativo em $Z_4$.

O algoritmo AES realiza todas as suas operações em bytes, ou seja, em elementos de 8 bits. Esses bytes são tratados como elementos de um **corpo finito** com $2^8 = 256$ elementos, denotado por ${F}_{2^8}$ ou $GF(2^8)$, também conhecido como **campo de Galois**.

## Breve Revisão de Corpos Finitos

Um **corpo** é uma estrutura algébrica onde é possível realizar as operações de adição, subtração, multiplicação e divisão (exceto por zero), e os resultados dessas operações permanecem dentro do conjunto. Para isso, é essencial que todo elemento não nulo possua um **inverso multiplicativo**.

O exemplo mais comum e intuitivo é o conjunto dos inteiros módulo $p$ onde $p$ é primo, denotado por ${Z}_p$. Esse conjunto contém os inteiros $\{0, 1, \dots, p-1\}$, com operações feitas módulo $p$. Por exemplo, em ${Z}_7$:

$$
2 \cdot 3 = 6 \mod 7 = 6
$$

Nesse corpo, o elemento $3$ tem inverso multiplicativo $5$, pois:

$$
3 \cdot 5 = 15 \equiv 1 \mod 7
$$

A condição fundamental é que $p$ seja primo. Se $n$ **não** for primo, como em ${Z}_{2^8}$, a estrutura não é um corpo. Por exemplo, em ${Z}_4$:

$$
2 \cdot 2 = 0 \mod 4
$$

Portanto, $2$ não tem inverso: é um **divisor de zero**.

## Extensão de Corpos

Vimos que, apesar de conveniente computacionalmente (8, 16, 32 bits), $Z_{2^n}$ **não é um corpo**. Então, para permitir inversão e garantir as propriedades de um corpo, o AES usa uma extensão de corpo finito, $F_{2^8}$.

Esse corpo é construído usando **polinômios sobre $F_2$** (isto é, coeficientes binários: 0 ou 1), reduzidos módulo um **polinômio irredutível** de grau 8. No AES, esse polinômio é:

$$
m(x) = x^8 + x^4 + x^3 + x + 1
$$

Um polinômio é **irredutível** se não pode ser fatorado como produto de dois polinômios de grau menor, o que garante que o quociente ${F}_2[x]/(m(x))$ é um corpo.

## Aritmética Modular com Polinômios

Na aritmética modular de inteiros (como ${Z}_7$), operamos com restos de divisão, informalmente conhecida como "aritmética do relógio".

Já na aritmética de ${F}_{2^8}$, operamos com **polinômios binários**. Por exemplo, o byte:

$$
\texttt{0x57} = 0101\,0111
$$

representa o polinômio:

$$
0\cdot x^7 + 1\cdot x^6 + 0\cdot x^5 + 1\cdot x^4 + 0\cdot x^3 + 1\cdot x^2 + 1\cdot x^1 + 1\cdot x^0 = x^6 + x^4 + x^2 + x + 1
$$

O valor binário indica a presença (1) ou ausência (0) de cada termo de grau correspondente.

### Exemplo de Multiplicação em ${F}_{2^8}$

Vamos multiplicar os bytes $\texttt{0x57}$ e $\texttt{0x83}$ em ${F}_{2^8}$, reduzido por $m(x) = x^8 + x^4 + x^3 + x + 1$.

#### Passo 1: Representar como polinômios

* $\texttt{0x57} = x^6 + x^4 + x^2 + x + 1$
* $\texttt{0x83} = x^7 + x + 1$

#### Passo 2: Multiplicação polinomial (sem módulo ainda)

$$
(x^6 + x^4 + x^2 + x + 1)(x^7 + x + 1)
$$

Expandindo até obter o polinômio resultante:

$$
x^{13} + x^{11} + x^9 + x^8 + x^7 + x^6 + x^5 + x^4 + x^2 + x + 1
$$

#### Passo 3: Redução módulo $m(x)$

Dividimos esse polinômio por $m(x)$ e obtemos o resto:

$$
\Rightarrow \text{Resultado final: } x^7 + x^6 + 1 = \texttt{0xC1}
$$

### Exemplo de Inversão (Divisão) em ${F}_{2^8}$

A divisão em ${F}_{2^8}$ é definida como multiplicação pelo inverso multiplicativo. Para encontrar o inverso de um elemento, usamos o algoritmo de Euclides (mdc) adaptado para polinômios sobre ${F}_2$.

Seja $a(x) = \texttt{0x53}$ (isto é, $x^6 + x^4 + x + 1$). Queremos encontrar $a^{-1}(x)$ tal que:

$$
a(x) \cdot a^{-1}(x) \equiv 1 \mod m(x)
$$

Aplicamos o algoritmo de Euclides estendido para polinômios binários. O resultado (omitido aqui por brevidade) é:

$$
a^{-1}(x) = \texttt{0xCA} = x^7 + x^6 + x^3 + x
$$

Verificação:

$$
\texttt{0x53} \cdot \texttt{0xCA} \mod m(x) = \texttt{0x01}
$$

## Implementando Operações em Go

A seguir um exemplo simples em Go para multiplicar dois bytes em ${F}_{2^8}$:

```go
func gfMul(a, b byte) byte {
    var res byte = 0
    for i := 0; i < 8; i++ {
        if (b & 1) != 0 {
            res ^= a
        }
        hiBitSet := (a & 0x80) != 0
        a <<= 1
        if hiBitSet {
            a ^= 0x1B // polinômio irredutível do AES sem o x^8
        }
        b >>= 1
    }
    return res
}
```

Para calcular o inverso multiplicativo, usamos o algoritmo de Euclides estendido para polinômios binários:

```go
func gfInv(a byte) byte {
    var t0, t1 byte = 0, 1
    var r0, r1 byte = 0x11B, a // 0x11B = m(x)

    for r1 != 0 {
        q := gfDiv(r0, r1)
        r0, r1 = r1, r0^gfMul(q, r1)
        t0, t1 = t1, t0^gfMul(q, t1)
    }

    return t0
}

func gfDegree(a byte) int {
    for i := 7; i >= 0; i-- {
        if (a & (1 << i)) != 0 {
            return i
        }
    }
    return -1
}

func gfDiv(dividend, divisor byte) byte {
    var quotient byte = 0

    degDividend := gfDegree(dividend)
    degDivisor := gfDegree(divisor)

    for degDividend >= degDivisor {
        shift := degDividend - degDivisor
        quotient ^= (1 << shift)
        dividend ^= (divisor << shift)
        degDividend = gfDegree(dividend)
    }

    return quotient
}

```

**Nota**: `gfDiv` aqui representa a operação de divisão entre polinômios binários, que é implementada por deslocamentos e subtrações (XOR) sucessivas.

## Para saber mais

Para compreender com profundidade o funcionamento de ${F}_{2^n}$, especialmente no contexto do AES, recomenda-se estudar os seguintes tópicos de matemática:

* Números primos e sua importância na teoria dos números
* Aritmética modular com inteiros (ex.: ${Z}_p$)
* Corpos finitos e suas extensões, como ${F}_{2^n}$
* Grupos, anéis e corpos na álgebra abstrata
* Polinômios irredutíveis e sua importância na construção de corpos
* Algoritmo de Euclides para inteiros (cálculo de mdc)
* Algoritmo de Euclides para polinômios (divisão em $\mathbb{F}_2[x]$)
* Algoritmo de Euclides estendido (para calcular inversos)
* Teoria de Galois e sua aplicação na estrutura dos corpos finitos
