# Fundamentos matemáticos do AES

## Introdução

O AES realiza todas as suas operações fundamentais sobre um corpo finito, especificamente o corpo $\mathbb{F}_{2^8}$. Compreender a aritmética nesse corpo é essencial para entender o comportamento das transformações internas do algoritmo, como **SubBytes**, **MixColumns** e a geração de subchaves (Key Schedule).

Nesta seção, faremos uma introdução rigorosa aos conceitos de corpos finitos, representação de polinômios binários, operações aritméticas com redução modular, e mostraremos detalhadamente como essas estruturas matemáticas se aplicam no projeto do AES.

---

## Corpos finitos: visão geral

Um **corpo finito** (ou campo finito) é um conjunto finito de elementos no qual estão definidas duas operações, **adição** e **multiplicação**, satisfazendo as propriedades:

* Existência de identidade aditiva e multiplicativa.
* Existência de inversos aditivo e multiplicativo.
* Associatividade, comutatividade e distributividade das operações.

Os corpos finitos são denotados por $\mathbb{F}_{p^n}$, onde:

* $p$ é um número primo, chamado de característica do corpo.
* $n \geq 1$ é um inteiro, representando a extensão.

O número total de elementos no corpo é $p^n$. No caso do AES:

* $p = 2$
* $n = 8$
* Portanto, $\mathbb{F}_{2^8}$ possui 256 elementos.

Esse corpo é amplamente utilizado em criptografia devido à eficiência de suas operações com bits.

---

## Construção de $\mathbb{F}_{2^8}$

A construção de $\mathbb{F}_{2^8}$ é feita como uma **extensão de corpo** de $\mathbb{F}_2$. Usamos polinômios de grau até 7 com coeficientes em $\mathbb{F}_2$. Para que essa estrutura seja um corpo (e não apenas um anel), definimos as operações módulo um **polinômio irreducível** de grau 8.

O polinômio irreducível usado no AES é:
$m(x) = x^8 + x^4 + x^3 + x + 1$

Esse polinômio não pode ser fatorado sobre $\mathbb{F}_2$, o que garante que o quociente $\mathbb{F}_2[x]/(m(x))$ seja um corpo.

---

## Representação de elementos

Cada elemento de $\mathbb{F}_{2^8}$ é um polinômio de grau $\leq 7$, com coeficientes em $\{0,1\}$. Podemos associar diretamente cada byte (8 bits) a um desses polinômios:

Exemplo:

```plaintext
01100011 = x^6 + x^5 + x + 1 → 0x63
```

A adição de elementos é feita com XOR bit a bit. Já a multiplicação exige manipulação polinomial com redução modular por $m(x)$.

---

## Operações em $\mathbb{F}_{2^8}$

### Adição

A operação de adição em $F_{2^8}$ é realizada por meio do operador XOR bit a bit, pois os coeficientes dos polinômios pertencem a $F_2 = {0,1}$. Isso significa que:

$$
a(x) + b(x) = a(x) \oplus b(x)
$$

A notação $a(x)$ deixa explícito que os operandos são **polinômios** com coeficientes binários. Cada byte representa um polinômio de grau até 7, e a operação XOR corresponde à soma de polinômios módulo 2, coeficiente a coeficiente.

Por exemplo:

```plaintext
0x57 = x⁶ + x⁴ + x² + x + 1
0x83 = x⁷ + x + 1
0x57 ⊕ 0x83 = x⁷ + x⁶ + x⁴ + x² → 0xD4
```

Essa operação é:

* **Comutativa**: $a \oplus b = b \oplus a$
* **Associativa**: $(a \oplus b) \oplus c = a \oplus (b \oplus c)$
* **Auto-inversa**: $a \oplus a = 0$

Não há necessidade de redução modular, pois a adição nunca altera o grau dos polinômios.


### Multiplicação

A multiplicação em $F_{2^8}$ envolve duas etapas principais: multiplicação polinomial binária e redução modular. Cada elemento do corpo é interpretado como um polinômio de grau no máximo 7 com coeficientes em $F_2$.

#### Passo 1: multiplicação binária

Multiplicamos os dois polinômios usando as regras normais de álgebra, mas com coeficientes módulo 2 (ou seja, toda soma entre coeficientes é feita com XOR).

Por exemplo, considere:

$$
a(x) = x^6 + x\\ 
b(x) = x^3 + 1
$$

A multiplicação direta resulta em:

$$
a(x) \cdot b(x) = (x^6 + x)(x^3 + 1) = x^9 + x^6 + x^4 + x
$$

Este polinômio tem grau 9, ou seja, não pertence a $F_{2^8}$.

#### Passo 2: redução módulo $m(x)$

Para manter o resultado dentro de $F_{2^8}$, realizamos a **divisão polinomial** do produto por um polinômio irreducível de grau 8. O AES define:

$$
m(x) = x^8 + x^4 + x^3 + x + 1
$$

Calculamos o resto da divisão do produto por $m(x)$, obtendo um polinômio de grau no máximo 7. Esse será o resultado da multiplicação em $F_{2^8}$.

#### Implementação eficiente

Na prática, a multiplicação por constantes fixas como 0x02, 0x03, 0x09 etc., usadas no AES (por exemplo em MixColumns), é implementada com:

* **Tabelas de lookup (log/antilog)** para multiplicações genéricas.
* **Algoritmos otimizados por bit shifting e XOR**, já que multiplicar por $x$ (ou seja, por 0x02) equivale a um deslocamento à esquerda seguido de redução condicional.

Por exemplo, para multiplicar por 0x02:

* Se o bit mais significativo (MSB) do byte é 0: apenas faça shift à esquerda.
* Se o MSB é 1: faça shift à esquerda e depois XOR com o byte ${1B}$ (que representa $x^4 + x^3 + x + 1$, parte do $m(x)$ sem o termo $x^8$).

Esse comportamento implementa a redução módulo $m(x)$ de forma eficiente sem realizar a divisão explícita.

#### Observação

Multiplicações em $F_{2^8}$ **não são comutativas** com relação a otimizações: multiplicar por ${02}$ é simples, mas multiplicar por ${0F}$ é bem mais custoso — por isso as matrizes do MixColumns foram escolhidas cuidadosamente para permitir implementação eficiente sem comprometer a difusão.

Essa estrutura matemática torna o AES eficiente em hardware e software, além de seguro contra criptoanálises baseadas em estruturas algébricas fracas.

---

## Aplicações no AES

### SubBytes: inverso multiplicativo + transformação afim

Para cada byte (exceto 0), calcula-se seu **inverso multiplicativo** em ${F}_{2^8}$. Por definição:
$a \cdot a^{-1} = 1 \mod m(x)$

Essa operação é essencial para garantir **não-linearidade forte**, resistindo à criptoanálise linear e diferencial. Após encontrar o inverso, aplica-se uma transformação afim bit a bit:

$S(x) = A \cdot x^{-1} + b$

onde:

* $A$ é uma matriz 8×8 sobre ${F}_2$
* $b$ é um vetor constante

Essa combinação gera a S-box do AES.

### MixColumns: multiplicação matricial

Cada coluna do *state* é considerada um vetor de 4 bytes, multiplicado por uma matriz fixa de coeficientes em ${F}_{2^8}$:

$$
M =
\begin{bmatrix}
02 & 03 & 01 & 01 \\
01 & 02 & 03 & 01 \\
01 & 01 & 02 & 03 \\
03 & 01 & 01 & 02
\end{bmatrix}
$$

Cada elemento da saída resulta de multiplicações e somas (XORs) em ${F}_{2^8}$. Isso garante **difusão vertical**, espalhando bits de um byte original em todos os bytes da coluna.

### Key Schedule: geração de subchaves

As subchaves são geradas com base na chave principal, através de:

* **RotWord:** rotação cíclica de uma palavra de 4 bytes.
* **SubWord:** aplicação da S-box aos 4 bytes (exige inversos em ${F}_{2^8}$).
* **Rcon\[i]:** vetor de constantes com valores $2^{i-1}$ em ${F}_{2^8}$, com redução por $m(x)$.

As palavras $w_i$ do key schedule são construídas recursivamente:
$w_i = w_{i-1} \oplus w_{i-N_k} \quad \text{(com modificações a cada múltiplo de } N_k)$

---

## Conclusão

A matemática por trás do AES — centrada na aritmética de corpos finitos ${F}_{2^8}$ — não é apenas formal, mas prática e essencial. Cada operação do algoritmo foi projetada para equilibrar eficiência, simetria e segurança. Compreender como os bytes se comportam como polinômios e como a redução por um polinômio irreducível define o corpo é indispensável para qualquer análise ou implementação séria do AES.

A profundidade matemática do AES mostra que criptografia moderna é inseparável da álgebra abstrata — uma das razões pelas quais o algoritmo permanece seguro e relevante mesmo décadas após sua padronização.
