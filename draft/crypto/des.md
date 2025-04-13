usa blocos de 64 bits
usa chaves de 56 bits
usa 16 rodadas de cifragem para cada bloco
cada rodada usa uma chave diferente derivada da chave principal
usa uma estrutura de Feistel (usada em vários algoritmos, mas não todos, p.ex.: AES não é uma cifra de Feistel)
vantagem do feistel: processo de cifrar e decifrar é praticamente o mesmo, mudando a ordem das subchaves
primeiro temo uma permutação inicial sobre o bloco de 64 bits dividido em duas partes: L0 e R0
L0 e R0 são a entrada da rede de Feistel que consiste em 16 rodadas
Ri é passado pela função f e o resultado é feito XOR com a outra metade esquerda Li
Então as metades são trocadas

```
Lᵢ = Rᵢ₋₁
Rᵢ = Lᵢ₋₁ ⊕ f(Rᵢ₋₁, kᵢ)

para i = 1, ..., 16.
```

```go
L[0] = L0
R[0] = R0

for i := 1; i <= 16; i++ {
    L[i] = R[i-1]
    R[i] = xor(L[i-1], f(R[i-1], K[i-1]))
}
```

cada rodada a chave ki é derivada da chave principal k de  56 bits via key schedule
a metade direita não é cifrada, ela é usada para cifrar a metade esquerda e se torna a metade esquerda da rodada seguinte

(a função f pode ser visto como um PRNG que recebe Ri e ki, se retorno serve para embaralhar Li via XOR. A saída de f precisa ser imprevisível)


A estrutura Feistel como um todo é uma função bijetora nos 64 bits de entrada:
Isso significa que cada entrada de 64 bits gera exatamente uma saída distinta de 64 bits, e toda saída pode ser revertida à entrada original (a função tem inversa)
A função f não é injetora, i.e. ela pode gerar a mesma saída para entradas Li e ki diferentes.



# A Função $f$

A função $f$ é o núcleo da transformação não linear do DES e é aplicada 16 vezes, uma em cada rodada da cifra. Seu papel é operar sobre a metade direita do bloco atual ($R_{n-1}$) em conjunto com uma subchave de 48 bits ($K_n$), produzindo um bloco intermediário de 32 bits que será misturado com a outra metade do bloco por meio de uma operação XOR.

A expressão geral da rodada é:

\[
R_n = L_{n-1} \oplus f(R_{n-1}, K_n)
\]

A força do DES está concentrada na complexidade da função $f$, que combina expansião de bits, operações XOR, substituição não linear (S-boxes) e uma permutação fixa.

## Etapas internas da função $f$

A função $f$ recebe uma entrada de 32 bits ($R_{n-1}$) e uma subchave de 48 bits ($K_n$). A seguir, aplica as seguintes transformações:

### 1. Expansão (E)
A primeira etapa da função $f$ é a expansão de $R_{n-1}$ de 32 para 48 bits, realizada pela chamada **E-box**. Essa transformação utiliza uma tabela fixa de 48 posições que reordena e duplica determinados bits da entrada. A expansão é projetada para criar sobreposições entre blocos adjacentes de 4 bits, aumentando a difusão e tornando mais eficaz a mistura posterior com a subchave:

- **Entrada**: bloco de 32 bits (metade direita do bloco original)
- **Tabela**: vetor de 48 índices no intervalo [1, 32]
- **Saída**: bloco de 48 bits, com repetição de alguns bits

Essa expansão permite que bits nos limites de grupos de 4 bits se repitam em mais de um bloco de 6 bits, o que aumenta a interdependência entre as partes processadas pelas S-boxes.

Como analogia simplificada, suponha uma entrada de 6 bits representada por $[A, B, C, D, E, F]$, e uma tabela de expansão:

```text
Tabela: [6, 1, 2, 3, 4, 5, 4, 5]
Saída: [F, A, B, C, D, E, D, E]
```

Note que alguns bits da entrada aparecem mais de uma vez na saída. Esse comportamento reflete o mesmo princípio da expansão real do DES. A tabela de expansão completa será apresentada em uma seção posterior.

### 2. XOR com a subchave
A saída de 48 bits da expansão é combinada bit a bit com a subchave $K_n$:

\[
B = E(R_{n-1}) \oplus K_n
\]

O resultado $B$ é um vetor de 48 bits, dividido em oito blocos de 6 bits cada: $B_1, B_2, …, B_8$.

### 3. Substituição (S-boxes)
Cada bloco $B_i$ é passado por uma caixa de substituição ($S_i$), que mapeia 6 bits de entrada em 4 bits de saída de forma não linear. Esse é o ponto central da não linearidade da cifra:

\[
S_i: \{0,1\}^6 \rightarrow \{0,1\}^4
\]

Ao final, os oito blocos de 4 bits são concatenados, resultando em um bloco de 32 bits.

### 4. Permutação (P)
O bloco de 32 bits gerado pelas S-boxes é então reorganizado segundo uma permutação fixa de 32 posições. Essa etapa finaliza a função $f$, garantindo maior difusão dos bits antes que o resultado seja combinado com $L_{n-1}$.

## Resumo do fluxo

```text
R_{n-1} (32 bits)
       │
   Expansão E
       ▼
E(R_{n-1}) (48 bits)
       │
XOR com K_n (48 bits)
       ▼
    B = 48 bits
┌────┬────┬────┬────┬────┬────┬────┬────┐
│ B₁ │ B₂ │ B₃ │ B₄ │ B₅ │ B₆ │ B₇ │ B₈ │
└────┴────┴────┴────┴────┴────┴────┴────┘
       │     (cada um com 6 bits)
       ▼
    S-boxes (S₁ a S₈)
       ▼
   8 × 4 bits = 32 bits
       │
 Permutação P
       ▼
Saída final de f (32 bits)
```

Esse processo assegura que pequenas mudanças em $R_{n-1}$ ou $K_n$ se propaguem de forma complexa e imprevisível, contribuindo significativamente para a segurança do DES. As definições precisas das tabelas de expansão, substituição e permutação são apresentadas nas próximas seções.

